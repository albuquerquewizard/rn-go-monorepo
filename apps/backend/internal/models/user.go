package models

// User represents a user in the system
type User struct {
	BaseModel
	Username  string `json:"username" gorm:"uniqueIndex;not null;size:100"`
	Password  string `json:"-" gorm:"not null;size:255"` // "-" means this field won't be included in JSON
	FirstName string `json:"first_name" gorm:"size:100"`
	LastName  string `json:"last_name" gorm:"size:100"`
	IsActive  bool   `json:"is_active" gorm:"default:true"`
}

// TableName specifies the table name for User model
func (User) TableName() string {
	return "users"
}
