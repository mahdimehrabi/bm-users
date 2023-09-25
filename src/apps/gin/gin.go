package gin

import "github.com/gin-gonic/gin"

func RunServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8100") // listen and serve on 0.0.0.0:8080
}
