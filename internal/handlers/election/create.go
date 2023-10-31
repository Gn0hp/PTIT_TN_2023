package election

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/pkg/rabbitMQ"
	"PTIT_TN/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateElection(c *gin.Context) {
	// TODO: Create New Election
	var createElectionBody CreateElectionBody
	if err := c.ShouldBindJSON(&createElectionBody); err != nil {
		_ = c.Error(err)
		return
	}
	createElectionBody.Election.Status = entities.ELECTION_STATUS_REGISTERING
	electionDto := entities.ElectionDto{
		Election: createElectionBody.Election,
	}
	rolesDto := createElectionBody.Roles
	if !electionDto.IsValid() || !entities.IsAllElectionRoleValid(rolesDto) {
		_ = c.Error(errors.New("[Election Handler] Invalid object value"))
		return
	}
	iEntity, _ := electionDto.ToEntity()
	entity := iEntity.(*entities.Election)
	err := h.repo.Database().Election().Create(c, entity)
	if err != nil {
		_ = c.Error(err)
		return
	}
	err = h.createElectionRelated(c, entity, rolesDto)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Election Handler] Create Election Related Error: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    entity,
	})
}

func (h *Handler) PushBlockchain(c *gin.Context) {
	// TODO: Push RabbitMQ save Blockchain
	// Update Election in off-chain db
	var electionDto entities.ElectionDto
	electionDto.Election.Status = entities.ELECTION_STATUS_OPENING
	var requestBody struct {
		Id uint64 `json:"id"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		_ = c.Error(err)
		return
	}
	iEntity, _ := electionDto.ToEntity()
	entity := iEntity.(*entities.Election)
	entity.ID = requestBody.Id
	err := h.repo.Database().Election().UpdateById(c, entity)
	if err != nil {
		_ = c.Error(err)
		return
	}
	h.mqService.Send(rabbitMQ.RabbitRequest{
		Type: "create_election",
		Data: CreateElectionReq{
			ID:           entity.ID,
			StartDate:    uint64(entity.DateStartElecting),
			Duration:     uint64(entity.Duration),
			NumCandidate: 0, // get count RoleElect in db
		},
	})

}

func (h *Handler) createElectionRelated(c *gin.Context, election *entities.Election, rolesDto []entities.ElectionRole) error {
	//ElectionCandidate, Election Result,
	electionCandidateDto := entities.ElectionCandidateDto{
		ElectionCandidate: entities.ElectionCandidate{
			ElectionId: election.ID,
			Note:       "",
		},
	}
	iEntity, _ := electionCandidateDto.ToEntity()
	entity := iEntity.(*entities.ElectionCandidate)
	if err := h.repo.Database().ElectionCandidate().Create(c, entity); err != nil {
		_ = c.Error(err)
		return err
	}
	// Create Election Role
	iRolesEntity, _ := entities.ElectionRoleToEntitiesWithEcId(rolesDto, entity.ID)
	if err := h.repo.Database().ElectionRole().BulkInsert(c, iRolesEntity); err != nil {
		_ = c.Error(err)
		return err
	}
	electionResultDto := entities.ElectionResultDto{
		ElectionResult: entities.ElectionResult{
			ElectionCandidateId: entity.ID,
			Note:                "",
		},
	}
	iEntity1, _ := electionResultDto.ToEntity()
	entity1 := iEntity1.(*entities.ElectionResult)
	if err := h.repo.Database().ElectionResult().Create(c, entity1); err != nil {
		_ = c.Error(err)
		return err
	}
	return nil
}
