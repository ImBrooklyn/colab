package form

type SendSmsForm struct {
    Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
    Biz    string `form:"biz" json:"biz" binding:"required"`
}
