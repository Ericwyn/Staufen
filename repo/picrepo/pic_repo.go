package picrepo

import (
	"github.com/Ericwyn/Staufen/repo"
	"github.com/Ericwyn/Staufen/util/log"
)

func SavePicture(pic repo.Picture) error {
	_, err := repo.SqlEngine.Insert(pic)
	if err != nil {
		return err
	}
	return nil
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
