package models

type Users_P3W2 struct {            
    Username string `json:"username" db:"username"` // Username
    Password string `json:"password" db:"password"` // Hashed password
}