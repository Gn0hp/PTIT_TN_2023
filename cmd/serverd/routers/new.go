package routers

import (
	"PTIT_TN/internal/handlers/admin"
	"PTIT_TN/internal/handlers/candidate"
	election2 "PTIT_TN/internal/handlers/election"
	"PTIT_TN/internal/handlers/electionRoles"
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
	candidateHandller := candidate.New(logger, repo)
	voterHandler := voter.New(logger, repo)
	electionHandler := election2.New(logger, repo, mqService)
	electionRoleHandler := electionRoles.New(logger, repo)

	api := r.Group("/api/v1")

	authUrl := api.Group("/auth")
	authUrl.POST("/register", userHandler.Register)
	authUrl.POST("/login", userHandler.Login)
	authUrl.GET("/logout", userHandler.Logout)

	adminUrl := api.Group("/admin")
	adminUrl.POST("/login", adminHandler.AdminLogin)
	adminUrl.GET("/pending-users", adminHandler.GetPendingUserList)
	adminUrl.POST("/verify-voter", adminHandler.VerifyVoter)
	adminUrl.POST("/verify-candidate", adminHandler.VerifyCandidate)

	voterUrl := api.Group("/voter")
	voterUrl.POST("/vote", voterHandler.Vote) //TODO
	voterUrl.POST("/register-candidate", voterHandler.RegisterCandidate)
	voterUrl.GET("/view-candidate", voterHandler.ViewCandidate)

	electionUrl := api.Group("/election")
	electionUrl.GET("/check", electionHandler.CheckElection)
	electionUrl.POST("/create", electionHandler.CreateElection) // TODO

	electionRoleUrl := api.Group("/election-role")
	electionRoleUrl.GET("/find-by-election-id", electionRoleHandler.FindRolesByElectionId)

	candidateUrl := api.Group("/candidate")
	candidateUrl.GET("/get-posts-by-candidate-id", candidateHandller.FindAllPost)
	candidateUrl.POST("/post-create", candidateHandller.CreatePost)
	candidateUrl.PATCH("/post-update", candidateHandller.UpdatePost)
	candidateUrl.DELETE("/post-delete", candidateHandller.DeletePost)

	return r
}
