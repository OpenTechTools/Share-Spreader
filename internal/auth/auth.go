package auth

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func ConfigureOauth() error {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// Fetch environment variables
	clientID := GetEnv("CLIENT_ID", "")
	clientSecret := GetEnv("CLIENT_SECRET", "")
	clientCallbackURL := GetEnv("CLIENT_CALLBACK_URL", "")

	if clientID == "" || clientSecret == "" || clientCallbackURL == "" {
		return errors.New("environment variables (CLIENT_ID, CLIENT_SECRET, CLIENT_CALLBACK_URL) are required")
	}

	// Configure OAuth provider
	goth.UseProviders(
		google.New(clientID, clientSecret, clientCallbackURL),
	)
	return nil
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func SignInWithProvider(c *gin.Context) {
	provider := c.Param("provider")
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func CallbackHandler(c *gin.Context) {
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

	// Opening a connection to the Spreader server
	if err == nil {
		log.Printf("Authenticated user: %+v", user)
	}

}
