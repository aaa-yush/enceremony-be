package models

import "github.com/golang-jwt/jwt/v5"

type JwtPayload struct {
	jwt.RegisteredClaims
	Uid          string `json:"uid"`
	FirstName    string `json:"fn"`
	LastName     string `json:"ln"`
	Email        string `json:"email"`
	UserId       string `json:"user_id"`
	GoogleUserId string `json:"google_user_id"`
	DOB          string `json:"dob"`
}
