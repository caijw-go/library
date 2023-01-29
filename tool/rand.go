package tool

import (
    "math"
    "math/rand"
    "time"
)

// RandNumberN 返回指定位数的随机数
func RandNumberN(len int) int {
    if len < 2 {
        len = 1
    }
    minLenAddOne := int(math.Pow10(len)) //获取len+1位的最小数
    minLen := minLenAddOne / 10          //获取len位的最小数
    //拿三位数举例
    //minLenAddOne=1000,minLen:=100
    //rand.Intn(900)即0-899的随机数 +100 即100到999的随机数
    return minLen + rand.New(rand.NewSource(time.Now().UnixNano())).Intn(minLenAddOne-minLen)
}
