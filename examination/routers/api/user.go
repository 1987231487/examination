package api

import (
	"examination/models"
	"examination/pkg/e"
	"examination/pkg/util"
	"examination/redis"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
	"strings"
)

type user struct {
	Username string `valid:"Required; MaxSize(20)"`
	Password string `valid:"Required; MaxSize(20)"`
	Level int
}
type cache struct{            //单核验证码
	Email string
	Code string
}
var a cache
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	valid := validation.Validation{}
	a := user{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckUser(username, password)
		if isExist {

			level:=models.GetUserPower(username)
			token, err := util.GenerateToken(username,level)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {data["token"] = token

			code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH //密码错误
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}
func Getcode(c *gin.Context) {
	email:=c.Query("email")
	mycode,err:=util.SentEmail(email)

	if err==nil{
	redis.Set(email,mycode)

	c.JSON(200,gin.H{"msg":"验证码已发送"})
}else{
		c.JSON(400,gin.H{"msg":"验证码发送失败"})

	}
}
func Register(c *gin.Context) {
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	//mycode:=a.Code
	//email:=a.Email
	email:=c.Query("email")
	usercode:=c.Query("code")
	name := c.Query("name")
	password := c.Query("password")
	level:=com.StrTo(c.Query("level")).MustInt()
	mycode,err:=redis.Get(email)
	if err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"code" : 301,
			"msg" : "邮箱未注册验证码",
			"data" : data,
		})
		return
	}
	if ok:=strings.EqualFold(mycode,usercode);!ok{
		code=10002
		c.JSON(http.StatusOK, gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			"data" : data,
		})
		return
	}
	valid := validation.Validation{}
	a := user{Username: name, Password: password}
	ok, _ := valid.Valid(&a)


	if ok {
		isExist := models.CheckName(name)
		if isExist {
			code=e.ERROR_REGISTER

		} else {
			models.CreateUser(name,password,email,level)
			code=e.SUCCESS
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}


	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

