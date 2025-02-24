package functions

import (
	"context"
	"fmt"

	"github.com/TheGeneral00/blog_aggregator/internal/config"
)

func handlerAggregate(state *state, cmd command) error {
        ctx := context.Background()
        rssFeed, err := fetchFeed(ctx, config.RSSFeedURL)
        if err != nil {
                return err
        }

        fmt.Printf("%+v\n", rssFeed)
        return nil
}
