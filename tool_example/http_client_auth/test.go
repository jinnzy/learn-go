package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)
// https://stackoverflow.com/questions/16673766/basic-http-auth-in-go 参考链接
func main()  {
	// 生成http client
	client := &http.Client{}
	// 生成http request
	req, err := http.NewRequest("GET","http://172.23.4.154:32277/api/v2/management/nodes",nil)
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

	fmt.Println(string(body))
}
