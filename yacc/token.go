package yacc

import "strings"

func Lookup(ident string) int {
	for tok := reserved_word_beg - INT + 3; tok < reserved_word_end-INT+3; tok++ {
		if strings.ToLower(yyToknames[tok]) == ident {
			return tok + INT - 3
		}
	}
	return ID
}

func Digit(ch byte) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch - '0')
	case 'a' <= ch && ch <= 'f':
		return int(ch - 'a' + 10)
	case 'A' <= ch && ch <= 'F':
		return int(ch - 'A' + 10)
	}
	return 16 // 其它字符都返回 16
}
