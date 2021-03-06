package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/ataboo/go-natm/pkg/api/routes"
	"github.com/ataboo/go-natm/pkg/common"
	"github.com/ataboo/go-natm/pkg/database"
	"github.com/ataboo/go-natm/pkg/oauth"
	"github.com/ataboo/go-natm/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/dig"
)

// migrate -source file://migrations -database postgres://username:pw@localhost:5432/gonatm up
// sqlboiler psql

// func serveWs(w http.ResponseWriter, r *http.Request) {
// 	ws, err := websocket.Upgrade(w, r)
// 	if err != nil {
// 		fmt.Fprintf(w, "%+V\n", err)
// 	}
// 	go websocket.Writer(ws)
// 	websocket.Reader(ws)
// }

// func setupRoutes() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "simpler server")
// 	})

// 	http.HandleFunc("/ws", serveWs)
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file", err)
	}

	common.AssertEnvVarsSet()

	router := gin.Default()
	container := buildContainer()

	err = container.Invoke(func(
		google *oauth.GoogleOAuthService,
		jwtService *oauth.JWTRouteService,
		db *sql.DB,
		userRepo *storage.UserRepository,
		projectRepo *storage.ProjectRepository,
		taskRepo *storage.TaskRepository,
		statusRepo *storage.StatusRepository,
		projectAssociationRepo *storage.ProjectAssociationRepository,
	) {
		defer db.Close()
		err := database.MigrateDB(db)
		if err != nil {
			log.Fatal(err)
		}

		router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello world!")
		})

		api := router.Group("/api/v1", AllowCrossSite(), jwtService.AuthJWTMiddleware())
		{
			api.GET("/", func(c *gin.Context) {
				userID, ok := c.Get("acting_user_id")
				if !ok {
					panic("UserID not in context!")
				}

				c.String(http.StatusOK, "Hello user: "+userID.(string))
			})

			api.GET("/userinfo", func(c *gin.Context) {
				userID, _ := c.Get("acting_user_id")
				user, err := userRepo.Find(userID.(string))
				if err != nil {
					c.AbortWithError(http.StatusInternalServerError, errors.New("failed to find user model"))
					return
				}

				userVM := data.UserRead{
					Email: user.Email,
					ID:    user.ID,
					Name:  user.Name,
				}

				c.JSON(http.StatusOK, userVM)
			})

			api.POST("/logout", func(c *gin.Context) {
				oauth.ClearJWTCookie(c)
				c.Status(http.StatusOK)
			})

			routes.RegisterRoutes(api, projectRepo, statusRepo, taskRepo, projectAssociationRepo)
		}
		if err != nil {
			log.Fatal(err)
		}

		authGroup := router.Group("/auth/")

		google.RegisterGoogleRoutes(authGroup)
		jwtService.RegisterJWTRoutes(authGroup)

		router.Run("localhost:8080")
	})
	if err != nil {
		log.Fatal(err)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()

	container.Provide(database.NewSqlDB)
	container.Provide(oauth.NewJWTFactory)
	container.Provide(storage.NewUserRepository)
	container.Provide(storage.NewProjectRepository)
	container.Provide(storage.NewStatusRepository)
	container.Provide(storage.NewTaskRepository)
	container.Provide(storage.NewProjectAssociationRepository)
	container.Provide(oauth.LoadJWTConfig)
	container.Provide(oauth.NewGooglOAuthHandler)
	container.Provide(oauth.NewJWTRouteService)

	return container
}

func AllowCrossSite() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv(common.EnvFrontendHostname))
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}
