package store_test

import (
	"testing"

	"github.com/ferikl/go-api/internal/app/model"
	"github.com/ferikl/go-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)

	tables := []string{"users"}

	defer teardown(tables)

	u, err := s.User().Create(&model.User{
		Email:             "test@test.com",
		EncryptedPassword: "asdasdasd",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
