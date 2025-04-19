package jwtToken

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rafaLino/couple-wishes-api/entities"
	valueObjects "github.com/rafaLino/couple-wishes-api/value-objects"
)

type TokenDataOutput struct {
	Token        string              `json:"token"`
	RefreshToken string              `json:"refreshToken,omitempty"`
	User         entities.UserOutput `json:"user,omitempty"`
}

func GenerateToken(user entities.User) (string, error) {
	key := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"id":        user.ID,
		"name":      user.Name,
		"username":  user.Username.String(),
		"couple_id": user.CoupleID,
		"phone":     user.Phone,
		"exp":       time.Now().Add(time.Hour * 36).Unix(),
	}

	token, err := signToken(key, claims)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateRefreshToken(username string, password string) (string, error) {
	key := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"username": username,
		"password": password,
		"exp":      time.Now().Add(time.Hour * 24 * 20).Unix(),
	}

	refreshToken, err := signToken(key, claims)

	return refreshToken, err
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

func VerifyRefreshToken(tokenString string) (*entities.UserInput, error) {
	key := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)

	user := &entities.UserInput{
		Username: (claims["username"].(string)),
		Password: (claims["password"].(string)),
	}
	return user, nil
}

func signToken(key string, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
