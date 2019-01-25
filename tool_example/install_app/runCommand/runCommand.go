package runCommand

import (
	"os/exec"
	"os"
	"log"
	"sync"
)

//func connect(user string, password string, key string, addr string) (*ssh.Session, error) {
//	var (
//		auth         []ssh.AuthMethod
//		clientConfig *ssh.ClientConfig
//		client       *ssh.Client
//		session      *ssh.Session
//		err          error
//	)
//	// get auth method
//	auth = make([]ssh.AuthMethod, 0)
//	auth = append(auth, ssh.Password(password))
//
//	clientConfig = &ssh.ClientConfig{
//		User:    user,
//		Auth:    auth,
//		Timeout: 30 * time.Second,
//		HostKeyCallback: func(host string, remote net.Addr, key ssh.PublicKey) error {
//			return nil
//		},
//		}
//	client, err = ssh.Dial("tcp", addr, clientConfig)
//	if err != nil {
//		log.Fatal("Failed to dial: ", err)
//	}
//
//	// create session
//	session, err = client.NewSession()
//	if err != nil {
//		log.Fatal("Failed to create session: ", err)
//	}
//	session.Stdout = os.Stdout
//	session.Stderr = os.Stderr
//	session.Stdin = os.Stdin
//	return session, nil
//}
//func main() {
//	session, err := connect("root", "", "172.23.3.118:22")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer session.Close()
//
//	session.Run("ls /; ls /abc")
//}
func RunCmd(arg string,wg *sync.WaitGroup)  {
	// 实时显示输出
	// 执行系统命令
	// 第一个参数是命令名称
	// 后面参数可以有多个，命令参数
	defer wg.Done()
	cmd := exec.Command("/bin/bash","-c",arg)
	// 获取实时输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal("运行错误:",err)
	}
	if err := cmd.Wait();err != nil {
		log.Fatal("等待错误:",err)
	}
}
func RunCmdOutPut(arg string) (outStr string) {
	// 有结果返回
	// 执行系统命令
	// 第一个参数是命令名称
	// 后面参数可以有多个，命令参数
	//defer wg.Done()
	cmd := exec.Command("/bin/bash","-c",arg)
	// 获取实时输出
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	out,err := cmd.CombinedOutput() // 运行命令并返回标准输出和标准错误
	if err != nil {
		log.Println("读取标准输入输出错误：",err)
	}
	outStr = string(out)
	return

}