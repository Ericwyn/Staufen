package piccontroller

import (
	"github.com/gin-gonic/gin"
)

const codeSuccess int = 2000
const codeParamError int = 4000
const codeTokenError int = 4001
const codeFileNotFound int = 4000
const codeServerError int = 5000

func resServerToken(msg string) gin.H {
	return gin.H{
		"code": codeServerError,
		"msg":  "server error " + msg,
	}
}

func resErrorToken(msg string) gin.H {
	return gin.H{
		"code": codeTokenError,
		"msg":  "token error " + msg,
	}
}

func resFileNotFound(msg string) gin.H {
	return gin.H{
		"code": codeFileNotFound,
		"msg":  "file not found " + msg,
	}
}

func resFileError(msg string) gin.H {
	return gin.H{
		"code": codeParamError,
		"msg":  "file error " + msg,
	}
}

func resUploadSuccess(fileId string) gin.H {
	return gin.H{
		"code":   codeSuccess,
		"fileId": fileId,
		"msg":    "upload success",
	}
}
