package services

import (
	"database/sql"
	"html/template"

	"asyons.com/models"
	"asyons.com/utils"
)

// Index queries the GitHub issue tracker.
func Index(terms []string) (*models.AskList, error) {
	var result models.AskList
	result.TotalCount = 15
	result.Items = []*models.Ask{}

	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/iPayask?charset=utf8")
	utils.CheckErr(err)
	rows, err := db.Query("select ID,Subject,NickName,Visited,Description,AddDate from Ask order by AddDate desc limit ?", 10)
	utils.CheckErr(err)

	for rows.Next() {
		var ask models.Ask
		err = rows.Scan(&ask.ID, &ask.Subject, &ask.NickName, &ask.Visited, &ask.Description, &ask.AddDate)
		utils.CheckErr(err)

		result.Items = append(result.Items, &ask)
	}

	return &result, nil
}

// Detail is for article
func Detail(id string) (*models.Ask, error) {
	var result models.Ask
	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/iPayask?charset=utf8")
	utils.CheckErr(err)
	rows, err := db.Query("select ID,Subject,NickName,Visited,Description,AddDate,Body from Ask where Id = ?", id)
	utils.CheckErr(err)

	for rows.Next() {
		var ask models.Ask
		var body string
		err = rows.Scan(&ask.ID, &ask.Subject, &ask.NickName, &ask.Visited, &ask.Description, &ask.AddDate, &body)
		utils.CheckErr(err)

		ask.Body = template.HTML(body)
		result = ask
	}

	return &result, nil
}

// Post is yes
func Post(ask models.Ask) (bool, error) {
	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/iPayask?charset=utf8")
	utils.CheckErr(err)
	stmt, err := db.Prepare("insert Ask set ID=?,Subject=?,Description=?,Body=?,UserID=?,NickName=?")
	utils.CheckErr(err)

	res, err := stmt.Exec(ask.ID, ask.Subject, ask.Description, string(ask.Body), ask.UserID, ask.NickName)
	utils.CheckErr(err)

	result, err := res.RowsAffected()
	utils.CheckErr(err)

	return result > 0, nil
}

// Login is yes
func Login(article models.User) (bool, error) {
	return true, nil
}

// List queries the GitHub issue tracker.
func List(terms []string) (*models.AskList, error) {
	var result models.AskList
	result.TotalCount = 15
	result.Items = []*models.Ask{}

	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/iPayask?charset=utf8")
	utils.CheckErr(err)
	rows, err := db.Query("select ID,Subject,NickName,Visited,Description,AddDate from Ask limit ?", 10)
	utils.CheckErr(err)

	for rows.Next() {
		var ask models.Ask
		err = rows.Scan(&ask.ID, &ask.Subject, &ask.NickName, &ask.Visited, &ask.Description, &ask.AddDate)
		utils.CheckErr(err)

		result.Items = append(result.Items, &ask)
	}

	return &result, nil
}

// User queries the GitHub issue tracker.
func User(terms []string) (*models.AskList, error) {
	var result models.AskList
	result.TotalCount = 15
	result.Items = []*models.Ask{}

	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/iPayask?charset=utf8")
	utils.CheckErr(err)
	rows, err := db.Query("select ID,Subject,NickName,Visited,Description,AddDate from Ask limit ?", 10)
	utils.CheckErr(err)

	for rows.Next() {
		var ask models.Ask
		err = rows.Scan(&ask.ID, &ask.Subject, &ask.NickName, &ask.Visited, &ask.Description, &ask.AddDate)
		utils.CheckErr(err)

		result.Items = append(result.Items, &ask)
	}

	return &result, nil
}

// Mine queries the GitHub issue tracker.
func Mine(terms []string) (*models.AskList, error) {
	var result models.AskList
	result.TotalCount = 15
	result.Items = []*models.Ask{}

	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/iPayask?charset=utf8")
	utils.CheckErr(err)
	rows, err := db.Query("select ID,Subject,NickName,Visited,Description,AddDate from Ask limit ?", 10)
	utils.CheckErr(err)

	for rows.Next() {
		var ask models.Ask
		err = rows.Scan(&ask.ID, &ask.Subject, &ask.NickName, &ask.Visited, &ask.Description, &ask.AddDate)
		utils.CheckErr(err)

		result.Items = append(result.Items, &ask)
	}

	return &result, nil
}

// Search queries the GitHub issue tracker.
func Search(terms []string) (*models.AskList, error) {
	var result models.AskList
	result.TotalCount = 15
	result.Items = []*models.Ask{}

	db, err := sql.Open("mysql", "root:hm3366@tcp(192.168.236.131:3306)/iPayask?charset=utf8")
	utils.CheckErr(err)
	rows, err := db.Query("select ID,Subject,NickName,Visited,Description,AddDate from Ask limit ?", 10)
	utils.CheckErr(err)

	for rows.Next() {
		var ask models.Ask
		err = rows.Scan(&ask.ID, &ask.Subject, &ask.NickName, &ask.Visited, &ask.Description, &ask.AddDate)
		utils.CheckErr(err)

		result.Items = append(result.Items, &ask)
	}

	return &result, nil
}
