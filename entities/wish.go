package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rafaLino/couple-wishes-api/infra/db"
)

type WishInput struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Price       string `json:"price"`
	CoupleID    int64  `json:"couple_id"`
	Completed   bool   `json:"completed"`
}

type WishUrlInput struct {
	Url string `json:"url"`
}

type WishOutput struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Completed   bool   `json:"completed"`
}

func MapToWish(wish WishInput) db.Wish {
	return db.Wish{
		Title:       wish.Title,
		Url:         pgtype.Text{String: wish.Url, Valid: true},
		Description: pgtype.Text{String: wish.Description, Valid: true},
		Price:       pgtype.Text{String: wish.Price, Valid: true},
		Completed:   pgtype.Bool{Bool: wish.Completed, Valid: true},
		CoupleID:    pgtype.Int8{Int64: wish.CoupleID, Valid: true},
	}
}

func MapToWishOutput(wish db.Wish) WishOutput {
	return WishOutput{
		Id:          wish.ID,
		Title:       wish.Title,
		Url:         wish.Url.String,
		Description: wish.Description.String,
		Price:       wish.Price.String,
		Completed:   wish.Completed.Bool,
	}
}

func MapToWishOutputs(wishes []db.Wish) []WishOutput {
	output := make([]WishOutput, len(wishes))
	for i, wish := range wishes {
		output[i] = MapToWishOutput(wish)
	}
	return output
}
