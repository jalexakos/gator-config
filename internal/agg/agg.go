package agg

import (
	"context"
	"fmt"
	"html"

	command "github.com/jalexakos/gator-config/internal/command"
	config "github.com/jalexakos/gator-config/internal/config"
	"github.com/jalexakos/gator-config/internal/utils"
)

func HandlerAgg(s *config.State, cmd command.Command) error {
	rssFeed, err := utils.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("Error Fetching RSS Feed")
	}
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i := 0; i < len(rssFeed.Channel.Item); i++ {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(rssFeed.Channel.Item[i].Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(rssFeed.Channel.Item[i].Description)
	}
	fmt.Print(rssFeed)
	return nil
}
