package auth_service

import (
	"api/config"
	"api/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthenticationTokens struct {
	AccessToken  string
	RefreshToken string
}

func CreateTokens(user models.User) (*AuthenticationTokens, error) {
	// TODO: Migrate to RegisteredClaims
	access_claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: (time.Now().Add(time.Minute * 45)).Unix(),
	})

	access_token, access_err := access_claims.SignedString([]byte(config.Get("ACCESS_SECRET")))

	if access_err != nil {
		return nil, access_err
	}

	refresh_claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: (time.Now().Add(time.Hour * 24 * 7)).Unix(),
	})

	refresh_token, refresh_err := refresh_claims.SignedString([]byte(config.Get("REFRESH_SECRET")))

	if refresh_err != nil {
		return nil, refresh_err
	}

	tokens := AuthenticationTokens{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}

	return &tokens, nil
}
