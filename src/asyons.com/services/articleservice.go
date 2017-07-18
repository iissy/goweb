package services

import (
	"database/sql"
	"html/template"

	"asyons.com/models"
)

// Index queries the GitHub issue tracker.
func Index(terms []string) (*models.ArticleListResult, error) {
	var result models.ArticleListResult
	result.TotalCount = 15
	result.Items = []*models.Article{}

	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.1.102:3306)/iPayask?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select ID,Subject,NickName,Visited,Description,AddDate from Ask limit ?", 10)
	checkErr(err)

	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.ID, &article.Subject, &article.NickName, &article.Visited, &article.Description, &article.AddDate)
		checkErr(err)

		result.Items = append(result.Items, &article)
	}

	return &result, nil
}

// Detail is for article
func Detail(id string) (*models.Article, error) {
	var result models.Article

	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.1.102:3306)/iPayask?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select ID,Subject,NickName,Visited,Description,AddDate,Body from Ask where Id = ?", id)
	checkErr(err)

	for rows.Next() {
		var article models.Article
		var body string
		err = rows.Scan(&article.ID, &article.Subject, &article.NickName, &article.Visited, &article.Description, &article.AddDate, &body)
		checkErr(err)

		article.Body = template.HTML(body)
		result = article
	}

	return &result, nil
}

// Post is yes
func Post(article models.Article) (bool, error) {
	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/asyons?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select ID,Keyword,Subject,NickName,Visited,Description,AddDate,Body from Article where Id = ?")
	checkErr(err)

	for rows.Next() {
		var article models.Article
		var body string
		err = rows.Scan(&article.ID, &article.Keyword, &article.Subject, &article.NickName, &article.Visited, &article.Description, &article.AddDate, &body)
		checkErr(err)

		article.Body = template.HTML(body)
	}

	return true, nil
}

// Login is yes
func Login(article models.User) (bool, error) {
	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/asyons?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select ID,Keyword,Subject,NickName,Visited,Description,AddDate,Body from Article where Id = ?")
	checkErr(err)

	for rows.Next() {
		var article models.Article
		var body string
		err = rows.Scan(&article.ID, &article.Keyword, &article.Subject, &article.NickName, &article.Visited, &article.Description, &article.AddDate, &body)
		checkErr(err)

		article.Body = template.HTML(body)
	}

	return true, nil
}

// List queries the GitHub issue tracker.
func List(terms []string) (*models.ArticleListResult, error) {
	var result models.ArticleListResult
	result.TotalCount = 15
	result.Items = []*models.Article{}

	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/iPayask?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select Id,Subject,NickName,Visited,Description,AddDate from Ask limit ?", 10)
	checkErr(err)

	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.ID, &article.Subject, &article.NickName, &article.Visited, &article.Description, &article.AddDate)
		checkErr(err)

		result.Items = append(result.Items, &article)
	}

	return &result, nil
}

// User queries the GitHub issue tracker.
func User(terms []string) (*models.ArticleListResult, error) {
	var result models.ArticleListResult
	result.TotalCount = 15
	result.Items = []*models.Article{}

	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/iPayask?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select Id,Subject,NickName,Visited,Description,AddDate from Ask limit ?", 10)
	checkErr(err)

	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.ID, &article.Subject, &article.NickName, &article.Visited, &article.Description, &article.AddDate)
		checkErr(err)

		result.Items = append(result.Items, &article)
	}

	return &result, nil
}

// Mine queries the GitHub issue tracker.
func Mine(terms []string) (*models.ArticleListResult, error) {
	var result models.ArticleListResult
	result.TotalCount = 15
	result.Items = []*models.Article{}

	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/iPayask?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select Id,Subject,NickName,Visited,Description,AddDate from Ask limit ?", 10)
	checkErr(err)

	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.ID, &article.Subject, &article.NickName, &article.Visited, &article.Description, &article.AddDate)
		checkErr(err)

		result.Items = append(result.Items, &article)
	}

	return &result, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
