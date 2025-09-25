package gormdb

import (
	"context"

	"gorm.io/gorm"

	"site.test/user"
)

type Users struct {
	db *gorm.DB
}

func (db *DB) Users() *Users {
	return &Users{db: db.db}
}

func (users *Users) Create(ctx context.Context, u *user.User) (created *user.User, err error) {
	err = users.db.WithContext(ctx).Create(u).Error
	return u, err
}

func (users *Users) FindAll(ctx context.Context) (loaded []*user.User, err error) {
	err = users.db.WithContext(ctx).Find(&loaded).Error
	return loaded, err
}
