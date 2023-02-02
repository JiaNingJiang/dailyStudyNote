package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File // ini配置文件对象

	RunMode string //运行模式，debug模式或其他工作模式

	HTTPPort     int //http对外暴露端口
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int    //每页数据的数量
	JwtSecret string // jwt密码
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()   //读取运行模式
	LoadServer() //读取http服务器配置信息(端口和最大读取和写入时间)
	LoadApp()    //读取app服务的配置信息(jwt密码、pagesize等。。。)
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug") //MustString()方法在对应value为空时，返回给定的默认值(即debug)
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
