package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Name     string `json:"Name" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	IsAdmin  bool   `json:"is_admin" gorm:"type: bool"`
	Password string `json:"password" gorm:"type: varchar(255)"`
}

func (User) TableName() string {
	return "users"
}
