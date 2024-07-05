package utils

import (
	"syscall"
	"unicode/utf16"
)

func FullPath(name string) (path string, err error) {
	p, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return "", err
	}
	n := uint32(100)
	for {
		buf := make([]uint16, n)
		n, err = syscall.GetFullPathName(p, uint32(len(buf)), &buf[0], nil)
		if err != nil {
			return "", err
		}
		if n <= uint32(len(buf)) {
			return syscall.UTF16ToString(buf[:n]), nil
		}
	}
}

func StringToCharPtr(str string) *uint8 {
	chars := append([]byte(str), 0) // null terminated
	return &chars[0]
}

func StringToUTF16Ptr(str string) *uint16 {
	wchars := utf16.Encode([]rune(str + "\x00"))
	return &wchars[0]
}
