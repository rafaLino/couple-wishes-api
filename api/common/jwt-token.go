package common

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rafaLino/couple-wishes-api/entities"
	valueObjects "github.com/rafaLino/couple-wishes-api/value-objects"
)

type JwtToken struct {
	Token string        `json:"token"`
	User  entities.User `json:"user"`
}

func NewJwtToken() *JwtToken {
	return &JwtToken{}

}

func (t *JwtToken) GenerateToken(user entities.User) (*JwtToken, error) {
	key := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"id":        user.ID,
		"name":      user.Name,
		"username":  user.Username.String(),
		"couple_id": user.CoupleID,
		"exp":       time.Now().Add(time.Hour * 36).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		return nil, err
	}
	t.Token = tokenString
	t.User = user
	return t, nil
}

func (t *JwtToken) VerifyToken(tokenString string) error {
	key := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invalid token")
	}
	claims := token.Claims.(jwt.MapClaims)
	t.Token = token.Raw
	t.User = entities.User{
		ID:       int64(claims["id"].(float64)),
		Name:     claims["name"].(string),
		Username: *valueObjects.NewUsername(claims["username"].(string)),
		CoupleID: int64(claims["couple_id"].(float64)),
	}
	return nil
}
