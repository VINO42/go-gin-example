package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vino42/go-gin-example/pkg/setting"
	"github.com/vino42/go-gin-example/routers/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("tag/:id", v1.GetTag)
		apiV1.GET("tags", v1.GetTags)
		apiV1.POST("addTag", v1.AddTag)
		apiV1.POST("editTag", v1.EditTag)
		apiV1.DELETE("tag/:id", v1.DeleteTag)

		apiV1.GET("/article/:id", v1.GetArticle)
		apiV1.GET("/articles/:id", v1.GetArticles)
		apiV1.POST("/articles", v1.AddArticle)
		apiV1.PUT("/articles/:id", v1.EditArticle)
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "testrouter",
		})

	})
	return r
}
