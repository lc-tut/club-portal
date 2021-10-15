package admins

import (
	models "github.com/lc-tut/club-portal/models/users"
	"github.com/lc-tut/club-portal/utils"
	"gorm.io/gorm"
)

type UserArgs struct {
	Email    string
	Name     string
	ClubUUID string
}

type AdminClubRepo interface {
	UpdateSpecifiedClub(userUUID string, args UserArgs) error
}

func (r *AdminRepository) UpdateSpecifiedClub(userUUID string, args UserArgs) error {
	user := models.GeneralUser{
		UserUUID: userUUID,
		Email:    args.Email,
		Name:     args.Name,
		ClubUUID: utils.StringToNullString(args.ClubUUID),
	}
	tx := r.db.Model(&user).Where("user_uuid = ?", userUUID).Updates(user)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	} else if tx.RowsAffected == 0 {
		err := gorm.ErrRecordNotFound
		r.logger.Info(err.Error())
		return err
	} else {
		return nil
	}
}
