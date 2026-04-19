package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"dbut.dev/x/vanity"
)

func main() {
	raw := os.Getenv("GO_MODULE")
	if raw == "" {
		log.Fatal("GO_MODULE is required (format: <module>=<repo>)")
	}
	module, repo, ok := strings.Cut(raw, "=")
	if !ok || module == "" || repo == "" {
		log.Fatalf("GO_MODULE must be in format <module>=<repo>, got %q", raw)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(vanity.MiddlewareFor(module, repo))
	r.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusNotFound)
	})

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
