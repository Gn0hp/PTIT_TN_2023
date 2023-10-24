package routers

import (
	"PTIT_TN/internal/handlers/admin"
	election2 "PTIT_TN/internal/handlers/election"
	"PTIT_TN/internal/handlers/user"
	"PTIT_TN/internal/handlers/voter"
	"PTIT_TN/internal/repositories"
	"PTIT_TN/internal/services/db"
	"PTIT_TN/internal/services/db/redis"
	"PTIT_TN/pkg/auth"
	"PTIT_TN/pkg/rabbitMQ"
	"PTIT_TN/pkg/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"logur.dev/logur"
)

func New(logger logur.LoggerFacade, database *db.GormDB, redisDb *redis.Client, mqService rabbitMQ.MqService, isDevEnv bool) *gin.Engine {
	r := gin.New()

	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/api/v1/liveliness"),
		gin.Recovery(),
	)
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	if isDevEnv {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = []string{"*"}
	}
	// Middlewares
	r.Use(
		cors.New(config),                 // CORS
		auth.Middleware(logger),          // Auth JWT
		utils.ResponseMiddleware(logger)) // Response
	//r.Use(web.RequestIdMiddleware())

	repo := repositories.New(logger, database, redisDb)
	userHandler := user.New(logger, repo)
	adminHandler := admin.New(logger, repo)
	voterHandler := voter.New(logger, repo)
	electionHandler := election2.New(logger, repo, mqService)

	api := r.Group("/api/v1")

	authUrl := api.Group("/auth")
	authUrl.POST("/register", userHandler.Register)
	authUrl.POST("/login", userHandler.Login)
	authUrl.GET("/logout", userHandler.Logout)

	adminUrl := api.Group("/admin")
	adminUrl.POST("/login", adminHandler.AdminLogin)
	adminUrl.GET("/pending-users", adminHandler.GetPendingUserList)

	voterUrl := api.Group("/voter")
	voterUrl.POST("/vote", voterHandler.Vote) //TODO

	electionUrl := api.Group("/election")
	electionUrl.POST("/create", electionHandler.CreateElection) // TODO

	return r
}
