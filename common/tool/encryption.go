package tool

import (
	"crypto/md5"
	"fmt"
	"io"
)

// Calculate MD5 hash value of string
func Md5ByString(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		panic(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

// Calculate MD5 hash value of bytes slice
func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}