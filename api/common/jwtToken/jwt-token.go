package jwtToken

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rafaLino/couple-wishes-api/entities"
	valueObjects "github.com/rafaLino/couple-wishes-api/value-objects"
)

type TokenData struct {
	Token string              `json:"token"`
	User  entities.UserOutput `json:"user"`
}

func GenerateToken(user entities.User) (string, *entities.User, error) {
	key := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"id":        user.ID,
		"name":      user.Name,
		"username":  user.Username.String(),
		"couple_id": user.CoupleID,
		"phone":     user.Phone,
		"exp":       time.Now().Add(time.Hour * 36).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", nil, err
	}

	return tokenString, &user, nil
}

func VerifyToken(tokenString string) (string, *entities.User, error) {
	key := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return "", nil, err
	}
	if !token.Valid {
		return "", nil, errors.New("invalid token")
	}
	claims := token.Claims.(jwt.MapClaims)
	user := &entities.User{
		ID:       int64(claims["id"].(float64)),
		Name:     claims["name"].(string),
		Username: *valueObjects.NewUsername(claims["username"].(string)),
		Phone:    claims["phone"].(string),
		CoupleID: int64(claims["couple_id"].(float64)),
	}
	return token.Raw, user, nil
}
