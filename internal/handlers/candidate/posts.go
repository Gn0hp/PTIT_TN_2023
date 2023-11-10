package candidate

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) CreatePost(c *gin.Context) {
	var postDto entities.PostDto
	if err := c.ShouldBindJSON(&postDto); err != nil {
		h.logger.Error(fmt.Sprintf("[Candidate Handler - Create Post] Bind Post failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	if !postDto.IsValid() {
		h.logger.Error("[Candidate Handler - Create Post] Invalid Post")
		_ = c.Error(errors.New("invalid post"))
		return
	}
	iEntity, _ := postDto.ToEntity()
	entity := iEntity.(*entities.Post)
	err := h.repo.Database().Post().Create(c, entity)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Candidate Handler - Create Post] Create Post failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    entity,
	})
}
func (h *Handler) DeletePost(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil || postID == 0 {
		h.logger.Error(fmt.Sprintf("[Candidate Handler - Delete Post] Invalid Post ID, detail: %v", err))
		_ = c.Error(err)
		return
	}
	err = h.repo.Database().Post().Delete(c, postID)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Candidate Handler - Delete Post] Delete Post failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
	})
}
func (h *Handler) UpdatePost(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil || postID == 0 {
		h.logger.Error(fmt.Sprintf("[Candidate Handler - Update Post] Invalid Post ID, detail: %v", err))
		_ = c.Error(err)
		return
	}
	var postDto entities.PostDto
	if err := c.ShouldBindJSON(&postDto); err != nil {
		h.logger.Error(fmt.Sprintf("[Candidate Handler - Update Post] Bind Post failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	fmt.Println(postDto)
	if !postDto.IsValid() {
		h.logger.Error("[Candidate Handler - Update Post] Invalid Post")
		_ = c.Error(errors.New("invalid post"))
		return
	}
	iEntity, _ := postDto.ToEntity()
	entity := iEntity.(*entities.Post)
	err = h.repo.Database().Post().Update(c, postID, entity)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Candidate Handler - Update Post] Update Post failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    entity,
	})
}
func (h *Handler) FindAllPost(c *gin.Context) {
	candidateId, _ := strconv.ParseInt(c.Query("candidate_id"), 10, 64)
	posts, err := h.repo.Database().Post().FindByCandidateId(c, uint64(candidateId))
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Candidate Handler - Find All Post] Find All Post failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    posts,
	})
}

func (h *Handler) WatchResult(c *gin.Context) {
	candidateId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || candidateId == 0 {
		h.logger.Error(fmt.Sprintf("[Candidate Handler - Watch Result] Invalid Candidate ID, detail: %v", err))
		_ = c.Error(err)
		return
	}
	roleElect, err := h.repo.Database().RoleElect().FindByCandidateId(c, uint64(candidateId))
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Candidate Handler - Watch Result] Find RoleElect failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	resDetail, err := h.repo.Database().BallotRoleElect().CountVoterByCandidateId(c, roleElect.ID)
	if err != nil {
		h.logger.Error(fmt.Sprintf("[Candidate Handler - Watch Result] Count Voter failed, detail: %v", err))
		_ = c.Error(err)
		return
	}
	utils.SetResponse(c, map[string]interface{}{
		"success": true,
		"data":    resDetail,
	})
}
