package catalog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	tvtime "com.aviebrantz.tvtime/pkg/api"
	badger "github.com/dgraph-io/badger/v3"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type badgerCatalogRepo struct {
	db *badger.DB
}

func NewBadgerCatalogRepository(db *badger.DB) Repository {
	return &badgerCatalogRepo{db: db}
}

func (r *badgerCatalogRepo) CreateCatalogEntry(ctx context.Context, param CreateCatalogEntryParam) (*CatalogEntry, error) {
	slug := slug.Make(param.Name)

	e, err := r.FindCatalogEntryBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	if e != nil {
		return nil, errors.New("duplicated entry")
	}

	txn := r.db.NewTransaction(true)

	key := fmt.Sprintf("catalog/%s", slug)
	entry := CatalogEntry{
		ID:       uuid.New().String(),
		Slug:     slug,
		Name:     param.Name,
		ImageURL: param.ImageURL,
		Type:     *param.Type,
	}

	data, err := json.Marshal(entry)
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
	return r.FindCatalogEntryBySlug(ctx, slug)
}

func (r *badgerCatalogRepo) FindCatalogEntryBySlug(ctx context.Context, slug string) (*CatalogEntry, error) {
	var entry CatalogEntry
	err := r.db.View(func(txn *badger.Txn) error {
		key := fmt.Sprintf("catalog/%s", slug)
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &entry)
		})
	})
	if errors.Is(err, badger.ErrKeyNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (r *badgerCatalogRepo) ListEntries(ctx context.Context, filter EntriesFilter) ([]CatalogEntry, int, error) {
	catalogEntries := []CatalogEntry{}
	total := 0
	err := r.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := "catalog/"
		bPrefix := []byte(prefix)
		for it.Seek(bPrefix); it.ValidForPrefix(bPrefix); it.Next() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				entry := CatalogEntry{}
				err := json.Unmarshal(val, &entry)
				if err != nil {
					return err
				}
				filtered := false
				if filter.Type != nil {
					if _, ok := tvtime.CatalogType_value[filter.Type.String()]; ok {
						if entry.Type != *filter.Type {
							filtered = true
						}
					}
				}

				if filter.Term != nil && len(*filter.Term) > 0 {
					if !strings.Contains(strings.ToLower(entry.Name), strings.ToLower(*filter.Term)) {
						filtered = true
					}
				}

				if !filtered {
					total++
					catalogEntries = append(catalogEntries, entry)
				}

				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, total, err
	}

	if filter.PerPage == 0 {
		filter.PerPage = 10
	}

	offset := filter.Page * filter.PerPage
	if offset > len(catalogEntries) {
		offset = len(catalogEntries)
	}
	limit := offset + filter.PerPage
	if limit > len(catalogEntries) {
		limit = len(catalogEntries)
	}

	return catalogEntries[offset:limit], total, nil
}

func (r *badgerCatalogRepo) ListByMovie(ctx context.Context, catalogEntryID string) ([]string, error) {
	users := []string{}
	err := r.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := fmt.Sprintf("favorites/%s/users/", catalogEntryID)
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

func (r *badgerCatalogRepo) AddFavorite(ctx context.Context, username, catalogEntryID string) error {
	txn := r.db.NewTransaction(true)
	key := fmt.Sprintf("users/%s/favorites/%s", username, catalogEntryID)
	reverseKey := fmt.Sprintf("favorites/%s/users/%s", catalogEntryID, username)
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

func (r *badgerCatalogRepo) RemoveFavorite(ctx context.Context, username, catalogEntryID string) error {
	txn := r.db.NewTransaction(true)
	key := fmt.Sprintf("users/%s/favorites/%s", username, catalogEntryID)
	reverseKey := fmt.Sprintf("favorites/%s/users/%s", catalogEntryID, username)

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
