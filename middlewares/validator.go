package middlewares

import (
	"errors"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

var (
	uni *ut.UniversalTranslator
)

// CustomValidator ..
type CustomValidator struct {
	Validator *validator.Validate
	Trans     ut.Translator
}

//Validate is func test valid object
func (cv *CustomValidator) Validate(i interface{}) error {

	err := cv.Validator.Struct(i)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			return errors.New(e.Translate(cv.Trans))
		}
	}
	return nil
}

//InitCustomValidator is func init CustomValid
func InitCustomValidator() *CustomValidator {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "Vui lòng nhập giá trị {0}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "Vui lòng nhập giá trị {0} lớn hơn {1} (kí tự/số).", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field(), fe.Param())

		return t
	})

	validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "Vui lòng nhập giá trị {0} nhỏ hơn {1} (kí tự/số).", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field(), fe.Param())

		return t
	})

	validate.RegisterTranslation("excludesall", trans, func(ut ut.Translator) error {
		return ut.Add("excludesall", "Vui lòng nhập giá trị {0} không tồn tại kí tự {1}.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("excludesall", fe.Field(), fe.Param())

		return t
	})

	validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "Giá trị {0} không hợp lệ.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())

		return t
	})
	return &CustomValidator{Validator: validate, Trans: trans}
}
