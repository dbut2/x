package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"dbut.dev/x/vanity"
)

func main() {
	r := gin.Default()
	r.Use(vanity.Middleware("github.com/dbut2/x/vanity-example"))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "vanity-example: pulsar.dbut.dev/ve")
	})
	log.Fatal(r.Run(":8080"))
}
