package routers

import (
	"gin.example/middleware/jwt"
	"gin.example/pkg/setting"
	"gin.example/routers/api"
	v1 "gin.example/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gin.example/docs"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	//r.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})

	/* swagger */
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")

	/* 使用jwt中间件 */
	apiv1.Use(jwt.JWT())
	{
		/* 标签方面 */

		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiv1.DELETE("tags/:id", v1.DeleteTag)

		/* 文章方面 */

		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 获取指定文章
		apiv1.GET("/article", v1.GetArticle)
		// 新建文章
		apiv1.POST("/articles", v1.AddArticle)
		// 更新指定文章
		apiv1.PUT("/article/:id", v1.EditArticle)
		// 删除指定文章
		apiv1.DELETE("/article/:id", v1.DeleteArticle)
	}

	return r
}
