package picmodel

import (
	"github.com/Ericwyn/Staufen/repo"
	"github.com/Ericwyn/Staufen/repo/picrepo"
	"github.com/Ericwyn/Staufen/storage"
	"github.com/Ericwyn/Staufen/util/gen"
	"github.com/Ericwyn/Staufen/util/strutil"
	"mime/multipart"
)

// 保存图片，并返回一个图片的 id
func SavePicFile(bucket repo.Bucket, picFile *multipart.FileHeader) (string, error) {
	picPath, err := storage.SavePic(picFile, bucket.Compress, storage.LocalFile)

	if err != nil {
		return "", err
	}

	picSave := repo.Picture{
		//Id:         0,
		Uuid:     gen.GeneralUuid(),
		FileName: picFile.Filename,
		FilePath: picPath,
		ExtName:  strutil.GetExtName(picFile.Filename),
		//CreateTime: time.Time{},
		//UpdateTime: time.Time{},
		BucketId: bucket.Id,
	}

	err = picrepo.SavePicture(picSave)
	if err != nil {
		return "", err
	}

	return picSave.Uuid, nil
}
