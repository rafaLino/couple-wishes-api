package ports

import "github.com/rafaLino/couple-wishes-api/entities"

type AIAdapter interface {
	Connect() error
	GenerateResponse(text string) (*entities.WishInput, error)
}
