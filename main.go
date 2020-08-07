package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ataboo/go-natm/v4/pkg/jwtauth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})

	router.GET("/test-token", func(c *gin.Context) {
		testToken, err := jwtauth.CreateJWTToken(loadJWTConfig(), "test-user-id")
		if err != nil {
			panic(err)
		}

		c.String(http.StatusOK, testToken)
	})

	api := router.Group("/api", jwtauth.AuthJWT(loadJWTConfig()))
	{
		api.GET("/", func(c *gin.Context) {
			userID, ok := c.Get("acting_user_id")
			if !ok {
				panic("UserID not in context!")
			}

			c.String(http.StatusOK, "Hello user: "+userID.(string))
		})
	}

	fmt.Println("Started server on :8080")

	router.Run(":8080")
}

func loadJWTConfig() jwtauth.JWTAuthConfig {
	expirationMins, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))

	return jwtauth.JWTAuthConfig{
		Audience:       os.Getenv("JWT_AUDIENCE"),
		ExpirationMins: expirationMins,
		Issuer:         os.Getenv("JWT_ISSUER"),
		Secret:         os.Getenv("JWT_SECRET"),
		Subject:        os.Getenv("JWT_SUBJECT"),
	}
}
