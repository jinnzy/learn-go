package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main()  {

	client := &http.Client{}
	url := fmt.Sprintf("http://192.168.30.75:8999/checkJobName?value=mavent12321est")
	fmt.Println(url)
	// 生成http request
	req, err := http.NewRequest("GET", url, nil)
	// 增加账号密码认证
	req.SetBasicAuth("root","1234qwer")
	if err != nil {
		// handle error
		fmt.Println("GET ERROR:%d", err)
	}
	// 提交请求
	resp, err := client.Do(req)
	if err != nil {
		resp.Body.Close()
		fmt.Println("BODY CLOSE ERR:", err)
	}
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println("READ BODY ERR:%d", err1)
	}
	fmt.Println(string(body))
}