package functions

import (
	"context"
	"fmt"

	"github.com/TheGeneral00/blog_aggregator/internal/database"
)

func handlerAddFeed(state *state, cmd command) error {
        if len(cmd.Args) != 2 {
                return fmt.Errorf("Usage: addfeed [name] [url]")
        }
        user, err := state.db.GetUser(context.Background(), state.config.CurrentUserName)
        if err != nil{
                return err
        }
        params := database.AddFeedParams{
                Name: cmd.Args[0],
                Url: cmd.Args[1],
                UserID: user.ID,
        }
        _, err = state.db.AddFeed(context.Background(), params)
        if err != nil {
                return err
        }
        fmt.Printf("Added '%v'\n", cmd.Args[0])
        return nil 
}
