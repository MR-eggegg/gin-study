package core

import (
	"FuckingVersion1/global"
	"FuckingVersion1/initialize"
	"fmt"
	"net/http"
	"time"
)

func RunWindowsServer() {
	if global.MyConfig.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.MyConfig.System.Addr)
	s := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	global.MyLog.Debug("server run success on ", address)

	global.MyLog.Error(s.ListenAndServe())
}

