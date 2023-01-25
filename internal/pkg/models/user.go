package models

type Users []*User

type User struct {
	BaseModel

	FirstName string
	LastName  string
	Email     string
}
