package ado

import (
	"database/sql"

	"iissy.com/models"
	"iissy.com/utils"
)

// Login is yes
func Login(user models.User) (*models.User, error) {
	db, err := sql.Open("mysql", utils.SQLDB)
	utils.CheckErr(err)

	err = db.QueryRow("select Id,UserId,UserName from Account where UserId = ? and Password = ? and Status = 1", user.UserID, user.Password).Scan(&user.ID, &user.UserID, &user.UserName)
	utils.CheckErr(err)
	return &user, nil
}

// Get is yes.
func Get(id int) (*models.User, error) {
	var result models.User

	db, err := sql.Open("mysql", utils.SQLDB)
	utils.CheckErr(err)

	err = db.QueryRow("select Id,UserId,UserName,RegDate,LastLoginDate,Status from Account where id=?", id).Scan(&result.ID, &result.UserID, &result.UserName, &result.RegDate, &result.LastLoginDate, &result.Status)
	utils.CheckErr(err)
	return &result, nil
}

// RegPost is yes.
func RegPost(user models.User) (bool, error) {
	db, err := sql.Open("mysql", utils.SQLDB)
	utils.CheckErr(err)

	res, err := db.Exec("insert Account set UserId=?,UserName=?,Password=?", user.UserID, user.UserName, user.Password)
	utils.CheckErr(err)

	result, err := res.RowsAffected()
	utils.CheckErr(err)

	return result > 0, nil
}
