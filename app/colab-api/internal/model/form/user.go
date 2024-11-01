package form

type SignupForm struct {
    Email    string `form:"email" json:"email" binding:"required,email"`
    Username string `form:"username" json:"username" binding:"required,username"`
    Nickname string `form:"nickname" json:"nickname" binding:"required"`

    Password string `form:"password" json:"password" binding:"required"`
    Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
    SmsCode  string `form:"sms_code" json:"sms_code" binding:"required"`
}

type SigninForm struct {
    SigninType int32  `form:"signin_type" json:"signin_type" binding:"required,oneof=0 1 2"`
    Identifier string `form:"identifier" json:"identifier" binding:"required"`
    Crypto     string `form:"crypto" json:"crypto" binding:"required"`
}
