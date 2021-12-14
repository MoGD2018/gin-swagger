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

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

// Create 创建文章模块
// @Summary 创建文章接口
// @Schemes
// @Description 创建文章模块
// @Tags 创建文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query dto.CreatePostRequest false "创建参数"
// @Success 200 {string} string "创建成功"
// @Failure 400 {string} string "数据验证错误"
// @Router /posts [post]
func (p PostController) Create(ctx *gin.Context) {
	var requestPost dto.CreatePostRequest
	// 数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		response.Fail(ctx, nil,"数据验证错误")
		return
	}

	// 获取登陆用户user
	user, _ := ctx.Get("user")

	// 创建post
	post := model.Post{
		UserID: user.(model.User).ID,
		CategoryID: requestPost.CategoryID,
		Title: requestPost.Title,
		HeadImg: requestPost.HeadImg,
		Content: requestPost.Content,
	}

	// 插入数据
	if err := p.DB.Create(&post).Error; err != nil {
		log.Println(err)
		return
	}

	response.Success(ctx, nil, "创建文章成功")
}

// Update 更新文章模块
// @Summary 更新文章接口
// @Schemes
// @Description 更新文章模块
// @Tags 更新文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path integer true "文章ID"
// @Param object query dto.CreatePostRequest false "查询参数"
// @Success 200 {string} string "修改成功"
// @Failure 400 {string} string "文章不存在"
// @Router /posts/{id} [put]
func (p PostController) Update(ctx *gin.Context) {
	var requestPost dto.CreatePostRequest
	// 数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Println(err)
		response.Fail(ctx, nil,"数据验证错误")
		return
	}


	// 获取path 中的id
	postID := ctx.Params.ByName("id")

	var post model.Post
	if err := p.DB.Where("id = ?", postID).First(&post).Error; err !=nil {
		response.Fail(ctx, nil,"文章不存在")
		return
	}

	// 判断当前用户是否为文章作者
	// 获取登陆用户user
	user, _ := ctx.Get("user")
	userID := user.(model.User).ID
	if userID != post.UserID {
		response.Fail(ctx, nil,"非文章作者，请勿操作")
		return
	}

	// 更新文章
	err := p.DB.Model(&post).
		Updates(model.Post{
			CategoryID: requestPost.CategoryID,
			Title: requestPost.Title,
			HeadImg: requestPost.HeadImg,
			Content: requestPost.Content,
		}).Error
	if  err != nil {
		response.Fail(ctx, nil,"文章更新失败")
		return
	}

	response.Success(ctx, gin.H{"post": post}, "更新成功")
}

// Show 查看文章模块
// @Summary 查看文章接口
// @Schemes
// @Description 查看文章模块
// @Tags 查看文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path integer true "文章ID"
// @Success 200 {string} string "查看成功"
// @Failure 400 {string} string "文章不存在"
// @Router /posts/{id} [get]
func (p PostController) Show(ctx *gin.Context) {
	// 获取path 中的id
	postID := ctx.Params.ByName("id")

	var post model.Post
	if err := p.DB.Preload("Category").Where("id = ?", postID).First(&post).Error; err !=nil {
		response.Fail(ctx, nil,"文章不存在")
		return
	}

	response.Success(ctx, gin.H{"post": post}, "查看文章成功")
}

// Delete 删除文章模块
// @Summary 删除文章接口
// @Schemes
// @Description 删除文章模块
// @Tags 删除文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path integer true "文章ID"
// @Success 200 {string} string "文章删除成功"
// @Failure 400 {string} string "删除失败"
// @Router /posts/{id} [delete]
func (p PostController) Delete(ctx *gin.Context) {
	// 获取path 中的id
	postID := ctx.Params.ByName("id")

	var post model.Post
	if err := p.DB.Where("id = ?", postID).First(&post).Error; err !=nil {
		response.Fail(ctx, nil,"文章不存在")
		return
	}

	// 判断当前用户是否为文章作者
	// 获取登陆用户user
	user, _ := ctx.Get("user")
	userID := user.(model.User).ID
	if userID != post.UserID {
		response.Fail(ctx, nil,"非文章作者，请勿操作")
		return
	}

	if err := p.DB.Delete(&post).Error; err != nil {
		response.Fail(ctx, nil,"文章删除失败")
		return
	}
	response.Success(ctx, nil, "删除文章成功")
}

// PageList 列出文章模块
// @Summary 列出文章接口
// @Schemes
// @Description 列出文章模块
// @Tags 列出文章
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query model.Post false "查询参数"
// @Success 200 {string} string "成功"
// @Failure 400 {string} string "失败"
// @Router /posts/{id} [delete]
func (p PostController) PageList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum","1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize","20"))

	// 分页
	var posts []model.Post
	p.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts);

	// 前端渲染分页需要知道总数
	var total int64
	p.DB.Model(model.Post{}).Count(&total)

	response.Success(ctx, gin.H{"data": posts, "total": total}, "成功")
}

func NewPostController() IPostController {
	db := dao.GetDB()
	db.AutoMigrate(model.Post{})
	return PostController{DB: db}
}