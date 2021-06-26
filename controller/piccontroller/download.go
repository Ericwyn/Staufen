package piccontroller

import (
	"github.com/Ericwyn/Staufen/repo/bucketrepo"
	"github.com/Ericwyn/Staufen/repo/picrepo"
	"github.com/Ericwyn/Staufen/storage"
	"github.com/gin-gonic/gin"
	"strings"
)

const maxAge = 3600 * 48 // 48小时过去

func downloadPic(ctx *gin.Context) {
	// 获取 query id
	// id 可能是 ucekmwgfmbhpa31y3t5lsqfptmuwhuhw 也可能是 ucekmwgfmbhpa31y3t5lsqfptmuwhuhw.jpg 之类的
	// 所以需要先去除后缀
	picUuid := ctx.Param("picUuid")
	picUuid = strings.Replace(picUuid, "/", "", -1)
	if strings.Index(picUuid, ".") >= 0 {
		picUuid = strings.Split(picUuid, ".")[0]
	}

	// 搜索 picUuid 是否存在
	picture := picrepo.GetPicture(picUuid)
	if picture == nil {
		ctx.JSON(200, resFileNotFound("id: "+picUuid))
		return
	}

	// 搜索 bucket 是否为 public
	bucket := bucketrepo.GetBucketById(picture.BucketId)
	if bucket != nil && !bucket.Public {
		ctx.JSON(200, resErrorToken("need reqToken to get file"))
		return
	}

	// 直接返回 IO 流
	picFileBytes := storage.GetPicBytes(picture.FilePath, storage.LocalFile)

	if picFileBytes == nil {
		ctx.JSON(200, resFileNotFound("read file error"))
		return
	}

	//ctx.Header("Content-Type", "application/octet-stream")
	//ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	//ctx.Header("Content-Transfer-Encoding", "binary")
	//os.ReadFile()
	//ctx.Data(200, "", io)

	//ctx.Header("Cache-Control", fmt.Sprint("max-age=",maxAge))

	ctx.Data(200, getContentType(picture.ExtName), picFileBytes)
}

func getContentType(extName string) string {
	extName = strings.ToLower(extName)
	if strings.Index(extName, "jpg") >= 0 {
		return "image/jpeg"
	} else if strings.Index(extName, "jpeg") >= 0 {
		return "image/jpeg"
	} else if strings.Index(extName, "png") >= 0 {
		return "image/png"
	} else if strings.Index(extName, "gif") >= 0 {
		return "image/gif"
	}
	return ""
}
