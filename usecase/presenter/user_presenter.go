package presenter

import "github.com/ulascansenturk/go-clean-architecture/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
