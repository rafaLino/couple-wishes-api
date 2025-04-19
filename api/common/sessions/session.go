package sessions

import (
	"net/http"
	"os"

	gsessions "github.com/gorilla/sessions"
)

var store *gsessions.CookieStore

func Start() {
	store = gsessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
}
func Get(req *http.Request) (*gsessions.Session, error) {
	return store.Get(req, "session")
}

func GetNamed(req *http.Request, name string) (*gsessions.Session, error) {
	return store.Get(req, name)
}
