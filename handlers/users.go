package users

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

func getUsers(c *gin.Context) {
	collection := db.Collection("users")
	cursor, err := collection.Find(context.TODO, bson.M{"is_active": true})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch the user's data",
		})
		return
	}
	defer cursor.Close(context.TODO())

	var users []User
	if err = cursor.All(context.TODO(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decode the users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	collection := db.Collection("users")
	var user User

	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}
	var updateData map[string]interface{}

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateData["updated_at"] = time.Now()

	collection := db.Collection("users")
	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{"_id", objectID}.
			bson.M{"$set", updateData},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated sucessfully",
	})
}
