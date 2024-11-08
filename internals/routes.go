package internals

import "github.com/gin-gonic/gin"

// This is a config structure where a shared instance which in our case is the pointer to the gin engine is used.
type Config struct {
    Router *gin.Engine
}

/*
    Just a normal method named as Routes. But here what is different is you have something that is called a reciever and basically this recievers job is to attach the method to that 
    particular instance. Here the instance is a pointer of Config struct and the app is just you know the name of this instance. So you are basically attaching it to the Config struct
    from above and you can call the method using the instance in our main method. This line is basically saying attach Routes to Config, so any Config instance (like app) can 
    use this method.
*/
func (app *Config) Routes() {
    //routes will come here. These are just basic routes of what function will be used in various routes.
    app.Router.GET("/", app.indexPageHandler())

    app.Router.GET("/api/temp", app.HandleStockData())
}