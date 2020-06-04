package main

import (
	"fmt"
	"gin.example/pkg/setting"
	"gin.example/routers"
	"net/http"
)

func main()  {
	router := routers.InitRouter()



	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:           router,
		ReadHeaderTimeout:  setting.ReadTimeOut,
		WriteTimeout:		setting.WriteTimeOut,
		MaxHeaderBytes: 	1 << 20,
	}

	// 开始监听
	s.ListenAndServe()
}