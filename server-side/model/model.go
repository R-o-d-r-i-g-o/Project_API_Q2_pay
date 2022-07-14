package model

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Cpf  string `json:"cpf"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
