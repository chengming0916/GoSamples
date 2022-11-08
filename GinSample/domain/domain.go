package domain

import (
	"log"
	"strconv"

	"GoSamples/GinSample/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb(cfg *config.Conifg) *gorm.DB {
	var err error
	dataSource := cfg.DbUser + ":" + cfg.DbPassword + "@tcp(" + cfg.DbHost + ":" + strconv.Itoa(int(cfg.DbPort)) + ")/" + cfg.DbName + "?charset=utf8mb4"
	db, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	if err != nil {
		// db.Close()
		log.Fatal("", err)
	}

	// 设置连接池
	// db.DB().SetMaxIdleConns(50)
	// 打开连接
	// db.DB().SetMaxOpenConns(100)

	// db.SingularTable(true)

	return db
}

type User struct {
	Id       int    `gorm:"column:id;paramary_key", json:"id"`
	Name     string `gorm:"column:name", json:"name"`
	Account  string `gorm:"column:account", json:"account"`
	Password string `gorm:"column:password", json:"password"`
}

func FindAll() {

}

func Create(user *User) error {
	if err := db.Create(user).Error; err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
