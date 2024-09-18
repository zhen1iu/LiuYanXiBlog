package routers

import (
	"LiuYanXiBlog/internal/middleware"
	v1 "LiuYanXiBlog/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	//url := ginSwagger.URL("docs/swagger.json")
	r := gin.New()
	r.Use(middleware.CoresMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	// Swagger API 文档路由go

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("tags", v1.NewTag().Create)
		apiv1.DELETE("tags/:id", v1.NewTag().Delete)
		apiv1.PUT("tags/:id", v1.NewTag().Update)
		apiv1.PATCH("tags/:id/state", v1.NewTag().Update)
		apiv1.GET("tags", v1.NewTag().List)

		apiv1.POST("articles", v1.NewArticles().Create)
		apiv1.DELETE("articles/:id", v1.NewArticles().Delete)
		apiv1.PUT("articles/:id", v1.NewArticles().Update)
		apiv1.PATCH("articles/:id/state", v1.NewArticles().Update)
		apiv1.GET("articles", v1.NewArticles().List)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
