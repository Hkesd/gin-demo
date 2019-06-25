package middleware

import (
	"encoding/json"
	"fmt"
	"gin-demo/common/applog"
	"gin-demo/common/error"
	"gin-demo/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

func RegisterBasicMiddleWare(r *gin.Engine) {
	if gin.IsDebugging() {
		r.Use(applog.GinLog(), DebugRecovery(recoveryHandler))
	} else {
		r.Use(applog.GinLog(), applog.GinRecovery(recoveryHandler))
	}
}

func DebugRecovery(f ...func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				panicInfo, _ := json.Marshal(map[string]interface{}{
					"error": err,
				})
				fmt.Println("[Debug Panic Recovery] " + string(panicInfo))
				debug.PrintStack()

				if len(f) == 1 {
					f[0](c)
				} else {
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		c.Next()
	}
}

func recoveryHandler(c *gin.Context) {
	c.JSON(controller.ReturnData(api_err.ErrorInternal, nil))
}
