package picrepo

import (
	"github.com/Ericwyn/Staufen/repo"
	"github.com/Ericwyn/Staufen/util/log"
	"xorm.io/xorm"
)

func SavePicture(pic repo.Picture) error {
	_, err := repo.SqlEngine.Insert(pic)
	return err
}

func GetPicture(picUuid string) *repo.Picture {
	pic := new(repo.Picture)
	res, err := repo.SqlEngine.Where("uuid=?", picUuid).Get(pic)
	if err != nil {
		log.E("get picture error : ", picUuid, err)
		return nil
	}
	if res {
		return pic
	}
	return nil
}

func ListPictureByBucket(bucket repo.Bucket) ([]repo.Picture, error) {
	//pic := new(repo.Picture)

	pics := make([]repo.Picture, 0)

	err := repo.SqlEngine.Where("bucket_id=?", bucket.Id).Find(&pics)
	if err != nil {
		return nil, err
	}
	return pics, nil
}

func DeletePictureByBucket(bucket repo.Bucket, session *xorm.Session) error {
	if session != nil {
		_, err := session.Delete(repo.Picture{
			BucketId: bucket.Id,
		})
		return err
	} else {
		_, err := repo.SqlEngine.Delete(repo.Picture{
			BucketId: bucket.Id,
		})
		return err
	}
}
