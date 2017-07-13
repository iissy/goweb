package github

import (
	"database/sql"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	var result IssuesSearchResult
	result.TotalCount = 25
	result.Items = []*Article{}

	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/asyons?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select ID,Keyword,Subject,NickName,Visited from Article limit ?", 10)
	checkErr(err)

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID, &article.Keyword, &article.Subject, &article.NickName, &article.Visited)
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
