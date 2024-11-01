package global

import (
    "colab/app/colab-api/conf"
    ut "github.com/go-playground/universal-translator"
)

var (
    Translator *ut.Translator
    Config     *conf.ApiServerConfig
)
