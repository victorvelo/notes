package models

type Note struct {
	Id          int    `json:"id" db:"id"`
	Description string `json:"description" validate:"required"`
	UserId      int    `json:"user_id" db:"user_id"`
}
