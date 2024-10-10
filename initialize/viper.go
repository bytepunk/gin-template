package initialize

import (
	"flag"
	"fmt"

	"github.com/bytepunk/gin-template/global"
	"github.com/spf13/viper"
)

var configFile string

func parseCommand()  {
	flag.StringVar(&configFile,"c","./config.yaml","配置")
	fmt.Println("configFile:",configFile)
	flag.Parse()
}


func LoadConfig() {
	parseCommand()
	if len(configFile)==0 {
		panic("no config file")
	}

	v:=viper.New()
    v.SetConfigFile(configFile)
	if err := v.ReadInConfig();err != nil {
		panic(fmt.Errorf("配置解析失败:%s\n",err))
	}
	// fmt.Printf("add address:%v", v.GetString("app.addr"))
	if err:=v.Unmarshal(&global.GlobalConfig);err!=nil {
		panic("配置加载失败")
	}

	// fmt.Println(global.GlobalConfig.App.Addr)
}