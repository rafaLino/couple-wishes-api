package common

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rafaLino/couple-wishes-api/entities"
)

type Controller struct {
}

func (c *Controller) SendJSON(w http.ResponseWriter, v interface{}, code int) {
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

func (c *Controller) GetContent(v interface{}, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	return err
}

func (c *Controller) GetParam(r *http.Request, param string) string {
	return mux.Vars(r)[param]
}

func (c *Controller) GetIntParam(r *http.Request, param string) (int64, error) {
	value := c.GetParam(r, param)
	if value == "" {
		return 0, errors.New("Parameter not found")
	}

	parsedValue, err := strconv.ParseInt(value, 10, 64)
	return parsedValue, err
}

func (c *Controller) HandleError(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}

	msg := map[string]string{
		"message": "An error occurred!",
	}

	c.SendJSON(w, &msg, http.StatusInternalServerError)
	return true
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
		jwtToken := NewJwtToken()
		err := jwtToken.VerifyToken(tokenString)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user", jwtToken.User)
		f(w, r.WithContext(ctx))
	})
}

func (c *Controller) GetUser(r *http.Request) (entities.User, bool) {
	user, ok := r.Context().Value("user").(entities.User)
	return user, ok
}

func (c *Controller) GenerateToken(user *entities.User) (*JwtToken, error) {
	jwtToken := NewJwtToken()
	return jwtToken.GenerateToken(*user)
}
