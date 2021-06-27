package bucketrepo

import (
	"github.com/Ericwyn/Staufen/repo"
	"github.com/Ericwyn/Staufen/util/log"
	"xorm.io/xorm"
)

func SaveBucket(bucket repo.Bucket) error {
	_, err := repo.SqlEngine.Insert(bucket)
	if err != nil {
		return err
	}
	return nil
}

func GetBucketById(id int64) *repo.Bucket {
	bucket := new(repo.Bucket)
	res, err := repo.SqlEngine.Where("id=?", id).Get(bucket)
	if err != nil {
		log.E("get bucket error : ", bucket, err)
		return nil
	}
	if res {
		return bucket
	} else {
		return nil
	}

}

func GetBucketByUuid(bucketUuid string) *repo.Bucket {
	bucket := new(repo.Bucket)
	res, err := repo.SqlEngine.Where("bucket_id=?", bucketUuid).Get(bucket)
	if err != nil {
		log.E("get bucket error : ", bucket, err)
		return nil
	}
	if res {
		return bucket
	} else {
		return nil
	}
}

func GetBucketByApiToken(apiToken string) *repo.Bucket {
	bucket := new(repo.Bucket)
	res, err := repo.SqlEngine.Where("api_token=?", apiToken).Get(bucket)
	if err != nil {
		log.E("get bucket error : ", bucket, err)
		return nil
	}
	if res {
		return bucket
	} else {
		return nil
	}
}

func UpdateBucket(bucket repo.Bucket) error {
	bucketSave := new(repo.Bucket)
	_, err := repo.SqlEngine.ID(bucket.Id).Get(&bucketSave)
	if err != nil {
		return err
	}

	// 替换属性, 后续智能修改公开/bucket 名称
	bucketSave.Name = bucket.Name
	//bucketSave.Name = bucket.Name
	//bucketSave.ApiToken = bucket.ApiToken
	//bucketSave.ReqToken = bucket.ReqToken
	bucketSave.Public = bucket.Public
	//bucketSave.Compress = bucket.Compress
	_, err = repo.SqlEngine.ID(bucketSave.Id).Update(bucketSave)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBucket(bucket repo.Bucket, session *xorm.Session) error {
	if session != nil {
		_, err := session.ID(bucket.Id).Delete(bucket)
		if err != nil {
			return err
		}
		return nil
	} else {
		_, err := repo.SqlEngine.ID(bucket.Id).Delete(bucket)
		if err != nil {
			return err
		}
		return nil
	}
}

func ListBucket() ([]repo.Bucket, error) {
	buckets := make([]repo.Bucket, 0)

	// TODO 分页
	err := repo.SqlEngine.Find(&buckets)

	if err != nil {
		return nil, err
	}
	return buckets, nil
}
