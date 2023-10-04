package routers

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	r := gin.New()

	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/api/v1/liveliness"),
		gin.Recovery(),
	)

	return r
}
