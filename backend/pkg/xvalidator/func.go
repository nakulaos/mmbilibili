package xvalidator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"unicode"
)

func registerPhoneValidation(validate *validator.Validate) {
	validate.RegisterValidation(
		"telephone",
		func(fl validator.FieldLevel) bool {
			regular := "^1\\d{2}\\d{8}$"
			reg := regexp.MustCompile(regular)
			return reg.MatchString(fl.Field().String())
		})

	validate.RegisterTranslation(
		"telephone",
		trans,
		func(ut ut.Translator) error {
			switch ut.Locale() {
			case "zh":
				return ut.Add("telephone", "{0}输入错误", false)
			case "en":
				return ut.Add("telephone", "{0} is invalid", false)
			case "fr":
				return ut.Add("telephone", "{0} est invalide", false)
			default:
				return nil
			}
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(fe.Tag(), fe.Field())
			return t
		})
}

func registerPasswordValidation(validate *validator.Validate) {
	validate.RegisterValidation(
		"password",
		func(fl validator.FieldLevel) bool {
			password := fl.Field().String()

			var hasUpper, hasLower, hasNumber, hasSpecial bool
			for _, char := range password {
				switch {
				case unicode.IsUpper(char):
					hasUpper = true
				case unicode.IsLower(char):
					hasLower = true
				case unicode.IsDigit(char):
					hasNumber = true
				case unicode.IsPunct(char) || unicode.IsSymbol(char):
					hasSpecial = true
				}
			}

			// 检查至少包含三种不同的字符类型
			typeCount := 0
			if hasUpper {
				typeCount++
			}
			if hasLower {
				typeCount++
			}
			if hasNumber {
				typeCount++
			}
			if hasSpecial {
				typeCount++
			}

			return typeCount >= 3
		})

	validate.RegisterTranslation(
		"password",
		trans,
		func(ut ut.Translator) error {
			switch ut.Locale() {
			case "zh":
				return ut.Add("password", "{0}必须包含大写字母、小写字母、数字、特殊字符中的三种", false)
			case "en":
				return ut.Add("password", "{0} must contain at least three of uppercase, lowercase, number, and special character", false)
			case "fr":
				return ut.Add("password", "{0} doit contenir au moins trois des éléments suivants : majuscule, minuscule, chiffre et caractère spécial", false)
			default:
				return nil
			}
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(fe.Tag(), fe.Field())
			return t
		})
}

func registerMinLenValidation(validate *validator.Validate) {
	validate.RegisterValidation(
		"min_len",
		func(fl validator.FieldLevel) bool {
			param, err := strconv.Atoi(fl.Param())
			if err != nil {
				return false
			}
			return len(fl.Field().String()) >= param
		})

	validate.RegisterTranslation(
		"min_len",
		trans,
		func(ut ut.Translator) error {
			switch ut.Locale() {
			case "zh":
				return ut.Add("min_len", "{0}长度不能少于{1}个字符", false)
			case "en":
				return ut.Add("min_len", "{0} must be at least {1} characters long", false)
			case "fr":
				return ut.Add("min_len", "{0} doit contenir au moins {1} caractères", false)
			default:
				return nil
			}
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(fe.Tag(), fe.Field(), fe.Param())
			return t
		})
}

func registerMaxLenValidation(validate *validator.Validate) {
	validate.RegisterValidation(
		"max_len",
		func(fl validator.FieldLevel) bool {
			param, err := strconv.Atoi(fl.Param())
			if err != nil {
				return false
			}
			return len(fl.Field().String()) <= param
		})

	validate.RegisterTranslation(
		"max_len",
		trans,
		func(ut ut.Translator) error {
			switch ut.Locale() {
			case "zh":
				return ut.Add("max_len", "{0}长度不能超过{1}个字符", false)
			case "en":
				return ut.Add("max_len", "{0} must be no more than {1} characters long", false)
			case "fr":
				return ut.Add("max_len", "{0} ne doit pas dépasser {1} caractères", false)
			default:
				return nil
			}
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(fe.Tag(), fe.Field(), fe.Param())
			return t
		})
}
