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
