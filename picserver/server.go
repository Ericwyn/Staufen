package picserver

import (
	"github.com/Ericwyn/Staufen/config"
	"github.com/Ericwyn/Staufen/piccontroller"
	"net/http"
	"time"
)

// 启动一个图片上传和下载服务器
func startServer(config config.PicServerConfig) {
	s := &http.Server{
		Addr:           ":" + config.Port,
		Handler:        piccontroller.NewMux(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
