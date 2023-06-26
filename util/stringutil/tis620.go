// ManyThanks https://github.com/tupkung/tis620
package stringutil

import "unicode/utf8"

const OFFSET = 0xD60
const WIDTH = 3 // because thai character in utf-8 use 3 bytes (start at : 0xE01 to 0xE5B)

func Tis620ToUtf8(tis620bytes []byte) []byte {
	output := make([]byte, 0)
	buffer := make([]byte, WIDTH)
	for _, c := range tis620bytes {
		if !isTis620ThaiChar(c) {
			output = append(output, c)
		} else {
			utf8.EncodeRune(buffer, rune(c)+OFFSET)
			output = append(output, buffer...)
		}
	}
	return output
}

func Utf8ToTis620(utf8String string) []byte {
	utf8Runes := []rune(utf8String)

	output := make([]byte, 0)

	for _, chr := range utf8Runes {
		if isUtfEngChar(chr) {
			output = append(output, byte(chr))
		}

		if isUtfThaiChar(chr) {
			output = append(output, byte(chr-OFFSET))
		}
	}

	return output
}

func IsTis620(data []byte) bool {

	for _, chr := range data {
		if !isTis620Char(chr) {
			return false
		}
	}

	return true
}

func isUtfThaiChar(chr rune) bool {
	return (chr >= 0xE01 && chr <= 0xE3A) || (chr >= 0xE3F && chr <= 0xE5B)
}

func isUtfEngChar(chr rune) bool {
	return (chr >= 0x00 && chr <= 0x7E)
}

func isTis620ThaiChar(c byte) bool {
	return (c >= 0xA1 && c <= 0xDA) || (c >= 0xDF && c <= 0xFB)
}

func isTis620Char(c byte) bool {
	return (c >= 0x00 && c <= 0x7E) || (c >= 0xA1 && c <= 0xDA) || (c >= 0xDF && c <= 0xFB)
}
