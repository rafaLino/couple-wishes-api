package ports

import "github.com/rafaLino/couple-wishes-api/entities"

type AIAdapter interface {
	Connect() error
	GenerateResponse(url string) (*entities.WishInput, error)
}
