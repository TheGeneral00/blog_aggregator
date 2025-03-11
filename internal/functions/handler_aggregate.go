package functions

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/TheGeneral00/blog_aggregator/internal/database"
	"github.com/lib/pq"

	"github.com/araddon/dateparse"
)

func handlerAggregate(s *state, cmd command, user database.User) error {
        if len(cmd.Args) != 0 {
                return fmt.Errorf("Usage: agg")
        }
        feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
        if err != nil {
                return err
        }
        err = scrapeFeeds(s, feeds)
        if err != nil {
                return err
        }
        return nil

}

func scrapeFeeds(s *state, feeds []database.GetFeedFollowsForUserRow) error {
        for _, item := range(feeds){
                feed, err := s.db.GetFeedByID(context.Background(), item.FeedID)     
                s.db.MarkFeedFetched(context.Background(), feed.ID)
                var feedRSS *RSSFeed
                feedRSS, err = fetchFeed(context.Background(), feed.Url)
                if err != nil {
                        return err
                }
                fmt.Printf("Fetched feed '%v'\n", feedRSS.Channel.Title)
                for _, rssItem := range feedRSS.Channel.Items {
                        var publishedAt time.Time
                        if rssItem.PubDate != "" {
                                parsedTime, err := parsePubDate(rssItem.PubDate)
                                if err != nil {
                                        fmt.Printf("Warning: couldn't parse date '%s' for post '%s': %v. Using current time.\n", 
                                        rssItem.PubDate, rssItem.Title, err)
                                        publishedAt = time.Now()
                                } else {
                                        publishedAt = parsedTime
                                }
                        } else {
                                publishedAt = time.Now()
                        }
                        createPostParams := database.CreatePostParams{
                                Title:          sql.NullString{
                                                        String: rssItem.Title,
                                                        Valid: rssItem.Title != "",
                                                },
                                Url:            rssItem.Link,
                                Description:    sql.NullString{
                                                        String: rssItem.Description,
                                                        Valid: rssItem.Description != "",
                                                },
                                PublishedAt:    sql.NullTime{
                                                        Time: publishedAt,
                                                        Valid: publishedAt.IsZero(),
                                                },
                                FeedID:         item.FeedID,
                        }
                        err = s.db.CreatePost(context.Background(), createPostParams)
                        if err != nil {
                                if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
                                        continue
                                }
                                fmt.Printf("Error saving post '%s': %v\n", rssItem.Title, err)
                        }
                }
        }
        return nil
}


func parsePubDate(dateStr string) (time.Time, error) {
    if dateStr == "" {
        return time.Now(), nil
    }
    t, err := dateparse.ParseAny(dateStr)
    if err != nil {
        return time.Time{}, fmt.Errorf("unable to parse date '%s': %w", dateStr, err)
    }
    
    return t, nil
}

