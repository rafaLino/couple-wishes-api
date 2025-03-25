package entities

import (
	"github.com/rafaLino/couple-wishes-api/infrastructure/db"
	valueObjects "github.com/rafaLino/couple-wishes-api/value-objects"
)

type User struct {
	ID       int64                 `json:"id"`
	Name     string                `json:"name"`
	Username valueObjects.Username `json:"username"`
	Password valueObjects.Password `json:"password,omitempty"`
	CoupleID int64                 `json:"couple_id"`
}

type UserInput struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdatePasswordInput struct {
	Password string `json:"password"`
}

type UserCreateCoupleInput struct {
	Username string `json:"username"`
}

type UserOutput struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Partner  string `json:"partner"`
	CoupleID int64  `json:"couple_id"`
}

func MapToUser(user db.GetUsersRow) User {
	return User{
		ID:       user.ID,
		Name:     user.Name,
		Username: *valueObjects.NewUsername(user.Username),
		CoupleID: user.CoupleID.Int64,
	}
}

func MapToUsers(users []db.GetUsersRow) []User {
	var mappedUsers []User
	for _, user := range users {
		mappedUsers = append(mappedUsers, MapToUser(user))
	}

	return mappedUsers
}

func MapGetUserRowToUser(user db.GetUserRow) User {
	return User{
		ID:       user.ID,
		Name:     user.Name,
		Username: *valueObjects.NewUsername(user.Username),
		CoupleID: user.CoupleID.Int64,
	}
}

func NewUser(input UserInput) User {
	return User{
		Name:     input.Name,
		Username: *valueObjects.NewUsername(input.Username),
		Password: *valueObjects.NewPassword(input.Password),
	}
}

func MapToUserOutput(user User, partner string) UserOutput {
	return UserOutput{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username.String(),
		Partner:  partner,
		CoupleID: user.CoupleID,
	}
}

func MapToUserOutputs(users []User) []UserOutput {
	var mappedUsers []UserOutput
	for _, user := range users {
		mappedUsers = append(mappedUsers, MapToUserOutput(user, ""))
	}

	return mappedUsers
}
