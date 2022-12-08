package controllers

import (
	"net/http"
	// "context"
	// "time"
	"github.com/UdumiziSolomon/Gopher-Backend/models"
	"github.com/UdumiziSolomon/Gopher-Backend/configs"
	
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func CreateUser() gin.HandlerFunc {

	return func(c *gin.Context){
		var user models.User
		
		// validate request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// create new user
		newUser := models.User{
			Id: 	      primitive.NewObjectID(),
			Name:   	  user.Name,
			Age:     	  user.Age,
			Gender: 	  user.Gender,
			Hobbies:      user.Hobbies,
		}

		result, err := userCollection.InsertOne(nil, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusCreated, result)
	}
}

func GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context){
		var user models.User
		userID := c.Param("id")

		objID, _ := primitive.ObjectIDFromHex(userID)

		err := userCollection.FindOne(nil, bson.M{ "id": objID }).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context){
		var users []models.User

		results, err := userCollection.Find(nil, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		// reading from DB (optimally)
		for results.Next(nil) {
			var singleUser models.User
			if err := results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}

			users = append(users, singleUser)
		}

		c.JSON(http.StatusOK, users)

		
	}
}

func DeleteUserByID() gin.HandlerFunc {
	return func(c *gin.Context){
		userID := c.Param("id")

		objID, _ := primitive.ObjectIDFromHex(userID)

		result, err := userCollection.DeleteOne(nil, bson.M{ "id": objID })
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, err.Error())
		}

		c.JSON(http.StatusOK, result)
	}
}

func UpdateUserByID() gin.HandlerFunc {
	return func(c *gin.Context){
		var user models.User
		userID := c.Param("id")

		// convert stringified userID into mongo iD format
		objID, _ := primitive.ObjectIDFromHex(userID)

		// check incoming data for JSON format
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// new data
		dataUpdates := bson.M{
			"name":       user.Name,
			"age": 	      user.Age,
			"gender":     user.Gender,
			"hobbies":    user.Hobbies,
		}

		// updating the DB with new data
		result, err := userCollection.UpdateOne(nil, bson.M{ "id": objID }, bson.M{ "$set": dataUpdates })
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// get updated data of user
		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(nil, bson.M{ "id": objID }).Decode(&updatedUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		c.JSON(http.StatusOK, result)
	}
}