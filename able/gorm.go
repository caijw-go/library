package able

import (
    "database/sql/driver"
    "github.com/caijw-go/library/consts"
    "time"
)

//这个文件定义了一些功能和结构体，用于gorm自定义数据类型
//主要是时间，由于正常的时间转成json以后，格式

type GormTime time.Time

func (t *GormTime) UnmarshalJSON(data []byte) (err error) {
    now, err := time.ParseInLocation(`"`+consts.TimeFormat+`"`, string(data), time.Local)
    *t = GormTime(now)
    return
}

func (t GormTime) MarshalJSON() ([]byte, error) {
    tt := time.Time(t)
    if &t == nil || tt.IsZero() {
        return []byte("null"), nil
    }
    b := make([]byte, 0, len(consts.TimeFormat)+2)
    b = append(b, '"')
    b = tt.AppendFormat(b, consts.TimeFormat)
    b = append(b, '"')
    return b, nil
}

func (t GormTime) String() string {
    return time.Time(t).Format(consts.TimeFormat)
}

func (t GormTime) Value() (driver.Value, error) {
    return t.String(), nil
}

type GormDate time.Time

func (t *GormDate) UnmarshalJSON(data []byte) (err error) {
    now, err := time.ParseInLocation(`"`+consts.TimeFormatDate+`"`, string(data), time.Local)
    *t = GormDate(now)
    return
}

func (t GormDate) MarshalJSON() ([]byte, error) {
    b := make([]byte, 0, len(consts.TimeFormatDate)+2)
    b = append(b, '"')
    b = time.Time(t).AppendFormat(b, consts.TimeFormatDate)
    b = append(b, '"')
    return b, nil
}

func (t GormDate) String() string {
    return time.Time(t).Format(consts.TimeFormatDate)
}

func (t GormDate) Value() (driver.Value, error) {
    return t.String(), nil
}
