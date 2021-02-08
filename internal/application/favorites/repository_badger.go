package favorites

import (
	"context"
	"fmt"
	"strings"

	badger "github.com/dgraph-io/badger/v3"
)

type badgerFavoriteRepo struct {
	db *badger.DB
}

func NewBadgerFavoriteRepository(db *badger.DB) Repository {
	return &badgerFavoriteRepo{db: db}
}

func (r *badgerFavoriteRepo) ListByUser(ctx context.Context, username string) ([]string, error) {
	catalogEntries := []string{}
	err := r.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := fmt.Sprintf("users/%s/favorites/", username)
		bPrefix := []byte(prefix)
		for it.Seek(bPrefix); it.ValidForPrefix(bPrefix); it.Next() {
			item := it.Item()
			k := item.Key()
			catalogEntryID := strings.TrimPrefix(string(k), prefix)
			catalogEntries = append(catalogEntries, catalogEntryID)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return catalogEntries, nil
}

func (r *badgerFavoriteRepo) ListByMovie(ctx context.Context, catalogEntrySlug string) ([]string, error) {
	users := []string{}
	err := r.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := fmt.Sprintf("favorites/%s/users/", catalogEntrySlug)
		bPrefix := []byte(prefix)
		for it.Seek(bPrefix); it.ValidForPrefix(bPrefix); it.Next() {
			item := it.Item()
			k := item.Key()
			userID := strings.TrimPrefix(string(k), prefix)
			users = append(users, userID)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *badgerFavoriteRepo) AddFavorite(ctx context.Context, username, catalogEntrySlug string) error {
	txn := r.db.NewTransaction(true)
	key := fmt.Sprintf("users/%s/favorites/%s", username, catalogEntrySlug)
	reverseKey := fmt.Sprintf("favorites/%s/users/%s", catalogEntrySlug, username)
	data := "ok"

	err := txn.Set([]byte(key), []byte(data))
	if err != nil {
		return err
	}

	err = txn.Set([]byte(reverseKey), []byte(data))
	if err != nil {
		return err
	}

	return txn.Commit()
}

func (r *badgerFavoriteRepo) RemoveFavorite(ctx context.Context, username, catalogEntrySlug string) error {
	txn := r.db.NewTransaction(true)
	key := fmt.Sprintf("users/%s/favorites/%s", username, catalogEntrySlug)
	reverseKey := fmt.Sprintf("favorites/%s/users/%s", catalogEntrySlug, username)

	err := txn.Delete([]byte(key))
	if err != nil {
		return err
	}

	err = txn.Delete([]byte(reverseKey))
	if err != nil {
		return err
	}

	return txn.Commit()
}
