package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"hwCalendar/user/storage"
	"hwCalendar/user/storage/postgres"
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
	user := &User{
		Username:  username,
		PassHash:  string(passHash),
		CreatedAt: now,
		UpdatedAt: now,
	}

	return user, nil
}

func (u *User) Add(ctx context.Context) (int, error) {
	var insertedid int
	err := pgStorage.QueryRowxContext(
		ctx,
		"INSERT INTO users (username, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id",
		u.Username, u.PassHash, u.CreatedAt, u.UpdatedAt,
	).Scan(&insertedid)
	if err != nil {
		return 0, err
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

func (u *User) UpdateUsername(ctx context.Context, tx *sqlx.Tx, newUsername string) error {
	if u.Username == newUsername {
		return nil
	}

	u.Username = newUsername
	u.UpdatedAt = time.Now()
	_, err := tx.NamedExecContext(
		ctx,
		"UPDATE users SET username = :username, updated_at = :updated_at WHERE id = :id",
		u,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) UpdatePassword(ctx context.Context, tx *sqlx.Tx, newPassword string) error {
	newPassHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if u.PassHash == string(newPassHash) {
		return nil
	}

	u.PassHash = string(newPassHash)
	u.UpdatedAt = time.Now()
	_, err = tx.NamedExecContext(
		ctx,
		"UPDATE users SET password_hash = :password_hash, updated_at = :updated_at WHERE id = :id",
		u,
	)
	if err != nil {
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

func (u *User) VerifyPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.PassHash), []byte(password))
	if err != nil {
		return ErrInvalidPassword
	}

	return nil
}
