package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/learn-go/cmdb/pkg/setting"
	"github.com/learn-go/cmdb/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	// 测试
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})


	// 路由
	apiv1 := r.Group("/api/v1")
	{
		// tag路由
		apiv1.GET("/tags",v1.GetTags)
		apiv1.POST("/tags",v1.AddTag)
		apiv1.PUT("/tags/:id",v1.EditTag)
		apiv1.DELETE("/tags/:id",v1.DleteTag)
		// 文章路由
		// 获取文章列表
		apiv1.GET("/articles",v1.GetArticles)
		// 获取指定文章
		apiv1.GET("articles/:id",v1.GetArticle)
		// 增加文章
		apiv1.POST("/articles",v1.AddArticle)
		// 更新指定文章
		apiv1.PUT("/articles/:id",v1.EditArticle)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}


	return r
}
