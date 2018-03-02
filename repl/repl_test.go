package repl

import (
	"testing"
)

func TestReplCmd(t *testing.T) {
	replServer := New()

	quitCmd := newInput(":quit")
	if ok, _ := replServer.runReplCmd(quitCmd); ok {
		t.Logf("Cmd quit expected=%v got=%v", false, ok)
	}
}

func TestRegister(t *testing.T) {
	// define a user command with info subcommand
	user := Command{
		Name: "user",
		SubCommands: []Command{
			{
				Name: "info",
				Action: func() error {
					t.Logf("User's name is %v", "pykmi")
					return nil
				},
			},
		},
	}

	replServer := New()
	replServer.Register(user)
	replServer.userCmds[user.Name].SubCommands[0].Action()
}