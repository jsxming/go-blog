package v1

import (
	"blog/internal/model"
	"blog/pkg/app"
	"blog/pkg/errorcode"
	"blog/pkg/response"
	"blog/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CreateArticleParams struct {
	Title       string `form:"title" binding:"required"`
	Content     string `form:"content" binding:"required"`
	Description string `form:"description" binding:"required"`
	Type      uint64 `form:"type"`
	Tag       uint64 `form:"tag"`
	UserId      uint64 `form:"userId"`
	CoverImgUrl string `form:"coverImgUrl"`
}

type UpdateArticleParams struct {
	Id uint64 `form:"id" binding:"required"`
	*CreateArticleParams
}

type ArticleSearchParams struct {
	Title     string `form:"title"`
	Type      uint64 `form:"type"`
	Tag       uint64 `form:"tag"`
	StartTime string `form:"startTime"`
	EndTime   string `form:"endTime"`
}

func CreateArticle(c *gin.Context) {
	params := CreateArticleParams{}
	errs := app.BindAndValid(c, &params)
	if errs != nil {
		response.Error(c, errorcode.InvalidParam.WithDetails(errs.Errors()...))
		return
	}
	u := app.GetTokenUser(c)
	article := model.Article{
		Model:       &model.Model{},
		Title:       params.Title,
		Content:     params.Content,
		Description: params.Description,
		TypeId:      params.Type,
		TagId:       params.Tag,
		UserId:      u.Id,
		CoverImgUrl: params.CoverImgUrl,
	}

	err := article.Create()
	if err != nil {
		response.Error(c, errorcode.CreateError)
		return
	}
	response.Success(c, article)
}


func FindOne(c *gin.Context){
	id,_:=strconv.Atoi(c.Query("id"))
	art:=model.Article{
		Model:&model.Model{
			Id: uint64(id),
		},
	}

	err:=art.FindOne()
	if err != nil {
		response.Error(c, errorcode.NotFound)
		return
	}
	response.Success(c, art)

}

func UpdateArticle(c *gin.Context) {
	params := UpdateArticleParams{}
	errs := app.BindAndValid(c, &params)
	if errs != nil {
		response.Error(c, errorcode.InvalidParam.WithDetails(errs.Errors()...))
		return
	}
	article := model.Article{
		Model: &model.Model{
			Id: params.Id,
			//UpdatedAt: model.JSONTime(time.Now()),
		},
		Title:       params.Title,
		Content:     params.Content,
		Description: params.Description,
		TypeId:      params.Type,
		TagId:       params.Tag,
		CoverImgUrl: params.CoverImgUrl,
	}
	err := article.Update()
	if err != nil {
		response.Error(c, errorcode.UpdateError)
		return
	}
	response.Success(c, nil)
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response.Error(c, errorcode.InvalidParam)
		return
	}
	article := model.Article{
		Model: &model.Model{
			Id:    util.UInt64(id),
			IsDel: 1,
		},
	}
	err := article.UpdateDel()
	if err != nil {
		response.Error(c, errorcode.DeleteError)
		return
	}
	response.Success(c, nil)
}

func SearchArticle(c *gin.Context) {
	params := ArticleSearchParams{}
	errs := app.BindAndValid(c, &params)
	if errs != nil {
		response.Error(c, errorcode.InvalidParam.WithDetails(errs.Errors()...))
		return
	}
	article := model.Article{
		Title:  params.Title,
		TypeId: params.Type,
		TagId:  params.Tag,
	}
	pager := util.NewPager(c)

	d, err := article.Search(pager, params.StartTime, params.EndTime)
	if err != nil {
		response.Error(c, errorcode.ServerError)
		return
	}
	response.Success(c, d)

}
