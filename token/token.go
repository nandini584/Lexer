package token 
type TokenType string 
type Token struct {
	Type TokenType   //tokentype's type is string 
	Literal string
}
var keywords = map[string]TokenType{
	"thou": DECLARE,
	"sayeth": PRINT,
	"listen": INPUT,
	"return": RETURN,
	"if": IF,
	"else": ELSE,
	"elseif": ELSEIF,
	"for": FOR,
}
func LookupIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok{
		return tok
	}
	return IDENT
}
const (
	//Misc
	ILLEGAL="ILLEGAL"
	EOF="EOF"
	NUL="NUL"

	// Identifiers + literals
	IDENT="IDENT"
	INT="INT"

	// Operators
	ASSIGN="="

	//Assignment operators
	PLUS="+"
	MINUS="-"
	MULTIPLY="*"
	DIVIDE="/"

	// Relational operators
	LESSTHAN="<"
	GREATERTHAN=">"
	EQUAL="=="
	NOTEQUAL="!="
    LEQUAL="<="
	GEQUAL=">="
	NOT="!"
	
	// Delimiters
	COMMA=","
	SEMICOLON=";"
	COLON=":"
	LPAREN="("
	RPAREN=")"
	LBRACE="{"
	RBRACE="}"
	LBRACKET="["
	RBRACKET="]"

	// Keywords
	DECLARE="THOU"
	PRINT="SAYETH"
	INPUT="LISTEN"
	RETURN="RETURN"
	IF="IF"
	ELSE="ELSE"
	ELSEIF="ELSEIF"
	FOR="FOR"
)