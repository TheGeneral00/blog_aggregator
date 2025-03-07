package functions

import (
	"context"
	"fmt"

	"github.com/TheGeneral00/blog_aggregator/internal/database"
)

func handlerFollowing(state *state, cmd command, user database.User) error{
        if len(cmd.Args) != 0 {
                return fmt.Errorf("The command takes no additional arguments")
        }

        feeds, err := state.db.GetFeedFollowsForUser(context.Background(), user.ID)
        if err != nil{
                return err
        }

        fmt.Printf("Current user: %v\n Followed feeds:\n", state.config.CurrentUserName)

        for _, feed := range(feeds){
                fmt.Printf("- %v\n", feed.FeedName)
        }
        return nil
}
