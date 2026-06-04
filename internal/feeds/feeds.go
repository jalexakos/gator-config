package feeds

import (
	"context"
	"fmt"

	"github.com/jalexakos/gator-config/internal/command"
	"github.com/jalexakos/gator-config/internal/config"
)

func HandlerFeeds(s *config.State, cmd command.Command) error {
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting feeds")
	}

	// NEED USER NAME
	for i := range len(feeds) {
		user, err := s.Db.GetUserById(context.Background(), feeds[i].UserID.UUID)
		if err != nil {
			return fmt.Errorf("Error getting user: %w", feeds[i].UserID)
		}
		fmt.Printf("Feed name: %s\n Feed URL: %s\n Created By: %s\n", feeds[i].Name, feeds[i].Url, user.Name)
	}
	return nil
}
