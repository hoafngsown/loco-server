package metadata_sql_store

import (
	"context"
	"rz-server/internal/app/user/infra/store"
	sql_store "rz-server/internal/app/user/infra/store/sql"
	repository "rz-server/internal/app/user/infra/store/sql/repository"

	"github.com/google/uuid"
)

var _ store.PreferenceMetadataStore = (*PreferenceMetadataStore)(nil)

type PreferenceMetadataStore struct {
	Queries *repository.Queries
}

func New(store *sql_store.Repository) *PreferenceMetadataStore {
	return &PreferenceMetadataStore{
		Queries: store.Queries,
	}
}

func (s *PreferenceMetadataStore) VerifyMetadataIDs(ids []uuid.UUID) bool {
	valid, err := s.Queries.VerifyMetadataIDs(context.Background(), ids)

	if err != nil {
		return false
	}

	return valid
}
