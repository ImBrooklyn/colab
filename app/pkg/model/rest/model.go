package rest

type Result struct {
    Code uint32 `json:"code"`
    Msg  string `json:"msg,omitempty"`
    Data any    `json:"data,omitempty"`
}

func Success(data any) *Result {
    return &Result{
        Code: 0,
        Msg:  "",
        Data: data,
    }
}

func Fail(code uint32, msg string) *Result {
    return &Result{
        Code: code,
        Msg:  msg,
        Data: nil,
    }
}

func InternalError() *Result {
    return Fail(0xFFFF, "internal error")
}
