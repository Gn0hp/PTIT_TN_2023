package candidate

import (
	"PTIT_TN/internal/entities"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i impl) FindByCccd(c *gin.Context, cccd string) (*entities.Candidate, error) {
	user := &entities.Candidate{}

	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Candidate{}).
		Where("cccd_id = ?", cccd)
	err := query.First(user).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Find Candidate failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New("candidate not found")
		}
		return nil, err
	}
	return user, nil

}
func (i impl) FindById(c *gin.Context, id uint64) (*entities.Candidate, error) {
	user := &entities.Candidate{}

	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Candidate{}).
		Where("id = ?", id)
	err := query.First(user).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Find Candidate failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New("candidate not found")
		}
		return nil, err
	}
	return user, nil
}
func (i impl) FindAll(c *gin.Context) ([]*entities.Candidate, error) {
	var candidates []*entities.Candidate
	query := i.db.Gdb().
		WithContext(c).Model(&entities.Candidate{})
	err := query.Find(&candidates).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Find All Candidate failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Candidate{}, nil
		}
		return nil, err
	}
	return candidates, nil
}
func (i impl) FindByOption(c *gin.Context, option entities.Candidate) ([]*entities.Candidate, error) {
	var candidates []*entities.Candidate
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Candidate{}).
		Where(&option)
	err := query.Find(&candidates).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Find Candidate failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Candidate{}, nil
		}
		return nil, err
	}
	return candidates, nil
}

func (i impl) FindCandidateByElectionRoleId(c *gin.Context, electionRoleId uint64) ([]*ViewCandidateDbFind, error) {
	var candidates []*ViewCandidateDbFind
	tx := i.db.Gdb().WithContext(c).Raw("select c.*, er.id as ElectionRoleId, er.name as erName, re.id as RoleElectId from ptit_tn.candidates c "+
		"join ptit_tn.role_elects re on re.candidate_id = c.id "+
		"join ptit_tn.election_roles er on er.id = re.election_role_id "+
		"where er.id = ?;", electionRoleId)
	err := tx.Scan(&candidates).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*ViewCandidateDbFind{}, nil
		}
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Find Candidate By ElectionRoleId Error: %v", err))
		return nil, err
	}
	return candidates, nil
}
func (i impl) DeactivateElectionCandidate(c *gin.Context, electionId uint64) error {
	var candidateIds []uint64
	tx := i.db.Gdb().WithContext(c).Begin()
	err := tx.Raw("select c.id from ptit_tn.candidates c "+
		"right join ptit_tn.role_elects re on re.candidate_id = c.id "+
		"join ptit_tn.election_roles er on er.id = re.election_role_id "+
		"join ptit_tn.election_candidates ec on ec.id =er.election_candidate_id "+
		"join ptit_tn.elections e on e.id = ec.election_id "+
		"where e.id=? and c.candidate_status=?;", electionId, entities.CANDIDATE_STATUS_VERIFIED).Scan(&candidateIds).Error
	if err != nil {
		tx.Rollback()
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil
		}
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Deactivate Election Candidate - Find candidateIds Error: %v", err))
		return err
	}
	if len(candidateIds) == 0 {
		tx.Rollback()
		return nil
	}
	err = tx.Model(&entities.Candidate{}).Where("id in ?", candidateIds).Update("candidate_status", entities.CANDIDATE_STATUS_INACTIVE).Error
	if err != nil {
		tx.Rollback()
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Deactivate Election Candidate - Update candidate status Error: %v", err))
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Deactivate Election Candidate - Commit Error: %v", err))
		return err
	}
	return nil
}

func (i impl) GetPendingCandidate(c *gin.Context) ([]*entities.Candidate, error) {
	var candidates []*entities.Candidate
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Candidate{}).
		Where("candidate_status = ?", entities.CANDIDATE_STATUS_PENDING)
	err := query.Find(&candidates).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[User Repo] Find Pending Candidate failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Candidate{}, nil
		}
		return nil, err
	}
	return candidates, nil
}

func (i impl) CheckRegistered(c *gin.Context, cccd string, roleElectId uint64) (bool, error) {
	var count int64
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Candidate{}).
		Where("cccd_id = ?", cccd).
		Where("role_elect_id = ?", roleElectId).
		Where("status != ?", entities.CANDIDATE_STATUS_INACTIVE).
		Count(&count)
	if query.Error != nil {
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Check Registered failed, detail: %v", query.Error))
		return false, query.Error
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (i impl) ViewResult(c *gin.Context, electionRoleId uint64) ([]*ViewResultDbFind, error) {
	var candidates []*ViewResultDbFind
	tx := i.db.Gdb().WithContext(c).Raw("select  c.* , er.id as ElectionRoleId, er.name as erName, re.id as RoleElectId, count(bre.role_elect_id) as NumVote from ptit_tn.candidates c "+
		"join ptit_tn.role_elects re on re.candidate_id = c.id "+
		"join ptit_tn.election_roles er on er.id = re.election_role_id "+
		"left join ptit_tn.ballot_role_elects bre on bre.role_elect_id = re.id "+
		"where er.id = ? "+
		"group by c.id order by NumVote desc;", electionRoleId)
	err := tx.Scan(&candidates).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*ViewResultDbFind{}, nil
		}
		i.logger.Error(fmt.Sprintf("[Candidate Repo] Find Candidate By ElectionRoleId Error: %v", err))
		return nil, err
	}
	return candidates, nil
}
