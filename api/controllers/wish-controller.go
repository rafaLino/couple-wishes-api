package controllers

import (
	"net/http"
	"strconv"

	"github.com/golobby/container/v3"
	"github.com/rafaLino/couple-wishes-api/api/common"
	"github.com/rafaLino/couple-wishes-api/entities"
	"github.com/rafaLino/couple-wishes-api/ports"
)

type WishController struct {
	common.Controller
}

func NewWishController() *WishController {
	return &WishController{}
}

func (c *WishController) GetAll(w http.ResponseWriter, r *http.Request) {
	var service ports.IWishService
	container.Resolve(&service)

	user, ok := c.GetUser(r)

	if !ok {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}
	output, err := service.GetAll(user.CoupleID)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
	}

	c.SendJSON(w, output, http.StatusOK)
}

func (c *WishController) Get(w http.ResponseWriter, r *http.Request) {
	var service ports.IWishService
	container.Resolve(&service)

	id := c.GetParam(r, "id")
	parsedID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}
	output, err := service.Get(parsedID)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	c.SendJSON(w, output, http.StatusOK)
}

func (c *WishController) Save(w http.ResponseWriter, r *http.Request) {
	var service ports.IWishService
	container.Resolve(&service)

	user, _ := c.GetUser(r)
	var input entities.WishInput
	err := c.GetContent(&input, r)
	input.CoupleID = user.CoupleID

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	output, err := service.Save(input)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	c.SendJSON(w, output, http.StatusCreated)
}

func (c *WishController) Update(w http.ResponseWriter, r *http.Request) {
	var service ports.IWishService
	container.Resolve(&service)

	id, err := c.GetIntParam(r, "id")

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}
	user, _ := c.GetUser(r)
	var input entities.WishInput
	err = c.GetContent(&input, r)
	input.CoupleID = user.CoupleID

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	err = service.Update(id, input)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	c.SendJSON(w, nil, http.StatusOK)
}

func (c *WishController) Delete(w http.ResponseWriter, r *http.Request) {
	var service ports.IWishService
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
		return
	}

	c.SendJSON(w, nil, http.StatusOK)
}

func (c *WishController) Create(w http.ResponseWriter, r *http.Request) {
	var service ports.IWishService
	container.Resolve(&service)

	user, _ := c.GetUser(r)
	var input entities.WishUrlInput
	err := c.GetContent(&input, r)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	output, err := service.Create(input.Url, user.CoupleID)

	if err != nil {
		c.SendJSON(w, nil, http.StatusBadRequest)
		return
	}

	c.SendJSON(w, output, http.StatusCreated)
}
