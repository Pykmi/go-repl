package repl

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"io"
	"os"
	"os/user"
)

const nl string = "\n"

type Command struct {
	Name string
	Action func() error
	SubCommands []Command
}

type repl struct {
	buf buffer
	replCmds ReplCommands
	userCmds UserCommands
}

type input struct {
	isRepl bool
	isUser bool
	in     string
}

type ReplCommands map[string]func() (bool, error)
type UserCommands map[string]Command

func (r *repl) run(in io.Reader, out io.Writer) {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	buf := newBuffer()
	buf.print()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc {
			break
		} else if key == keyboard.KeySpace {
			buf.add('\u0020')
		} else if key == keyboard.KeyBackspace2 {
			buf.add('\u0008')
			buf.add('\u0020')
			buf.add('\u0008')
			//buf.add('\b')
		} else if key == keyboard.KeyEnter {
			if "hello" == buf.toString() {
				fmt.Println("YES!")
			}
		}

		buf.add(char)
		buf.print()
	}




	/*scanner := bufio.NewScanner(in)

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

		input := newInput(text)
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
	//}
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
	//i.history = map[int]string{}

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
