package dto


type CreateCategoryRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

