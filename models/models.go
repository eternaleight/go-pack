package models

import "time"

// Post model
type Post struct {
	ID        int `gorm:"primaryKey"`
	Content   string
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	AuthorID  int       `gorm:"column:authorId"`
	Author    User
}

// TableName overrides the table name
func (Post) TableName() string {
	return "Post"
}

// User model
type User struct {
	ID       int `gorm:"primaryKey"`
	Username string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Posts    []Post `gorm:"foreignKey:AuthorID"`
	Profile  Profile
}

// TableName overrides the table name
func (User) TableName() string {
	return "User"
}

// Profile model
type Profile struct {
	ID              int    `gorm:"primaryKey"`
	Bio             string `gorm:"size:1000"`
	ProfileImageUrl string `gorm:"column:profileImageUrl"`
	UserID          int    `gorm:"column:userId"`
	User            *User  `gorm:"foreignKey:UserID"`
}

// TableName overrides the table name
func (Profile) TableName() string {
	return "Profile"
}

// Product represents an item or content that can be sold on the platform.
type Product struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Description string
	Price       float64
	ImageURL    string
	VideoURL    string
	CreatedAt   time.Time
	SellerID    uint
	Seller      User
}

func (Product) TableName() string {
	return "Product"
}

// Purchase represents a transaction where a user buys a product.
type Purchase struct {
	ID              uint `gorm:"primarykey"`
	ProductID       uint
	BuyerID         uint
	PurchaseDate    time.Time
	StripePaymentID string
}

func (Purchase) TableName() string {
	return "Purchase"
}
