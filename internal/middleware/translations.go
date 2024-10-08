package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var trans  ut.Translator
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		local := c.GetHeader("local")
		trans, err := uni.GetTranslator(local)
		if err != true {
			trans, _ = uni.GetTranslator("en")
		}
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch local {
			case "zh":
				zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				en_translations.RegisterDefaultTranslations(v, trans)
				break
			default:
				zh_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}
