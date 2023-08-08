package Controller

import (
	"net/http"
	"taskup/Model"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func ChangePassword(context *gin.Context) {

	var input User

	//parse the token to function getauthenticateid

	token := context.Request.Header.Get("Authorization")

	id, err := Model.GetAuthenticatedID(token)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//add id to user struct

	user := Model.User{
		Password: input.Password,
		Email:    input.Email,
		Name:     input.Name,
		Role:     input.Role,
	}

	_, err = user.Update(id)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User Data changed successfully!"})

}
