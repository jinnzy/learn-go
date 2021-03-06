package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/learn-go/gitlab-webhook/pkg/logging"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)
// 解析配置文件
type conf struct {
	JenkinsUrl string `yaml:"jenkinsUrl"`
	JenkinsUser string `yaml:"jenkinsUser"`
	JenkinsPassword string `yaml:"jenkinsPassword"`
}

func (c *conf) getConf() *conf {
	yamlFile,err := ioutil.ReadFile("gitlab-webhook/conf.yaml")
	if err!= nil{
		log.Printf("yamlFile.Get err #%v", err)
	}
	err = yaml.Unmarshal(yamlFile,c)
	if err != nil {
		log.Fatal("Unmarshal:%v",err)
	}
	return c
}


// 检查job是否存在
func checkJobExist(project_name string,tag_name string,url string,user string,password string) (exitstFlag int) {
	client := &http.Client{}

	fmt.Println(url)
	// 生成http request
	req, err := http.NewRequest("GET",url,nil)
	// 增加账号密码认证
	req.SetBasicAuth(user,password)
	if err != nil {
		// handle error
		logging.Error("NewRequest error",zap.Error(err))
		logging.Error("NewRequest",zap.Error(err))
	}
	// 提交请求
	resp,err := client.Do(req)
	if err != nil {
		resp.Body.Close()
		logging.Error("submit a request error",zap.Error(err))
	}
	defer resp.Body.Close()
	body,err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		logging.Error("READ BODY ERR",zap.Error(err))
	}
	//fmt.Println(string(body)) // 如果存在会有个 aleady exsiting 的提示，可以取这个用来判断
	return strings.Index(string(body),"exists") //  如果不存在exists则返回-1

}

// post请求
func postJenkins(url string,mavenConf string,user string,password string) {
	req, err := http.NewRequest("POST",url,bytes.NewBuffer([]byte(mavenConf)))  // post二进制数据，要转为byte类型，然后用bytes,NewBuffer发送
	if err != nil {
		logging.Error("NewRequest error",zap.Error(err))
	}
	client := &http.Client{}
	// 增加账号密码认证
	req.SetBasicAuth(user,password)
	req.Header.Set("Content-Type","text/xml")
	resp,err := client.Do(req)
	if err != nil {
		//panic(err) // 出错误终端程序
		logging.Error("submit a request error",zap.Error(err))
	}
	defer resp.Body.Close()
}
func CreateJob(w http.ResponseWriter, r *http.Request) {
	var c conf
	c.getConf() // 获取配置信息

	var config GitConf//webhook
	result, err := ioutil.ReadAll(r.Body)
	log.Printf("接收gitlab push请求: %d",string(result))
	if err != nil {
		fmt.Println(err)
	} else {
		err := json.Unmarshal([]byte(result),&config)
		if err != nil{
			fmt.Println(err)
		}
		namespace := config.Project.Namespace
		projectName := config.Project.Name
		tagSlice := strings.Split(config.Ref,"/")

		tagName := tagSlice[2]
		url :=  c.JenkinsUrl + "/checkJobName?value=" + projectName + tagName // 这个http:// 不要少加了。检查job存在的url
		mavenConf := Maven(tagName,projectName,namespace) //webhook
		// 这个位置还没加判断job是否已经存在，后续添加
		logging.Info("",zap.String("projectName",projectName),zap.String("tagName",tagName),zap.String("jenkinsUser",c.JenkinsUser),zap.String("jenkinsPassword",c.JenkinsPassword))
		if checkJobExist(projectName,tagName,url,c.JenkinsUser,c.JenkinsPassword) == -1 {
			// 发送post请求携带xml文件，创建job
			url = c.JenkinsUrl + "/createItem?name=" + projectName + tagName
			postJenkins(url,mavenConf,c.JenkinsUser,c.JenkinsPassword)
			// 创建完job，调用jenkins构建api
			url = c.JenkinsUrl +  "/job/" + projectName + tagName + "/build"
			postJenkins(url,mavenConf,c.JenkinsUser,c.JenkinsPassword)
		}else {
			url = c.JenkinsUrl + "/job"+ projectName + tagName + "/config.xml"
			postJenkins(url,mavenConf,c.JenkinsUser,c.JenkinsPassword)
			// 更新完job，调用jenkins构建api
			url = c.JenkinsUrl + "/job/" + projectName + tagName + "/build"
			postJenkins(url,mavenConf,c.JenkinsUser,c.JenkinsPassword)
		}
	}
}
