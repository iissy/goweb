package ado

import (
	"database/sql"

	"iissy.com/models"
	"iissy.com/utils"
)

// NewsIndex is yes.
func NewsIndex() (*models.Pictures, error) {
	db, err := sql.Open("mysql", utils.SQLDB)
	utils.CheckErr(err)
	list, err := db.Query("select ID,Subject,Picture from Article where PostType=? order by AddDate desc limit ?", "新闻", 30)
	utils.CheckErr(err)

	pictures := models.Pictures{}
	pictures.Items = []*models.Picture{}
	for list.Next() {
		picture := models.Picture{}
		err = list.Scan(&picture.ID, &picture.Subject, &picture.URL)
		utils.CheckErr(err)

		pictures.Items = append(pictures.Items, &picture)
	}

	list.Close()
	return &pictures, nil
}

// TechIndex is yes.
func TechIndex() (*models.Course, error) {
	db, err := sql.Open("mysql", utils.SQLDB)
	utils.CheckErr(err)
	list, err := db.Query("select ID,Subject,Picture,Description from Article order by AddDate desc limit ?", 30)
	utils.CheckErr(err)
	course := models.Course{}
	course.ArticleItems = []*models.Article{}
	for list.Next() {
		var article models.Article
		err = list.Scan(&article.ID, &article.Subject, &article.Picture, &article.Description)
		utils.CheckErr(err)

		course.ArticleItems = append(course.ArticleItems, &article)
	}

	list.Close()
	return &course, nil
}
