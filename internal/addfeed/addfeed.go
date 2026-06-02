package addfeed

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/jalexakos/gator-config/internal/command"
	"github.com/jalexakos/gator-config/internal/config"
	"github.com/jalexakos/gator-config/internal/database"
)

func HandlerAddFeed(s *config.State, cmd command.Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("no feed name or URL provided; please provide one")
	}
	feedName := cmd.Args[0]
	url := cmd.Args[1]

	userName := s.Cfg.CurrentUserName

	user, err := s.Db.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       url,
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
	}

	feed, err := s.Db.CreateFeed(context.Background(), feedParams)

	if err != nil {
		return err
	}

	fmt.Println("feed added:", feed.ID, feed.CreatedAt, feed.UpdatedAt, feed.Name, feed.Url, feed.UserID)

	return nil
}
