package webhook

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"encoding/json"
		"bytes"
	"github.com/pkg/errors"
	"log"
	"gopkg.in/yaml.v2"
)

type toJenkins struct {
	projectName string
	tagName string
	url string
	user string
	password string
	mavenConf string
	jenkinsUser string
	jenkinsPassword string
	r *http.Request
}

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
func (p *toJenkins) checkJobExist() (exitstFlag int) {
	client := &http.Client{}

	fmt.Println(p.url)
	// 生成http request
	req, err := http.NewRequest("GET",p.url,nil)
	// 增加账号密码认证
	req.SetBasicAuth(p.user,p.password)
	if err != nil {
		// handle error
		fmt.Println(errors.Wrap(err,"NewRequest error"))
	}
	// 提交请求
	resp,err := client.Do(req)
	if err != nil {
		resp.Body.Close()
		fmt.Println(errors.Wrap(err,"submit a request error"))
	}
	defer resp.Body.Close()
	body,err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println(errors.Wrap(err,"READ BODY ERR"))
	}
	//fmt.Println(string(body)) // 如果存在会有个 aleady exsiting 的提示，可以取这个用来判断
	return strings.Index(string(body),"exists") //  如果不存在exists则返回-1

}

// post请求
func (p *toJenkins)postJenkins() {
	req, err := http.NewRequest("POST",p.url,bytes.NewBuffer([]byte(p.mavenConf)))  // post二进制数据，要转为byte类型，然后用bytes,NewBuffer发送
	if err != nil {
		fmt.Println(errors.Wrap(err,"NewRequest error"))
	}
	client := &http.Client{}
	// 增加账号密码认证
	req.SetBasicAuth(p.user,p.password)
	req.Header.Set("Content-Type","text/xml")
	resp,err := client.Do(req)
	if err != nil {
		//panic(err) // 出错误终端程序
		fmt.Println(errors.Wrap(err,"submit a request error"))
	}
	defer resp.Body.Close()
}
func (p.)CreateJob(w http.ResponseWriter, r *http.Request) {
	var c conf
	c.getConf() // 获取配置信息

	var config Git_conf //webhook
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
		project_name := config.Project.Name
		tagSlice := strings.Split(config.Ref,"/")

		tag_name := tagSlice[2]
		url :=  c.JenkinsUrl + "/checkJobName?value=" + project_name +tag_name // 这个http:// 不要少加了。检查job存在的url
		mavenConf := Maven(tag_name,project_name,namespace) //webhook
		// 这个位置还没加判断job是否已经存在，后续添加
		if checkJobExist(project_name,tag_name,url,c.JenkinsUser,c.JenkinsPassword) == -1 {
			// 发送post请求携带xml文件，创建job
			url = c.JenkinsUrl + "/createItem?name=" + project_name + tag_name
			postJenkins(url,mavenConf,c.JenkinsUser,c.JenkinsPassword)
			// 创建完job，调用jenkins构建api
			url = c.JenkinsUrl +  "/job/" + project_name + tag_name + "/build"
			postJenkins(url,mavenConf,c.JenkinsUser,c.JenkinsPassword)
		}else {
			url = c.JenkinsUrl + "/job"+ project_name + tag_name + "/config.xml"
			postJenkins(url,mavenConf,c.JenkinsUser,c.JenkinsPassword)
			// 更新完job，调用jenkins构建api
			url = c.JenkinsUrl + "/job/" + project_name + tag_name + "/build"
			postJenkins(url,mavenConf,c.JenkinsUser,c.JenkinsPassword)
		}
	}
}
