package controller

import (
	"gin-swagger/dao"
	"gin-swagger/dto"
	"gin-swagger/model"
	"gin-swagger/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := dao.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{DB:db}
}

// Create 创建类别模块
// @Summary 创建类别接口
// @Schemes
// @Description 创建类别模块
// @Tags 创建类别
// @Accept application/json
// @Produce application/json
// @Param object query dto.CreateCategoryRequest false "创建参数"
// @Success 200 {string} string "分类创建成功"
// @Failure 400 {string} string "数据验证错误，分类名称必填"
// @Router /categories [post]
func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory dto.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil,"数据验证错误，分类名称必填")
		return
	}

	category := model.Category{Name: requestCategory.Name}
	log.Println(category)
	c.DB.Create(&category)
	response.Success(ctx, gin.H{"category": requestCategory}, "分类创建成功")
}

// Update 更新类别模块
// @Summary 更新类别接口
// @Schemes
// @Description 更新类别模块
// @Tags 更新类别
// @Accept application/json
// @Produce application/json
// @Param id path integer true "类别ID"
// @Param object query dto.CreateCategoryRequest false "查询参数"
// @Success 200 {string} string "修改分类成功"
// @Failure 400 {string} string "分类不存在"
// @Router /categories/{id} [put]
func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory dto.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil,"数据验证错误，分类名称必填")
		return
	}

	// 获取path中的参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var updateCategory model.Category
	err := c.DB.First(&updateCategory, categoryID).Error
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	// 更新分类
	// map, struct, name value
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	response.Success(ctx, gin.H{"category": updateCategory}, "修改分类成功")
}

// Show 查看类别模块
// @Summary 查看类别接口
// @Schemes
// @Description 查看类别模块
// @Tags 查看类别
// @Accept application/json
// @Produce application/json
// @Param id path integer true "类别ID"
// @Success 200 {string} string "分类查看成功"
// @Failure 400 {string} string "分类不存在"
// @Router /categories/{id} [get]
func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var category model.Category
	err := c.DB.First(&category, categoryID).Error
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{"category": category}, "查询成功")
}

// Delete 删除类别模块
// @Summary 删除类别接口
// @Schemes
// @Description 删除类别模块
// @Tags 删除类别
// @Accept application/json
// @Produce application/json
// @Param id path integer true "类别ID"
// @Success 200 {string} string "分类删除成功"
// @Failure 400 {string} string "删除失败，请重试"
// @Router /categories/{id} [delete]
func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryID, _ := strconv.Atoi(ctx.Params.ByName("id"))

	err := c.DB.Delete(model.Category{}, categoryID).Error
	if err != nil {
		response.Fail(ctx, nil, "删除失败，请重试")
		return
	}
	response.Success(ctx, nil, "删除成功")
}

