package auth_sql_store

import (
	"context"
	"rz-server/internal/app/user/infra/store"
	sql_store "rz-server/internal/app/user/infra/store/sql"
	auth_store_data "rz-server/internal/app/user/infra/store/sql/auth/data"
	repository "rz-server/internal/app/user/infra/store/sql/repository"
	"time"

	"github.com/google/uuid"
)

var _ store.AuthStore = (*AuthStore)(nil)

type AuthStore struct {
	Queries *repository.Queries
}

func New(store *sql_store.Repository) *AuthStore {
	return &AuthStore{
		Queries: store.Queries,
	}
}

func (s *AuthStore) SaveRefreshToken(body auth_store_data.RefreshTokenBody) *auth_store_data.Data {
	auth, err := s.Queries.SaveRefreshToken(context.Background(), repository.SaveRefreshTokenParams{
		UserID:    body.UserID,
		Token:     body.Token,
		ExpiredAt: body.ExpiredAt,
	})

	if err != nil {
		return nil
	}

	return &auth_store_data.Data{
		ID:       auth.ID,
		UserID:   auth.UserID,
		Token:    auth.Token,
		ExpireAt: auth.ExpiredAt,
	}
}

func (s *AuthStore) UpdateRefreshTokenExpiredAt(id uuid.UUID, expiredAt time.Time) error {
	err := s.Queries.UpdateRefreshTokenExpiredAt(context.Background(), repository.UpdateRefreshTokenExpiredAtParams{
		ID:        id,
		ExpiredAt: expiredAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *AuthStore) GetRefreshTokenByToken(token string) *auth_store_data.Data {
	auth, err := s.Queries.GetRefreshTokenByToken(context.Background(), token)

	if err != nil {
		return nil
	}

	return &auth_store_data.Data{
		ID:       auth.ID,
		UserID:   auth.UserID,
		Token:    auth.Token,
		ExpireAt: auth.ExpiredAt,
	}
}

func (s *AuthStore) GetRefreshTokenByUserID(userID uuid.UUID) *auth_store_data.Data {
	auth, err := s.Queries.GetRefreshTokenByUserID(context.Background(), userID)

	if err != nil {
		return nil
	}

	return &auth_store_data.Data{
		ID:       auth.ID,
		UserID:   auth.UserID,
		Token:    auth.Token,
		ExpireAt: auth.ExpiredAt,
	}
}

func (s *AuthStore) DeleteRefreshTokenByUserID(userID uuid.UUID) error {
	err := s.Queries.DeleteRefreshTokenByUserID(context.Background(), userID)

	if err != nil {
		return err
	}

	return nil
}
