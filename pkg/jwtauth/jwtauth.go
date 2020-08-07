package jwtauth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

type JWTAuthConfig struct {
	Secret         string
	Issuer         string
	Audience       string
	Subject        string
	ExpirationMins int
}

func AuthJWT(config JWTAuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(config.Secret))
			return b, nil
		})

		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("failed to parse token"))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !claimsAreValid(config, claims) {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("claims invalid"))
		}

		c.Set("acting_user_id", claims["jti"].(string))
	}
}

func CreateJWTToken(config JWTAuthConfig, userId string) (string, error) {
	claims := jwt.StandardClaims{
		Audience:  config.Audience,
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.ExpirationMins)).Unix(),
		Id:        userId,
		IssuedAt:  time.Now().Unix(),
		Issuer:    config.Issuer,
		NotBefore: 0,
		Subject:   config.Subject,
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(config.Secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func claimsAreValid(config JWTAuthConfig, claims jwt.MapClaims) bool {
	if claims.Valid() != nil {
		return false
	}

	if !claims.VerifyIssuer(config.Issuer, true) || !claims.VerifyAudience(config.Audience, true) {
		return false
	}

	if _, ok := claims["jti"]; !ok {
		return false
	}

	return true
}
