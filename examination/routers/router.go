package routers

import (
	"examination/middleware/jwt"
	"examination/middleware/power"
	"examination/pkg/setting"
	"examination/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/getcode",api.Getcode)//获取验证码

	r.GET("/register",api.Register)//注册

	r.GET("/login", api.Login) //登录

	apiv:=r.Group("")
	apiv.Use(jwt.JWT()) //token验证中间件
	{
		apiv0 := apiv.Group("/level/0")
		{
			//获取文章列表
			apiv0.GET("/articles", api.GetArticles)
			//获取指定文章
			apiv0.GET("/articles/:id", api.GetArticle)
			//获取文章所有评论
			apiv0.GET("/comment/:id", api.GetComment)
			//给文章添加评论
			apiv0.POST("/comment/:id", api.AddComment)
		}

		/*管理员可用*/
		apiv1 := apiv.Group("/level/1")
		apiv1.Use(power.Power()) //权限中间件
		{
			//新建文章
			apiv1.POST("/articles", api.AddArticle)
			//更新指定文章
			apiv1.PUT("/articles/:id", api.EditArticle)
			//删除指定文章
			apiv1.DELETE("/articles/:id", api.DeleteArticle)
			//设置文章是否可被评论
			apiv1.POST("/articles/:id", api.ChangeState)
			//删除评论
			apiv1.DELETE("/comment/:id", api.DeleteComment)
		}
	}

	return r
}