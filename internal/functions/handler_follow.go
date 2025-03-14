package functions

import (
	"context"
	"fmt"

	"github.com/TheGeneral00/blog_aggregator/internal/database"
)

func handlerFollow(state *state, cmd command, user database.User) error{
        if len(cmd.Args) != 1 {
                return fmt.Errorf("Usage: follow [URL]\n")
        }
        feed, err := state.db.GetFeedByURL(context.Background(), cmd.Args[0])
        if err != nil {
                return err
        }
        var feedFollowParams database.CreateFeedFollowParams
        feedFollowParams.FeedID = feed.ID
        feedFollowParams.UserID = user.ID
        _, err = state.db.CreateFeedFollow(context.Background(), feedFollowParams)
        if err != nil{
                return err
        }
        fmt.Printf("%v is now following %v\n", user.Name, feed.Name)
        return nil
}
