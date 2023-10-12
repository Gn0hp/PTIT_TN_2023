package user

import "PTIT_TN/internal/entities"

func (i impl) Create(user *entities.User) error {
	return i.db.Gdb().Model(entities.User{}).Create(user).Error
}

func (i impl) Update(id uint64, user *entities.User) error {
	panic("implement me")
}
