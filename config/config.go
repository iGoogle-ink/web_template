package config

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/pkg/errors"
)

var (
	filePath string
)

// 解析配置文件
func ParseConfig() error {
	if filePath == "" {
		return errors.New("load config path fail")
	}
	fmt.Println("filePath:", filePath)
	return nil
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&filePath, "config", "", "config file path")
	flag.Parse()
}
