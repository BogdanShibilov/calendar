package user

import (
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"hwCalendar/storage"
	"hwCalendar/storage/postgres"
	"time"
)

type User struct {
	Id        int       `db:"id"`
	Username  string    `db:"username"`
	PassHash  string    `db:"password_hash"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

var pgStorage = postgres.GetDb()

func New(username, pass string) (*User, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	return &User{
		Username:  username,
		PassHash:  string(passHash),
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (u *User) Add(ctx context.Context) (int, error) {
	if err := validate(u); err != nil {
		return -1, err
	}

	var insertedid int
	err := pgStorage.QueryRowxContext(
		ctx,
		"INSERT INTO users (username, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id",
		u.Username, u.PassHash, u.CreatedAt, u.UpdatedAt,
	).Scan(&insertedid)
	if err != nil {
		return -1, err
	}

	u.Id = insertedid

	return insertedid, nil
}

func (u *User) Update(ctx context.Context, newUsername, newPass string) error {
	newPassHash, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Username = newUsername
	u.PassHash = string(newPassHash)

	err = validate(u)
	if err != nil {
		return err
	}

	u.UpdatedAt = time.Now()
	_, err = pgStorage.NamedExecContext(
		ctx,
		"UPDATE users SET username = :username, password_hash = :password_hash, updated_at = :updated_at WHERE id = :id",
		u,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.ErrNotFound
		}
		return err
	}

	return nil
}

func (u *User) Delete(ctx context.Context) error {
	_, err := pgStorage.NamedExecContext(ctx, "DELETE FROM users WHERE id = :id", u)
	if err != nil {
		return err
	}

	return nil
}

func validate(user *User) error {
	if user.Username == "" {
		return ErrEmptyUsername
	}
	if user.PassHash == "" {
		return ErrEmptyPassword
	}
	return nil
}
