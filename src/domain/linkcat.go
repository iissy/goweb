package domain

import (
	"github.com/juju/errors"
	"hrefs.cn/src/model"
)

func GetCatOptions() ([]*model.LinkCat, error) {
	result := make([]*model.LinkCat, 0)
	sql := "SELECT id,catname FROM linkcat ORDER BY id desc"
	_, err := dbMap.Select(&result, sql)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, nil
}

func GetLinkCat(id string) (*model.LinkCat, error) {
	result := new(model.LinkCat)
	sql := "SELECT id,catname FROM linkcat where id = ?"
	err := dbMap.SelectOne(&result, sql, id)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, nil
}
