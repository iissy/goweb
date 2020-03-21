package domain

import (
	"fmt"
	"github.com/juju/errors"
	"hrefs.cn/src/model"
	"hrefs.cn/src/utils"
)

func ListLinks(cat_id string) ([]*model.Link, error) {
	result := make([]*model.Link, 0)
	sql := "select id,title,visited,brief,linktype from link where catid = ? ORDER BY visited desc"
	_, err := dbMap.Select(&result, sql, cat_id)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}

func IndexLinks() ([]*model.Link, error) {
	result := make([]*model.Link, 0)
	sql := "select icon,title,id,catid,linkType,brief FROM link ORDER BY linkType asc,visited desc"
	_, err := dbMap.Select(&result, sql)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}

func GetLinkUrl(id string) (string, error) {
	sql := "select url FROM link where id = ?"
	result, err := dbMap.SelectStr(sql, id)
	if err != nil {
		return "", errors.Trace(err)
	}

	return result, err
}

func LinkVisitedCount() (int64, error) {
	sql := "select sum(visited) from link"
	result, err := dbMap.SelectInt(sql)
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result, err
}

func GetLink(id string) (*model.Link, error) {
	result := new(model.Link)
	sql := "select id,title,visited,brief,linktype,icon,catid,url from link where id = ?"
	err := dbMap.SelectOne(&result, sql, id)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}

func GetLinkList(page int, size int, search *model.Search) (*model.LinkList, error) {
	result := new(model.LinkList)
	list := make([]*model.Link, 0)
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
	sql := "select id,title,visited,brief,linktype,icon,catid,url,createtime from link %s order by createtime desc limit ?, ?"
	sql = fmt.Sprintf(sql, where)
	args, _ := utils.BuildSqlArgs(params, (page-1)*size, size)
	_, err := dbMap.Select(&list, sql, args...)
	if err != nil {
		return nil, errors.Trace(err)
	}

	sql = "select count(*) from link %s"
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

func DeleteLink(id string) (int64, error) {
	sql := "delete from link where id = ?"
	result, err := dbMap.Exec(sql, id)
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result.RowsAffected()
}

func SaveLink(link *model.Link) (int64, error) {
	sql := "INSERT INTO link(id,icon,catid,linktype,title,url,brief) values (?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE icon = ?, catid = ?, linktype = ?, title = ?, url = ?, brief = ?;"
	result, err := dbMap.Exec(sql, link.Id, link.Icon, link.Catid, link.LinkType, link.Title, link.Url, link.Brief, link.Icon, link.Catid, link.LinkType, link.Title, link.Url, link.Brief)
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result.RowsAffected()
}

func UpdateLinkVisited(id string) error {
	sql := "update link set visited = visited + 1 where id = ?"
	_, err := dbMap.Exec(sql, id)
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}
