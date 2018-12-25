package main

import (
		"net/http"
		"fmt"
	"io/ioutil"
		"encoding/json"
	"strings"
)

const (
	path = "/webhooks"
)

type git_conf struct {
	Object_kind string `json:"object_kind"`
	Event_name string `json:"event_name"`
	Before string `json:"before"`
	After string `json:"after"`
	Ref string `json:"ref"`
	Checkout_sha string `json:"checkout_sha"`
	Message string `json:"message"`
	User_id int `json:"user_id"`
	User_name string `json:"user_name"`
	User_username string `json:"user_username"`
	User_email string `json:"user_email"`
	User_avatar string `json:"user_avatar"`
	Project_id int `json:"project_id"`
	Project projectStruct `json:"project"`
	Commits interface{} `json:"commits"`
	Total_commits_count int `json:"total_commits_count"`
	Repository interface{} `json:"repository"`
}
type projectStruct struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Web_url string `json:"web_url"`
	Avatar_url interface{} `json:avatar_url"`
	Git_ssh_url string `json:"git_ssh_url"`
	Git_http_url string `json:"git_http_url"`
	Namespace string `json:"namespace"`
	Visibility_level int `json:"visibility_level"`
	Path_with_namespace string `json:"path_with_namespace"`
	Default_branch string `json:"default_branch"`
	Ci_config_path interface{} `json:"ci_config_path"`
	Homepage string `json:"homepage"`
	Url string `json:"url"`
	Ssh_url string `json:"ssh_url"`
	Http_url string `json:"http_url"`
}

func maven (tag_name string,project_name string,namespace string) (mavenConf string) {
	mavenConf = `
<maven2-moduleset plugin="maven-plugin@3.2">
<actions/>
<description/>
<keepDependencies>false</keepDependencies>
<properties>
<com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty plugin="gitlab-plugin@1.5.11">
<gitLabConnection/>
</com.dabsquared.gitlabjenkins.connection.GitLabConnectionProperty>
</properties>
<scm class="hudson.plugins.git.GitSCM" plugin="git@3.9.1">
<configVersion>2</configVersion>
<userRemoteConfigs>
<hudson.plugins.git.UserRemoteConfig>
<url>http://192.168.30.75:8998/root/java-test.git</url>
</hudson.plugins.git.UserRemoteConfig>
</userRemoteConfigs>
<branches>
<hudson.plugins.git.BranchSpec>
<name>*/master</name>
</hudson.plugins.git.BranchSpec>
</branches>
<doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
<submoduleCfg class="list"/>
<extensions/>
</scm>
<canRoam>true</canRoam>
<disabled>false</disabled>
<blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
<blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
<triggers/>
<concurrentBuild>false</concurrentBuild>
<aggregatorStyleBuild>true</aggregatorStyleBuild>
<incrementalBuild>false</incrementalBuild>
<ignoreUpstremChanges>true</ignoreUpstremChanges>
<ignoreUnsuccessfulUpstreams>false</ignoreUnsuccessfulUpstreams>
<archivingDisabled>false</archivingDisabled>
<siteArchivingDisabled>false</siteArchivingDisabled>
<fingerprintingDisabled>false</fingerprintingDisabled>
<resolveDependencies>false</resolveDependencies>
<processPlugins>false</processPlugins>
<mavenValidationLevel>-1</mavenValidationLevel>
<runHeadless>false</runHeadless>
<disableTriggerDownstreamProjects>false</disableTriggerDownstreamProjects>
<blockTriggerWhenBuilding>true</blockTriggerWhenBuilding>
<settings class="jenkins.mvn.DefaultSettingsProvider"/>
<globalSettings class="jenkins.mvn.DefaultGlobalSettingsProvider"/>
<reporters/>
<publishers/>
<buildWrappers/>
<prebuilders/>
<postbuilders>
<hudson.tasks.Shell>
<command>
docker build --pull -t reg.test.cn/{namespace}/{project_name}:{tag_name} .
</command>
</hudson.tasks.Shell>
</postbuilders>
<runPostStepsIfResult>
<name>UNSTABLE</name>
<ordinal>1</ordinal>
<color>YELLOW</color>
<completeBuild>true</completeBuild>
</runPostStepsIfResult>
</maven2-moduleset>
`
	mavenConf = strings.Replace(mavenConf,"{tag_name}",tag_name,-1)
	mavenConf = strings.Replace(mavenConf,"{namespace}",namespace,-1)
	mavenConf = strings.Replace(mavenConf,"{project_name}",project_name,-1)

	fmt.Println(mavenConf)
	return mavenConf
}



func main() {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		var config git_conf
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
			maven(tag_name,project_name,namespace)
			fmt.Println("ref: ", config.Ref)
			fmt.Println("projectname: ", config.Project.Name)
		}
	})
	http.ListenAndServe(":9000", nil)
}