# go-blog-aggregator or gator

A multi-player command line tool for aggregating RSS feeds and viewing the posts.

## Installation

You need to have *postgresql* and [*go*](https://go.dev/doc/install) installed to run *gator* cli tool.

**Debian/Ubuntu:**
```bash
sudo apt install postgresql
```

**Arch:**
```bash
sudo pacman -S postgresql
```

**go-blog-aggregator:**
```bash
go install github.com/triobant/go-blog-aggregator@latest
```

## Set up config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
    "db_url": "postgres://username:password@localhost:5432/database?sslmode=disable"
}
```

Replace values with your database connection string.

## Usage

Create a new user:
```bash
go-blog-aggregator register <name>
```

Add a feed:
```bash
go-blog-aggregator addfeed <url>
```

Start the aggregator:
```bash
go-blog-aggregator agg 30s
```

View the posts:
```bash
go-blog-aggregator browse [limit]
```

There are a few other commands you'll need as well:
- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in DB
- `gator unfollow <url>` - Unfollow a feed that already exists in DB

## Extending


* [ ] Add sorting and filtering options to the browse command
* [ ] Add pagination to the browse command
* [ ] Add concurrency to the agg command so that it can fetch more frequently
* [ ] Add a search command that allows for fuzzy searching of posts
* [ ] Add bookmarking or liking posts
* [ ] Add a TUI that allows you to select a post in the terminal and view it in a more readable format (either in the terminal or open in a browser)
* [ ] Add an HTTP API (and authentication/authorization) that allows other users to interact with the service remotely
* [ ] Write a service manager that keeps the agg command running in the background and restarts it if it crashes
