package ports

import "github.com/rafaLino/couple-wishes-api/entities"

type IUserService interface {
	Create(input entities.UserInput) (*entities.UserOutput, error)
	GetAll() ([]entities.UserOutput, error)
	Get(id int64) (*entities.UserOutput, error)
	Update(id int64, input entities.UserInput) error
	Delete(id int64) error
	CheckUsername(username string) (bool, error)
	CheckPassword(username, password string) (*entities.User, error)
	ChangePassword(id int64, password string) error
	CreateCouple(user entities.User, username string) (*entities.UserOutput, error)
	DeleteCouple(id int64) error
}
