package api

import (
	"examination/models"
	"examination/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetComment(c *gin.Context){
	data := make(map[string]interface{})

	id:=com.StrTo(c.Param("id")).MustInt()


	data["comments"]=models.GetComments(id)
	code:=e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}
func AddComment(c *gin.Context){
	article_id:=com.StrTo(c.Param("id")).MustInt()
	user_id:=com.StrTo(c.Query("user_id")).MustInt()
	text:=c.Query("text")
	created_by:=c.Query("created_by")
	code := e.INVALID_PARAMS
	state:=models.GetState(article_id)
	if state==1{
		code=e.ERROR_COMMENT
		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			"data" : make(map[string]interface{}),
		})
		return
	}


	data := make(map[string]interface {})
	data["article_id"] = article_id
	data["user_id"] = user_id
	data["text"] = text
	data["created_by"] = created_by
	models.AddComment(data)
	code = e.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})
}
func DeleteComment(c * gin.Context){
	id:=com.StrTo(c.Param("id")).MustInt()
	models.DeletedComment(id)
	code:=e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})
}