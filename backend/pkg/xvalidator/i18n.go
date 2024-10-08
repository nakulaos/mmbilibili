package xvalidator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fr"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	frTranslations "github.com/go-playground/validator/v10/translations/fr"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var trans ut.Translator

const labelName = "label"

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	// 初始化三个语言的支持
	uni := ut.New(en.New(), zh.New(), fr.New())

	// 根据需要选择语言，这里以中文为例
	lang := "zh" // 你可以根据实际情况修改为 "en" 或 "fr"
	trans, _ = uni.GetTranslator(lang)

	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get(labelName)
		if label == "" {
			return field.Name
		}
		return label
	})

	customError(Validate)

	// 注册默认翻译
	switch lang {
	case "zh":
		_ = zhTranslations.RegisterDefaultTranslations(Validate, trans)
	case "en":
		_ = enTranslations.RegisterDefaultTranslations(Validate, trans)
	case "fr":
		_ = frTranslations.RegisterDefaultTranslations(Validate, trans)
	}
}

// Translate 校验异常
func Translate(errs error) string {
	errors, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs.Error()
	}
	return errors[0].Translate(trans)
}
func customError(validate *validator.Validate) {
	registerPhoneValidation(validate)
	registerPasswordValidation(validate)
	registerMinLenValidation(validate)
	registerMaxLenValidation(validate)
}
