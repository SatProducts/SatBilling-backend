package model

import (
	"gorm.io/gorm"
)

func (u *User) BeforeDelete(db *gorm.DB) error {
	if u.Permissions == ADMINISTRATOR {
		return DeleteAdminError
	}
	return nil
}