package main

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
)

type Person struct {
	Age     int    `form:"age" validate:"required,gt=10"`
	Name    string `form:"name" validate:"required"`
	Address string `form:"address" validate:"required"`
}

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func multiLangBindingHandler(c *gin.Context) {
	local := c.DefaultQuery("local", "en")
	tans, _ := uni.GetTranslator(local)
	switch local {
	case "en":
		en_translations.RegisterDefaultTranslations(validate, tans)
	case "zh":
		zh_translations.RegisterDefaultTranslations(validate, tans)
	default:
		en_translations.RegisterDefaultTranslations(validate, tans)
	}
	var person Person

	if err := c.ShouldBind(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := validate.Struct(person); err != nil {

		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			// can translate each error one at a time.
			sliceErrs = append(sliceErrs, e.Translate(tans))
		}
		c.JSON(http.StatusBadRequest, gin.H{
			// translate all error at once
			// returns a map with key = namespace & value = translated error
			// NOTICE: 2 errors are returned and you'll see something surprising
			// translations are i18n aware!!!!
			// eg. '10 characters' vs '1 character'
			"message":   errs.Translate(tans),
			"sliceErrs": sliceErrs,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"personInfo": person,
	})
}

func main() {
	zh := zh2.New()
	en := en2.New()
	uni = ut.New(en, zh)

	validate = validator.New()
	router := gin.Default()
	router.GET("/testMultiLangBinding", multiLangBindingHandler)
	router.Run(":9999")
}
