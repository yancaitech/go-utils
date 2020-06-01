package utils

import (
	"math"
	"time"
)

// Remove func
func Remove(slice []interface{}, elem interface{}) []interface{} {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if v == elem {
			slice = append(slice[:i], slice[i+1:]...)
			return Remove(slice, elem)
		}
	}
	return slice
}

// RemoveZero func
func RemoveZero(slice []interface{}) []interface{} {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if IfZero(v) {
			slice = append(slice[:i], slice[i+1:]...)
			return RemoveZero(slice)
		}
	}
	return slice
}

// IfZero func
func IfZero(arg interface{}) bool {
	if arg == nil {
		return true
	}
	switch v := arg.(type) {
	case int, int32, int16, int64:
		if v == 0 {
			return true
		}
	case float32:
		r := float64(v)
		return math.Abs(r-0) < 0.0000001
	case float64:
		return math.Abs(v-0) < 0.0000001
	case string:
		if v == "" || v == "%%" || v == "%" {
			return true
		}
	case *string, *int, *int64, *int32, *int16, *int8, *float32, *float64, *time.Time:
		if v == nil {
			return true
		}
	case time.Time:
		return v.IsZero()
	default:
		return false
	}
	return false
}

// StringSliceEqual func
func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// StringSliceContain func
func StringSliceContain(a []string, b string) (c bool) {
	if a == nil {
		return false
	}

	c = false

	for _, v := range a {
		if v == b {
			c = true
			break
		}
	}

	return c
}

// ByteSliceEqual func
func ByteSliceEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// ByteSliceCopy func
func ByteSliceCopy(src []byte, srcOffset int, dst []byte, dstOffset, count int) bool {
	srcLen := len(src)
	if srcOffset > srcLen || count > srcLen || srcOffset+count > srcLen {
		return false
	}
	dstLen := len(dst)
	if dstOffset > dstLen || count > dstLen || dstOffset+count > dstLen {
		return false
	}
	index := 0
	for i := srcOffset; i < srcOffset+count; i++ {
		dst[dstOffset+index] = src[srcOffset+index]
		index++
	}
	return true
}
