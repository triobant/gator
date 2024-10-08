package main

import (
    "fmt"
    "log"

    "github.com/triobant/go-blog-aggregator/internal/config"
)

func main() {
    cfg, err := config.Read()
    if err != nil {
        log.Fatalf("error reading config: %v", err)
    }
    fmt.Printf("Read config: %+v\n", cfg)

    err = cfg.SetUser("triobant")

    cfg, err = config.Read()
    if err != nil {
        log.Fatalf("error reading config: %v", err)
    }
    fmt.Printf("Read config again: %+v\n", cfg)
}
