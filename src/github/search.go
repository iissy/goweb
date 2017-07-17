package github

import (
	"database/sql"
	"html/template"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	var result IssuesSearchResult
	result.TotalCount = 25
	result.Items = []*Article{}

	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/asyons?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select ID,Keyword,Subject,NickName,Visited,Description,AddDate,Picture from Article limit ?", 10)
	checkErr(err)

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID, &article.Keyword, &article.Subject, &article.NickName, &article.Visited, &article.Description, &article.AddDate, &article.Picture)
		checkErr(err)

		result.Items = append(result.Items, &article)
	}

	return &result, nil
}

// Detail is for article
func Detail(id string) (*Article, error) {
	var result Article

	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/asyons?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select ID,Keyword,Subject,NickName,Visited,Description,AddDate,Body from Article where ID = ?", id)
	checkErr(err)

	for rows.Next() {
		var article Article
		var body string
		err = rows.Scan(&article.ID, &article.Keyword, &article.Subject, &article.NickName, &article.Visited, &article.Description, &article.AddDate, &body)
		checkErr(err)

		article.Body = template.HTML(body)
		result = article
	}

	return &result, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
