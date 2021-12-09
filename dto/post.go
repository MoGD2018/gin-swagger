package dto

type CreatePostRequest struct {
	CategoryID uint `json:"category_id" form:"category_id" binding:"required"`
	Title string `json:"title" form:"title" binding:"required,max=10"`
	HeadImg string `json:"head_img" form:"head_img"`
	Content string `json:"content" form:"content" binding:"required"`
}
