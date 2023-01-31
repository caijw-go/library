package types

func InArray[T comparable](needle T, arr []T) bool {
    for _, item := range arr {
        if item == needle {
            return true
        }
    }
    return false
}
