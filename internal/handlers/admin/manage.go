package admin

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) GetVoterList(c *gin.Context) {
	voters, err := h.repo.Database().Voter().FindAll(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    voters,
	})
}
func (h *Handler) UpdateVoter(c *gin.Context) {
	voterId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler UpdateVoter] Invalid voter_id, detail: %v", err))
		_ = c.Error(err)
		return
	}
	var voterDto entities.VoterDto
	if err := c.ShouldBindJSON(&voterDto); err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler UpdateVoter] Invalid request body: %v", err))
		_ = c.Error(err)
		return
	}
	voterDto.Voter.ID = uint64(voterId)
	iEntity, err := voterDto.ToEntity()
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler UpdateVoter] Convert to entity failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	entity := iEntity.(*entities.Voter)
	err = h.repo.Database().Voter().UpdateById(c, entity)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler UpdateVoter] Update voter failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
	})
}
func (h *Handler) DeleteVoter(c *gin.Context) {
	voterId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler DeleteVoter] Invalid voter_id, detail: %v", err))
		_ = c.Error(err)
		return
	}
	err = h.repo.Database().Voter().Delete(c, uint64(voterId))
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler DeleteVoter] Delete voter failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
	})
}
func (h *Handler) GetVoterDetail(c *gin.Context) {
	voterId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler GetVoterDetail] Invalid voter_id, detail: %v", err))
		_ = c.Error(err)
		return
	}
	voterExist, err := h.repo.Database().Voter().FindById(c, uint64(voterId))
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler GetVoterDetail] Find voter failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    voterExist,
	})
}
func (h *Handler) GetCandidateDetail(c *gin.Context) {
	//candidateId := c.Param("id")
	candidateId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler GetCandidateDetail] Invalid id, detail: %v", err))
		_ = c.Error(err)
		return
	}
	candidateExist, err := h.repo.Database().Candidate().FindById(c, uint64(candidateId))
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler GetCandidateDetail] Find candidate failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    candidateExist,
	})
}
func (h *Handler) UpdateCandidate(c *gin.Context) {
	candidateId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler UpdateCandidate] Invalid voter_id, detail: %v", err))
		_ = c.Error(err)
		return
	}
	var candidateDto entities.CandidateDto
	if err := c.ShouldBindJSON(&candidateDto); err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler UpdateCandidate] Invalid request body: %v", err))
		_ = c.Error(err)
		return
	}
	iEntity, err := candidateDto.ToEntity()
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler UpdateCandidate] Convert to entity failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	entity := iEntity.(*entities.Candidate)
	err = h.repo.Database().Candidate().UpdateById(c, uint64(candidateId), entity)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler UpdateCandidate] Update candidate failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
	})
}
func (h *Handler) DeleteCandidate(c *gin.Context) {
	candidateId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler DeleteCandidate] Invalid voter_id, detail: %v", err))
		_ = c.Error(err)
		return
	}
	err = h.repo.Database().Candidate().Delete(c, uint64(candidateId))
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Admin Handler DeleteCandidate] Delete candidate failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
	})
}
func (h *Handler) GetCandidateList(c *gin.Context) {
	candidates, err := h.repo.Database().Candidate().FindAll(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    candidates,
	})
}
