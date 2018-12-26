package webhook

import (
	"net/http"
	"github.com/learn-go/test/gitlab-webhook/webhook"
	"io/ioutil"
	"fmt"
	"strings"
	"encoding/json"
)

//func CreateJob(mavenConf string)  {
//func CheckJobExists(tag_name string,project_name string)  {
//	fmt.Println("CheckJobExistsTag_name",tag_name)
//	fmt.Println("CheckJobExistsProject_name",project_name)
//	//req,err := http.Get("http://192.168.30.75:8999/checkJobName?value=test")
//	//
//	//if err != nil {
//	//	fmt.Println("GET ERR:",err)
//	//}
//	//body,err := ioutil.ReadAll(req.Body)
//	//fmt.Println(string(body))
//	// 生成http client
//	client := &http.Client{}
//
//	url := fmt.Sprintf("http://192.168.30.75:8999/checkJobName?value=%d%d",project_name,tag_name) // 这个http:// 不要少加了
//	fmt.Println(url)
//	// 生成http request
//	req, err := http.NewRequest("GET",url,nil)
//	// 增加账号密码认证
//	//req.SetBasicAuth("root","1234qwer")
//	if err != nil {
//		// handle error
//		fmt.Println("GET ERROR:%d",err)
//	}
//	// 提交请求
//	resp,err := client.Do(req)
//	if err != nil {
//		resp.Body.Close()
//		fmt.Println("BODY CLOSE ERR:",err)
//	}
//	defer resp.Body.Close()
//	body,err1 := ioutil.ReadAll(resp.Body)
//	if err1 != nil {
//		fmt.Println("READ BODY ERR:%d",err1)
//	}
//	fmt.Println(string(body))
//}
func CreateJob(w http.ResponseWriter, r *http.Request) {
	var config webhook.Git_conf
	result, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	} else {
		err := json.Unmarshal([]byte(result),&config)
		if err != nil{
			fmt.Println(err)
		}
		namespace := config.Project.Namespace
		project_name := config.Project.Name
		tagSlice := strings.Split(config.Ref,"/")
		tag_name := tagSlice[2]
		temp := webhook.Maven(tag_name,project_name,namespace)
		fmt.Sprintf("temp:",temp)
		//fmt.Println("ref: ", config.Ref)
		//fmt.Println("projectname: ", config.Project.Name)
		//go webhook.CheckJobExists(tag_name,project_name)
		////////////////////////////////////////
		client := &http.Client{}

		url := "http://192.168.30.75:8999/checkJobName?value=" + project_name +tag_name // 这个http:// 不要少加了
		fmt.Println(url)
		// 生成http request
		req, err := http.NewRequest("GET",url,nil)
		// 增加账号密码认证
		req.SetBasicAuth("root","1234qwer")
		if err != nil {
			// handle error
			fmt.Println("GET ERROR:%d",err)
		}
		// 提交请求
		resp,err := client.Do(req)
		if err != nil {
			resp.Body.Close()
			fmt.Println("BODY CLOSE ERR:",err)
		}
		defer resp.Body.Close()
		body,err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			fmt.Println("READ BODY ERR:%d",err1)
		}
		fmt.Println(string(body))
	}
}

//func ApplyJob() {
//
//}