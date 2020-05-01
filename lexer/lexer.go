package lexer

type Lexer struct {
	input        string
	position     int  //char position
	readPosition int  //reading position in input after current char
	ch           byte //char being examinated
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 //sets ASCII for NUL char
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
