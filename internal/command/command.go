package command

import (
	"fmt"

	"github.com/jalexakos/gator-config/internal/config"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Commands map[string]Command
	Handlers map[string]func(*config.State, Command) error
}

func (c *Commands) Run(s *config.State, cmd Command) error {
	if handler, ok := c.Handlers[cmd.Name]; ok {
		return handler(s, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.Name)
}

func (c *Commands) Register(name string, f func(*config.State, Command) error) {
	if _, ok := c.Handlers[name]; ok {
		fmt.Printf("Command %s already registered\n", name)
		return
	}
	c.Handlers[name] = f
}
