package jwt

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"hwCalendar/jwt/storage/redis"
	"os"
	"strconv"
	"time"
)

var refreshTTL = os.Getenv("REFRESH_TTL")

var refreshDb = redis.GetDb(redis.RefreshTokenDb)

type RefreshTokenClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func NewRefreshTokenClaims(id int) *RefreshTokenClaims {
	now := time.Now()
	ttl, _ := time.ParseDuration(refreshTTL)

	return &RefreshTokenClaims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
}

func (c *RefreshTokenClaims) GenerateToken(ctx context.Context) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	idString := strconv.Itoa(c.Id)
	expiresIn := c.ExpiresAt.Sub(time.Now())
	refreshDb.Set(ctx, idString, tokenString, expiresIn)

	return tokenString, nil
}

func IsRefreshTokenInRedis(ctx context.Context, id int, givenToken string) bool {
	idString := strconv.Itoa(id)
	tokenString := refreshDb.Get(ctx, idString).Val()
	if tokenString != givenToken {
		return false
	}
	return true
}

func RemoveRefreshToken(ctx context.Context, id int) {
	_ = refreshDb.Del(ctx, strconv.Itoa(id))
}

func ParseRefreshToken(tokenString string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, ErrMalformedToken
		}
		return nil, err
	}

	if !token.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*RefreshTokenClaims); ok {
		return claims, nil
	}
	return nil, ErrUnknownClaims
}
