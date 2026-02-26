package models

// Modelo de usuario

type User struct {
	ID       int    `bson:"_id,omitempty" json:"id"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"-"`
}
