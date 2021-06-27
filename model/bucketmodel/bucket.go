package bucketmodel

import (
	"github.com/Ericwyn/Staufen/repo"
	"github.com/Ericwyn/Staufen/repo/bucketrepo"
	"github.com/Ericwyn/Staufen/repo/picrepo"
	"github.com/Ericwyn/Staufen/storage"
	"github.com/Ericwyn/Staufen/util/gen"
	"github.com/Ericwyn/Staufen/util/log"
)

// 创建一个 bucket，返回一个
func CreateBucket(bucketName string, public bool, compress bool) (*repo.Bucket, error) {
	bucket := repo.Bucket{
		//Id:         0,
		BucketId: gen.GeneralUuid(),
		Name:     bucketName,
		ApiToken: gen.GeneralUuid(),
		ReqToken: gen.GeneralUuid(),
		Public:   public,
		Compress: compress,
		//CreateTime: time.Time{},
		//UpdateTime: time.Time{},
	}
	err := bucketrepo.SaveBucket(bucket)
	if err != nil {
		return nil, err
	}
	return &bucket, nil
}

func DeleteBucket(bucket repo.Bucket) error {

	session := repo.SqlEngine.NewSession()
	defer session.Close()
	// add Begin() before any action
	err := session.Begin()

	// 删除 bucket
	err = bucketrepo.DeleteBucket(bucket, session)
	if err != nil {
		return err
	}

	// 删除 picture 表
	// TODO 后续修改为软删除 + cron 定时任务处理，删除表 + 本地数据
	err = picrepo.DeletePictureByBucket(bucket, session)
	if err != nil {
		_ = session.Rollback()
		return err
	}
	err = session.Commit()

	// 删除本地文件
	pics, err := picrepo.ListPictureByBucket(bucket)
	for _, picture := range pics {
		// 删除图片
		err = storage.DeletePic(picture.FilePath, storage.LocalFile)
		if err != nil {
			log.E("delete file fail", err)
		}
	}

	return nil
}
