package able

import (
    "database/sql/driver"
    "encoding/json"
    "errors"
    "fmt"
    "github.com/caijw-go/library/consts"
    "time"
)

//这个文件定义了一些功能和结构体，用于gorm自定义数据类型

// GormJsonScanner 将传入的类型解析为当前类型（用于数据库查询后得到的json字符串，解析为对应的结构体）
func GormJsonScanner(t interface{}, value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New(fmt.Sprint("Failed to struct value:", value))
    }
    return json.Unmarshal(bytes, t)
}

// GormJsonValuer 将当前类型解析为传入的类型（用于将结构体保存到数据库时，转换为字符串进行保存）[不能使用*,避免转数据失败]
func GormJsonValuer(t interface{}) (driver.Value, error) {
    b, err := json.Marshal(t) //转换成JSON返回的是byte[]
    return string(b), err
}

type GormTime time.Time

func (t *GormTime) UnmarshalJSON(data []byte) (err error) {
    now, err := time.ParseInLocation(`"`+consts.TimeFormat+`"`, string(data), time.Local)
    *t = GormTime(now)
    return
}

func (t GormTime) MarshalJSON() ([]byte, error) {
    b := make([]byte, 0, len(consts.TimeFormat)+2)
    b = append(b, '"')
    b = time.Time(t).AppendFormat(b, consts.TimeFormat)
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
    now, err := time.ParseInLocation(`"`+consts.DateFormat+`"`, string(data), time.Local)
    *t = GormDate(now)
    return
}

func (t GormDate) MarshalJSON() ([]byte, error) {
    b := make([]byte, 0, len(consts.DateFormat)+2)
    b = append(b, '"')
    b = time.Time(t).AppendFormat(b, consts.DateFormat)
    b = append(b, '"')
    return b, nil
}

func (t GormDate) String() string {
    return time.Time(t).Format(consts.DateFormat)
}

func (t GormDate) Value() (driver.Value, error) {
    return t.String(), nil
}
