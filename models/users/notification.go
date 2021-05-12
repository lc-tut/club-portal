package users

type Notification struct {
	NotificationID string `gorm:"type:char(36);not null;primaryKey"`
	Title          string `gorm:"type:varchar(63);not null"`
}

type NotificationContent struct {
	NotificationID string `gorm:"type:char(36);not null;primaryKey"`
	Title          string `gorm:"->"`
	Content        string `gorm:"type:mediumtext;not null"`
}

type UserNotification struct {
	UserUUID       string `gorm:"type:char(36);not null;primaryKey"`
	NotificationID string `gorm:"type:char(36);not null;primaryKey"`
	UnreadCount    uint32
}
