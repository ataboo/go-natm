package oauth

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTAuthConfig struct {
	Secret         string
	Issuer         string
	Audience       string
	Subject        string
	IssueExpMins   int
	RefreshExpMins int
}

type JWTFactory struct {
	config *JWTAuthConfig
}

func NewJWTFactory(config *JWTAuthConfig) *JWTFactory {
	return &JWTFactory{
		config: config,
	}
}

func (f *JWTFactory) CreateAccessToken(userID string) string {
	claims := jwt.StandardClaims{
		Audience:  f.config.Audience,
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(f.config.IssueExpMins)).Unix(),
		Id:        userID,
		IssuedAt:  time.Now().Unix(),
		Issuer:    f.config.Issuer,
		NotBefore: 0,
		Subject:   f.config.Subject,
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(f.config.Secret))
	if err != nil {
		log.Fatal(err)
	}

	return token
}

func (f *JWTFactory) CreateRefreshToken() string {
	claims := jwt.StandardClaims{
		Audience:  f.config.Audience,
		Issuer:    f.config.Issuer,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(f.config.RefreshExpMins)).Unix(),
		Subject:   f.config.Subject,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := refreshToken.SignedString([]byte(f.config.Secret))
	if err != nil {
		log.Fatal(err)
	}

	return token
}

type JWTRouteService struct {
	config     *JWTAuthConfig
	jwtFactory *JWTFactory
}

func NewJWTRouteService(config *JWTAuthConfig, jwtFactory *JWTFactory) *JWTRouteService {
	service := JWTRouteService{
		config:     config,
		jwtFactory: jwtFactory,
	}

	return &service
}

func (s *JWTRouteService) RegisterJWTRoutes(parent *gin.RouterGroup) {
	g := parent.Group("/jwt")
	{
		g.POST("/refresh/", s.handleJWTRefresh)
	}
}

func (s *JWTRouteService) handleJWTRefresh(e *gin.Context) {
	e.String(http.StatusForbidden, "Todo")
}

func (s *JWTRouteService) AuthJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		strToken, err := c.Cookie("jwt-token")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(s.config.Secret), nil
		})
		if err != nil {
			ClearJWTCookie(c)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !claimsAreValid(*s.config, claims) {
			ClearJWTCookie(c)
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("claims invalid"))
			return
		}

		c.Set("acting_user_id", claims["jti"].(string))
	}
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

func LoadJWTConfig() *JWTAuthConfig {
	issueExpMins, _ := strconv.Atoi(os.Getenv("JWT_ISSUE_EXP_MINS"))
	refreshExpMins, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_EXP_MINS"))

	return &JWTAuthConfig{
		Audience:       os.Getenv("JWT_AUDIENCE"),
		IssueExpMins:   issueExpMins,
		RefreshExpMins: refreshExpMins,
		Issuer:         os.Getenv("JWT_ISSUER"),
		Secret:         os.Getenv("JWT_SECRET"),
		Subject:        os.Getenv("JWT_SUBJECT"),
	}
}

func ClearJWTCookie(c *gin.Context) {
	c.SetCookie("jwt-token", "", 0, "", "", true, true)
}
