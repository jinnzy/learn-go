package main

import (
	"fmt"
	"net"
)

func main()  {
	en0, err := net.InterfaceByName("enp0s3")
	if err != nil {
		fmt.Println(err)
	}
	a,err := en0.Addrs()
	fmt.Println(a[0].String())
	addrs, err := net.InterfaceAddrs()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(Ips())
	for _, value := range addrs{
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback(){
			if ipnet.IP.To4() != nil{
				fmt.Println(ipnet.IP.String())
				return
			}
		}
	}

}
func Ips() (map[string]string, error) {

	ips := make(map[string]string)
	//返回 interface 结构体对象的列表，包含了全部网卡信息
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	//遍历全部网卡
	for _, i := range interfaces {

		// Addrs() 方法返回一个网卡上全部的IP列表
		address, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		//遍历一个网卡上全部的IP列表，组合为一个字符串，放入对应网卡名称的map中
		for _, v := range address {
			if ipnet, ok := v.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ips[i.Name] = v.String()
				}
			}
		}
	}
	return ips, nil
}

//func getEncryptedPasswd(passwd string) chromedp.Tasks {
//	return chromedp.Tasks{
//
//		chromedp.Navigate(""),
//		//chromedp.WaitVisible(`#username`, chromedp.ByID), // ByID 意思是查找对应id标签名为#input1的
//		//chromedp.WaitVisible(`#password`, chromedp.ByID), // 等待#textarea1 标签渲染成功
//		//chromedp.WaitVisible(`btn-submit`,chromedp.NodeVisible),
//		////chromedp.Sleep(1 * gettime.Second),
//		//// 输入falcon账号
//		//chromedp.SendKeys(`#username`, "15011424102", chromedp.ByID), // 向指定的html元素内输入内容
//		////chromedp.Value(`#input1`, val1, chromedp.ByID), // 获取#input1的值
//		////chromedp.SetValue(`#input2`, "test3", chromedp.ByID), // 改变对应元素中地value值，例子中会把next input框改为test3 input框
//		//// 输入密码
//		//chromedp.SendKeys(`#password`, "", chromedp.ByID),
//		//// 点击登录
//
//		//chromedp.Click("btn-submit",chromedp.NodeVisible),
//		//chromedp.WaitVisible(`input[name=ServerSearch\[hostname\]]`),
//		chromedp.SendKeys(`input[name=ServerSearch\[hostname\]]`, hostName+kb.Enter), // 找到inout标签中 class是input-sm的
//		chromedp.Sleep(100 * time.Microsecond),
//		chromedp.OuterHTML(`html`, &resCmdb, chromedp.ByQueryAll), //获取改 tbody标签的html
//		//chromedp.Text(`tbody`,&res,chromedp.ByQueryAll),
//		//chromedp.Sleep(60000 * time.Second),
//	}
//}
