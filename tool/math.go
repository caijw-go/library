package tool

import (
    "math"
)

// IsSameNum 判断数字是否为相同的数，如111，2222等，返回bool,长度，这个相同的数字
func IsSameNum(num uint) (bool, uint, uint8) {
    length, bit := 1, num%10

    if bit == 0 { //多位数的情况下，如果bit为0，则一定不是同位数
        return false, 0, 0
    }

    for {
        if num /= 10; num == 0 {
            break
        }
        if tempBit := num % 10; tempBit != bit {
            return false, 0, 0
        }
        length++
    }
    return true, uint(length), uint8(bit)
}

func GetTotalPage(totalCount, pageSize uint) uint {
    return uint(math.Ceil(float64(totalCount) / float64(pageSize)))
}
