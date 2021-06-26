package piccontroller

import (
	"github.com/Ericwyn/Staufen/util/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func NewMux() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//store := cookie.NewStore(getCookieKey())

	router.Use(gin.Logger())
	//router.LoadHTMLGlob("static/*.html")

	initApi(router)
	return router
}

func initApi(router *gin.Engine) {
	// 图片上传

	router.POST("/api/v1/pic/upload", uploadPic)
	router.GET("/pic/*picUuid", downloadPic)

}

// 启动一个图片上传和下载服务器
func StartPicHttpServer() {
	serverRunPort := viper.GetString("server.port")
	log.I("try to start http service in :" + serverRunPort)

	s := &http.Server{
		Addr:           ":" + serverRunPort,
		Handler:        NewMux(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
