package database

import (
	"api/model"
	"fmt"

	"gorm.io/gorm"
)

func UserSeed(db *gorm.DB) error {
	var count int64

	// main.goが実行される度にレコードが生成されないようにする。
	db.Model(&model.User{}).Count(&count)
	if count > 0 {
		return nil
	}
	users := model.User{Name: "Yamada Taro", Email: "mailmail@gmail.com", Password: "password"}

	if err := db.Create(&users).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	return nil
}
