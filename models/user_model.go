package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id		    primitive.ObjectID	`json:"id, omitempty"`
	Email		string				`json: "email, omitempty" validate: "required"`
	Name 	    string				`json: "name, omitempty" validate: "required"`
	Age    		int				    `json: "age, omitempty" validate: "required"`
	Gender 	 	string				`json: "gender, omitempty" validate: "required"`
	Hobbies     []string    		`json: "hobbies, omniempty"`
}