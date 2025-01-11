package models

type User struct {            
    Username string `json:"username" db:"username"` // Username
    Password string `json:"password" db:"password"` // Hashed password
}