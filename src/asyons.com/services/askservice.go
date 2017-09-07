package services

import (
	"database/sql"
	"html/template"

	"asyons.com/models"
	"asyons.com/utils"
)

const (
	sqldb = "root:hm3366@tcp(192.168.236.132:3306)/iPayask?charset=utf8"
)

// Index is yes.
func Index(terms []string) (*models.AskList, error) {
	var result models.AskList
	result.TotalCount = 15
	result.Items = []*models.Ask{}

	db, err := sql.Open("mysql", sqldb)
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
	db, err := sql.Open("mysql", sqldb)
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

	rows.Close()
	return &result, nil
}

// Post is yes
func Post(ask models.Ask) (bool, error) {
	db, err := sql.Open("mysql", sqldb)
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
func Login(user models.User) (*models.User, error) {
	db, err := sql.Open("mysql", sqldb)
	utils.CheckErr(err)
	rows, err := db.Query("select Id,UserId,UserName from Account where UserId = ? and Password = ?", user.UserID, user.Password)
	utils.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.UserID, &user.UserName)
		utils.CheckErr(err)
		break
	}

	return &user, nil
}

// List is yes.
func List(terms []string) (*models.AskList, error) {
	var result models.AskList
	result.TotalCount = 15
	result.Items = []*models.Ask{}

	db, err := sql.Open("mysql", sqldb)
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

// User is yes.
func User(id string) (*models.User, error) {
	var result models.User

	db, err := sql.Open("mysql", sqldb)
	utils.CheckErr(err)
	rows, err := db.Query("select Id,UserId,UserName,RegDate,LastLoginDate,Status from Account where id=?", id)
	utils.CheckErr(err)

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.UserID, &result.UserName, &result.RegDate, &result.LastLoginDate, &result.Status)
		utils.CheckErr(err)
		break
	}

	return &result, nil
}

// RegPost is yes.
func RegPost(user models.User) (bool, error) {
	db, err := sql.Open("mysql", sqldb)
	utils.CheckErr(err)
	stmt, err := db.Prepare("insert Account set UserId=?,UserName=?,Password=?")
	utils.CheckErr(err)

	res, err := stmt.Exec(user.UserID, user.UserName, user.Password)
	utils.CheckErr(err)

	result, err := res.RowsAffected()
	utils.CheckErr(err)

	return result > 0, nil
}

// Search is yes.
func Search(terms []string) (*models.AskList, error) {
	var result models.AskList
	result.TotalCount = 15
	result.Items = []*models.Ask{}

	db, err := sql.Open("mysql", sqldb)
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
