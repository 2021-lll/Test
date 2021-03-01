package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)
//var filePath = "../src/ipPro/cfgfile/config.yaml"

func ReadConfig(path string, name string)  {
	//viper.AddConfigPath("cfgfile")
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	//viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//log.Panicln("no such config file")
		}else {
			//log.Println("read config error")
		}
		log.Fatal(err)
	}

	fmt.Println("==================")
	fmt.Println(viper.AllSettings())
	fmt.Println("==================")
	//fmt.Println("用户密码",viper.Get("data.password"))
}





