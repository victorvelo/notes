package service

import (
	"github.com/victorvelo/notes/internal/models"
	"github.com/victorvelo/notes/pkg/repository"
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	Add(user models.User) error
	Get(login, password string) (*models.User, error)
	CreateToken(login, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Notes interface {
	Add(userId int, note models.Note) error
	ListAll(userId int) ([]models.Note, error)
	Delete(id int) error
	Update(id int, note *models.Note) error
}

type Service struct {
	Authorization
	Notes
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewUserService(repository.Authorization),
		Notes:         NewNoteService(repository.Notes),
	}
}
