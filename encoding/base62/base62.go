package base62

import (
	"bytes"
	"github.com/kasvith/urlsh/math"
	"strings"
)

const alphabet = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_@-"
const base = len(alphabet)

func Encode(n int) string {
	if n < 0 {
		return ""
	}
	if n == 0 {
		return string(alphabet[0])
	}

	buffer := &bytes.Buffer{}
	for n > 0 {
		buffer.WriteByte(alphabet[n%base])
		n /= base
	}
	data := buffer.Bytes()
	for left, right := 0, len(data)-1; left < right; left, right = left+1, right-1 {
		data[left], data[right] = data[right], data[left]
	}
	return string(data)
}

func Decode(s string) int {
	n := 0
	l := len(s)
	for pos, char := range s {
		p := l - pos - 1
		i := strings.IndexByte(s, byte(char))
		n += i * math.IntPow(base, p)
	}
	return n
}

