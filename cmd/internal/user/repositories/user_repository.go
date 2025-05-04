package repositories

import (
	"database/sql"
	"errors"
	"tenant-management/cmd/internal/user/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// Save implements the UserRepository interface
func (r *userRepository) Save(user domain.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (id, name, email, password, role, created_at) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID, user.Name, user.Email, user.PasswordHash, user.Role, user.CreatedAt,
	)
	return err
}

// FindByEmail implements the UserRepository interface
func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id, name, email, password, role, created_at FROM users WHERE email=$1", email)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
