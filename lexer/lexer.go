package lexer 
import (
	"testing"
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
func (l *Lexer) NextToken() token.Token{
	var tok token.Token
	l.skipWhitespace()
	switch l.ch{
	case '=':
		if l.peekChar()== '='{
			ch=l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch)+string(l.ch)}
		}else{
			tok = newToken(token.ASSIGN, l.ch)

		}
	case '+':
		tok= newToken(token.PLUS, l.ch)
	case ';':
		tok= newToken(token.SEMICOLON, l.ch)
	case '(':
		tok= newToken(token.LPAREN, l.ch)
	case ')':
		tok= newToken(token.RPAREN, l.ch)
	case '{':
		tok= newToken(token.LBRACE, l.ch)
	case '}':
		tok= newToken(token.RBRACE, l.ch)
	case ',':
		tok= newToken(token.COMMA, l.ch)
	case '-':
		tok= newToken(token.MINUS, l.ch)
	case '/':
		tok= newToken(token.SLASH, l.ch)
	case '*':
		tok= newToken(token.ASTERISK, l.ch)
	case '>' :
		if l.peekChar()== '='{
			ch=l.ch
			l.readChar()
			tok = token.Token{Type: token.GREATER_EQ, Literal: string(ch)+string(l.ch)}
		}else{
		tok= newToken(token.GREATER, l.ch)}
	case '<':
		if l.peekChar()== '='{
			ch=l.ch
			l.readChar()
			tok = token.Token{Type: token.LESS_EQ, Literal: string(ch)+string(l.ch)}
		}else{
		tok= newToken(token.LESS, l.ch)}
	case '[':
		tok= newToken(token.LBRACKET, l.ch)
	case ']':
		tok= newToken(token.RBRACKET, l.ch)
	case '!':
		if l.peekChar()== '='{
			ch=l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch)+string(l.ch)}
		}else{
		tok= newToken(token.BANG, l.ch)}
	case ':':
		tok= newToken(token.COLON, l.ch)
	case 0:
		tok.Literal=""
		tok.Type=token.EOF
	default:
		if isLetter(l.ch){
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
		else if isDigit(l.ch){
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}else{
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}
func (l *Lexer) skipWhitespace(){
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r'{
		l.readChar()
	}
}
func isDigit(ch byte) bool{
	return '0' <= ch && ch <= '9'
}
func (l *Lexer) readNumber() string{
	position=l.position
	for isDigit(l.ch){
		l.readChar()
	}
	return l.input[position:l.position]
}
func isLetter(ch byte) bool{
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' 
}
func (l *Lexer) readIdentifier() string{
	position := l.position
	for isLetter(l.ch){
		l.readChar()
	}
	return l.input[position:l.position]
}
func (l *Lexer) peekChar() byte{
	if l.readPosition >= len(l.input){
		return 0
	}else{
		return l.input[l.readPosition]
	}
}
func newToken(tokenType token)
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
		{token.EOF,""},
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