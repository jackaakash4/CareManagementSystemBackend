package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
    ID primitive.ObjectID       `bson: "_id, omitempty" json:"id"`
    Email string                `bson: "email" json: "email" binding: "required, email"`
    Password string             `bson: "password" json: "-"`
    Role string                 `bson: "role" json: "role" binding:"required"`
    IsActive bool               `bson: "is_active" json: "is_active"`
    CreatedAt time.Time         `bson: "created_at" json: "created_at"`
    UpdatedAt time.Time         `bson: "updated_at" json: "updated_at"`

    //Profile references
    AdminID *primitive.ObjectID     `bson: "admin_id, omitempty" json: "admin_id, omitempty"`
    StaffID *primitive.ObjectID     `bson: "staff_id, omitempty" json: "staff_id, omitempty"`
    ClientID *primitive.ObjectID    `bson: "client_id, omitempty" json: "client_id, omitempty"`
}

