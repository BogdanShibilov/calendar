package user

import (
	"context"
	"database/sql"
	"errors"
	"hwCalendar/storage"
)

func ById(ctx context.Context, id int) (*User, error) {
	var user User
	err := pgStorage.GetContext(ctx, &user, "SELECT id, username, password_hash, created_at, updated_at FROM users WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func All(ctx context.Context) ([]User, error) {
	var users []User
	err := pgStorage.SelectContext(ctx, &users, "SELECT id, username, password_hash, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func Delete(ctx context.Context, id int) error {
	deletedUser := &User{Id: id}
	return deletedUser.Delete(ctx)
}
