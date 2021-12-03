package main

import (
	"fmt"
	"fzuhelper_launch_screen/models"
	"fzuhelper_launch_screen/mongo"
	"fzuhelper_launch_screen/pkg/setting"
	"fzuhelper_launch_screen/pkg/util"
	"fzuhelper_launch_screen/routers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	setting.Setup()
	models.Setup()
	util.Setup()
	mongo.Setup()
}


func main() {
	logrus.SetLevel(logrus.DebugLevel)
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	logrus.Infoln("start http server listening ", endPoint)

	_ = server.ListenAndServe()
}
