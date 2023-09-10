package service

import (
	"github.com/victorvelo/notes/internal/models"
	"github.com/victorvelo/notes/pkg/repository"
)

type NoteService struct {
	repo repository.Notes
}

func NewNoteService(repo repository.Notes) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) Add(userId int, note models.Note) error {
	return s.repo.Add(userId, note)
}

func (s *NoteService) ListAll(userId int) ([]models.Note, error) {
	return s.repo.ListAll(userId)
}

func (s *NoteService) Update(id int, note *models.Note) error {
	return s.repo.Update(id, note)
}

func (s *NoteService) Delete(id int) error {
	return s.repo.Delete(id)
}
