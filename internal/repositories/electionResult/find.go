package electionResult

import (
	"PTIT_TN/internal/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i impl) FindAll(c *gin.Context) ([]*entities.ElectionResult, error) {
	var electionResults []*entities.ElectionResult
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.ElectionResult{})
	err := query.Find(&electionResults).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.ElectionResult{}, nil
		}
		i.logger.Error(fmt.Sprintf("[ElectionResult Repo] Find All ElectionResult failed, detail: %v", err))
		return nil, err
	}
	return electionResults, nil
}
func (i impl) FindById(c *gin.Context, id uint64) (*entities.ElectionResult, error) {
	var electionResult *entities.ElectionResult
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.ElectionResult{}).
		Where("id = ?", id)
	err := query.First(electionResult).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[ElectionResult Repo] Find ElectionResult failed, detail: %v", err))
		return nil, err
	}
	return electionResult, nil
}
func (i impl) FindByElectionCandidateId(c *gin.Context, electionCandidateId uint64) ([]*entities.ElectionResult, error) {
	var electionResults []*entities.ElectionResult
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.ElectionResult{}).
		Where("election_candidate_id = ?", electionCandidateId)
	err := query.Find(&electionResults).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.ElectionResult{}, nil
		}
		i.logger.Error(fmt.Sprintf("[ElectionResult Repo] Find ElectionResult failed, detail: %v", err))
		return nil, err
	}
	return electionResults, nil
}
func (i impl) FindByOption(c *gin.Context, option entities.ElectionResult) ([]*entities.ElectionResult, error) {
	var electionResults []*entities.ElectionResult
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.ElectionResult{}).
		Where(&option)
	err := query.Find(&electionResults).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.ElectionResult{}, nil
		}
		i.logger.Error(fmt.Sprintf("[ElectionResult Repo] Find ElectionResult failed, detail: %v", err))
		return nil, err
	}
	return electionResults, nil
}
