package repository

import (
	"fmt"
	"form/model"
)

// import (
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

func InsertFormInf(newUser model.User) string {

	DataBase.Table("tb_form_users").Select("Name", "Age", "CreatedAt").Create(FormUser)
	return fmt.Sprintf(`data: {User %s was created successfuly!}`, newUser.Name)
}

func GetFromInf() []model.User {

	DataBase.Find(Users)
	return Users
}

func UpdateFormInf() {

}

func DeleteFormInf() {

}
