package ports

import "github.com/rafaLino/couple-wishes-api/entities"

type IWishService interface {
	Save(input entities.WishInput) (*entities.WishOutput, error)
	GetAll(coupleId int64) ([]entities.WishOutput, error)
	Get(id int64) (*entities.WishOutput, error)
	Update(id int64, input entities.WishInput) error
	Delete(id int64) error
	Create(url string, coupleId int64) (*entities.WishOutput, error)
	CreateFromWhatsApp(req entities.WhatsAppMessage) error
}
