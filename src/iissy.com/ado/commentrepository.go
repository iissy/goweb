package ado

import (
	"database/sql"
	"html/template"

	"iissy.com/models"
	"iissy.com/utils"
)

// List is yes.
func List(askid int) (*models.Comments, error) {
	var result models.Comments
	result.TotalCount = 15
	result.Items = []*models.Comment{}

	db, err := sql.Open("mysql", utils.SQLDB)
	utils.CheckErr(err)
	rows, err := db.Query("select ID,Userid,NickName,Body,AddDate from Comment where askid = ? order by Id", askid)
	utils.CheckErr(err)

	for rows.Next() {
		item := models.Comment{}
		var body string
		err = rows.Scan(&item.ID, &item.UserID, &item.NickName, &body, &item.AddDate)
		utils.CheckErr(err)
		item.Body = template.HTML(body)

		result.Items = append(result.Items, &item)
	}

	rows.Close()
	return &result, nil
}

// CommentPost is yes
func CommentPost(comment models.Comment) (bool, error) {
	db, err := sql.Open("mysql", utils.SQLDB)
	utils.CheckErr(err)

	res, err := db.Exec("insert Comment set AskId=?,UserID=?,NickName=?,Body=?", comment.AskID, comment.UserID, comment.NickName, string(comment.Body))
	utils.CheckErr(err)

	result, err := res.RowsAffected()
	utils.CheckErr(err)

	return result > 0, nil
}
