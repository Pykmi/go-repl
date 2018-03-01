package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"
)

const nl string = "\n"

type Command struct {
	Name string
	Action func() error
	SubCommands []Command
}

type repl struct {
	replCmds ReplCommands
	userCmds UserCommands
}

type input struct {
	isRepl bool
	isUser bool
	in     string
}

type parser struct {
	commands UserCommands
	items []string
	pos int
}

type ReplCommands map[string]func() (bool, error)
type UserCommands map[string]Command

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

func (r *repl) run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		r.prompt()
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		text := strings.TrimSpace(scanner.Text())

		if len(text) < 1 {
			continue
		}

		input := newInput(scanner.Text())

		if input.isRepl {
			ok, _ := r.runReplCmd(input)
			if !ok {
				break
			}
		} else if input.isUser {
			parser := newParser(input, r.userCmds)
			parser.run()
		}

		/*if err != nil {
			// error handling
		}*/

		//fmt.Print(input.in)
	}
}

func (r *repl) prompt() {
	fmt.Printf("%v> ", nl)
}

func (r *repl) Register(cmd Command) {
	r.userCmds[cmd.Name] = cmd
}

func (r *repl) runReplCmd(in *input) (bool, error) {
	if in.in == "" {
		return true, nil
	}

	cmd := in.in[1:]
	if _, ok := r.replCmds[cmd]; ok {
		return r.replCmds[cmd]()
	}

	return true, nil
}

func (r *repl) Start() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!\n", u.Username)
	fmt.Printf("Feel free to type in commands\n")
	r.run(os.Stdin, os.Stdout)
}

func New() *repl {
	i := repl{}

	i.replCmds = map[string]func() (bool, error){}
	i.userCmds = UserCommands{}

	// register repl commands
	i.replCmds["quit"] = func() (bool, error) {
		return false, nil
	}

	return &i
}

func newInput(i string) *input {
	in := input{isRepl: false, isUser: false, in: i}

	if in.in[:1] == ":" {
		in.isRepl = true
	} else {
		in.isUser = true
	}

	return &in
}

func newParser(in *input, c UserCommands) *parser {
	items := strings.Split(in.in, " ")
	p := parser{c,items, 0}

	return &p
}
