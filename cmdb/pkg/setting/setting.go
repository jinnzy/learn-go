package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
	"fmt"
)

var (
	Cfg *ini.File
	RunMode string

	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JwtSecret string
)
func init() {
	var err error
	Cfg,err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}
func LoadBase(){
	// 由 Must 开头的方法名允许接收一个相同类型的参数来作为默认值，
	// 当键不存在或者转换失败时，则会直接返回该默认值。
	// 但是，MustString 方法必须传递一个默认值。
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
func LoadServer() {
	// 选择app.ini的server分区
	sec,err := Cfg.GetSection("server")
	if err != nil {
		// 打印输出内容，退出应用程序，defer函数不执行
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	fmt.Println("http port:",HttpPort)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec,err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal()
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}


















