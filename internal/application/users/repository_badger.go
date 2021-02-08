package users

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/google/uuid"
)

type badgerUserRepo struct {
	db *badger.DB
}

func NewBadgerUserRepository(db *badger.DB) Repository {
	return &badgerUserRepo{db: db}
}

func (r *badgerUserRepo) CreateUser(ctx context.Context, param CreateUserParam) (*User, error) {
	txn := r.db.NewTransaction(true)
	key := fmt.Sprintf("users/%s", param.Username)
	user := User{
		ID:       uuid.New().String(),
		Username: param.Username,
		Password: param.Password,
	}

	data, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	err = txn.Set([]byte(key), []byte(data))
	if err != nil {
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}
	return r.FindUserByUsername(ctx, param.Username)
}

func (r *badgerUserRepo) FindUserByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	err := r.db.View(func(txn *badger.Txn) error {
		key := fmt.Sprintf("users/%s", username)
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &user)
		})
	})
	if errors.Is(err, badger.ErrKeyNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
