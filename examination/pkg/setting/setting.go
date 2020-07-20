package setting
//这个包会读取  ---配置数据---


import (
	"log"
	"time"

"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string    //gin执行模式

	HTTPPort int   //端口号
	ReadTimeout time.Duration   //
	WriteTimeout time.Duration  //

	PageSize int
	JwtSecret string
)

func init() {  //调用包默认执行此函数
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {    //载入RunMode
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("release")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")  //获取server组
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)

}

func LoadApp() {
	sec, err := Cfg.GetSection("app")//获取app组
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("wdnmd")

}