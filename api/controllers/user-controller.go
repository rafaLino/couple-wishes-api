package controllers

import (
	"net/http"
	"strconv"

	"github.com/golobby/container/v3"
	"github.com/rafaLino/couple-wishes-api/api/common"
	"github.com/rafaLino/couple-wishes-api/api/common/jwtToken"
	"github.com/rafaLino/couple-wishes-api/api/common/sessions"
	"github.com/rafaLino/couple-wishes-api/entities"
	"github.com/rafaLino/couple-wishes-api/ports"
)

type UserController struct {
	common.Controller
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)
	output, err := service.GetAll()

	if err != nil {
		c.SendError(nil, http.StatusBadRequest, w)
		return
	}

	c.SendJSON(w, output, http.StatusOK)
}

func (c *UserController) Get(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	id, _ := c.GetIntParam(r, "id")
	output, err := service.Get(id)

	if err != nil {
		c.SendError(nil, http.StatusBadRequest, w)
		return
	}

	c.SendJSON(w, output, http.StatusOK)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	var input entities.UserInput
	err := c.GetContent(&input, r)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	output, err := service.Create(input)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	c.SendJSON(w, output, http.StatusCreated)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	id := c.GetParam(r, "id")
	parsedID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	var input entities.UserInput
	err = c.GetContent(&input, r)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	err = service.Update(parsedID, input)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	c.SendJSON(w, nil, http.StatusOK)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	id, err := c.GetIntParam(r, "id")

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	err = service.Delete(id)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	c.SendJSON(w, nil, http.StatusOK)
}

func (c *UserController) CheckUsername(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	username := c.GetParam(r, "username")
	exists, err := service.CheckUsername(username)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	c.SendJSON(w, exists, http.StatusOK)
}

func (c *UserController) CreateCouple(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	user, _ := c.GetUser(r)

	var input entities.UserCreateCoupleInput
	c.GetContent(&input, r)

	output, err := service.CreateCouple(*user, input.Username)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	c.SendJSON(w, output, http.StatusCreated)
}

func (c *UserController) DeleteCouple(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	coupleId, err := c.GetIntParam(r, "id")

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	err = service.DeleteCouple(coupleId)

	if err != nil {
		c.SendError(nil, http.StatusBadRequest, w)
		return
	}

	c.SendJSON(w, nil, http.StatusOK)
}

func (c *UserController) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	user, ok := c.GetUser(r)

	if !ok {
		c.SendError(nil, http.StatusUnauthorized, w)
		return
	}

	var input entities.UserUpdatePasswordInput
	err := c.GetContent(&input, r)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	err = service.ChangePassword(user.ID, input.Password)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	c.SendJSON(w, nil, http.StatusOK)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	var input entities.UserLoginInput
	err := c.GetContent(&input, r)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	user, err := service.CheckPassword(input.Username, input.Password)

	if err != nil {
		c.SendError(nil, http.StatusUnauthorized, w)
		return
	}
	token, err := jwtToken.GenerateToken(*user)

	if err != nil {
		c.SendError(err, http.StatusInternalServerError, w)
	}

	refreshToken, err := jwtToken.GenerateRefreshToken(input.Username, input.Password)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	session, err := sessions.Get(r)

	if err != nil {
		c.SendError(err, http.StatusInternalServerError, w)
		return
	}

	session.Values["refreshToken"] = refreshToken
	err = session.Save(r, w)

	if err != nil {
		c.SendError(err, http.StatusInternalServerError, w)
		return
	}

	userOutput := entities.MapToUserOutput(*user, "")
	res := jwtToken.TokenDataOutput{
		Token:        token,
		RefreshToken: refreshToken,
		User:         userOutput,
	}

	c.SendJSON(w, res, http.StatusOK)
}

func (c *UserController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Get(r)

	if err != nil {
		c.SendError(err, http.StatusInternalServerError, w)
		return
	}

	if refreshToken, ok := session.Values["refreshToken"].(string); !ok || refreshToken == "" {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	refreshToken := session.Values["refreshToken"].(string)

	input, err := jwtToken.VerifyRefreshToken(refreshToken)

	if err != nil {
		c.SendError(err, http.StatusUnauthorized, w)
		return
	}

	var service ports.IUserService
	container.Resolve(&service)

	user, err := service.CheckPassword(input.Username, input.Password)

	if err != nil {
		c.SendError(nil, http.StatusUnauthorized, w)
		return
	}

	token, err := jwtToken.GenerateToken(*user)

	if err != nil {
		c.SendError(err, http.StatusBadRequest, w)
		return
	}

	userOutput := entities.MapToUserOutput(*user, "")
	response := &jwtToken.TokenDataOutput{
		Token: token,
		User:  userOutput,
	}

	c.SendJSON(w, response, http.StatusOK)
}
