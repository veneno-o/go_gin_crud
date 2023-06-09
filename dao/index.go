package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const username, password, host, port, Dbname = "root", "123456", "127.0.0.1", 3306, "todotask"

var DB *gorm.DB
var err error

func Connect() {
	//1.连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("数据库连接失败")
		panic(err)
	}

	//2.迁移表结构
	DB.Migrator().AutoMigrate(&Task{})
}
