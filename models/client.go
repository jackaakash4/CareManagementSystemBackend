package client

import (
	"time",
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmergencyContact struct{
	Name 		string 	`bson: "name" json:"name"`
	Relationship	string	`bson: "relationship" json: "relationship"`
	Phone		string 	`bson: "phone" json: "phone"`
	Email 		string	`bson: "email" json: "email"`
	IsPrimary	bool	`bson: "is_primary" json: "is_primary"`
}
