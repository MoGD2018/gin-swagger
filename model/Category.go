package model

type Category struct {
	ID uint `json:"id" form:"id" gorm:"primary_key"`
	Name string `json:"name" form:"name" gorm:"type:varchar(50);not null;unique"`
	CreatedAt Time `json:"created_at" form:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time `json:"updated_at" form:"updated_at" gorm:"type:timestamp"`
}