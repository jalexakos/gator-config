package follow

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/jalexakos/gator-config/internal/command"
	"github.com/jalexakos/gator-config/internal/config"
	"github.com/jalexakos/gator-config/internal/database"
)

func HandlerFollow(s *config.State, cmd command.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("no URL provided; please provide URL to follow")
	}
	url := cmd.Args[0]

	userName := s.Cfg.CurrentUserName

	user, err := s.Db.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}

	feed, err := s.Db.GetFeedByURL(context.Background(), url)

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feed_follow, err := s.Db.CreateFeedFollow(context.Background(), feedFollowParams)

	if err != nil {
		return err
	}

	fmt.Printf("%s followed feed %s", feed_follow.UserName, feed_follow.FeedName)

	return nil
}
