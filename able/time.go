package able

import (
    "database/sql/driver"
    "github.com/caijw-go/library/consts"
    "time"
)

//这个文件重新定义了time，主要是为了解决当time作为字段需要被json序列化的情况
// 默认的time RFC3339，使用本文件的结构体可以保证json序列化是可读的
// 当没赋值时，json序列化以后会变成null，所以使用本文件的结构体本身就是指针类型，不需要额外加*了

type jsonTime time.Time
type jsonDate time.Time

type JsonTime = *jsonTime
type JsonDate = *jsonDate

func marshalJSON[T jsonTime | jsonDate](format string, t T) ([]byte, error) {
    tt := time.Time(t)
    if &tt == nil || tt.IsZero() {
        return []byte("null"), nil
    }
    b := make([]byte, 0, len(format)+2)
    b = append(b, '"')
    b = tt.AppendFormat(b, format)
    b = append(b, '"')
    return b, nil
}

func (t *jsonTime) UnmarshalJSON(data []byte) (err error) {
    now, err := time.ParseInLocation(`"`+consts.TimeFormat+`"`, string(data), time.Local)
    *t = jsonTime(now)
    return
}

func (t jsonTime) MarshalJSON() ([]byte, error) {
    tt := time.Time(t)
    if &tt == nil || tt.IsZero() {
        return []byte("null"), nil
    }
    b := make([]byte, 0, len(consts.TimeFormat)+2)
    b = append(b, '"')
    b = tt.AppendFormat(b, consts.TimeFormat)
    b = append(b, '"')
    return b, nil
}

func (t jsonTime) String() string {
    return time.Time(t).Format(consts.TimeFormat)
}

func (t jsonTime) Value() (driver.Value, error) {
    return t.String(), nil
}

func (t *jsonDate) UnmarshalJSON(data []byte) (err error) {
    now, err := time.ParseInLocation(`"`+consts.DateFormat+`"`, string(data), time.Local)
    *t = jsonDate(now)
    return
}

func (t jsonDate) MarshalJSON() ([]byte, error) {
    tt := time.Time(t)
    if &tt == nil || tt.IsZero() {
        return []byte("null"), nil
    }
    b := make([]byte, 0, len(consts.DateFormat)+2)
    b = append(b, '"')
    b = tt.AppendFormat(b, consts.DateFormat)
    b = append(b, '"')
    return b, nil
}

func (t jsonDate) String() string {
    return time.Time(t).Format(consts.DateFormat)
}

func (t jsonDate) Value() (driver.Value, error) {
    return t.String(), nil
}
