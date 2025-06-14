package orm

type Amplua struct {
	AmpluaID int `gorm:"primaryKey"`
	Name     string
	Players  []Player
}

// type Amplua struct {
// 	AmpluaID int    `gorm:"primaryKey" json:"amplua_id"`
// 	Name     string `gorm:"size:45" json:"name"`

// 	// Players []Player `gorm:"foreignKey:AmpluaID"`
// }
