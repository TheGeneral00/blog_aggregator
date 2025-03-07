package functions

import (
	"context"
	"fmt"
	"time"

	"github.com/TheGeneral00/blog_aggregator/internal/database"
)

func handlerAggregate(s *state, cmd command, user database.User) error {
        if len(cmd.Args) != 1 {
                return fmt.Errorf("Usage: agg [time interval]\n The time interval has tbe written as a string like: 1m")
        }
        timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
        if err != nil {
                return err
        }
        ticker := time.NewTicker(timeBetweenRequests)
        for ;; <-ticker.C{
                scrapeFeeds(s)
        }
        fmt.Println("Aggregation command has been stopped")
        return nil
}

func scrapeFeeds(s *state) {
        
}
