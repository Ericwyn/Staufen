package piccontroller

import (
	"github.com/Ericwyn/Staufen/model/authmodel"
	"github.com/Ericwyn/Staufen/model/picmodel"
	"github.com/Ericwyn/Staufen/repo/bucketrepo"
	"github.com/Ericwyn/Staufen/util/log"
	"github.com/gin-gonic/gin"
)

func uploadPic(ctx *gin.Context) {

	// 上传的临时 token
	uploadToken := ctx.Query("token")
	bucketUuid := ctx.Query("bucketUuid")

	// token 和 bucket 校验
	if uploadToken == "" {
		ctx.JSON(200, resErrorToken("token nil"))
		return
	}

	bucket := bucketrepo.GetBucketByUuid(bucketUuid)
	if bucket == nil {
		ctx.JSON(200, resErrorToken("bucket uuid error"))
		return
	}

	if !authmodel.CheckUploadToken(*bucket, uploadToken) {
		ctx.JSON(200, resErrorToken("token error or expired"))
		return
	}

	// 文件校验
	file, err := ctx.FormFile("pic")
	if err != nil {
		ctx.JSON(200, resFileError("file nil"))
		return
	}

	fileId, err := picmodel.SavePicFile(*bucket, file)
	// 保存到数据库当中

	if err != nil {
		log.E(err)
		ctx.JSON(200, resServerToken(""))
	}

	ctx.JSON(200, resUploadSuccess(fileId))
}
