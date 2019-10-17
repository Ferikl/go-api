package store

import "github.com/ferikl/go-api/internal/app/model"

// UserRepository ..
type UserRepository struct {
	store *Store
}

// Create ..
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	stmt, err := r.store.db.Prepare("INSERT INTO users (email, password) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(u.Email, u.EncryptedPassword)
	if err != nil {
		return nil, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = lastID

	return u, nil
}

// FindByEmail ..
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
