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
	input := `thou x = 5;
	thou y = 10;
	thou add = thou(x, y) {
		return x + y;
	};
	thou result = add(x, y);
	sayeth(result);
	`
	tests:= []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.DECLARE, "thou"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT,"5"},
		{token.SEMICOLON,";"},
		{token.DECLARE, "thou"},
		{token.IDENT, "y"},
		{token.ASSIGN, "="},
		{token.INT,"10"},
		{token.SEMICOLON,";"},
		{token.DECLARE,"thou"},
		{token.IDENT,"add"},
		{token.ASSIGN,"="},
		{token.DECLARE,"thou"},
		{token.LPAREN,"("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN,")"},
		{token.LBRACE,"{"},
		{token.RETURN,"return"},
		{token.IDENT,"x"},
		{token.PLUS,"+"},
		{token.IDENT,"y"},
		{token.SEMICOLON,";"},
		{token.RBRACE,"}"},
		{token.SEMICOLON,";"},
		{token.DECLARE, "thou"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT,"add"}
		{token.LPAREN,"("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN,")"},
		{token.PRINT,"sayenth"},
		{token.LPAREN,"("},
		{token.IDENT,"result"},
		{token.RPAREN,")"},
		{token.SEMICOLON,";"},
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