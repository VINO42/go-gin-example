package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/vino42/go-gin-example/middleware/zlog"
	"github.com/vino42/go-gin-example/pkg/setting"
	"github.com/vino42/go-gin-example/routers"
	"net/http"
	"syscall"
)

func main() {
	//router := gin.Default()
	restartGracefull()

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	zlog.Info("starting server....   at port %d", setting.HTTPPort)
	err := s.ListenAndServe()
	if err != nil {
		zlog.Info("Server err: %v", err)
	}
}

func restartGracefull() {
	// 热更新是采取创建子进程后，将原进程退出的方式
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		zlog.Info("Actual pid is %d", syscall.Getpid())
	}
}
