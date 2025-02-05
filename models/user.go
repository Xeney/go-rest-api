package models

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var UsersBase = make(map[string]User)
