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

var (
	secret    = []byte(config.Get().Secret)
	accessTTL = config.Get().AccessTTL
)

var accessDb = redis.GetDb(redis.AccessTokenDb)

type AccessTokenClaims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewAccessTokenClaims(id int, username string) *AccessTokenClaims {
	now := time.Now()
	ttl, _ := time.ParseDuration(accessTTL)

	return &AccessTokenClaims{
		UserId:   id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
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

	key := strconv.Itoa(c.UserId) + ":" + c.ID
	expiresIn := c.ExpiresAt.Sub(time.Now())
	accessDb.Set(ctx, key, tokenString, expiresIn)

	return tokenString, nil
}

func IsAccessTokenInRedis(ctx context.Context, userId int, tokenId, givenToken string) bool {
	key := strconv.Itoa(userId) + ":" + tokenId
	tokenString := accessDb.Get(ctx, key).Val()
	if tokenString != givenToken {
		return false
	}
	return true
}

func RemoveAllAccessTokensFor(ctx context.Context, userId int) {
	keyPattern := strconv.Itoa(userId) + ":" + "*"
	keys := accessDb.Keys(ctx, keyPattern).Val()

	for _, k := range keys {
		_ = accessDb.Del(ctx, k)
	}
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
