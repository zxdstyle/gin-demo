package routers

import (
	api "gin-demo/api/v1"
	"gin-demo/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	v1 := r.Group("/api/v1")
	{
		//获取标签列表
		v1.GET("/tags", api.GetTags)
		//新建标签
		v1.POST("/tags", api.AddTag)
		//更新指定标签
		v1.PUT("/tags/:id", api.EditTag)
		//删除指定标签
		v1.DELETE("/tags/:id", api.DeleteTag)

		//获取文章列表
		v1.GET("/articles", api.GetArticles)
		//获取指定文章
		v1.GET("/articles/:id", api.GetArticle)
		//新建文章
		v1.POST("/articles", api.AddArticle)
		//更新指定文章
		v1.PUT("/articles/:id", api.EditArticle)
		//删除指定文章
		v1.DELETE("/articles/:id", api.DeleteArticle)
	}
	return r
}