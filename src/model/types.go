package model

import "html/template"

type LinkCat struct {
	ID      string `json:"id"`
	CatName string `json:"cat_name"`
}

type Article struct {
	Id         string        `json:"id"`
	Title      string        `json:"title"`
	Icon       string        `json:"icon"`
	Brief      string        `json:"brief"`
	Body       template.HTML `json:"body"`
	UserID     int           `json:"user_id"`
	NickName   string        `json:"nick_name"`
	Origin     string        `json:"origin"`
	CreateTime string        `json:"create_time"`
	Visited    int           `json:"visited"`
	Catalog    string        `json:"catalog"`
}

type ArticleList struct {
	List  []*Article `json:"list"`
	Total int64      `json:"total"`
}

type AccountList struct {
	List  []*Account `json:"list"`
	Total int64      `json:"total"`
}

type LinkList struct {
	List  []*Link `json:"list"`
	Total int64   `json:"total"`
}

type CusLinkList struct {
	List  []*CusLink `json:"list"`
	Total int64      `json:"total"`
}

type CusLink struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Url        string `json:"url"`
	Status     int    `json:"status"`
	LinkType   string `json:"link_type"`
	Catid      string `json:"cat_id"`
	Adddate    string `json:"add_date"`
	Updatedate string `json:"update_date"`
	Visited    int    `json:"visited"`
}

type Link struct {
	Id         string        `json:"id"`
	Title      string        `json:"title"`
	Icon       string        `json:"icon"`
	Catid      string        `json:"cat_id"`
	LinkType   string        `json:"link_type"`
	Visited    int           `json:"visited"`
	Brief      template.HTML `json:"brief"`
	Url        string        `json:"url"`
	CreateTime string        `json:"create_time"`
}

type LinkGroup struct {
	Active *OneGroup
	Rests  []*OneGroup
}

type Index struct {
	Common   []*Link
	Info     []*Link
	CusLinks []*CusLink
	Articles []*Article
	Other    *LinkGroup
	Backend  *LinkGroup
	Frontend *LinkGroup
	Hot      *LinkGroup
}

type Account struct {
	ID            int    `json:"id"`
	UserId        string `json:"userid"`
	Password      string `json:"password"`
	UserName      string `json:"user_name"`
	RegDate       string `json:"reg_date"`
	LastLoginDate string `json:"last_login_date"`
	Status        int    `json:"status"`
}

type OneGroup struct {
	CatId string
	Name  string
	Items []*Link
	Size  int
}

type Search struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	CatId string `json:"cat_id"`
	Url   string `json:"url"`
}

type OneGroups []*OneGroup

func (s OneGroups) Len() int           { return len(s) }
func (s OneGroups) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s OneGroups) Less(i, j int) bool { return s[i].Size > s[j].Size }

type LinkItems struct{ Items []*Link }
type ArticleItems struct{ Items []*Article }
type CusLinkItems struct{ Items []*CusLink }
