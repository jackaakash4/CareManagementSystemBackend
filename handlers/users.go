package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	fmt.Println("vim-go")
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Hash password
	user.ID = primitive.NewObjectID()
	user.IsActive = true
	user.CreatedAt = time.Now()
	user.UpdateAt = time.Now()

	collection := db.Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternerServerError, gin.H{
			"error": "Failed to insert the user",
		})
		return
	}
	c.JSON(http.StatusCreated, user)

}
