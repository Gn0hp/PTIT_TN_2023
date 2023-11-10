package election

import (
	"PTIT_TN/internal/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (i impl) FindAll(c *gin.Context) ([]*entities.Election, error) {
	var elections []*entities.Election
	query := i.db.Gdb().
		WithContext(c).Model(&entities.Election{})
	err := query.Find(&elections).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Election{}, nil
		}
		i.logger.Error(fmt.Sprintf("[Election Repo] Find All Election failed, detail: %v", err))
		return nil, err
	}
	return elections, nil
}

func (i impl) FindById(c *gin.Context, id uint64) (*entities.Election, error) {
	elections := &entities.Election{}
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Election{}).
		Where("id = ?", id)
	err := query.First(elections).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[Election Repo] Find Election failed, detail: %v", err))
		return nil, err
	}
	return elections, nil
}

func (i impl) FindByOption(c *gin.Context, option entities.Election) ([]*entities.Election, error) {
	var elections []*entities.Election
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Election{}).
		Where(&option)
	err := query.Find(&elections).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.Election{}, nil
		}
		i.logger.Error(fmt.Sprintf("[Election Repo] Find Election failed, detail: %v", err))
		return nil, err
	}
	return elections, nil
}

func (i impl) FindOpenElection(c *gin.Context) (*entities.Election, error) {
	var election entities.Election
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Election{}).
		Where("status = ?", entities.ELECTION_STATUS_REGISTERING).
		Or("status = ?", entities.ELECTION_STATUS_OPENING)
	err := query.First(&election).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		i.logger.Error(fmt.Sprintf("[Election Repo] Find Open Election failed, detail: %v", err))
		return nil, err
	}
	return &election, nil
}

func (i impl) CountErsByElectionId(c *gin.Context, electionId uint64) (uint64, error) {
	var count uint64
	query := i.db.Gdb().
		WithContext(c).
		Raw("SELECT COUNT(*) FROM ptit_tn.role_elects re "+
			"LEFT JOIN ptit_tn.election_roles er on er.id = re.election_role_id "+
			"LEFT JOIN ptit_tn.election_candidates ec on ec.id = er.election_candidate_id "+
			"LEFT JOIN ptit_tn.elections e on e.id = ec.election_id "+
			"WHERE e.id = ?;", electionId)
	err := query.Scan(&count).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Election Repo] Count By Election Id failed, detail: %v", err))
		return 0, err
	}
	return count, nil
}

func (i impl) FindLatestElection(c *gin.Context) (uint64, error) {
	var id uint64
	query := i.db.Gdb().
		WithContext(c).
		Model(&entities.Election{}).
		Order("created_at desc").Select("id").Limit(1)
	err := query.First(&id).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[Election Repo] Find Latest Election failed, detail: %v", err))
		return id, err
	}
	return id, nil
}
