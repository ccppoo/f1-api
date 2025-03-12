package api

import (
	"github.com/gin-gonic/gin"
)

// @Summary hello
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /hello [get]
func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "hello"})
}
