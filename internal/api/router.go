package api

import (
	"blog/global"
	"blog/internal/api/common"
	v1 "blog/internal/api/v1"
	"blog/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(middleware.Translations())
	r.Use(middleware.Cors())
	r.Use(middleware.AccessLog())
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "test close http server")
	})
	apiV1 := r.Group("/blog/v1")
	apiV1.Use(middleware.CheckToken())
	{
		apiV1.POST("/upload", common.UploadFile)
		apiV1.GET("/qiniutoken", common.CreateQinuToken)
		userApi := apiV1.Group("/user")
		//userApi.POST("/create", v1.CreateUser)
		userApi.POST("/update", v1.UpdateUser)
		userApi.GET(":id", v1.FindUserOne)

		ArticleApi := apiV1.Group("/article")
		ArticleApi.POST("/create", v1.CreateArticle)
		ArticleApi.POST("/update", v1.UpdateArticle)
		ArticleApi.POST("/delete/:id", v1.DeleteArticle)
		ArticleApi.GET("/search", v1.SearchArticle)

		ArticleApi.POST("/type/create", v1.CreateArticleType)
		ArticleApi.GET("/type/all", v1.QueryAllArticleTypes)
		ArticleApi.POST("/type/update", v1.UpdateArticleType)

		ArticleApi.POST("/tag/create", v1.CreateArticleTag)
		ArticleApi.GET("/tag/all", v1.QueryAllArticleTags)
		ArticleApi.POST("/tag/update", v1.UpdateArticleTag)

	}

	apiNoToken := r.Group("/blog/t")
	{
		apiNoToken.POST("/login", common.Login)
		apiNoToken.POST("/register", v1.CreateUser)
		apiNoToken.GET("/articles", v1.SearchArticle)
		apiNoToken.GET("/article", v1.FindOne)
		apiNoToken.GET("/types", v1.QueryAllArticleTypes)
		//apiNoToken.GET("/article", v1.FindOne)
	}

	return r
}
