package login

import (
	"context"
	"fmt"

	command "github.com/jalexakos/gator-config/internal/command"
	config "github.com/jalexakos/gator-config/internal/config"
)

func HandlerLogin(s *config.State, cmd command.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no username provided; please provide one")
	}
	name := cmd.Args[0]

	user, err := s.Db.GetUser(context.Background(), name)
	if err != nil {
		println("user not found: ", err.Error())
		return err
	}
	if err := s.Cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Println("user has been set:", user)
	return nil
}
