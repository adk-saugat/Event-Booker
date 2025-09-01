package routes

import (
	"net/http"

	"github.com/event-booker/models"
	"github.com/gin-gonic/gin"
)


func signUp(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnot parse the request data!"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot save user!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}