package route

import (
	"gin-demo/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(controller.ReturnSuccessData(map[string]interface{}{
			"hello": "world",
		}))
	})

	//registerUserRoute(r)
}

//func registerUserRoute(r *gin.Engine) {
//	r.GET("/user/login")
//}
