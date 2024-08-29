package storage

import (
	"github.com/SaidovZohid/deposit-project/storage/postgres"
	"github.com/SaidovZohid/deposit-project/storage/repo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StorageI interface {
	User() repo.UserStorageI
}

type storage struct {
	userRepo repo.UserStorageI
}

func New(db *pgxpool.Pool) StorageI {
	return &storage{
		userRepo: postgres.NewUser(db),
	}
}

func (s *storage) User() repo.UserStorageI {
	return s.userRepo
}
