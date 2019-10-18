package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User ..
type User struct {
	ID                int64
	Email             string
	Password          string
	EncryptedPassword string
}

// Validate ..
func (u *User) Validate() error {
	eRule := validation.Field(&u.Email, validation.Required, is.Email)
	pRule := validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100))

	return validation.ValidateStruct(u, eRule, pRule)
}

// BeforeCreate ..
func (u *User) BeforeCreate() error {

	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

func encryptString(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
