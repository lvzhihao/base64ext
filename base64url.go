package base64ext

import (
	"encoding/base64"
	"strings"
)

type Base64ExtURL struct {
	Base64Ext
}

var URL Base64ExtURL

func (this Base64ExtURL) Encode(src []byte) string {
	s := base64.StdEncoding.EncodeToString(src)
	return strings.Map(func(r rune) rune {
		switch r {
		case '+':
			return '-'
		case '/':
			return '_'
		case '=':
			return 0
		default:
			return r
		}
	}, s)
}

func (this Base64ExtURL) Decode(des string) ([]byte, error) {
	str := strings.Map(func(r rune) rune {
		switch r {
		case '-':
			return '+'
		case '_':
			return '/'
		default:
			return r
		}
	}, des)
	if len(str)%4 != 0 {
		str += "="
	}
	return base64.StdEncoding.DecodeString(str)
}
