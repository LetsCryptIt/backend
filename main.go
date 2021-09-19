package main

import (
	"firebase.google.com/go/v4/auth"

	"github.com/LetsCryptIt/backend/api"
	"github.com/LetsCryptIt/backend/config"
	"github.com/LetsCryptIt/backend/middleware"
	"github.com/gin-gonic/gin"
)

var client *auth.Client

func main() {
	// initialize new gin engine (for server)
	r := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	client = config.SetupFirebase()

	r.Use(func(c *gin.Context) {
		c.Set("firebaseAuth", client)
	})

	// Public Api endpoints
	r.POST("/api/login", api.Login)

	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(middleware.AuthMiddleware)
	{
		authorized.GET("/api/test", api.Test)
	}

	// start the server
	r.Run(":8000")
}
