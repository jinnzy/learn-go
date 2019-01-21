package main

import (
	"net/http"
	"fmt"
	"github.com/learn-go/blog/pkg/setting"
	"github.com/learn-go/blog/routers"
	)

func main() {
	router := routers.InitRouter()
	//a := models.GetArticle(1	)
	//fmt.Println(a)

	//maps := make(map[string]interface{})
	//maps["state"] = 1
	//a := models.GetArticles(0,2,maps)
	//fmt.Println(a)

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