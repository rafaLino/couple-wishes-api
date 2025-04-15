package common

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rafaLino/couple-wishes-api/api/common/jwtToken"
	"github.com/rafaLino/couple-wishes-api/entities"
)

type Controller struct {
}

func (c *Controller) SendJSON(w http.ResponseWriter, v any, code int) {
	w.Header().Set("Content-Type", "application/json")

	result := &Result{
		Data: v,
	}
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(result)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "Internal server error"})
	}
}

func (c *Controller) SendError(err error, code int, w http.ResponseWriter) {
	var content string
	if err != nil {
		content = err.Error()
	} else {
		content = "Something went wrong!"
	}
	msg := map[string]string{
		"message": content,
	}

	c.SendJSON(w, msg, code)
}

func (c *Controller) GetContent(v interface{}, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	return err
}

func (c *Controller) GetParam(r *http.Request, param string) string {
	return mux.Vars(r)[param]
}

func (c *Controller) GetQuery(r *http.Request, param string) string {
	return r.URL.Query().Get(param)
}

func (c *Controller) GetIntParam(r *http.Request, param string) (int64, error) {
	value := c.GetParam(r, param)
	if value == "" {
		return 0, errors.New("parameter not found")
	}

	parsedValue, err := strconv.ParseInt(value, 10, 64)
	return parsedValue, err
}

func (c *Controller) ProtectedHandler(f func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString = string(tokenString[len("Bearer "):])
		_, user, err := jwtToken.VerifyToken(tokenString)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		f(w, r.WithContext(ctx))
	})
}

func (c *Controller) GetUser(r *http.Request) (*entities.User, bool) {
	user, ok := r.Context().Value("user").(*entities.User)
	return user, ok
}

func (c *Controller) GenerateToken(user *entities.User) (string, *entities.User, error) {
	return jwtToken.GenerateToken(*user)
}
