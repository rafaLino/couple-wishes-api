package ioc

import (
	"github.com/golobby/container/v3"
	"github.com/rafaLino/couple-wishes-api/adapters"
	dbclient "github.com/rafaLino/couple-wishes-api/infra/db-client"
	"github.com/rafaLino/couple-wishes-api/infra/repositories"
	"github.com/rafaLino/couple-wishes-api/ports"
	"github.com/rafaLino/couple-wishes-api/services"
)

type Container struct {
	IContainer
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) RegisterDependencies() error {
	container.Singleton(func() *dbclient.DbContext {
		return dbclient.NewDBContext()
	})
	container.Singleton(func() ports.AIAdapter {
		return adapters.NewGeminiAIAdapter()
	})

	container.Singleton(func() ports.WhatsAppApiAdapter {
		return adapters.NewWhatsAppApiAdapter()
	})

	container.Transient(func(dbContext *dbclient.DbContext) ports.IWishRepository {
		repo, _ := repositories.NewWishRepository(dbContext)
		return repo
	})

	container.Transient(func(dbContext *dbclient.DbContext) ports.IUserRepository {
		repo, _ := repositories.NewUserRepository(dbContext)
		return repo
	})

	container.Transient(func(repository ports.IWishRepository, iaAdapter ports.AIAdapter, userRepository ports.IUserRepository) ports.IWishService {
		return services.NewWishService(repository, iaAdapter, userRepository)
	})

	container.Transient(func(repository ports.IUserRepository) ports.IUserService {
		return services.NewUserService(repository)
	})

	return nil
}
