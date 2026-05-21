package reset

import (
	"context"

	command "github.com/jalexakos/gator-config/internal/command"
	config "github.com/jalexakos/gator-config/internal/config"
)

func HandlerReset(s *config.State, cmd command.Command) error {
	ctx := context.Background()
	if err := s.Db.DeleteAllUsers(ctx); err != nil {
		return err
	}
	return nil
}
