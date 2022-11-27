package convert

import (
    "fmt"
    "strconv"
)

func MBSizeFormat(size interface{}) string {
    res, err := strconv.ParseFloat(fmt.Sprintf("%v", size), 0)
    if err != nil {
        fmt.Println(err)
        return "0 Bytes"
    }
    sizeConfig := []string{
        0: "MB", 1: "GB", 2: "TB",
    }

    var resIndex int

    for res >= 1024 {
        res = res / 1024
        resIndex++
        if resIndex == 4 {
            break
        }
    }
    return fmt.Sprintf("%.1f %v", res, sizeConfig[resIndex])
}

func SizeFormat(size interface{}) string {
    res, err := strconv.ParseFloat(fmt.Sprintf("%v", size), 0)
    if err != nil {
        fmt.Println(err)
        return "0 Bytes"
    }
    sizeConfig := []string{
        0: "Bytes", 1: "KB", 2: "MB", 3: "GB", 4: "TB",
    }

    var resIndex int

    for res >= 1024 {
        res = res / 1024
        resIndex++
        if resIndex == 4 {
            break
        }
    }
    return fmt.Sprintf("%.1f %v", res, sizeConfig[resIndex])
}

func SizeFormatToMBWithoutUnit(size interface{}) string {
    res, err := strconv.ParseFloat(fmt.Sprintf("%v", size), 0)
    if err != nil {
        fmt.Println(err)
        return "0"
    }
    res = res / 1024 / 1024
    return fmt.Sprintf("%.2f", res)
}
