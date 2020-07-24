package utils

import "unsafe"

func Int64toint(source int64) int {
	return *(*int)(unsafe.Pointer(&source))
}
