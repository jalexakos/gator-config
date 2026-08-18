package following

import (
	"context"
	"fmt"

	"github.com/jalexakos/gator-config/internal/command"
	"github.com/jalexakos/gator-config/internal/config"
)

func HandlerFollowing(s *config.State, cmd command.Command) error {

	userName := s.Cfg.CurrentUserName

	user, err := s.Db.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}

	feeds, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("%s is following %s", user.Name, feed.FeedName)
	}

	return nil
}
