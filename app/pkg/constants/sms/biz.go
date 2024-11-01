package sms

import (
    "colab/app/pkg/constants/keys"
    "fmt"
)

type Biz string

func (b *Biz) Key(mobile string) string {
    return fmt.Sprintf("%s-%s-%s", keys.SmsCodeKeyPrefix, *b, mobile)
}

const (
    BizSignup Biz = "signup"
    BizSignin Biz = "signin"
)
