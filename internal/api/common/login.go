package common

import (
	"blog/internal/model"
	"blog/pkg/app"
	"blog/pkg/errorcode"
	"blog/pkg/response"
	"blog/pkg/util"
	"github.com/gin-gonic/gin"
)

type LoginDto struct {
	Account  string `form:"account" binding:"required,len=11"`
	Password string `form:"password" binding:"required,min=6,max=20"`
}

func Login(c *gin.Context) {
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
	}
	err := u.FindOne()
	if err != nil {
		response.Error(c, errorcode.NotFoundUser)
		return
	}

	if u.Password != util.EncodeMD5(params.Password) {
		response.Error(c, errorcode.InvalidPassword)
		return
	}

	tokenStr, _ := app.CreateToken(&u)

	response.Success(c, map[string]interface{}{
		"token": tokenStr,
		"user":  u,
	})

}
