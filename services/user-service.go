package services

import (
	"errors"

	"github.com/rafaLino/couple-wishes-api/entities"
	"github.com/rafaLino/couple-wishes-api/ports"
	valueObjects "github.com/rafaLino/couple-wishes-api/value-objects"
)

type UserService struct {
	repository ports.IUserRepository
	ports.IUserService
}

func NewUserService(repository ports.IUserRepository) ports.IUserService {
	return &UserService{repository: repository}
}

func (s *UserService) Create(input entities.UserInput) (*entities.UserOutput, error) {
	user := entities.NewUser(input)

	if !user.Username.IsValid() {
		return nil, errors.New("invalid username")
	}
	id, err := s.repository.Create(&user)

	output := &entities.UserOutput{
		ID:       id,
		Name:     user.Name,
		Username: user.Username.String(),
		Partner:  "",
	}
	return output, err
}

func (s *UserService) GetAll() ([]entities.UserOutput, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	output := entities.MapToUserOutputs(users)

	return output, nil
}

func (s *UserService) Get(id int64) (*entities.UserOutput, error) {
	user, err := s.repository.Get(id)

	if err != nil {
		return nil, err
	}

	partnerName, err := s.repository.GetPartnerUsername(user.CoupleID, id)

	if err != nil {
		return nil, err
	}
	output := entities.MapToUserOutput(*user, partnerName)

	return &output, nil
}

func (s *UserService) Update(id int64, input entities.UserInput) error {
	user := entities.NewUser(input)
	user.ID = id

	return s.repository.Update(&user)
}

func (s *UserService) Delete(id int64) error {
	return s.repository.Delete(id)
}

func (s *UserService) CheckUsername(username string) (bool, error) {
	return s.repository.CheckUsername(username)
}

func (s *UserService) CheckPassword(usernameString string, passwordString string) (*entities.User, error) {
	username := *valueObjects.NewUsername(usernameString)
	password := *valueObjects.NewPassword(passwordString)
	return s.repository.CheckPassword(username, password)
}

func (s *UserService) ChangePassword(id int64, spassword string) error {
	password := *valueObjects.NewPassword(spassword)
	return s.repository.ChangePassword(id, password)
}

func (s *UserService) CreateCouple(user entities.User, usernameString string) (*entities.UserOutput, error) {
	username := *valueObjects.NewUsername(usernameString)

	if !username.IsValid() {
		return nil, errors.New("invalid username")
	}

	partner, err := s.repository.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if user.Username.Equals(partner.Username) {
		return nil, errors.New("user cannot create a couple with themselves")
	}

	if user.CoupleID != 0 || partner.CoupleID != 0 {
		return nil, errors.New("user already has a couple")
	}

	coupleID, err := s.repository.CreateCouple(user.ID, partner.ID)

	output := &entities.UserOutput{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username.String(),
		Partner:  partner.Username.String(),
		CoupleID: coupleID,
	}

	return output, err
}

func (s *UserService) DeleteCouple(coupleId int64) error {
	return s.repository.DeleteCouple(coupleId)
}
