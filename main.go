package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/music_room/adminpanel"
)

func main() {
	server := gin.Default()

	adminPanel := adminpanel.New()



	v1 := server.Group("/v1")
	{
		v1.GET("/welcome", func(c *gin.Context) {
			firstname := c.DefaultQuery("firstname", "Guest")
			lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

			c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
		})
		v1.GET("/admin", gin.WrapH(adminPanel))

	}

	server.Run("localhost:8888")

}
