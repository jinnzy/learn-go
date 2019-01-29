package setting

import (
	"time"
	"github.com/go-ini/ini"
	"github.com/prometheus/common/log"
)

// 配置文件管理

type App struct {
	JwtSecret string
	PageSize int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}
var AppSetting = &App{}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func Init()  {
	Cfg,err := ini.Load("conf/app.ini")
	if err != nil {

	}
}