package users

type FavoriteClub struct {
	UserUUID string `gorm:"type:char(36);not null;primaryKey"`
	ClubUUID string `gorm:"type:char(36);not null;primaryKey"`
}
