package repositories

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rafaLino/couple-wishes-api/infra/db"
	dbclient "github.com/rafaLino/couple-wishes-api/infra/db-client"
	"github.com/rafaLino/couple-wishes-api/ports"
)

type WishRepository struct {
	context *dbclient.DbContext
	ports.IWishRepository
}

func NewWishRepository(c *dbclient.DbContext) (ports.IWishRepository, error) {
	return &WishRepository{context: c}, nil
}

func (r *WishRepository) Get(id int64) (*db.Wish, error) {
	client, err := r.context.GetClient()
	wish, err := client.GetWish(r.context.GetContext(), id)

	return &wish, err
}

func (r *WishRepository) GetAll(coupleID int64) ([]db.Wish, error) {
	client, err := r.context.GetClient()

	wishes, err := client.GetWishes(r.context.GetContext(), pgtype.Int8{Int64: coupleID, Valid: true})

	return wishes, err
}

func (r *WishRepository) Create(wish *db.Wish) (int64, error) {
	client, err := r.context.GetClient()
	wishId, err := client.CreateWish(r.context.GetContext(), db.CreateWishParams{
		Title:       wish.Title,
		Url:         wish.Url,
		Description: wish.Description,
		Price:       wish.Price,
		Completed:   wish.Completed,
		CoupleID:    wish.CoupleID,
	})

	return wishId, err
}

func (r *WishRepository) Update(wish *db.Wish) error {
	client, err := r.context.GetClient()
	client.UpdateWish(r.context.GetContext(), db.UpdateWishParams{
		ID:          wish.ID,
		Title:       wish.Title,
		Url:         wish.Url,
		Description: wish.Description,
		Price:       wish.Price,
		Completed:   wish.Completed,
	})

	return err
}

func (r *WishRepository) Delete(id int64) error {
	client, err := r.context.GetClient()
	client.DeleteWish(r.context.GetContext(), id)

	return err
}
