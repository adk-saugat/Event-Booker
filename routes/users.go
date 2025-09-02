package routes

import (
	"net/http"

	"github.com/event-booker/models"
	"github.com/event-booker/utils"
	"github.com/gin-gonic/gin"
)

func login(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnot parse the request data!"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Couldnot authorize!"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot generate token!"} )
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token} )
}


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