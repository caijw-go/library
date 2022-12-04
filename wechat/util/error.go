package util

type CommonError struct {
    ErrCode int64  `json:"errcode"`
    ErrMsg  string `json:"errmsg"`
}

func (e *CommonError) IsError() bool {
    return e != nil && e.ErrCode != 0
}

//NewError 主动构建一个错误，主动都是有error错误的
func NewError(errCode int64, msg string, err error) *CommonError {
    return &CommonError{
        ErrCode: errCode,
        ErrMsg:  msg + err.Error(),
    }
}
