package power

import (
	"examination/pkg/e"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Power() gin.HandlerFunc {
	return func(c *gin.Context) {
		le,err:=c.Get("level")
		if err {
			log.Println("get level error")
		}
		level:=le.(int)
		if level==0{
			code:=e.ERROR_POWER

			c.JSON(http.StatusOK, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : make(map[string]string),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}