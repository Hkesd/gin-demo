package main

import (
	"bytes"
	"context"
	"fmt"
	"gin-demo/common/applog"
	"gin-demo/config"
	"gin-demo/middleware"
	"gin-demo/model"
	"gin-demo/route"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 配置gin是否是Debug模式
	if config.C.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 连接mysql
	err := model.InitDB(getConnectMysqlStr())
	if err != nil {
		return
	}
	defer func() {
		err := model.DB.Close()
		if err != nil {
			logrus.Error("Failed to close mysql")
		}
	}()

	server := gin.New()
	// 注册基本中间件
	middleware.RegisterBasicMiddleWare(server)
	// 注册路由
	route.RegisterRoute(server)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", config.C.Port),
		Handler: server,
	}

	// 异步启动server
	go func() {
		logrus.Infof("starting server at :%d", config.C.Port)
		// 使用http启动server
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalln("start server failed.", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Infoln("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Errorln("server shutdown failed:", err)
	}
	logrus.Info("server exit")
}

func init() {
	// 初始化日志设置
	applog.ConfigLocalFilesystemLogger(config.C.Log.Path, "app")
}

// 生成mysql连接字符串
func getConnectMysqlStr() string {
	var buffer bytes.Buffer
	buffer.WriteString(config.C.MysqlConf.User)
	buffer.WriteString(":")
	buffer.WriteString(config.C.MysqlConf.Pass)
	buffer.WriteString("@tcp")
	buffer.WriteString("(")
	buffer.WriteString(config.C.MysqlConf.Host)
	buffer.WriteString(":")
	buffer.WriteString(config.C.MysqlConf.Port)
	buffer.WriteString(")/")
	buffer.WriteString(config.C.MysqlConf.DbName)
	buffer.WriteString("?charset=utf8&parseTime=True&loc=Local")

	return buffer.String()
}
