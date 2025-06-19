package orm

import "time"

type UserLoginLog struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Token  string
	Time   time.Time `gorm:"autoCreateTime"`

	User User `gorm:"constraint:OnDelete:CASCADE"`
}

func (UserLoginLog) TableName() string {
	return "user_login_log"
}
