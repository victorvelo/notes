package service

import "github.com/victorvelo/notes/pkg/repository"

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
