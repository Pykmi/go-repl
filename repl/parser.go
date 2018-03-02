package repl

import (
	"fmt"
	"strings"
)

type parser struct {
	commands UserCommands
	items []string
	pos int
}

func (p *parser) run() {
	cmd := p.items[p.pos]
	next := p.pos+1

	if !p.isCommand(cmd) {
		fmt.Printf("Invalid command: %v", cmd)
		return
	}

	if len(p.items) > next {
		if idx, ok := p.isSubCommand(cmd, p.items[next]); ok {
			// exec sub command
			p.commands[cmd].SubCommands[idx].Action()
			return
		}
	}

	// exec command
	p.commands[cmd].Action()
}

func (p *parser) isSubCommand(cmd string, subcmd string) (int, bool) {
	for idx, C := range p.commands[cmd].SubCommands {
		if C.Name == subcmd {
			return idx, true
		}
	}

	return -1, false
}

func (p *parser) isCommand(cmd string) bool {
	if _, ok := p.commands[cmd]; ok {
		return true
	}

	return false
}

func newParser(in *input, c UserCommands) *parser {
	items := strings.Split(in.in, " ")
	p := parser{c,items, 0}

	return &p
}