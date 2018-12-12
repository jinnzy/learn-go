package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"

	)

func HttpGetMqtt() []byte {
	// 生成http client
	client := &http.Client{}
	// 生成http request
	req, err := http.NewRequest("GET","http://172.23.4.154:32453/api/v2/monitoring/stats",nil)
	// 增加账号密码认证
	req.SetBasicAuth("admin","public")
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	// 提交请求
	resp,err := client.Do(req)
	defer resp.Body.Close()
	// 处理返回结果
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

type test struct {
	Code int `json:"code"`
	Result interface{} `json:"result"`
}

// https://stackoverflow.com/questions/16673766/basic-http-auth-in-go 参考链接
func main()  {


	var f test
	body := HttpGetMqtt()
	//fmt.Println(string(body))
	err := json.Unmarshal(body,&f)
	if err != nil {
		fmt.Println("json err",err)
	}
	switch vv := f.Result.(type) {
	case interface{}:
		fmt.Println(vv)
		bb :=
	default:
		fmt.Println("is of a type I don’t know how to handle")
	}
	}