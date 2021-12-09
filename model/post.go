package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Post 文章结构体
type Post struct {
	ID uuid.UUID `json:"id" form:"id" gorm:"type:char(36);primary_key"`
	UserID uint `json:"user_id" form:"user_id" gorm:"not null"`
	CategoryID uint `json:"category_id" form:"category_id" gorm:"not null"`
	Category *Category
	Title string `json:"title" form:"title" gorm:"type:varchar(50);not null"`
	HeadImg string `json:"head_img" form:"head_img"`
	Content string `json:"content" form:"content" gorm:"type:text;not null"`
	CreatedAt Time `json:"created_at" form:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time `json:"updated_at" form:"updated_at" gorm:"type:timestamp"`
}

func (post *Post) BeforeCreate(db *gorm.DB) error {
	post.ID = uuid.NewV4()
	return nil
}