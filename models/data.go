// Package models handle data process
package models

// User defines user information
type User struct {
	Username string
	Password string
}

// UserOnline record users online
var UserOnline User
