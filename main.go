package main

import (
	"fmt"
	"os"

	command "github.com/jalexakos/gator-config/internal/command"
	config "github.com/jalexakos/gator-config/internal/config"
	login "github.com/jalexakos/gator-config/internal/login"
)

func main() {
	configFile, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}
	state := config.State{
		Cfg: configFile,
	}
	commands := command.Commands{
		Handlers: make(map[string]func(*config.State, command.Command) error),
	}
	commands.Register("login", login.HandlerLogin)
	cmds := os.Args
	if len(cmds) < 2 {
		fmt.Errorf("Please provide a command")
		os.Exit(1)
	}
	cmdName := cmds[1]
	cmdArgs := cmds[2:]
	cmd := command.Command{
		Name: cmdName,
		Args: cmdArgs,
	}
	err = commands.Run(&state, cmd)
	if err != nil {
		fmt.Println("Error running command:", err)
		os.Exit(1)
	}
}
