package conf

import (
	"flag"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	env      string
	filePath string
	Conf     = &Config{}
)

type Config struct {
	ProjectName string
	DB          *DB
	Redis       *Redis
	HttpServer  *HttpServer
	WeChatPay   *WeChatPay
	AliPay      *AliPay
	QQPay       *QQPay
	ReloadTime  int
	RedisExpire int
	NotifyURL   []string
}

type DB struct {
	Driver  string
	Dsn     string
	Idle    int
	ShowSQL bool
}

type Redis struct {
	Addr     string
	Db       int
	Password string
}

type HttpServer struct {
	Port string
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

type QQPay struct {
	MchId          string
	ApiKey         string
	CertFilePath   string
	KeyFilePath    string
	Pkcs12FilePath string
}

func init() {
	flag.StringVar(&env, "env", "", "env")
	flag.StringVar(&filePath, "conf", "", "conf file path")
}

// 解析配置文件
func ParseConfig() error {
	flag.Parse()
	if filePath == "" {
		return errors.New("load conf path fail")
	}
	viper.SetConfigFile(filePath)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return viper.UnmarshalKey(env, Conf)
}
