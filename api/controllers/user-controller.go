package controllers

import (
	"net/http"
	"strconv"

	"github.com/golobby/container/v3"
	"github.com/rafaLino/couple-wishes-api/api/common"
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
		c.SendJSON(w, nil, http.StatusBadRequest)
	}

	c.SendJSON(w, output, http.StatusOK)
}

func (c *UserController) Get(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	id := c.GetParam(r, "id")
	parsedID, err := strconv.ParseInt(id, 10, 64)
	output, err := service.Get(parsedID)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
	}

	c.SendJSON(w, output, http.StatusOK)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	var input entities.UserInput
	err := c.GetContent(&input, r)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
	}

	output, err := service.Create(input)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
	}

	c.SendJSON(w, output, http.StatusCreated)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	id := c.GetParam(r, "id")
	parsedID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	var input entities.UserInput
	err = c.GetContent(&input, r)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	err = service.Update(parsedID, input)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
	}

	c.SendJSON(w, nil, http.StatusOK)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	id := c.GetParam(r, "id")
	parsedID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	err = service.Delete(parsedID)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
	}

	c.SendJSON(w, nil, http.StatusOK)
}

func (c *UserController) CheckUsername(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	username := c.GetParam(r, "username")
	exists, err := service.CheckUsername(username)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	c.SendJSON(w, exists, http.StatusOK)
}

func (c *UserController) CreateCouple(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	user, _ := c.GetUser(r)

	var input entities.UserCreateCoupleInput
	err := c.GetContent(&input, r)

	output, err := service.CreateCouple(user, input.Username)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	c.SendJSON(w, output, http.StatusCreated)
}

func (c *UserController) DeleteCouple(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	coupleId, err := c.GetIntParam(r, "id")

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	err = service.DeleteCouple(coupleId)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	c.SendJSON(w, nil, http.StatusOK)
}

func (c *UserController) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var service ports.IUserService
	container.Resolve(&service)

	user, ok := c.GetUser(r)

	if !ok {
		c.SendJSON(w, nil, http.StatusUnauthorized)
		return
	}

	var input entities.UserUpdatePasswordInput
	err := c.GetContent(&input, r)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	err = service.ChangePassword(user.ID, input.Password)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
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
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	user, err := service.CheckPassword(input.Username, input.Password)

	if err != nil {
		c.SendJSON(w, nil, http.StatusUnauthorized)
		return
	}
	token, err := c.GenerateToken(user)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	c.SendJSON(w, token, http.StatusOK)
}
