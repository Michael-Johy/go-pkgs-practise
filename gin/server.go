package main

import (
	"fmt"
	"github.com/Michael-Johy/go-pkgs-practise/gin/pkg/setting"
	"github.com/Michael-Johy/go-pkgs-practise/gin/routers"
	"log"
	"net/http"
	"time"
)

func main() {
	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPORT),
		Handler:        router,
		ReadTimeout:    time.Duration(setting.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(setting.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("server start fail ...")
	}

}
