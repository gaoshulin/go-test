package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID        uint64    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"type:varchar(30);not null" json:"name" form:"name"`
	Age       int64     `gorm:"type:tinyint(3);not null;default:0"  json:"age"  form:"age"`
	Gender    string    `gorm:"type:varchar(6);not null"  json:"gender" form:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Database *gorm.DB

func init() {
	var database = "root:root@(127.0.0.1:3306)/galen?charset=utf8mb4"
	db, err := gorm.Open("mysql", database)
	if err != nil {
		panic(any(err))
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&UserInfo{})

	Database = db
}

func main() {
	// select
	var user UserInfo
	// 获取第一条记录 按逐渐排序
	Database.First(&user)
	// 获取最后一条记录，按逐渐排序
	Database.Last(&user)
	// 查询所有记录
	Database.Select("id", "name", "age", "gender").Where("id > ?", 1).Find(&user)

	fmt.Println(user)

	// update
	//user.Name = "later"
	//user.Age = 28
	//Database.Save(&user)
	//// 更新
	//Database.Model(&user).Update("name", "alan")
	//// 批量更新
	//Database.Model(&user).Updates(map[string]interface{}{"name":"jinzhu2", "age":30});
	//
	//// delete
	//Database.Delete(&user)
	//// 批量删除
	//Database.Where("name like ?", "later").Delete(UserInfo{})
}
