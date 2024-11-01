package initialize

import (
    "colab/app/colab-api/internal/global"
    "github.com/gin-gonic/gin/binding"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    "regexp"
)

func registerMobileValidator(v *validator.Validate) {
    _ = v.RegisterValidation("mobile", func(fld validator.FieldLevel) bool {
        mobile := fld.Field().String()
        ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|166|7[1-35-8]|9[189])\d{8}$`, mobile)
        return ok
    })
    _ = v.RegisterTranslation("mobile", *global.Translator,
        func(ut ut.Translator) error {
            return ut.Add("mobile", "{0} Invalid mobile number", true) // see universal-translator for details
        },
        func(ut ut.Translator, fe validator.FieldError) string {
            t, _ := ut.T("mobile", fe.Field())
            return t
        },
    )
}

func registerUsernameValidator(v *validator.Validate) {
    _ = v.RegisterValidation("username", func(fld validator.FieldLevel) bool {
        username := fld.Field().String()
        ok, _ := regexp.MatchString(`^[A-Za-z0-9_]+$`, username)
        return ok && len(username) <= 16
    })
    _ = v.RegisterTranslation("username", *global.Translator,
        func(ut ut.Translator) error {
            return ut.Add("username", "{0} Invalid username", true) // see universal-translator for details
        },
        func(ut ut.Translator, fe validator.FieldError) string {
            t, _ := ut.T("username", fe.Field())
            return t
        },
    )
}

func InitValidator() {
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        registerMobileValidator(v)
        registerUsernameValidator(v)
    }
}
