package sysutil

import "strings"

func Rune2String(data []rune) string {
	builder := strings.Builder{}
	for _, v := range data {
		if v == 0 {
			break
		}
		builder.WriteRune(v)
	}
	return builder.String()
}

func Byte2String(data []byte) string {
	builder := strings.Builder{}
	for _, v := range data {
		if v == 0 {
			break
		}
		builder.WriteByte(v)
	}
	return builder.String()
}
