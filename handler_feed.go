package main

import (
    "context"
    "fmt"
    "time"

    "github.com/triobant/go-blog-aggregator/internal/database"
    "github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
    user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
    if err != nil {
        return err
    }

    if len(cmd.Args) != 2 {
        return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
    }

    name := cmd.Args[0]
    url := cmd.Args[1]

    feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
        ID:         uuid.New(),
        CreatedAt:  time.Now().UTC(),
        UpdatedAt:  time.Now().UTC(),
        UserID:     user.ID,
        Name:       name,
        Url:        url,
    })
    if err != nil {
        return fmt.Errorf("couldn't create feed: %w", err)
    }

    fmt.Println("Feed created successfully!")
    printFeed(feed)
    fmt.Println()
    fmt.Println("=====================================")

    return nil
}

func printFeed(feed database.Feed) {
    fmt.Printf(" * ID:          %v\n", feed.ID)
    fmt.Printf(" * Created:     %v\n", feed.CreatedAt)
    fmt.Printf(" * Updated:     %v\n", feed.UpdatedAt)
    fmt.Printf(" * Name:        %v\n", feed.Name)
    fmt.Printf(" * URL:         %v\n", feed.Url)
    fmt.Printf(" * UserID:      %v\n", feed.UserID)
}
