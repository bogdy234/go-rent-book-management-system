package config

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
	Trans    ut.Translator
)

func InitValidate() {
	en := en.New()
	Uni = ut.New(en, en)

	Trans, _ = Uni.GetTranslator("en")
	Validate = validator.New()
	en_translations.RegisterDefaultTranslations(Validate, Trans)
}
