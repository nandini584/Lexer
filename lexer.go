package lexer 
import (
	"testing"
	"lexer/token"
)
type Lexer struct{
	input string
    position int 
	readPosition int 
	ch byte
}
func New(input string) *Lexer{
	l :=&Lexer{input: input}
	return l
}
func (l *Lexer) readChar(){
	if l.readPosition >= len(l.input) {
		l.ch=0;
	} else{
		l.ch= l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
func TestNextToken( t *testing.T){
	input := `THOU x = 5;
	THOU y = 10;
	THOU add = THOU(x, y) {
		RETURN x + y;
	};
	THOU result = add(x, y);
	SAYETH(result);
	`
	tests:= []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.DECLARE, "THOU"},
		{token.PRINT, "SAYETH"},
		{TOKEN.INPUT, "LISTEN"},
		{token.RETURN, "RETURN"},
		{token.IF, "IF"},
		{token.ELSE, "ELSE"},
		{token.ELSEIF, "ELSEIF"},
		{token.FOR, "FOR"},
		{token.COMMA, ','},
		{token.SEMICOLON, ";"},
		{token.COLON, ":"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.MULTIPLY, "*"},
		{token.DIVIDE, "/"},
		{token.LESSTHAN, "<"},
		{token.GREATERTHAN, ">"},
		{token.EQUAL, "=="},
		{token.NOTEQUAL, "!="},
		{token.LEQUAL, "<="},
		{token.GEQUAL, ">="},
		//identifiers
		//integers
	}
	l:=New(input)

	for i, test:=range tests{
		tok:=NextToken()

		if tok.Type != test.expectedType{
			t.Fatalf("tests[%d] - Error: Art thou mistaken? %q was excpected but got %q instead", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral{
			t.Fatalf("tests[%d] - Error: Art thou mistaken? %q was excpected but got %q instead", i, test.expectedLiteral, tok.Literal)
		}
	}
}