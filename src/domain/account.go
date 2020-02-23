package domain

import (
	"github.com/juju/errors"
	"hrefs.cn/src/model"
)

func Login(user *model.Account) (*model.Account, error) {
	result := new(model.Account)
	sql := "select id,userid,username,regdate from account where UserId = ? and Password = ?"
	err := dbMap.SelectOne(&result, sql, user.UserId, user.Password)
	if err != nil {
		return nil, errors.Trace(err)
	}

	sql = "update account set lastlogindate = ? where id = ?"
	dbMap.Exec(sql, user.LastLoginDate, result.ID)

	return result, err
}

func GetAccount(id int) (*model.Account, error) {
	result := new(model.Account)
	sql := "select id,userid,username,regdate from account where id = ?"
	err := dbMap.SelectOne(&result, sql, id)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}

func GetrAccountList(page int, size int) (*model.AccountList, error) {
	result := new(model.AccountList)
	list := make([]*model.Account, 0)
	sql := "select id,userid,username,regdate,lastlogindate from account order by regdate desc limit ?, ?"
	_, err := dbMap.Select(&list, sql, (page-1)*size, size)
	if err != nil {
		return nil, errors.Trace(err)
	}

	total, err := dbMap.SelectInt("select count(*) from account")
	if err != nil {
		return nil, errors.Trace(err)
	}

	result.List = list
	result.Total = total
	return result, nil
}
