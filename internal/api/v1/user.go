package v1

import (
	"blog/internal/model"
	"blog/pkg/app"
	"blog/pkg/errorcode"
	"blog/pkg/response"
	"blog/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginDto struct {
	Account  string `form:"account" binding:"required,len=11"`
	Password string `form:"password" binding:"required,min=6,max=20"`
	NickName string `form:"nickName"`
}

func CreateUser(c *gin.Context) {
	params := LoginDto{}
	errs := app.BindAndValid(c, &params)
	if errs != nil {
		response.Error(c, errorcode.InvalidParam.WithDetails(errs.Errors()...))
		return
	}
	u := model.User{
		Model: &model.Model{
			Id: util.UInt64(params.Account),
		},
		Password: util.EncodeMD5(params.Password),
		NickName: params.NickName,
	}

	err := u.FindOne()
	if err == nil {
		fmt.Println(err,u)
		//global.Logger.Errorf("svc.FindUserOne err:v%", err)
		response.Error(c, errorcode.ExistError)
		return
	}

	err = u.Create()
	if err != nil {
		//global.Logger.Errorf("svc.CreateUser err:v%", err)
		response.Error(c, errorcode.CreateError)
		return
	}
	response.Success(c, u)
}

func FindUserOne(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, errorcode.NotFound)
		return
	}
	u := model.User{
		Model: &model.Model{
			Id: util.UInt64(id),
		},
	}
	err := u.FindOne()
	if err != nil {
		response.Error(c, errorcode.NotFound)
		return
	}

	response.Success(c, u)
}

func UpdateUser(c *gin.Context) {
	u := model.User{}
	errs := app.BindAndValid(c, &u)
	if errs != nil {
		response.Error(c, errorcode.InvalidParam.WithDetails(errs.Errors()...))
		return
	}

	err := u.FindOne()
	if err != nil {
		//global.Logger.Errorf("svc.FindUserOne err:v%", err)
		response.Error(c, errorcode.NotFound)
		return
	}

	err = u.Update()
	if err != nil {
		//global.Logger.Errorf("app.dao.update errs :v%", errs)
		response.Error(c, errorcode.UpdateError)
		return
	}

	response.Success(c, u)
}
