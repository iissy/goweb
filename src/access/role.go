package access

import (
	"math"

	"iissy.com/src/models"
	"iissy.com/src/utils"
)

// RoleList is yes.
func RoleList(userid int, page int, size int) (*models.Roles, error) {
	var result models.Roles
	result.Items = []*models.Role{}

	rows, err := db.Query("select Id,RoleName,Status,CreateTime,UpdatedTime from Roles order by CreateTime desc limit ?, ?", (page-1)*size, size)
	utils.CheckErr(err)

	for rows.Next() {
		item := models.Role{}
		err = rows.Scan(&item.ID, &item.RoleName, &item.Status, &item.CreateTime, &item.UpdatedTime)
		utils.CheckErr(err)

		result.Items = append(result.Items, &item)
	}

	total := 0
	err = db.QueryRow("select count(*) from Roles").Scan(&total)
	utils.CheckErr(err)

	pageCount := int(math.Ceil(float64(total) / float64(size)))
	result.PageArgs = models.PageArgs{PageNumber: page, TotalCount: total, PageSize: size, PageCount: pageCount}
	rows.Close()
	return &result, nil
}

// PostRole is yes
func PostRole(role models.Role) (bool, error) {
	var result int64
	if role.ID > 0 {
		res, err := db.Exec("update Roles set RoleName=?,Status=? where id=?", role.RoleName, role.Status, role.ID)
		utils.CheckErr(err)
		result, err = res.RowsAffected()
		utils.CheckErr(err)
	} else {
		res, err := db.Exec("insert Roles set RoleName=?,Status=?", role.RoleName, role.Status)
		utils.CheckErr(err)
		result, err = res.RowsAffected()
		utils.CheckErr(err)
	}

	return result > 0, nil
}

// GetRole is for article
func GetRole(id string) (*models.Role, error) {
	var result models.Role

	err := db.QueryRow("select Id,RoleName,Status from Roles where Id = ?", id).Scan(&result.ID, &result.RoleName, &result.Status)
	utils.CheckErr(err)
	return &result, nil
}

// FunctionList is yes.
func FunctionList(userid int, page int, size int) (*models.Functionalities, error) {
	var result models.Functionalities
	result.Items = []*models.Functionality{}

	rows, err := db.Query("select Id,Funname,FunType,Controller,CreateTime,UpdatedTime from Functionality order by CreateTime desc limit ?, ?", (page-1)*size, size)
	utils.CheckErr(err)

	for rows.Next() {
		item := models.Functionality{}
		err = rows.Scan(&item.ID, &item.Funname, &item.FunType, &item.Controller, &item.CreateTime, &item.UpdatedTime)
		utils.CheckErr(err)

		result.Items = append(result.Items, &item)
	}

	total := 0
	err = db.QueryRow("select count(*) from Functionality").Scan(&total)
	utils.CheckErr(err)

	pageCount := int(math.Ceil(float64(total) / float64(size)))
	result.PageArgs = models.PageArgs{PageNumber: page, TotalCount: total, PageSize: size, PageCount: pageCount}
	rows.Close()
	return &result, nil
}

// PostFunction is yes
func PostFunction(fun models.Functionality) (bool, error) {
	var result int64
	if fun.ID > 0 {
		res, err := db.Exec("update Functionality set Funname=?,FunType=?,Controller=? where id=?", fun.Funname, fun.FunType, fun.Controller, fun.ID)
		utils.CheckErr(err)
		result, err = res.RowsAffected()
		utils.CheckErr(err)
	} else {
		res, err := db.Exec("insert Functionality set Funname=?,FunType=?,Controller=?", fun.Funname, fun.FunType, fun.Controller)
		utils.CheckErr(err)
		result, err = res.RowsAffected()
		utils.CheckErr(err)
	}

	return result > 0, nil
}

// GetFunction is for article
func GetFunction(id string) (*models.Functionality, error) {
	var result models.Functionality

	err := db.QueryRow("select Id,Funname,FunType,Controller from Functionality where Id = ?", id).Scan(&result.ID, &result.Funname, &result.FunType, &result.Controller)
	utils.CheckErr(err)
	return &result, nil
}

// FunctionGroup is yes.
func FunctionGroup() (map[string][]*models.Functionality, error) {
	result := make(map[string][]*models.Functionality)

	rows, err := db.Query("select Id,Funname,FunType,Controller from Functionality order by FunType desc")
	utils.CheckErr(err)

	defer rows.Close()
	for rows.Next() {
		item := models.Functionality{}
		err = rows.Scan(&item.ID, &item.Funname, &item.FunType, &item.Controller)
		utils.CheckErr(err)

		funlist, ok := result[item.FunType]

		if ok {
			funlist = append(funlist, &item)
			result[item.FunType] = funlist
		} else {
			items := []*models.Functionality{}
			items = append(items, &item)
			result[item.FunType] = items
		}
	}

	return result, nil
}

// MappingPost is yes
func MappingPost(mapping models.RoleFunctionMapping) (bool, error) {
	var sql string
	if mapping.Toggle {
		sql = "delete from RoleFunctionMapping where FunID=? and RoleID=?"
	} else {
		sql = "insert RoleFunctionMapping set FunID=?,RoleID=?"
	}

	res, err := db.Exec(sql, mapping.FunID, mapping.RoleID)
	utils.CheckErr(err)

	result, err := res.RowsAffected()
	utils.CheckErr(err)

	return result > 0, nil
}

// GetRoleFunction is yes.
func GetRoleFunction(roleid int) ([]int, error) {
	result := make([]int, 50)

	rows, err := db.Query("select FunID from RoleFunctionMapping where RoleID=?", roleid)
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		funid := 0
		err = rows.Scan(&funid)
		utils.CheckErr(err)

		result = append(result, funid)
	}

	return result, nil
}

// GetFunctionNames is yes.
func GetFunctionNames(roleid int) ([]string, error) {
	result := make([]string, 50)

	rows, err := db.Query("SELECT Controller FROM Functionality fun join RoleFunctionMapping map on fun.Id = map.FunID where RoleID = ?", roleid)
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		controller := ""
		err = rows.Scan(&controller)
		utils.CheckErr(err)

		result = append(result, controller)
	}

	return result, nil
}
