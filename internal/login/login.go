package login

import (
	"fmt"

	command "github.com/jalexakos/gator-config/internal/command"
	config "github.com/jalexakos/gator-config/internal/config"
)

func HandlerLogin(s *config.State, cmd command.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no username provided; please provide one")
	}
	user := cmd.Args[0]
	if err := s.Cfg.SetUser(user); err != nil {
		return err
	}

	fmt.Println("user has been set:", user)
	return nil
}
