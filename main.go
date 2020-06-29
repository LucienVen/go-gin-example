package main

import (
	"context"
	"fmt"
	"gin.example/pkg/setting"
	"gin.example/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	// 开始监听
	//s.ListenAndServe()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Serve ... ")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exitting")
}