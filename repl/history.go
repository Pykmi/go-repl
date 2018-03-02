package repl

import "fmt"

const T string = "\r"

type history struct {
	store []string
	pos int
}

func(h *history) goUp() {
	if h.pos+1 >= len(h.store) {
		return
	}

	h.pos++
	fmt.Printf("%v%v", T, h.store[h.pos])
}

func(h *history) save(in *input) {
	h.store = append(h.store, in.in)
}

func newHistory() *history {
	h := history{pos: 0}
	return &h
}