package usecase

import (
	"context"

	"github.com/DarkSoul94/golang-template/app"
)

// Usecase ...
type usecase struct {
	repo app.Repository
}

// NewUsecase ...
func NewUsecase(repo app.Repository) app.Usecase {
	return &usecase{
		repo: repo,
	}
}

// HelloWorld ...
func (u *usecase) HelloWorld(c context.Context) {
	println("Hello")
}
