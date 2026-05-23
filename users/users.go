package users

import (
	"context"
	"fmt"

	"github.com/jalexakos/gator-config/internal/command"
	config "github.com/jalexakos/gator-config/internal/config"
)

func HandlerUsers(s *config.State, cmd command.Command) error {
	if users, err := s.Db.GetUsers(context.Background()); err != nil {
		return err
	} else {
		for i := 0; i < len(users); i++ {
			if users[i].Name == s.Cfg.CurrentUserName {
				fmt.Println(users[i].Name, "(current)")
			} else {
				fmt.Println(users[i].Name)
			}
		}
		return nil
	}
}
