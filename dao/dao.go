package dao

import (
	
	"ginblog/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//声明全局变量
type manager struct {
	db *gorm.DB
}

type Manager interface{
	AddUser(user * model.User)
	Login(username string) model.User


	//博客操作
	AddPost(post *model.Post)
	GetAllPost() []model.Post
	GetPost(pid int) model.Post

}

var Mgr Manager


func init(){
	//使用gorm连接mysql数据库
	dsn := "root:123456hh@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil{
		log.Fatal("Failed to init db:", err)
	}

	Mgr = &manager{db:db}
	//自动迁移，会自己创建表，如果没有的话
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})

}

func (mgr *manager)AddUser(user *model.User){
	mgr.db.Create(user)
}

func (mgr *manager)Login(username string) model.User{
	var user model.User
	mgr.db.Where("username = ?", username).First(&user)
	return user
}

func (mgr *manager)AddPost(post *model.Post){
	mgr.db.Create(post)
}

func (mgr *manager)GetAllPost() []model.Post{
	var posts []model.Post
	mgr.db.Find(&posts)
	return posts

}

func (mgr *manager)GetPost(pid int) model.Post{
	var post model.Post
	mgr.db.Where("id = ?", pid).First(&post)
	return post
}
	