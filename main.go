package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ataboo/go-natm/v4/pkg/oauth"
	"github.com/ataboo/go-natm/v4/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/dig"
)

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

	err = container.Invoke(func(google *oauth.GoogleOAuthService, jwtService *oauth.JWTRouteService, db *sql.DB) {
		defer db.Close()

		router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello world!")
		})

		api := router.Group("/api", jwtService.AuthJWTMiddleware())
		{
			api.GET("/", func(c *gin.Context) {
				userID, ok := c.Get("acting_user_id")
				if !ok {
					panic("UserID not in context!")
				}

				c.String(http.StatusOK, "Hello user: "+userID.(string))
			})
		}

		google.RegisterGoogleRoutes(router)
		jwtService.RegisterJWTRoutes(router)

		router.Run(":8080")
	})
}

func buildContainer() *dig.Container {
	container := dig.New()

	container.Provide(storage.NewSqlDB)
	container.Provide(oauth.NewJWTFactory)
	container.Provide(storage.NewUserRepository)
	container.Provide(oauth.LoadJWTConfig)
	container.Provide(oauth.NewGooglOAuthHandler)
	container.Provide(oauth.NewJWTRouteService)

	return container
}
