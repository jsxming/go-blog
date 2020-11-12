package v1

import (
	"blog/global"
	"blog/internal/model"
	"blog/pkg/app"
	"blog/pkg/errorcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

type CreateArticleTagParams struct {
	Name string `form:"name" binding:"required"`
}

type UpdateArticleTagParams struct {
	Id   uint64 `form:"id" binding:"required"`
	Name string `form:"name" binding:"required"`
}

func CreateArticleTag(c *gin.Context) {
	params := CreateArticleTagParams{}
	errs := app.BindAndValid(c, &params)
	if errs != nil {
		response.Error(c, errorcode.InvalidParam.WithDetails(errs.Errors()...))
		return
	}

	aType := model.ArticleTag{
		Name: params.Name,
	}

	err := aType.Create()
	if err != nil {
		global.Logger.Errorf("svc.CreateArticleType err:v%", err)
		response.Error(c, errorcode.CreateError)
		return
	}

	response.Success(c, aType)

}

func UpdateArticleTag(c *gin.Context) {
	params := UpdateArticleTagParams{}
	errs := app.BindAndValid(c, &params)
	if errs != nil {
		response.Error(c, errorcode.InvalidParam.WithDetails(errs.Errors()...))
		return
	}

	aType := model.ArticleTag{
		Id:   params.Id,
		Name: params.Name,
	}

	err := aType.FindOne()
	if err != nil {
		response.Error(c, errorcode.NotFound)
		return
	}
	aType.Name = params.Name

	err = aType.Update()
	if err != nil {
		response.Error(c, errorcode.UpdateError)
		return
	}

	response.Success(c, aType)
}

func QueryAllArticleTags(c *gin.Context) {
	list, err := model.ArticleTag{}.All()
	if err != nil {
		response.Error(c, errorcode.ServerError)
		return
	}
	response.Success(c, list)
}
