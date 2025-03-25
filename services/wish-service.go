package services

import (
	"github.com/rafaLino/couple-wishes-api/entities"
	"github.com/rafaLino/couple-wishes-api/ports"
)

type WishService struct {
	repository ports.IWishRepository
	aiAdapter  ports.AIAdapter
	ports.IWishService
}

func NewWishService(repository ports.IWishRepository, aiAdapter ports.AIAdapter) ports.IWishService {
	return &WishService{repository: repository, aiAdapter: aiAdapter}
}

func (s *WishService) Save(input entities.WishInput) (*entities.WishOutput, error) {
	wish := entities.MapToWish(input)

	wishId, err := s.repository.Create(&wish)
	if err != nil {
		return nil, err
	}
	wish.ID = wishId
	output := entities.MapToWishOutput(wish)
	return &output, err
}

func (s *WishService) GetAll(coupleId int64) ([]entities.WishOutput, error) {
	wishes, err := s.repository.GetAll(coupleId)
	if err != nil {
		return nil, err
	}

	output := entities.MapToWishOutputs(wishes)

	return output, nil
}

func (s *WishService) Get(id int64) (*entities.WishOutput, error) {
	wish, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}

	output := entities.MapToWishOutput(*wish)

	return &output, nil
}

func (s *WishService) Update(id int64, input entities.WishInput) error {
	wish := entities.MapToWish(input)
	wish.ID = id
	return s.repository.Update(&wish)
}

func (s *WishService) Delete(id int64) error {
	return s.repository.Delete(id)
}

func (s *WishService) MaskAsCompleted(id int64) error {
	return s.repository.MaskAsCompleted(id)
}

func (s *WishService) Create(url string, coupleId int64) (*entities.WishOutput, error) {
	input, err := s.aiAdapter.GenerateResponse(url)

	if err != nil {
		return nil, err
	}

	input.CoupleID = coupleId

	wish := entities.MapToWish(*input)

	wishId, err := s.repository.Create(&wish)

	wish.ID = wishId
	output := entities.MapToWishOutput(wish)
	return &output, err
}
