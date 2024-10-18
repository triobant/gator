package main

import (
    "context"
    "fmt"
)

func handlerAgg(s *state, cmd command) error {
    feed, err := fetchFeed(context.Background(), "https://wagslane.dev/index.xml")
    if err != nil {
        return fmt.Errorf("couldn't fetch feed: %w", err)
    }
    fmt.Printf("Feed: %+v\n", feed)
    return nil
}
