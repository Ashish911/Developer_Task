package internals

import "github.com/gin-gonic/gin"

type Config struct {
    Router *gin.Engine
}

func (app *Config) Routes() {
    //routes will come here
    app.Router.GET("/", app.indexPageHandler())

    app.Router.GET("/api/temp", app.HandleStockData())
}