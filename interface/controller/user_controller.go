package controller

import (
	"net/http"

	"github.com/ulascansenturk/go-clean-architecture/domain/model"

	"github.com/ulascansenturk/go-clean-architecture/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{
		userInteractor: us,
	}
}

func (uc *userController) GetUsers(c Context) error {
	var u []*model.User
	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)

}
