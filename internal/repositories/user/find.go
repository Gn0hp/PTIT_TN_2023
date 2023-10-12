package user

import (
	"PTIT_TN/internal/entities"
	"PTIT_TN/internal/query_params"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (i impl) FindAll(c context.Context, params query_params.GetUserParams, lock bool) ([]*entities.User, error) {
	if !params.IsValid() {
		return nil, errors.New("Invalid request: invalid query params")
	}
	var res []*entities.User

	query := i.db.Gdb().WithContext(c).Model(entities.User{})
	query = filterUser(query, params)

	err := query.Find(&res).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[User Repo] Find User failed, details: %v", err))
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return []*entities.User{}, nil
		}
		return nil, err
	}
	return res, nil
}

func (i impl) FindById(id int) (*entities.User, error) {
	panic("implement me")
}
func (i impl) FindByCccd(c context.Context, cccd string) (*entities.User, error) {
	user := &entities.User{}

	query := i.db.Gdb().
		WithContext(c).
		Model(entities.User{}).
		Where("cccd_id = ?", cccd)
	err := query.First(user).Error
	if err != nil {
		i.logger.Error(fmt.Sprintf("[User Repo] Find User failed, detail: %v", err))

		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	return user, nil
}

func filterUser(db *gorm.DB, params query_params.GetUserParams) *gorm.DB {
	limit := params.CommonQueryParams.Limit
	if limit < 20 {
		limit = 20
	}
	if limit > 50 {
		limit = 50
	}
	db.Limit(int(limit))
	//db.Where()
	return db
}
