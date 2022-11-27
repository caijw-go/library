package consts

import "time"

const (
    TimeFormat = "2006-01-02 15:04:05"
    DateFormat = "2006-01-02"
    TimeDay    = 24 * time.Hour
    TimeMonth  = 30 * TimeDay  //按30天算，不精准的算法时可以使用
    TimeYear   = 365 * TimeDay //按365天算，不精准的算法时可以使用
)
