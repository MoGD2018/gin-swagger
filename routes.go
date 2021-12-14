package main

import (
	"gin-swagger/controller"
	"gin-swagger/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Cors())
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld",controller.Helloworld)
		}
	}
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware() , controller.Info)

	categoryRoutes := r.Group("/categories")
	{
		categoryController := controller.NewCategoryController()
		categoryRoutes.POST("",categoryController.Create)
		categoryRoutes.PUT("/:id", categoryController.Update)
		categoryRoutes.GET("/:id", categoryController.Show)
		categoryRoutes.DELETE("/:id", categoryController.Delete)
	}

	postRoutes := r.Group("/posts")
	{
		postRoutes.Use(middleware.AuthMiddleware())
		postController := controller.NewPostController()
		postRoutes.POST("",postController.Create)
		postRoutes.PUT("/:id", postController.Update)
		postRoutes.GET("/:id", postController.Show)
		postRoutes.DELETE("/:id", postController.Delete)
		postRoutes.DELETE("page/list", postController.PageList)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
