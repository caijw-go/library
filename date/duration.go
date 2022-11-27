package date

import (
    "fmt"
    "time"
)

/*
 * go time Duration不满足需求，需要获取到年的间隔，所以自定义一个
 * 开始时间到结束时间为时间间隔，格式:Y-m-d H:i:s
 */

// Duration /**
type Duration struct {
    d      time.Duration
    Year   int //时间间隔中的Y
    Month  int //时间间隔中的m
    Day    int //时间间隔中的d
    Hour   int //时间间隔中的H
    Minute int //时间间隔中的m
    Second int //时间间隔中的s
}

func isLeap(year int) bool {
    return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func getMonthDays(year int, month time.Month) int {
    days := 31
    switch month {
    case time.February: //2
        if isLeap(year) {
            days = 29
        } else {
            days = 28
        }
    case time.April: //4
        fallthrough
    case time.June: //6
        fallthrough
    case time.September: //9
        fallthrough
    case time.November: //11
        days = 30
    }
    return days
}

func InitDuration(timestamp time.Time) *Duration {
    return initDuration(timestamp, time.Now())
}

func initDuration(startTime time.Time, endTime time.Time) *Duration {
    if startTime.Unix() > endTime.Unix() {
        startTime, endTime = endTime, startTime
    }

    duration := endTime.Sub(startTime)

    y := endTime.Year() - startTime.Year()
    m := int(endTime.Month() - startTime.Month())
    d := endTime.Day() - startTime.Day()
    h := endTime.Hour() - startTime.Hour()
    i := endTime.Minute() - startTime.Minute()
    s := endTime.Second() - startTime.Second()

    if s < 0 { //处理秒
        s += 60
        i -= 1
    }
    if i < 0 {
        i += 60
        h -= 1
    }
    if h < 0 {
        h += 24
        d -= 1
    }
    //如果天数<0，则需要向endTime的月份借天数，能借到多少天取决于endTime的月份的上一个月有多少天
    if d < 0 {
        lastM := endTime.Month() - 1 //获得上一个月

        if lastM < 1 { //如果是1月-1，就设置为12月，按说年份也应该-1，但12月用不到年份
            lastM = time.December
        }
        d += getMonthDays(endTime.Year(), lastM)
        m -= 1
    }

    if m < 0 {
        m += 12
        y -= 1
    }

    if y < 0 { //这种不可能出现，如果出现了，说明开始时间小于结束时间，直接设置为0
        y, m, d, h, i, s = 0, 0, 0, 0, 0, 0
    }

    return &Duration{
        d:      duration,
        Year:   y,
        Month:  m,
        Day:    d,
        Hour:   h,
        Minute: i,
        Second: s,
    }
}

func (t *Duration) TotalDays() uint {
    return t.TotalHour() / 24
}

func (t *Duration) TotalHour() uint {
    return uint(t.d / time.Hour)
}

func (t *Duration) IsWholeYear() bool {
    //判断整年只要到天就行了，这24小时的任意时间都算整年
    return t.Month == 0 && t.Day == 0
}

func (t *Duration) FullTime() string {
    return fmt.Sprintf("%d年%d个月%d天【共%d天】%d时%d分", t.Year, t.Month, t.Day, t.TotalDays(), t.Hour, t.Minute)
}
