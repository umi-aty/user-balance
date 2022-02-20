package entities

type User struct {
	ID       uint   `gorm:"primary_key:auto_increment" json:"id"`
	Username string `gorm:"type:varchar(255)" json:"username"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
