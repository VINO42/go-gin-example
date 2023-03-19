package main

import (
	"fmt"
	"github.com/vino42/go-gin-example/middleware/zlog"
	"github.com/vino42/go-gin-example/pkg/setting"
	"github.com/vino42/go-gin-example/routers"
	"net/http"
)

func main() {
	//router := gin.Default()

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	zlog.Info("starting server....   at port %d", setting.HTTPPort)
	s.ListenAndServe()
}
