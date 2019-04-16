package main

import (
	"strings"
	"sync"
	"github.com/learn-go/tool_example/install_app/runCommand"
	"fmt"
	"flag"
)

func execShFile(arg string,shName string,wg *sync.WaitGroup) {
	// 循环本地hosts，获取ip
	//arg := "scp ./main.go 172.23.3.118:/tmp/;echo 123"
	out := runCommand.RunCmdOutPut(arg)
	b := strings.Split(string(out),"\n")
	for _,value := range b {
		if value != "" {
			arg = fmt.Sprintf("scp ./%s %s:/tmp/; ssh %s \"bash /tmp/%s\"",shName,value,value,shName)
			//arg = fmt.Sprintf("scp ./test.sh %s:/tmp/test.sh",value)
			wg.Add(1)
			//go runCommand.RunCmd(arg,wg) // 这里的wg已经是指针类型了，所以不用加&
			go runCommand.RunCmdTest(arg,wg,value)
		}
	}
}
func execSh(arg string,shName string,wg *sync.WaitGroup) {
	// 循环本地hosts，获取ip
	//arg := "scp ./main.go 172.23.3.118:/tmp/;echo 123"
	out := runCommand.RunCmdOutPut(arg)
	b := strings.Split(string(out),"\n")
	for _,value := range b {
		if value != "" {
			arg = fmt.Sprintf("ssh %s \"%s\"",value,shName)
			//arg = fmt.Sprintf("scp ./test.sh %s:/tmp/test.sh",value)
			wg.Add(1)
			//go runCommand.RunCmd(arg,wg) // 这里的wg已经是指针类型了，所以不用加&
			go runCommand.RunCmdTest(arg,wg,value)
		}
	}
}

func main()  {
	var wg sync.WaitGroup //此处也申明为指针变量，声明为值变量，直接传。  声明为值变量用&传地址就行，两种方法
	var service string
	var execShell string
	flag.StringVar(&service,"s","","输出需要安装的服务名如 init sshkey etcd k8s ceph app cassandra等")
	flag.StringVar(&execShell,"a","","输入需要执行的命令")
	flag.Parse() // 解析生效
	switch {
	case service == "init":
		arg := "cat /etc/hosts|awk '/##########/{while(getline)if($0!~/###########/)print;else exit}'|egrep -v \"localhost|xiaoneng\"|egrep -v \"^$|^#\"|awk -F ' ' '{print $1}'"
		shName := service + ".sh"
		execShFile(arg,shName,&wg)
	case service == "ceph":
		arg := "cat /etc/hosts|awk '/##########/{while(getline)if($0!~/###########/)print;else exit}'|grep \"master-01\"|egrep -v \"localhost|xiaoneng\"|egrep -v \"^$|^#\"|awk -F ' ' '{print $1}'"
		shName := service + ".sh"
		execShFile(arg,shName,&wg)
	case service == "etcd":
		fmt.Println("etcd")
		arg := "cat /etc/hosts|awk '/##########/{while(getline)if($0!~/###########/)print;else exit}'|grep \"master\"|egrep -v \"localhost|xiaoneng\"|egrep -v \"^$|^#\"|awk -F ' ' '{print $1}'"
		shName := service + ".sh"
		execShFile(arg,shName,&wg)
	case service == "k8s":
		fmt.Println("k8s")
		arg := "cat /etc/hosts|awk '/##########/{while(getline)if($0!~/###########/)print;else exit}'|grep \"master\"|egrep -v \"localhost|xiaoneng\"|egrep -v \"^$|^#\"|awk -F ' ' '{print $1}'"
		shName := service + ".sh"
		execShFile(arg,shName,&wg)
	default:
		fmt.Println("输入错误，请输入正确的选项如-s init -s etcd  -s k8s -a xxxshell命令")
	}
	if execShell != "" {
		arg := "cat /etc/hosts|awk '/##########/{while(getline)if($0!~/###########/)print;else exit}'|egrep -v \"localhost|xiaoneng\"|egrep -v \"^$|^#\"|awk -F ' ' '{print $1}'"
		execSh(arg,execShell,&wg)
	}
	wg.Wait()
}