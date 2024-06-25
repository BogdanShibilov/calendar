package jwt

import (
	"context"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GeneratePair(ctx context.Context, id int, username string) (*TokenPair, error) {
	accessToken, err := NewAccessTokenClaims(id, username).GenerateToken(ctx)
	if err != nil {
		return nil, err
	}
	refreshToken, err := NewRefreshTokenClaims(id).GenerateToken(ctx)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func RemovePair(ctx context.Context, id int) {
	RemoveAccessToken(ctx, id)
	RemoveRefreshToken(ctx, id)
}
