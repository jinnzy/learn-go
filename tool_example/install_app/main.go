package main

import (
		"strings"
	"sync"
	"github.com/learn-go/tool_example/install_app/runCommand"
	"fmt"
)

func initSh(wg *sync.WaitGroup) {
	// 循环本地hosts，获取ip
	arg := "scp ./main.go 172.23.3.118:/tmp/;echo 123"
	arg = "cat /etc/hosts|egrep -v \"localhost|xiaoneng\"|egrep -v \"^$\"|awk -F ' ' '{print $1}'"
	out := runCommand.RunCmdOutPut(arg)

	b := strings.Split(string(out),"\n")
	for _,value := range b {
		if value != "" {
			arg = fmt.Sprintf("scp ./init.sh %s:/root/init.sh; ssh %s \"bash /root/init.sh\"",value,value)
			//arg = fmt.Sprintf("scp ./test.sh %s:/tmp/test.sh",value)
			wg.Add(1)
			go runCommand.RunCmd(arg,wg) // 这里的wg已经是指针类型了，所以不用加&
		}
	}
}


func main()  {
	var wg sync.WaitGroup //此处也申明为指针变量，声明为值变量，直接传。  声明为值变量用&传地址就行，两种方法
	initSh(&wg)
	wg.Wait()
}