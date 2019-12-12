package conf

import (
	"flag"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	filePath string
	Conf     = &Config{}
)

type Config struct {
	ProjectName string
	DB          *DB
	Redis       *Redis
	HTTP        *HTTPServer
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

type HTTPServer struct {
	Port string
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
	return viper.Unmarshal(Conf)
}

func init() {
	flag.StringVar(&filePath, "conf", "", "conf file path")
}
