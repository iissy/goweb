package domain

import (
	"fmt"
	"github.com/juju/errors"
	"hrefs.cn/src/model"
	"hrefs.cn/src/utils"
)

func ListCusLinks() ([]*model.CusLink, error) {
	result := make([]*model.CusLink, 0)
	sql := "select id,title,url,linktype from cuslink ORDER BY id desc"
	_, err := dbMap.Select(&result, sql)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}

func ListCusLinksByCatId(catid string) ([]*model.CusLink, error) {
	result := make([]*model.CusLink, 0)
	sql := "select id,title from cuslink where catid = ? ORDER BY id desc"
	_, err := dbMap.Select(&result, sql, catid)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}

func TopCusLinks() ([]*model.CusLink, error) {
	result := make([]*model.CusLink, 0)
	sql := "select id,title from cuslink ORDER BY id desc limit 15"
	_, err := dbMap.Select(&result, sql)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}

func GetCusLinkUrl(id int) (string, error) {
	sql := "select url FROM cuslink where id = ?"
	result, err := dbMap.SelectStr(sql, id)
	if err != nil {
		return "", errors.Trace(err)
	}

	return result, err
}

func GetCusLink(id int) (*model.CusLink, error) {
	result := new(model.CusLink)
	sql := "select id,title,url,linktype,catid from cuslink where id = ?"
	err := dbMap.SelectOne(&result, sql, id)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}

func GetCusLinkList(page int, size int, search *model.Search) (*model.CusLinkList, error) {
	result := new(model.CusLinkList)
	list := make([]*model.CusLink, 0)
	params := make([]string, 0)
	where := "where 1=1"
	if search != nil && len(search.CatId) > 0 {
		where = "where catid = ?"
		params = append(params, search.CatId)
	}
	if search != nil && len(search.Title) > 0 {
		where += " and title like ?"
		params = append(params, "%"+search.Title+"%")
	}
	if search != nil && len(search.Url) > 0 {
		where += " and url like ?"
		params = append(params, "%"+search.Url+"%")
	}
	sql := "select id,title,url,linktype,visited,adddate,updatedate from cuslink %s order by id desc limit ?, ?"
	sql = fmt.Sprintf(sql, where)
	args, _ := utils.BuildSqlArgs(params, (page-1)*size, size)
	_, err := dbMap.Select(&list, sql, args...)
	if err != nil {
		return nil, errors.Trace(err)
	}

	sql = "select count(*) from cuslink %s"
	sql = fmt.Sprintf(sql, where)
	args, _ = utils.BuildSqlArgs(params)
	total, err := dbMap.SelectInt(sql, args...)
	if err != nil {
		return nil, errors.Trace(err)
	}

	result.List = list
	result.Total = total
	return result, nil
}

func DeleteCusLink(id int) (int64, error) {
	sql := "delete from cuslink where id = ?"
	result, err := dbMap.Exec(sql, id)
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result.RowsAffected()
}

func SaveCusLink(cuslink *model.CusLink) (int64, error) {
	sql := "INSERT INTO cuslink(id, title,url,status,catid,linktype,adddate) values (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE title = ?, url = ?, catid = ?, linktype = ?, updatedate = ?;"
	result, err := dbMap.Exec(sql, cuslink.Id, cuslink.Title, cuslink.Url, cuslink.Status, cuslink.Catid, cuslink.LinkType, cuslink.Adddate,
		cuslink.Title, cuslink.Url, cuslink.Catid, cuslink.LinkType, cuslink.Updatedate)
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result.RowsAffected()
}

func UpdateCusLinkVisited(id int) error {
	sql := "update cuslink set visited = visited + 1 where id = ?"
	_, err := dbMap.Exec(sql, id)
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}
