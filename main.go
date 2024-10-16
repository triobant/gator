package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/triobant/go-blog-aggregator/internal/config"
    "github.com/triobant/go-blog-aggregator/internal/database"
    _ "github.com/lib/pq"
)

type state struct {
    db  *database.Queries
    cfg *config.Config
}

func main() {
    cfg, err := config.Read()
    if err != nil {
        log.Fatalf("error reading config: %v", err)
    }

    db, err := sql.Open("postgres", cfg.DBURL)
    if err != nil {
        log.Fatalf("error connecting to database: %v", err)
    }
    defer db.Close()
    dbQueries := database.New(db)

    programState := &state{
        db:  dbQueries,
        cfg: &cfg,
    }

    cmds := commands{
        registeredCommands: make(map[string]func(*state, command) error),
    }
    cmds.register("login", handlerLogin)
    cmds.register("register", handlerRegister)
    cmds.register("reset", handlerReset)

    if len(os.Args) < 2 {
        fmt.Println("Usage: cli <command> [args...]")
        os.Exit(1)
    }

    cmdName := os.Args[1]
    cmdArgs := os.Args[2:]

    err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
    if err != nil {
        log.Fatal(err)
    }
}
