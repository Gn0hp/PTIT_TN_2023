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
	adminHandler := admin.New(logger, repo, mqService)
	candidateHandler := candidate.New(logger, repo)
	voterHandler := voter.New(logger, repo, mqService)
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
	adminUrl.POST("/reject-voter", adminHandler.RejectVoter)
	adminUrl.POST("/reject-candidate", adminHandler.RejectCandidate)
	adminUrl.GET("/voter-list", adminHandler.GetVoterList)
	adminUrl.GET("/voter-detail/:id", adminHandler.GetVoterDetail)
	adminUrl.PATCH("/update-voter/:id", adminHandler.UpdateVoter)
	adminUrl.DELETE("/delete-voter/:id", adminHandler.DeleteVoter)
	adminUrl.GET("/candidate-list", adminHandler.GetCandidateList)
	adminUrl.GET("/candidate-detail/:id", adminHandler.GetCandidateDetail)
	adminUrl.PATCH("/update-candidate/:id", adminHandler.UpdateCandidate)
	adminUrl.DELETE("/delete-candidate/:id", adminHandler.DeleteCandidate)
	adminUrl.GET("/stat-by-candidate/:election_id", adminHandler.StatByCandidate)
	adminUrl.GET("/stat-ballot-by-candidate/:candidate_id", adminHandler.StatBallotByCandidate)
	adminUrl.GET("/stat-by-ballot/:ballot_id", adminHandler.StatByBallotId)

	adminBlockchainUrl := adminUrl.Group("/blockchain")
	adminBlockchainUrl.GET("/get-nonce", adminHandler.GetBlockchainNonce)
	//adminBlockchainUrl.GET("/get-election-result", adminHandler.GetElectionResult)
	//adminBlockchainUrl.GET("/get-election-to-result", adminHandler.GetElectionToResutl)

	voterUrl := api.Group("/voter")
	voterUrl.POST("/vote", voterHandler.Vote)
	voterUrl.POST("/register-candidate", voterHandler.RegisterCandidate)
	voterUrl.GET("/view-candidate", voterHandler.ViewCandidate)

	electionUrl := api.Group("/election")
	electionUrl.GET("/check", electionHandler.CheckElection)
	electionUrl.POST("/create", electionHandler.CreateElection)
	electionUrl.GET("/view-result", electionHandler.ViewResult)
	electionUrl.POST("/push-blockchain/:id", electionHandler.PushBlockchain)
	electionUrl.POST("/close/:id", electionHandler.CloseElection)

	electionRoleUrl := api.Group("/election-role")
	electionRoleUrl.GET("/find-by-election-id", electionRoleHandler.FindRolesByElectionId)

	candidateUrl := api.Group("/candidate")
	candidateUrl.GET("/get-posts-by-candidate-id", candidateHandler.FindAllPost)
	candidateUrl.POST("/post-create", candidateHandler.CreatePost)
	candidateUrl.PATCH("/post-update", candidateHandler.UpdatePost)
	candidateUrl.DELETE("/post-delete", candidateHandler.DeletePost)
	candidateUrl.GET("/watch-result/:id", candidateHandler.WatchResult)

	return r
}
