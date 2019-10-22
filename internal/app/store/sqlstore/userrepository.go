package sqlstore

import (
	"database/sql"

	"github.com/ferikl/go-api/internal/app/model"
	"github.com/ferikl/go-api/internal/app/store"
)

// UserRepository ..
type UserRepository struct {
	store *Store
}

// Create ..
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	stmt, err := r.store.db.Prepare("INSERT INTO users (email, password) VALUES (?, ?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, u.EncryptedPassword)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = lastID

	return nil
}

// FindByEmail ..
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	q := "SELECT id, email, password FROM users where email = ?"

	err := r.store.db.QueryRow(q, email).Scan(&u.ID, &u.Email, &u.EncryptedPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}
