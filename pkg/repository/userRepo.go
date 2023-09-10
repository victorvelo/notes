package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/victorvelo/notes/internal/models"
)

var (
	tableUsers = "users"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Add(user *models.User) error {
	query := fmt.Sprintf("INSERT INTO %s (login, password) values ($1, $2)", tableUsers)
	if _, err := r.db.Query(query, user.Login, user.Password); err != nil {
		return fmt.Errorf("addiction is failed: %w", err)
	}
	return nil
}

func (r *userRepo) Get(login, password string) (*models.User, error) {
	var u models.User
	query := fmt.Sprintf("SELECT id, login, password FROM %s where login = $1 AND password = $2", tableUsers)

	if err := r.db.QueryRow(query, login, password).Scan(&u.Id, &u.Login, &u.Password); err != nil {
		return nil, fmt.Errorf("geting is failed %w", err)
	}
	return &u, nil
}

func (r *userRepo) GetUserById(id int) (*models.User, error) {
	var u models.User
	query := fmt.Sprintf("SELECT * FROM %s users WHERE users.id=$1", tableUsers)
	err := r.db.Get(&u, query, id)
	return &u, err
}

// func (r *userRepo) ListAll() ([]user.User, error) {
// 	var allUser []user.User
// 	rows, err := r.db.Query("SELECT * FROM authtable")
// 	if err != nil {
// 		return nil, fmt.Errorf("select all from db is failed %w", err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var u user.User
// 		if err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.Login, &u.Password, &u.Email); err != nil {
// 			return nil, fmt.Errorf("scan is faile %w", err)
// 		}
// 		allUser = append(allUser, u)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("scan is faile %w", err)
// 	}
// 	return allUser, nil
// }

// func (r *userRepo) Delete(id int) error {
// 	_, err := r.db.Exec("DELETE from authtable WHERE id=?", id)

// 	if err != nil {
// 		return fmt.Errorf("deletion is failed %w", err)
// 	}
// 	return nil
// }

// func (r *userRepo) Update(id int, u *user.User) error {
// 	if _, err := r.db.Exec("UPDATE authtable SET name =?, age= ?, login= ?, password= ?, email= ? WHERE id=?", u.Name, u.Age, &u.Login, &u.Password, &u.Email, u.Id); err != nil {
// 		return fmt.Errorf("updation is failed %w", err)
// 	}
// 	return nil
// }

// func (r *userRepo) Get(login, password string) (*user.User, error) {
// 	var u user.User
// 	if err := r.db.QueryRow("SELECT id, name, age, login, password, email FROM authtable where login = ? AND password = ?", login, password).Scan(&u.Id, &u.Name, &u.Age, &u.Login, &u.Password, &u.Email); err != nil {
// 		return nil, fmt.Errorf("geting is failed %w", err)
// 	}
// 	return &u, nil
// }
