package routers

import (
	"PTIT_TN/internal/handlers/user"
	"PTIT_TN/internal/repositories"
	"PTIT_TN/internal/services/db"
	"PTIT_TN/internal/services/db/redis"
	"PTIT_TN/pkg/auth"
	"PTIT_TN/pkg/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"logur.dev/logur"
)

func New(logger logur.LoggerFacade, database *db.GormDB, redisDb *redis.Client, isDevEnv bool) *gin.Engine {
	r := gin.New()

	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/api/v1/liveliness"),
		gin.Recovery(),
	)
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	if isDevEnv {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = []string{"*"}
	}
	// Middlewares
	r.Use(cors.New(config))                 // CORS
	r.Use(auth.Middleware(logger))          // Auth JWT
	r.Use(utils.ResponseMiddleware(logger)) // Response
	//r.Use(web.RequestIdMiddleware())

	repo := repositories.New(logger, database, redisDb)
	userHandler := user.New(logger, repo)

	api := r.Group("/api/v1")

	authUrl := api.Group("/auth")
	authUrl.POST("/register", userHandler.Register)
	authUrl.GET("/login", userHandler.Login)
	authUrl.GET("/logout", userHandler.Logout)

	return r
}
