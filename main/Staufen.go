package main

import (
	"flag"
	"fmt"
	"github.com/Ericwyn/Staufen/cmd"
	"github.com/Ericwyn/Staufen/controller/piccontroller"
	"github.com/Ericwyn/Staufen/model/bucketmodel"
	"github.com/Ericwyn/Staufen/repo"
	"github.com/Ericwyn/Staufen/repo/bucketrepo"
	"github.com/Ericwyn/Staufen/util/log"
	"github.com/spf13/viper"
	"os"
)

var picServer = flag.Bool("ps", false, "start picture server")
var cli = flag.Bool("cli", false, "cli")

// 入口类
func main() {
	flag.Parse()

	initConfig()

	if *picServer {
		//CreateNewBucket()
		//ListAllBucketNow()
		// 启动 http 服务器
		piccontroller.StartPicHttpServer()
		return
	}
	if *cli {
		cmd.StartCli()
		return
	}

}

func initConfig() {
	// 载入配置
	//viper.AddConfigPath("./.conf") // 设置读取路径：就是在此路径下搜索配置文件。
	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	viper.SetConfigFile("./.conf/pic_server.yaml") // 设置被读取文件的全名，包括扩展名。
	//viper.SetConfigName("server") // 设置被读取文件的名字： 这个方法 和 SetConfigFile实际上仅使用一个就够了
	err := viper.ReadInConfig() // 读取配置文件： 这一步将配置文件变成了 Go语言的配置文件对象包含了 map，string 等对象。
	if err != nil {
		log.E(err)
		os.Exit(-1)
	}
	// 初始化数据库
	repo.InitDb()
}

func CreateNewBucket() {
	bucket, err := bucketmodel.CreateBucket("公开bucket", true, false)
	if err != nil {
		log.E(err)
	} else {
		fmt.Println(bucket.Name, bucket.BucketId, bucket.ApiToken)
	}
}

func ListAllBucketNow() {
	bucketList, err := bucketrepo.ListBucket()
	if err != nil {
		log.E(err)
	} else {
		for _, bucket := range bucketList {
			fmt.Println(bucket.Name, bucket.BucketId, bucket.ApiToken, bucket.Public)
		}
	}
}
