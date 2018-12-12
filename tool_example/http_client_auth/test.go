package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
)
// https://stackoverflow.com/questions/16673766/basic-http-auth-in-go 参考链接
func main()  {
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

	fmt.Println(string(body))
	var f interface{}
	err = json.Unmarshal(body,&f)
	if err != nil {
		fmt.Println(err)
	}
	m := f.(map[string]interface{})
	//fmt.Println(m)
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
		//fmt.Println(vv)
			for _, u := range vv {
				a := u.(map)
				for _,data := range a{
					println(data)
				}
				//ua := u.(map[string]interface{})
				//fmt.Println(ua["node_status"]) // 取到节点的运行状态
			}
		default:
			fmt.Println(k, "is of a type I don’t know how to handle")
		}
	}
}