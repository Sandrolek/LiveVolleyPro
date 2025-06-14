package orm

type User struct {
	UserID   int `gorm:"primaryKey"`
	Name     string
	Password string
	Email    string
	Teams    []Team
}

// type User struct {
// 	gorm.Model
// 	UserID   int    `gorm:"primaryKey" json:"user_id"`
// 	Name     string `gorm:"size:45" json:"name"`
// 	Password string `gorm:"size:255" json:"password"`
// 	Email    string `gorm:"size:255" json:"email"`

// 	// Teams []Team `gorm:"foreignKey:UserID"`
// 	// Teams []Team `gorm:"foreignKey:UserID" json:"teams"`
// }

type CreateUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
