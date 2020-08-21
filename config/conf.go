package config

import (
	"gotil/orm"
	"gotil/web"
)

var Conf = &Config{}

type Config struct {
	// project name
	Name string
	// appkeys
	AppKey []string
	// cpu number
	NumCPU int
	// mysql config
	MySQL *orm.MySQLConfig
	// redis config
	Redis *orm.RedisConfig
	// web config
	Web *web.Config
	// wechat config
	WeChatPay *WeChatPay
	// alipay config
	AliPay *AliPay
}

type WeChatPay struct {
	Appid          string
	MchId          string
	ApiKey         string
	IsProd         bool
	CertFilePath   string
	KeyFilePath    string
	Pkcs12FilePath string
}

type AliPay struct {
	Appid          string
	PrivateKey     string
	IsProd         bool
	AppCertPath    string
	RootCertPath   string
	PublicCertPath string
}
