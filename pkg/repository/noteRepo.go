package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/victorvelo/notes/internal/models"
)

var (
	noteUsers = "note"
)

type noteRepo struct {
	db *sqlx.DB
}

func NewNoteRepo(db *sqlx.DB) *noteRepo {
	return &noteRepo{db: db}
}

func (r *noteRepo) Add(userId int, note models.Note) error {
	query := fmt.Sprintf("INSERT INTO %s (description, user_Id) values ($1, $2)", noteUsers)
	if _, err := r.db.Query(query, note.Description, userId); err != nil {
		return fmt.Errorf("addiction is failed: %w", err)
	}
	return nil
}

func (r *noteRepo) ListAll(userId int) ([]models.Note, error) {
	var allNotes []models.Note

	query := fmt.Sprintf("SELECT * FROM %s where user_Id = $1", noteUsers)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("select all from db is failed %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var n models.Note
		if err := rows.Scan(&n.Id, &n.Description, &n.UserId); err != nil {
			return nil, fmt.Errorf("scan is faile %w", err)
		}
		allNotes = append(allNotes, n)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("scan is faile %w", err)
	}
	return allNotes, nil
}

func (r *noteRepo) Delete(id int) error {
	query := fmt.Sprintf("DELETE from %s WHERE id=$1", noteUsers)
	_, err := r.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("deletion is failed %w", err)
	}
	return nil
}

func (r *noteRepo) Update(id int, note *models.Note) error {
	query := fmt.Sprintf("UPDATE %s SET description= $1 WHERE id=$2", noteUsers)
	if _, err := r.db.Exec(query, note.Description, note.Id); err != nil {
		return fmt.Errorf("updation is failed %w", err)
	}
	return nil
}
