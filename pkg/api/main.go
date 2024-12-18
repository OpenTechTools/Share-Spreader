package main

import (
	"log"

	"github.com/OpenTechTools/Share-Spreader/internal/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	// Engine instance intialize the gin engine
	r := gin.Default()

	r.GET("/", auth.Home)

	err := auth.ConfigureOauth()
	if err == nil {
		r.GET("/auth/:provider", auth.SignInWithProvider)
		r.GET("/auth/:provider/callback", auth.CallbackHandler)
	}

	r.GET("/", auth.Home)

	r.GET("/success", auth.Success)

	// Start the server
	port := auth.GetEnv("PORT", "5000")
	log.Printf("Starting server on port %s...", port)
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
