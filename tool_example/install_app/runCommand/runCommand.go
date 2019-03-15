package runCommand

import (
	"os/exec"
	"os"
	"log"
	"sync"
	"bufio"
	"io"
	"fmt"
	"strings"
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
	// func (c *Cmd) Run() error　　　　　　　　　　//开始指定命令并且等待他执行结束，如果命令能够成功执行完毕，则返回nil，否则的话边会产生错误
	//func (c *Cmd) Start() error　　　　　　　　　　//使某个命令开始执行，但是并不等到他执行结束，这点和Run命令有区别．然后使用Wait方法等待命令执行完毕并且释放响应的资源
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
func RunCmdTest(arg string,wg *sync.WaitGroup,ip string)  {
	// 实时显示输出的基础上增加了输出信息
	// 执行系统命令
	// 第一个参数是命令名称
	// 后面参数可以有多个，命令参数
	defer wg.Done()
	cmd := exec.Command("/bin/bash","-c",arg)
	fmt.Println(cmd.Args)
	// 获取实时输出
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr,_ := cmd.StderrPipe()
	// 运行命令
	// func (c *Cmd) Run() error　　　　　　　　　　//开始指定命令并且等待他执行结束，如果命令能够成功执行完毕，则返回nil，否则的话边会产生错误
	//func (c *Cmd) Start() error　　　　　　　　　　//使某个命令开始执行，但是并不等到他执行结束，这点和Run命令有区别．然后使用Wait方法等待命令执行完毕并且释放响应的资源
	if err := cmd.Start(); err != nil {
		log.Fatal("运行错误:",err)
	}
	//stdin.Write([]byte("go text for grep\n"))
	//stdin.Write([]byte("go test text for grep\n"))
	//stdin.Close()
	// 缓存输出 https://studygolang.com/articles/4004
	readerOut := bufio.NewReader(stdout)
	// 处理标准输出
	for {
		// ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串。如果ReadString方法在读取到delim之前遇到了错误，它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。当且仅当ReadString方法返回的切片不以delim结尾时，会返回一个非nil的错误。
		line, err := readerOut.ReadString('\n')
		// 上面的是遇到\n则为一行，但是并不作处理，输出的时候会多空出一行，在下面处理去掉\n换行符
		line = strings.Trim(line,"\n")
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println("["+ip+"]"+":"+line)
	}
	// 处理标准错误
	readerErr := bufio.NewReader(stderr)
	for {

		// ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串。如果ReadString方法在读取到delim之前遇到了错误，它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。当且仅当ReadString方法返回的切片不以delim结尾时，会返回一个非nil的错误。
		line, err := readerErr.ReadString('\n')
		line = strings.Trim(line,"\n")
		// break 要在下面，因为读到最后一行的没有\n会出错，要在break之前打印出来
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println("["+ip+"]"+":"+line)
	}
	if err := cmd.Wait();err != nil {
		fmt.Printf("[%s]执行命令出错: %s\n",ip,err)
	}
}