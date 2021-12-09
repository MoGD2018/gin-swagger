package middleware

import (
	"gin-swagger/dao"
	"gin-swagger/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware 中间件验证
// @Summary 中间件验证接口
// @Schemes
// @Description 中间件验证模块
// @Tags 中间件验证
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {string} string "登陆成功"
// @Failure 400 {string} string "登陆失败"
// @Router /auth/info [get]
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		}

		tokenString = tokenString[7:]

		token, claims, err := dao.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 验证通过，获取Claims中的userID
		userID := claims.UserID
		DB := dao.GetDB()
		var user model.User
		DB.First(&user, userID)


		// 判断用户是否存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 用户存在， 将user信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
