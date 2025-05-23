package routers

import (
	"github.com/EDDYCJY/go-gin-example/middleware/jwt"
	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"github.com/EDDYCJY/go-gin-example/pkg/qrcode"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/upload"
	"github.com/EDDYCJY/go-gin-example/routers/api"
	v1 "github.com/EDDYCJY/go-gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	_ "github.com/EDDYCJY/go-gin-example/docs"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/auth", api.GetAuth)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/upload", api.UploadImage)

	//导出标签
	r.POST("/tags/export", v1.ExportTag)

	//导入标签
	r.POST("/tags/import", v1.ImportTag)

	// 生成二维码
	r.POST("/poster/generate", v1.GenerateArticlePoster)

	// 生成合并二维码
	r.POST("/poster/generate/merge", v1.GenerateArticlePosterMerged)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT()) // 添加中间件
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag) // id 是动态params
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag) // id 是动态params

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	}

	return r
}
