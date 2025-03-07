package functions

import (
	"context"
	"fmt"

	"github.com/TheGeneral00/blog_aggregator/internal/database"
)

func handlerAddFeed(state *state, cmd command, user database.User) error {
        if len(cmd.Args) != 2 {
                return fmt.Errorf("Usage: addfeed [name] [url]")
        }
        params := database.AddFeedParams{
                Name: cmd.Args[0],
                Url: cmd.Args[1],
                UserID: user.ID,
        }
        feed, err := state.db.AddFeed(context.Background(), params)
        if err != nil {
                return err
        }
        fmt.Printf("Added '%v'\n", cmd.Args[0])

        user, err = state.db.GetUser(context.Background(), state.config.CurrentUserName)
        if err != nil{
                return err
        }
        var createFeedFollowParams database.CreateFeedFollowParams
        createFeedFollowParams.UserID = user.ID
        createFeedFollowParams.FeedID = feed.ID 
        _, err = state.db.CreateFeedFollow(context.Background(), createFeedFollowParams)
        if err != nil {
                return err
        }
        return nil 
}
