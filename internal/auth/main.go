package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found. Falling back to system environment variables")
	}

	// Fetch environment variables
	clientID := getEnv("CLIENT_ID", "")
	clientSecret := getEnv("CLIENT_SECRET", "")
	clientCallbackURL := getEnv("CLIENT_CALLBACK_URL", "")

	if clientID == "" || clientSecret == "" || clientCallbackURL == "" {
		log.Fatal("Environment variables (CLIENT_ID, CLIENT_SECRET, CLIENT_CALLBACK_URL) are required")
	}

	// Configure OAuth provider
	goth.UseProviders(
		google.New(clientID, clientSecret, clientCallbackURL),
	)

	// Initialize Gin
	r := gin.Default()

	// Load templates
	r.LoadHTMLGlob("templates/*")

	// Define routes
	r.GET("/", home)
	r.GET("/auth/:provider", signInWithProvider)
	r.GET("/auth/:provider/callback", callbackHandler)
	r.GET("/success", success)

	// Start the server
	port := getEnv("PORT", "5000")
	log.Printf("Starting server on port %s...", port)
	err = r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func signInWithProvider(c *gin.Context) {
	provider := c.Param("provider")
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func callbackHandler(c *gin.Context) {
	provider := c.Param("provider")
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Printf("Error during authentication: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Redirect to success route with user data
	log.Printf("Authenticated user: %+v", user)
	c.Redirect(http.StatusTemporaryRedirect, "/success")
}

func success(c *gin.Context) {
	htmlContent := `
      <div style="
          background-color: #fff;
          padding: 40px;
          border-radius: 8px;
          box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
          text-align: center;
      ">
          <h1 style="
              color: #333;
              margin-bottom: 20px;
          ">You have successfully signed in!</h1>
      </div>
  `
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
}
