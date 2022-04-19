package interactor

import (
	"github.com/ulascansenturk/go-clean-architecture/domain/model"
	"github.com/ulascansenturk/go-clean-architecture/usecase/presenter"
	"github.com/ulascansenturk/go-clean-architecture/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{
		UserRepository: r,
		UserPresenter:  p,
	}
}

func (us *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}
	return us.UserPresenter.ResponseUsers(u), nil

}
