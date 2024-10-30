package main

import (
	"go_backend/internals"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    //initialize config
    app := internals.Config{Router: router}

    //routes
    app.Routes()

    router.Static("/static", "../static")

    router.Run(":8080")

}