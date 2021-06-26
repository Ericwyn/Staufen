package piccontroller

import "github.com/gin-gonic/gin"

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

}
