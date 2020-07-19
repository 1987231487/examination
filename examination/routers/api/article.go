package api

import (
	"examination/models"
	"examination/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface {}
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

//获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	/*var state int = 0
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var userid int = -1
	if arg := c.Query("user_id"); arg != "" {
		userid = com.StrTo(arg).MustInt()
		maps["user_id"] = userid

		valid.Min(userid, 1, "tag_id").Message("用户ID必须大于0")
	}*/

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetArticles( maps)
		data["total"] = models.GetArticleTotal(maps)  //文章总数

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

//新增文章
func AddArticle(c *gin.Context) {
	userId := com.StrTo(c.Query("user_id")).MustInt()
	title := c.Query("title")
	text := c.Query("text")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createby:=c.Query("created_by")

	valid := validation.Validation{}
	valid.Min(userId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(text, "content").Message("内容不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {

			data := make(map[string]interface {})
			data["user_id"] = userId
			data["title"] = title
			data["text"] = text
			data["state"] = state
			data["created_by"]=createby
			models.AddArticle(data)
			code = e.SUCCESS

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})
}

//修改文章
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	userId := com.StrTo(c.Query("user_id")).MustInt()
	title := c.Query("title")
	text := c.Query("text")
	updatedBy:=c.Query("updated_by")

	var state int = 0
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(text, 65535, "text").Message("内容最长为65535字符")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {

				data := make(map[string]interface {})
				if userId > 0 {
					data["user_id"] = userId
				}
				if title != "" {
					data["title"] = title
				}
				data["state"]=state
				if text != "" {
					data["text"] =  text
				}
				data["updated_by"]=updatedBy

				models.EditArticle(id, data)
				code = e.SUCCESS

		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}

func ChangeState(c *gin.Context){
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	var state int =0
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {

			data := make(map[string]interface {})
			data["state"]=state
			models.EditArticle(id, data)
			code = e.SUCCESS

		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}

 }
 c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}