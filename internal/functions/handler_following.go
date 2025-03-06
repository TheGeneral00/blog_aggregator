package functions

import (
	"context"
	"fmt"
)

func handlerFollowing(state *state, cmd command) error{
        if len(cmd.Args) != 0 {
                return fmt.Errorf("The command takes no additional arguments")
        }

        currUser, err := state.db.GetUser(context.Background(), state.config.CurrentUserName)
        if err != nil{
                return err
        } 

        feeds, err := state.db.GetFeedFollowsForUser(context.Background(), currUser.ID)
        if err != nil{
                return err
        }

        fmt.Printf("Current user: %v\n Followed feeds:\n", state.config.CurrentUserName)

        for _, feed := range(feeds){
                fmt.Printf("- %v\n", feed.FeedName)
        }
        return nil
}
