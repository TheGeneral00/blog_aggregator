package functions

import (
	"context"
	"fmt"
)

func handlerFeeds(state *state, cmd command) error {
        if len(cmd.Args) != 0 { 
                return fmt.Errorf("This command dosent take extra arguments. Usage: feeds ")
        }
        feeds, err := state.db.GetFeeds(context.Background())
        if err != nil{
                return err
        }

        for _, feed := range(feeds){
                user, err := state.db.GetUserByID(context.Background(), feed.UserID)
                if err != nil {
                        return err
                }
                fmt.Printf("Name: %v, URL: %v, Added By: %v \n", feed.Name, feed.Url, user.Name)
        } 
        return nil
}
