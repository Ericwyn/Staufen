package piccontroller

import (
	"github.com/Ericwyn/Staufen/repo/bucketrepo"
	"github.com/Ericwyn/Staufen/repo/picrepo"
	"github.com/Ericwyn/Staufen/storage"
	"github.com/gin-gonic/gin"
	"strings"
)

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
	// TODO 此处无法获取 bucket
	bucket := bucketrepo.GetBucketById(picture.BucketId)
	if bucket != nil && !bucket.Public {
		ctx.JSON(200, resErrorToken("need reqToken to get file"))
		return
	}

	picFile := storage.GetPic(picture.FilePath)
	if picFile == nil {
		ctx.JSON(200, resFileNotFound("read file error"))
		return
	}

	//ctx.Header("Content-Type", "application/octet-stream")
	//ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	//ctx.Header("Content-Transfer-Encoding", "binary")
	//os.ReadFile()
	//ctx.Data(200, "", io)
	// TODO 文件路径修改
	ctx.File(picture.FilePath)
}
