package controller

import (
	model "form/model"
	"form/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paemuri/brdoc"
)

// ----------------------------------------< Requests >---------------------------------------- \\
var newUser model.User

// var users = []model.User{
// 	{Cpf: "283.639.620-23", Name: "John Coltrane", Age: 56},
// 	{Cpf: "507.255.470-81", Name: "Gerry Mulligan", Age: 17},
// 	{Cpf: "461.178.870-94", Name: "Sarah Vaughan", Age: 39},
// }

func GetUsers(c *gin.Context) {

	users := repository.GetFromInf()
	c.IndentedJSON(http.StatusOK, users) // 200
}

func PostUser(c *gin.Context) {

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "wrong data inserted") // 406
		return
	}

	if brdoc.IsCPF(newUser.Cpf) && CPFisApproved(newUser) {
		// users = append(users, newUser)
		menssage := repository.InsertFormInf(newUser)
		c.IndentedJSON(http.StatusCreated, menssage) // 201

	} else {
		c.IndentedJSON(http.StatusNotAcceptable, newUser) // 406

	}
}

// func UpdateUser(c *gin.Context) {

// 	c.IndentedJSON(http.StatusOK, users)

// }

// func DeleteUser(c *gin.Context) {

// 	c.IndentedJSON(http.StatusOK, users)

// }

// ----------------------------------------< Aux Func's >---------------------------------------- \\
func CPFisApproved(u model.User) (boolean bool) {
	const NUM_OF_CHARCTER = 14
	var cont int

	for _, value := range users {
		if value.Cpf == u.Cpf {
			cont++
		}
	}

	if len(u.Cpf) == NUM_OF_CHARCTER && cont == 0 {
		return true

	} else {
		return false
	}
}
