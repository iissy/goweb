package access

import (
	"math"

	"iissy.com/models"
	"iissy.com/utils"
)

// Login is yes
func Login(user models.User) (*models.User, error) {
	err := db.QueryRow("select Id,UserId,UserName from Account where UserId = ? and Password = ? and Status = 1", user.UserID, user.Password).Scan(&user.ID, &user.UserID, &user.UserName)
	utils.CheckErr(err)
	return &user, nil
}

// AccountList is yes.
func AccountList(page, size int) (*models.Users, error) {
	var result models.Users
	result.Items = []*models.User{}

	rows, err := db.Query("select Id,UserId,UserName,RegDate,LastLoginDate,Status from Account order by RegDate desc limit ?, ?", (page-1)*size, size)
	utils.CheckErr(err)

	for rows.Next() {
		item := models.User{}
		err = rows.Scan(&item.ID, &item.UserID, &item.UserName, &item.RegDate, &item.LastLoginDate, &item.Status)
		utils.CheckErr(err)

		result.Items = append(result.Items, &item)
	}

	total := 0
	err = db.QueryRow("select count(*) from Account").Scan(&total)
	utils.CheckErr(err)

	pageCount := int(math.Ceil(float64(total) / float64(size)))
	result.PageArgs = models.PageArgs{PageNumber: page, TotalCount: total, PageSize: size, PageCount: pageCount}
	rows.Close()
	return &result, nil
}

// Get is yes.
func Get(id int) (*models.User, error) {
	var result models.User
	err := db.QueryRow("select Id,UserId,UserName,RegDate,LastLoginDate,Status from Account where id=?", id).Scan(&result.ID, &result.UserID, &result.UserName, &result.RegDate, &result.LastLoginDate, &result.Status)
	utils.CheckErr(err)
	return &result, nil
}

// RegPost is yes.
func RegPost(user models.User) (bool, error) {
	res, err := db.Exec("insert Account set UserId=?,UserName=?,Password=?", user.UserID, user.UserName, user.Password)
	utils.CheckErr(err)

	result, err := res.RowsAffected()
	utils.CheckErr(err)

	return result > 0, nil
}
