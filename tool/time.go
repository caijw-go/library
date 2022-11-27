package tool

import (
    "caijw/library/consts"
    "time"
)

// ParseTime 解析时间
func ParseTime(t string) time.Time {
    res, _ := time.ParseInLocation(consts.TimeFormat, t, time.Local)
    return res
}
