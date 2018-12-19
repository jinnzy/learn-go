package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/bitly/go-simplejson"
	"encoding/json"
	"reflect"
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
	//{"code":0,"result":[{"emqtt@10.1.2.223":{"clients/count":8,"clients/max":467,"retained/count":9,"retained/max":9,"routes/count":93,"routes/max":2226,"sessions/count":8,"sessions/max":468,"subscribers/count":8,"subscribers/max":135,"subscriptions/count":8,"subscriptions/max":135,"topics/count":93,"topics/max":2226},"emqtt@10.1.4.75":{"clients/count":12,"clients/max":459,"retained/count":9,"retained/max":9,"routes/count":93,"routes/max":2226,"sessions/count":13,"sessions/max":459,"subscribers/count":37,"subscribers/max":850,"subscriptions/count":37,"subscriptions/max":850,"topics/count":93,"topics/max":2226},"emqtt@10.1.1.205":{"clients/count":10,"clients/max":474,"retained/count":9,"retained/max":9,"routes/count":93,"routes/max":2226,"sessions/count":11,"sessions/max":473,"subscribers/count":21,"subscribers/max":623,"subscriptions/count":21,"subscriptions/max":623,"topics/count":93,"topics/max":2226},"emqtt@10.1.3.79":{"clients/count":14,"clients/max":468,"retained/count":9,"retained/max":9,"routes/count":93,"routes/max":2225,"sessions/count":14,"sessions/max":469,"subscribers/count":27,"subscribers/max":637,"subscriptions/count":27,"subscriptions/max":637,"topics/count":93,"topics/max":2225}}]}
	// 上面的是json例子
	//var f test
	body := HttpGetMqtt()
	// 使用simple.json 反序列化
	res,err := simplejson.NewJson(body)
	if err != nil{
		fmt.Println(err)
	}
	//fmt.Println(res.Get("result").Array())
	// 获取result下的数组
	a,err := res.Get("result").Array()
	// 遍历result下的数组，结果为0对应map[emqtt....]
	for _,v := range a {
		fmt.Printf("v: %s",v)
	}



	////fmt.Println(string(body))
	//err := json.Unmarshal(body,&f)
	//if err != nil {
	//	fmt.Println("json err",err)
	//}
	//switch vv := f.Result.(type) {
	//case interface{}:
	//	fmt.Println(vv)
	//	bb :=
	//default:
	//	fmt.Println("is of a type I don’t know how to handle")
	//}
	}

	fmt.Println(f.Result)
	//switch vv := f.Result.(type) {
	//case interface{}:
	//	//for k,v := range vv.(map[string]interface{}){
	//	//	fmt.Println(k)
	//	//	fmt.Println(v)
	//	//}
	fmt.Println(reflect.TypeOf(f.Result))
	//t1 := make([]interface{},0)
	//var t1 intertest in
	//t1 := f.Result
	//fmt.Println(reflect.TypeOf(t1))
	//fmt.Println(f.Result[0])
	//	//fmt.Println(vv[0])
	//	var interfaceSlice []interface{} = make([]interface{}, 1)
	//
	//default:
	//	fmt.Println("is of a type I don’t know how to handle")
	}

