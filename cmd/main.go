package main

import (
	"web_template/config"
)

func main() {
	err := config.ParseConfig()
	if err != nil {
		panic(err)
	}

}
