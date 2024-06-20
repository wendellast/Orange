package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wendellast/Orange/src/controler"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userID", controler.FindUserById)
	r.GET("/getUserByEmail/:userID", controler.FindUserByEmail)
	r.POST("/createUser", controler.CreateUser)
	r.PUT("/updateUser/:userID", controler.UpdateUser)
	r.PUT("/deleteUser/:userID", controler.DeleteUser)

}
