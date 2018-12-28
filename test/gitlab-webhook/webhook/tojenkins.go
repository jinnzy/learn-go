package webhook

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"encoding/json"
		"bytes"
)

func checkJobExist(project_name string,tag_name string,url string) {
	client := &http.Client{}

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
	fmt.Println(string(body)) // 如果存在会有个 aleady exsiting 的提示，可以取这个用来判断

}

func post_jenkins(url string,mavenConf string) {
	req, err := http.NewRequest("POST",url,bytes.NewBuffer([]byte(mavenConf)))  // post二进制数据，要转为byte类型，然后用bytes,NewBuffer发送
	if err != nil {
		fmt.Println("request err:",err)
	}
	client := &http.Client{}
	// 增加账号密码认证
	req.SetBasicAuth("root","1234qwer")
	req.Header.Set("Content-Type","text/xml")
	resp,err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
func CreateJob(w http.ResponseWriter, r *http.Request) {

	var config Git_conf //webhook
	result, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(result))
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
		url := "http://192.168.30.75:8999/checkJobName?value=" + project_name +tag_name // 这个http:// 不要少加了。检查job存在的url
		mavenConf := Maven(tag_name,project_name,namespace) //webhook
		// 这个位置还没加判断job是否已经存在，后续添加
		checkJobExist(project_name,tag_name,url)

		// 发送post请求携带xml文件，创建job
		url = "http://192.168.30.75:8999/createItem?name=" + project_name + tag_name
		post_jenkins(url,mavenConf)
		// 创建完job，调用jenkins构建api
		url = "http://192.168.30.75:8999/job/" + project_name + tag_name + "/build"
		post_jenkins(url,mavenConf)


	}
}



// 逻辑处理函数 CreateJob换这个，  然后重写createjob， 用逻辑处理函数调用createjob ， updatejob  ，判断job是否存在等
//func ApplyJob() {
//
//}