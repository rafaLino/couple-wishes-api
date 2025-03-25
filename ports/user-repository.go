package ports

import (
	"github.com/rafaLino/couple-wishes-api/entities"
	valueObjects "github.com/rafaLino/couple-wishes-api/value-objects"
)

type IUserRepository interface {
	GetAll() ([]entities.User, error)
	Get(id int64) (*entities.User, error)
	CheckUsername(username string) (bool, error)
	CheckPassword(username valueObjects.Username, password valueObjects.Password) (*entities.User, error)
	Create(user *entities.User) (int64, error)
	Update(user *entities.User) error
	ChangePassword(id int64, password valueObjects.Password) error
	Delete(id int64) error
	GetPartnerUsername(id int64, userId int64) (string, error)
	CreateCouple(userId, partnerId int64) (int64, error)
	DeleteCouple(coupleId int64) error
	GetByUsername(username valueObjects.Username) (*entities.User, error)
}
