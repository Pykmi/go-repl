package repl

import "fmt"

type buffer struct {
	pos int
	store []rune
}

func(b *buffer) add(char rune) {
	b.store = append(b.store, char)
}

func(b *buffer) print() {
	fmt.Printf("\r> %v", string(b.store))
}

func(b *buffer) backspace() {

}

func(b *buffer) toString() string {
	return string(b.store)
}

func newBuffer() *buffer {
	buf := buffer{pos: -1}
	return &buf
}