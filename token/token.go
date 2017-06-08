package token

type Offset int

const (
	EOF int = iota
	ILLEGAL
	INT
	ID
	STRING
	/* punctuation mark */
	COMMA
	COLON
	SEMICOLON
	LPAREN
	RPAREN
	LBRACK
	RBRACK
	LBRACE
	RBRACE
	DOT
	PLUS
	MINUS
	TIMES
	DIVIDE
	EQ
	NEQ
	LT
	LE
	GT
	GE
	AND
	OR
	ASSIGN
	/* reserved word */
	reserved_word_beg
	WHILE
	FOR
	TO
	BREAK
	LET
	IN
	END
	FUNCTION
	VAR
	TYPE
	ARRAY
	IF
	THEN
	ELSE
	DO
	OF
	NIL
	reserved_word_end
)

// yyToknames
var yyToknames = [...]string{
	"EOF",
	"ILLEGAL",
	"INT",
	"ID",
	"STRING",
	"COMMA",
	"COLON",
	"SEMICOLON",
	"LPAREN",
	"RPAREN",
	"LBRACK",
	"RBRACK",
	"LBRACE",
	"RBRACE",
	"DOT",
	"PLUS",
	"MINUS",
	"TIMES",
	"DIVIDE",
	"EQ",
	"NEQ",
	"LT",
	"LE",
	"GT",
	"GE",
	"AND",
	"OR",
	"ASSIGN",
	"reserved_word_beg", // reserved_word_beg
	"WHILE",
	"FOR",
	"TO",
	"BREAK",
	"LET",
	"IN",
	"END",
	"FUNCTION",
	"VAR",
	"TYPE",
	"ARRAY",
	"IF",
	"THEN",
	"ELSE",
	"DO",
	"OF",
	"NIL",
	"reserved_word_end", // reserved_word_end
}

func Lookup(ident string) int {
	for tok := reserved_word_beg; tok < reserved_word_end-1; tok++ {
		if yyToknames[tok] == ident {
			return tok
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
