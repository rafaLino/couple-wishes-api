package ports

import "github.com/rafaLino/couple-wishes-api/infra/db"

type IWishRepository interface {
	Get(id int64) (*db.Wish, error)
	GetAll(coupleID int64) ([]db.Wish, error)
	Create(wish *db.Wish) (int64, error)
	Update(wish *db.Wish) error
	Delete(id int64) error
}
