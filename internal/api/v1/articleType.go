package v1

import (
	"blog/global"
	"blog/internal/model"
	"blog/pkg/app"
	"blog/pkg/errorcode"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

type CreateArticleTypeParams struct {
	Name string `form:"name" binding:"required"`
}

type UpdateArticleTypeParams struct {
	Id   uint64 `form:"id" binding:"required"`
	Name string `form:"name" binding:"required"`
}

func CreateArticleType(c *gin.Context) {
	params := CreateArticleTypeParams{}
	errs := app.BindAndValid(c, &params)
	if errs != nil {
		response.Error(c, errorcode.InvalidParam.WithDetails(errs.Errors()...))
		return
	}

	aType := model.ArticleType{
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

func UpdateArticleType(c *gin.Context) {
	params := UpdateArticleTypeParams{}
	errs := app.BindAndValid(c, &params)
	if errs != nil {
		response.Error(c, errorcode.InvalidParam.WithDetails(errs.Errors()...))
		return
	}

	aType := model.ArticleType{
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

func QueryAllArticleTypes(c *gin.Context) {
	list, err := model.ArticleType{}.All()
	if err != nil {
		response.Error(c, errorcode.ServerError)
		return
	}
	response.Success(c, list)
}
