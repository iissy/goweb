package domain

import (
	"fmt"
	"github.com/juju/errors"
	"hrefs.cn/src/model"
	"hrefs.cn/src/utils"
)

func ListArticles() ([]*model.Article, error) {
	result := make([]*model.Article, 0)
	sql := "select id,title,icon,brief,createTime from article ORDER BY createTime desc"
	_, err := dbMap.Select(&result, sql)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}

func TopArticles() ([]*model.Article, error) {
	result := make([]*model.Article, 0)
	sql := "select id,title from article ORDER BY createTime desc limit 3"
	_, err := dbMap.Select(&result, sql)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}

func GetArticle(id string) (*model.Article, error) {
	result := new(model.Article)
	sql := "select id,title,body,icon,brief,createTime from article where id = ?"
	err := dbMap.SelectOne(&result, sql, id)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}

func GetArticleList(page int, size int, search *model.Search) (*model.ArticleList, error) {
	result := new(model.ArticleList)
	list := make([]*model.Article, 0)
	params := make([]string, 0)
	where := "where 1=1"
	if search != nil && len(search.Id) > 0 {
		where = "where id = ?"
		params = append(params, search.Id)
	}

	if search != nil && len(search.Title) > 0 {
		where += " and title like ?"
		params = append(params, "%"+search.Title+"%")
	}
	sql := "select id,title,icon,brief,createTime,visited from article %s order by createTime desc limit ?, ?"
	sql = fmt.Sprintf(sql, where)
	args, _ := utils.BuildSqlArgs(params, (page-1)*size, size)
	_, err := dbMap.Select(&list, sql, args...)
	if err != nil {
		return nil, errors.Trace(err)
	}

	sql = "select count(*) from article %s"
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

func DeleteArticle(id string) (int64, error) {
	sql := "delete from article where id = ?"
	result, err := dbMap.Exec(sql, id)
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result.RowsAffected()
}

func SaveArticle(article *model.Article) (int64, error) {
	sql := "INSERT INTO article(id,icon,title,brief,body,createTime) values (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE icon = ?, title = ?, brief = ?, body = ?;"
	result, err := dbMap.Exec(sql, article.Id, article.Icon, article.Title, article.Brief, article.Body, article.CreateTime, article.Icon, article.Title, article.Brief, article.Body)
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result.RowsAffected()
}

func UpdateArticleVisited(id string) error {
	sql := "update article set visited = visited + 1 where id = ?"
	_, err := dbMap.Exec(sql, id)
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}
