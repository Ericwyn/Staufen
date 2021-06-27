package piccontroller

import (
	"github.com/Ericwyn/Staufen/model/authmodel"
	"github.com/Ericwyn/Staufen/repo/bucketrepo"
	"github.com/gin-gonic/gin"
)

// 各个 SDK Client 通过 API Token 获取 upload Token
type UploadTokenRequest struct {
	T string `json:"t"`
}

func GetUploadToken(ctx *gin.Context) {
	request := new(UploadTokenRequest)
	ctx.BindJSON(request)

	bucket := bucketrepo.GetBucketByApiToken(request.T)
	if bucket == nil {
		ctx.JSON(200, resErrorToken(" : api token error"))
		return
	}

	token := authmodel.GetUploadToken(*bucket)

	ctx.JSON(200, gin.H{
		"uploadToken": token,
	})
}
