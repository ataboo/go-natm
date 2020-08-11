package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ataboo/go-natm/v4/pkg/database"
	"github.com/ataboo/go-natm/v4/pkg/oauth"
	"github.com/ataboo/go-natm/v4/pkg/storage"
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

	router := gin.Default()
	container := buildContainer()

	err = container.Invoke(func(google *oauth.GoogleOAuthService, jwtService *oauth.JWTRouteService, db *sql.DB, userRepo *storage.UserRepository) {
		defer db.Close()
		err := database.MigrateDB(db)
		if err != nil {
			log.Fatal(err)
		}

		router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello world!")
		})

		api := router.Group("/api", AllowCrossSite(), jwtService.AuthJWTMiddleware())
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
				c.JSON(http.StatusOK, map[string]string{"name": user.Name})

				fmt.Printf("Found user: %+v\n", user)
			})

			api.POST("/logout", func(c *gin.Context) {
				oauth.ClearJWTCookie(c)
				c.Status(http.StatusOK)
			})
		}

		authGroup := router.Group("/auth/")

		google.RegisterGoogleRoutes(authGroup)
		jwtService.RegisterJWTRoutes(authGroup)

		router.Run(":8080")
	})
}

func buildContainer() *dig.Container {
	container := dig.New()

	container.Provide(database.NewSqlDB)
	container.Provide(oauth.NewJWTFactory)
	container.Provide(storage.NewUserRepository)
	container.Provide(oauth.LoadJWTConfig)
	container.Provide(oauth.NewGooglOAuthHandler)
	container.Provide(oauth.NewJWTRouteService)

	return container
}

func AllowCrossSite() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}
