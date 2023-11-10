package ballotRoleElect

import (
	"PTIT_TN/internal/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i impl) FindAll(c *gin.Context) ([]*entities.BallotRoleElect, error) {
	var result []*entities.BallotRoleElect
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.BallotRoleElect{})
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[BallotRoleElect Repo] Find BallotRoleElect failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindById(c *gin.Context, id uint64) (*entities.BallotRoleElect, error) {
	result := &entities.BallotRoleElect{}
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.BallotRoleElect{}).
		Where("id = ?", id)
	err := query.First(result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[BallotRoleElect Repo] Find BallotRoleElect failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindByOption(c *gin.Context, option entities.BallotRoleElect) ([]*entities.BallotRoleElect, error) {
	var result []*entities.BallotRoleElect
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.BallotRoleElect{}).
		Where(&option)
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[BallotRoleElect Repo] Find BallotRoleElect failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}

func (i impl) FindByBallotId(c *gin.Context, ballotId uint64) ([]*StatByBallotResponse, error) {
	var result []*StatByBallotResponse
	tx := i.db.Gdb().WithContext(c).Raw(
		"select v.full_name as voter_name, er.name as role_name, c.full_name as candidate_name, c.cccd_id as candidate_cccd from ptit_tn.voters v "+
			"join ptit_tn.ballots b on v.id = b.voter_id "+
			"join ptit_tn.ballot_role_elects bre on bre.ballot_id = b.id "+
			"join ptit_tn.role_elects re on bre.role_elect_id = re.id "+
			"join ptit_tn.candidates c on c.id = re.candidate_id "+
			"join ptit_tn.election_roles er on er.id = re.election_role_id "+
			"where b.id=?;", ballotId)
	err := tx.Scan(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[BallotRoleElect Repo] Find BallotRoleElect failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}
func (i impl) FindByRoleElectId(c *gin.Context, roleElectId uint64) ([]*entities.BallotRoleElect, error) {
	var result []*entities.BallotRoleElect
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.BallotRoleElect{}).
		Where("role_elect_id = ?", roleElectId)
	err := query.Find(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[BallotRoleElect Repo] Find BallotRoleElect failed, detail: %v", err))
		return nil, err
	}
	return result, nil
}

func (i impl) CountVoterByCandidateId(c *gin.Context, roleElectId uint64) (int64, error) {
	var result int64
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.BallotRoleElect{}).
		Where("role_elect_id = ?", roleElectId).
		Count(&result)
	err := query.Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return 0, nil
		}
		i.logger.Error(fmt.Sprintf("[BallotRoleElect Repo] Count Voter By Candidate Id failed, detail: %v", err))
		return 0, err
	}
	return result, nil
}
func (i impl) StatVoterByRoleElectId(c *gin.Context, roleElectId uint64) ([]*StatBreByCandidateDbFind, error) {
	var result []*StatBreByCandidateDbFind
	tx := i.db.Gdb().WithContext(c).Raw(
		"select v.id as voter_id , v.full_name as voter_name , v.cccd_id, b.id as ballot_id, re.candidate_id, bre.created_at as vote_time from ptit_tn.voters v "+
			"join ptit_tn.ballots b on b.voter_id = v.id "+
			"join ptit_tn.ballot_role_elects bre on bre.ballot_id = b.id "+
			"join ptit_tn.role_elects re on re.id = bre.role_elect_id "+
			"where re.id = ?;", roleElectId)
	err := tx.Scan(&result).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*StatBreByCandidateDbFind{}, nil
		}
		i.logger.Error(fmt.Sprintf("[BallotRoleElect Repo] Stat Voter By Role Elect Id Error: %v", err))
		return nil, err
	}
	return result, nil
}
