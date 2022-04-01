package model

import(
	
	"gorm.io/gorm"
)

type User struct{
	//先继承gorm的模板结构体
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}


//博客的结构体
type Post struct{
	gorm.Model
	Title string `json:"title"`
	Content string `gorm:"type:text" `
	Tag string
}