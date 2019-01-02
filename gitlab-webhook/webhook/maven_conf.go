package webhook

import (
	"strings"
	)

func Maven (tag_name string,project_name string,namespace string) (mavenConf string) {
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

	//fmt.Println(mavenConf)
	return mavenConf
}