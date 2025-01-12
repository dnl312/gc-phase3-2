package model

type Claims struct {
	Exp    string  `json:"exp"`
	UserID string `json:"userid"`
}