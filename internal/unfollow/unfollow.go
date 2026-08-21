package unfollow

import (
	"context"
	"fmt"

	"github.com/jalexakos/gator-config/internal/command"
	"github.com/jalexakos/gator-config/internal/config"
	"github.com/jalexakos/gator-config/internal/database"
)

func HandlerUnfollow(s *config.State, cmd command.Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("no URL provided; please provide URL to unfollow")
	}

	url := cmd.Args[0]

	feed, err := s.Db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to get feed by URL: %w", err)
	}

	err = s.Db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("failed to delete feed follow: %w", err)
	}

	fmt.Println("unfollowed feed:", feed.Name)
	return nil
}
