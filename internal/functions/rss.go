package functions

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	)

type RSSFeed struct {
        Channel struct {
                Title           string          `xml:"title"`
                Link            string          `xml:"link"`
                Description     string          `xml:"description"`
                Items            []RSSItem       `xml:"item"`
        } `xml:"channel"`
}

type RSSItem struct {
        Title           string  `xml:"title"`
        Link            string  `xml:"link"`
        Description     string  `xml:"description"`
        PubDate         string  `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
        
        req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
        if err != nil {
                return nil, err
        }
        req.Header.Set("User-Agent", "gator")
        
        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
                return nil, err
        }
        defer resp.Body.Close()

        if resp.StatusCode < 200 || resp.StatusCode >= 300 {
                return nil, fmt.Errorf("Request to RSS failed with code: %v", resp.Status)
        }
        body, err := io.ReadAll(resp.Body)
        if err != nil {
                return nil, err
        }
        var rssFeed RSSFeed
        err = xml.Unmarshal(body, &rssFeed)
        if err != nil {
                return nil, err
        }

        rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
        rssFeed.Channel.Link = html.UnescapeString(rssFeed.Channel.Link)

        for i := range rssFeed.Channel.Items {
                rssFeed.Channel.Items[i].Title = html.UnescapeString(rssFeed.Channel.Items[i].Title)
                rssFeed.Channel.Items[i].Link = html.UnescapeString(rssFeed.Channel.Items[i].Link)
        }
        return &rssFeed, nil
}
