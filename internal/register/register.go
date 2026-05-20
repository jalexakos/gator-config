package register

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	command "github.com/jalexakos/gator-config/internal/command"
	config "github.com/jalexakos/gator-config/internal/config"
	"github.com/jalexakos/gator-config/internal/database"
)

func HandlerRegister(s *config.State, cmd command.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no name provided; please provide one")
	}
	name := cmd.Args[0]
	ctx := context.Background()
	if user, err := s.Db.CreateUser(ctx, database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name}); err != nil {
		return err
	} else {
		fmt.Println("user has been created:")
		fmt.Println("Name", user.Name, "ID", user.ID, "CreatedAt", user.CreatedAt, "UpdatedAt", user.UpdatedAt)
	}
	if err := s.Cfg.SetUser(name); err != nil {
		return err
	}
	return nil
}
