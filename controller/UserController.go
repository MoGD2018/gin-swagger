package controller

import (
	"fmt"
	"gin-swagger/dao"
	"gin-swagger/dto"
	"gin-swagger/model"
	"gin-swagger/response"
	"gin-swagger/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Helloworld 测试用例
// @Summary example
// @Schemes
// @Description swagger测试用例
// @Tags swagger的列子
// @Accept json
// @Produce json
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string	"false"
// @Router /api/v1/example/helloworld [get]
func Helloworld(g *gin.Context)  {
	g.JSON(http.StatusOK,"helloworld")
}


// Register 用户注册模块
// @Summary 用户注册接口
// @Schemes
// @Description 用户注册模块
// @Tags 用户注册
// @Accept application/json
// @Produce application/json
// @Param object query model.User false "查询参数"
// @Success 200 {string} string "注册成功"
// @Failure 400 {string} string "注册失败"
// @Router /api/auth/register [post]
func Register(ctx *gin.Context) {
	DB := dao.GetDB()
	//使用map获取请求参数
	var requestUser model.User
	ctx.Bind(&requestUser)
	//获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	fmt.Println(telephone, "手机号码长度", len(telephone))
	//name := ctx.PostForm("name")
	//telephone := ctx.PostForm("telephone")
	//password := ctx.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	// 如果名称没有传，给一个10位的随机字符串
	if len(name) ==0 {
		name = util.RandomString(10)
	}

	log.Println(name, telephone, password)
	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}

	// 密码加密
	hasepassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	// 创建用户
	newUser := model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hasepassword),
	}
	DB.Create(&newUser)

	// 返回结果
	response.Success(ctx, nil, "注册成功")
}

// Login 用户登陆模块
// @Summary 用户登陆接口
// @Schemes
// @Description 用户登陆模块
// @Tags 用户登陆
// @Accept application/json
// @Produce application/json
// @Param telephone query string true "手机号"
// @Param password query string true "密码"
// @Success 200 {string} string "登陆成功"
// @Failure 400 {string} string "登陆失败"
// @Router /api/auth/login [post]
func Login(ctx *gin.Context)  {
	DB := dao.GetDB()

	// 使用map获取请求参数
	var requestUser model.User
	ctx.Bind(&requestUser)
	//获取参数
	telephone := requestUser.Telephone
	password := requestUser.Password

	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password) ); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	// 发放token
	token, err := dao.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error ： %v", err)
		return
	}
	
	// 返回结果
	response.Success(ctx, gin.H{"token": token}, "注册成功")

}

// Info 获取用户信息模块
func Info(ctx  *gin.Context)  {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{ "user": dto.ToUserDto(user.(model.User)) }, "Token授权成功")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}