package models

type Post struct {
	ID      int `gorm:"primaryKey"`
	Title   string
	Content string
	UserID  int
	User    User `gorm:"foreignKey:UserID"`
}
