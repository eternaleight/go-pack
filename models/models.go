package models

import "time"

type Post struct {
	ID        int `gorm:"primaryKey"`
	Content   string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	AuthorID  int
	Author    User
}

type User struct {
	ID       int `gorm:"primaryKey"`
	Username string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Posts    []Post `gorm:"foreignKey:AuthorID"`
	Profile  Profile
}

type Profile struct {
	ID              int    `gorm:"primaryKey"`
	Bio             string `gorm:"size:1000"`
	ProfileImageUrl string
	UserID          int
	User            *User `gorm:"foreignKey:UserID"`
}
