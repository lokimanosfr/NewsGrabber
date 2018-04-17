package router

import (
	"net/http"

	"github.com/projects/newsgrabber/news"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}

func collectHandler(c *gin.Context) {
	category := c.Param("category")

	news.Collect(category)
	c.String(http.StatusOK, "Search is init")
}
func resultHandler(c *gin.Context) {
	category := c.Param("category")

	topics := news.Result(category)
	c.JSON(http.StatusOK, topics)
}
