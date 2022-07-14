package all

import (
	"form/controller"

	"github.com/gin-gonic/gin"
)

func AvaibleRoutes(r *gin.Engine) {

	form := r.Group("/form")
	{
		form.GET("/", controller.GetUsers)
		form.POST("/", controller.PostUser)
		// form.PATCH("/", controller.UpdateUser)
		// form.DELETE("/", controller.DeleteUser)
	}

}
