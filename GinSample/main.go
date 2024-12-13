package main

import (
	"GoSamples/GinSample/assets"
	"GoSamples/GinSample/config"
	"GoSamples/GinSample/routers"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

func main() {
	r := gin.Default()

	var config *config.Conifg

	// 加载配置文件
	file, err := ioutil.ReadFile("./resources/application.yaml")

	if err != nil {
		log.Fatal("加载配置文件失败，err = ", err)
	}

	if err = yaml.Unmarshal(file, &config); err != nil {
		log.Fatal("配置文件解析失败，err = ", err)
	}

	// 初始化路由
	routers.InitRouter(r)

	// 初始化数据库
	// domain.InitDb(config)
	// defer db.Close()

	assets.Load("./assets/server/static")
	r.StaticFS("/static", assets.FileSystem)

	port := ":" + strconv.Itoa(int(config.Port))
	if err := r.Run(port); err != nil {
		log.Fatal("程序启动失败: ", err)
	}
}
