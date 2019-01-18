package main

import (
	"net/http"
	"fmt"
	"github.com/learn-go/blog/pkg/setting"
	"github.com/learn-go/blog/routers"
)

func main() {
	router := routers.InitRouter()
	fmt.Println(setting.HttpPort,setting.ReadTimeout,setting.WriteTimeout)
	s := &http.Server{
		Addr: fmt.Sprintf(":%d",setting.HttpPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}