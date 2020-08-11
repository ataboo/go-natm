package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ataboo/go-natm/v4/pkg/models"
	"github.com/ataboo/go-natm/v4/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauthapi "google.golang.org/api/oauth2/v2"
)

type GoogleOAuthService struct {
	config     *oauth2.Config
	userRepo   *storage.UserRepository
	jwtFactory *JWTFactory
}

func NewGooglOAuthHandler(userRepo *storage.UserRepository, jwtFactory *JWTFactory) *GoogleOAuthService {
	handler := GoogleOAuthService{
		config:     loadOAuthConfig(),
		userRepo:   userRepo,
		jwtFactory: jwtFactory,
	}

	return &handler
}

func (h *GoogleOAuthService) RegisterGoogleRoutes(parent *gin.RouterGroup) {
	g := parent.Group("/google/")
	{
		g.GET("/", h.handleAuthGet)
		g.GET("/callback", h.handleAuthCallback)
	}
}

func (h *GoogleOAuthService) handleAuthGet(c *gin.Context) {
	stateBytes := make([]byte, 16)
	rand.Read(stateBytes)
	state := base64.URLEncoding.EncodeToString(stateBytes)

	c.SetCookie("oauthstate", state, 3600, "", c.Request.Host, true, true)
	authCodeURL := h.config.AuthCodeURL(state)

	c.Redirect(http.StatusTemporaryRedirect, authCodeURL)
}

func (h *GoogleOAuthService) handleAuthCallback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")
	scope := c.Query("scope")

	if code == "" || state == "" || scope == "" {
		c.AbortWithError(401, errors.New("missing parameters"))
	}

	stateCookie, err := c.Cookie("oauthstate")
	if err != nil || stateCookie != state {
		c.AbortWithError(401, errors.New("invalid state"))
		return
	}

	token, err := h.config.Exchange(context.Background(), code)
	if err != nil {
		c.AbortWithError(401, errors.New("failed to get token"))
		return
	}

	client := h.config.Client(context.Background(), token)
	service, err := oauthapi.New(client)

	uiService := oauthapi.NewUserinfoService(service)

	userInfo, err := uiService.Get().Do()
	if err != nil {
		c.AbortWithError(401, errors.New("failed to get profile"))
	}

	fmt.Println("Email: " + userInfo.Email)
	fmt.Println("Name: " + userInfo.Name)

	user := h.userRepo.FindByEmail(userInfo.Email)
	creating := user == nil
	if creating {
		user = &models.User{
			ID:    uuid.New().String(),
			Email: userInfo.Email,
			Name:  userInfo.Name,
		}

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	if creating {
		h.userRepo.Create(user)
	} else {
		h.userRepo.Update(user)
	}

	accessToken := h.jwtFactory.CreateAccessToken(user.ID)
	c.SetCookie("jwt-token", accessToken, h.jwtFactory.config.IssueExpMins*60, "", "", true, true)

	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
}

func loadOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes:       []string{oauthapi.UserinfoProfileScope, oauthapi.UserinfoEmailScope},
	}
}
