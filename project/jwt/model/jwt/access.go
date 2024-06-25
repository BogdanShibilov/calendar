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

var (
	secret    = []byte(os.Getenv("SECRET"))
	accessTTL = os.Getenv("ACCESS_TTL")
)

var accessDb = redis.GetDb(redis.AccessTokenDb)

type AccessTokenClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewAccessTokenClaims(id int, username string) *AccessTokenClaims {
	now := time.Now()
	ttl, _ := time.ParseDuration(accessTTL)

	return &AccessTokenClaims{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
}

func (c *AccessTokenClaims) GenerateToken(ctx context.Context) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	idString := strconv.Itoa(c.Id)
	expiresIn := c.ExpiresAt.Sub(time.Now())
	accessDb.Set(ctx, idString, tokenString, expiresIn)

	return tokenString, nil
}

func IsAccessTokenInRedis(ctx context.Context, id int, givenToken string) bool {
	idString := strconv.Itoa(id)
	tokenString := accessDb.Get(ctx, idString).Val()
	if tokenString != givenToken {
		return false
	}
	return true
}

func RemoveAccessToken(ctx context.Context, id int) {
	_ = accessDb.Del(ctx, strconv.Itoa(id))
}

func ParseAccessToken(tokenString string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
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

	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrUnknownClaims
}
