package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"dbut.dev/x/vanity"
)

func main() {
	r := gin.Default()
	r.Use(vanity.Middleware("github.com/dbut2/x"))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "pulsar.dbut.dev/vanity-example — append ?go-get=1 to any path to see the import metadata")
	})
	log.Fatal(r.Run(":8080"))
}
