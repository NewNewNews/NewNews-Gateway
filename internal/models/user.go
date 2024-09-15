package models

type User struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	HashedPassword string `json:"-"`
	IsAdmin        bool   `json:"is_admin"`
}
