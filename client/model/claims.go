package model

type Claims struct {
	Exp    float64  `json:"exp"`
	UserID string `json:"userid"`
}