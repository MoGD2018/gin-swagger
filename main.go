package main

import (
	"gin-swagger/dao"
	docs "gin-swagger/docs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)


// @title swagger API
// @version 1.0
// @description 简单的后端登陆注册和文章分类API

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /

func main()  {
	InitConfig()
	dao.InitDB()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())

}

func InitConfig()  {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

