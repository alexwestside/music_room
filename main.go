package main

import (
	"github.com/music_room/config"
	"github.com/gin-gonic/gin"
	"github.com/music_room/adminpanel"
	"net/http"
)

func main() {
	config.Load()

	server := gin.Default()

	mux := adminpanel.New()

	server.GET("/admin/", gin.WrapH(mux))

	v1 := server.Group("/v1")
	{
		v1.GET("/welcome", func(c *gin.Context) {
			firstname := c.DefaultQuery("firstname", "Guest")
			lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

			c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
		})
	}

	server.Run("localhost:8888")

}
