package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golobby/container/v3"
	"github.com/rafaLino/couple-wishes-api/api/common"
	"github.com/rafaLino/couple-wishes-api/api/controllers"
	"github.com/rafaLino/couple-wishes-api/api/ioc"
	dbclient "github.com/rafaLino/couple-wishes-api/infrastructure/db-client"
	"github.com/rafaLino/couple-wishes-api/ports"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type App struct {
	router *mux.Router
}

func NewApp() *App {
	return &App{}
}

func (a *App) Initialize() *App {
	godotenv.Load()
	ioc.NewContainer().RegisterDependencies()
	initDependencies()
	return a
}

func (a *App) StartupDatabase() *App {
	var dbContext *dbclient.DbContext
	container.Resolve(&dbContext)
	dbContext.Connect()
	return a
}

func (a *App) ConfigEndpoints() *App {
	a.router = mux.NewRouter()
	s := a.router.PathPrefix("/api").Subrouter()

	for _, bundle := range initBundles() {
		for _, route := range bundle.GetRoutes() {
			s.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	http.Handle("/", a.router)
	return a
}

func (a *App) Run() *App {
	port := os.Getenv("APP_PORT")

	if port == "" {
		port = ":3201"
	}

	originString := os.Getenv("ORIGIN_ALLOWED")

	originsAllowed := strings.Split(originString, ",")
	c := cors.New(cors.Options{
		AllowedOrigins:   originsAllowed,
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(a.router)

	fmt.Printf("couple-wishes-api Running on Port: %s\n", port)

	log.Fatal(http.ListenAndServe(port, handler))

	return a
}

func initBundles() []common.Bundle {
	return []common.Bundle{
		controllers.NewWishRouter(),
		controllers.NewUserRouter(),
	}
}

func initDependencies() {
	var aiAdapter ports.AIAdapter
	container.Resolve(&aiAdapter)
	aiAdapter.Connect()
}
