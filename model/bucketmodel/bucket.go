package bucketmodel

import (
	"github.com/Ericwyn/Staufen/repo"
	"github.com/Ericwyn/Staufen/repo/bucketrepo"
	"github.com/Ericwyn/Staufen/util/gen"
)


// 创建一个 bucket，返回一个
func CreateBucket(bucketName string, public bool, compress bool) (*repo.Bucket,error) {
	bucket := repo.Bucket{
		//Id:         0,
		BucketId:   gen.GeneralUuid(),
		Name:       bucketName,
		ApiToken:   gen.GeneralUuid(),
		ReqToken:   gen.GeneralUuid(),
		Public:     public,
		Compress:   compress,
		//CreateTime: time.Time{},
		//UpdateTime: time.Time{},
	}
	err := bucketrepo.SaveBucket(bucket)
	if err != nil {
		return nil, err
	}
	return &bucket, nil
}