package orm

type Team struct {
	TeamID  int `gorm:"primaryKey"`
	Name    string
	UserID  int
	User    User
	Players []Player
}
