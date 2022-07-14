package repository

import (
	"fmt"
	"form/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DataBase *gorm.DB
var FormUser *model.User
var Users []model.User

func ConnectToDataBase() {

	var con model.Config
	con.GetConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%smb4&parseTime=True&loc=Local",

		con.Username, con.Password, con.Host,
		con.Port, con.Name, con.Charset,
	)

	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {

		panic("failed to connect database!")
	} else {

		PrepareTables(db)
	}
}

func PrepareTables(db *gorm.DB) {

	db.Table("tb_form_users").Create(FormUser)

	DataBase = db
}
