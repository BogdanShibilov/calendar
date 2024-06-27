package jwt

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"hwCalendar/jwt/config"
	"hwCalendar/jwt/storage/redis"
	"strconv"
	"time"
)

var refreshTTL = config.Get().RefreshTTL

var refreshDb = redis.GetDb(redis.RefreshTokenDb)

type RefreshTokenClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func NewRefreshTokenClaims(userId int) *RefreshTokenClaims {
	now := time.Now()
	ttl, _ := time.ParseDuration(refreshTTL)

	return &RefreshTokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
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

	key := strconv.Itoa(c.UserId) + ":" + c.ID
	expiresIn := c.ExpiresAt.Sub(time.Now())
	refreshDb.Set(ctx, key, tokenString, expiresIn)

	return tokenString, nil
}

func IsRefreshTokenInRedis(ctx context.Context, userId int, tokenId, givenToken string) bool {
	key := strconv.Itoa(userId) + ":" + tokenId
	tokenString := refreshDb.Get(ctx, key).Val()
	if tokenString != givenToken {
		return false
	}
	return true
}

func RemoveAllRefreshTokensFor(ctx context.Context, userId int) {
	keyPattern := strconv.Itoa(userId) + ":" + "*"
	keys := refreshDb.Keys(ctx, keyPattern).Val()

	for _, k := range keys {
		_ = refreshDb.Del(ctx, k)
	}
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
