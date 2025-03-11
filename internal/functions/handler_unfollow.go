package functions

import (
	"context"
	"fmt"

	"github.com/TheGeneral00/blog_aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
        if len(cmd.Args) != 1 {
                return fmt.Errorf("Usage: unfollow [URL]")
        }

        deleteFollowParams := database.DeleteFollowParams{
                UserID: user.ID,
                Url: cmd.Args[0],
        }
        feedName, err := s.db.DeleteFollow(context.Background(), deleteFollowParams)
        if err != nil {
                return err
        }
        fmt.Printf("Unfollowed Feed '%v'\n", feedName)
        return nil
}
