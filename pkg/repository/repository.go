package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/victorvelo/notes/internal/models"
)

type Authorization interface {
	Add(user *models.User) error
	Get(login, password string) (*models.User, error)
}

type Notes interface {
	Add(userId int, note models.Note) error
	ListAll(userId int) ([]models.Note, error)
	Update(id int, note *models.Note) error
	Delete(id int) error
}

type Repository struct {
	Authorization
	Notes
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewUserRepo(db),
		Notes:         NewNoteRepo(db),
	}
}
