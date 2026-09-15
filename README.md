# SMDB — Simple Movie DataBrowser

A terminal user interface for browsing movies powered by [TMDB](https://www.themoviedb.org/).

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green)

## Features

- Browse **Now Playing**, **Popular**, **Top Rated**, and **Upcoming** movies
- Detailed movie view with synopsis, rating, and metadata
- Keyboard-driven navigation (arrow keys + vim keys)
- Responsive terminal layout with adaptive colors

## Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- A free [TMDB account](https://www.themoviedb.org/signup)
- A [TMDB API Read Access Token](https://www.themoviedb.org/settings/api)

## Setup

1. **Clone the repository**

   ```bash
   git clone https://github.com/Utkarsh932006/SMDB.git
   cd SMDB
   ```

2. **Set your TMDB API token**

   ```bash
   export TMDB_API_TOKEN="your_read_access_token_here"
   ```

   > You can also add this to your shell profile (`~/.bashrc`, `~/.zshrc`, etc.)

3. **Run the app**

   ```bash
   go run .
   ```

   Or build and run:

   ```bash
   go build -o smdb .
   ./smdb
   ```

## Keybindings

| Key                | Action                    |
| ------------------ | ------------------------- |
| `↑` / `k`          | Move up                   |
| `↓` / `j`          | Move down                 |
| `Enter`            | Open movie details        |
| `Esc` / `Backspace`| Go back                   |
| `Tab` / `1`–`4`    | Switch category           |
| `?`                | Toggle help               |
| `q` / `Ctrl+C`     | Quit                      |

## Project Structure

```
SMDB/
├── main.go                  # Entry point
├── internal/
│   ├── tui/
│   │   ├── model.go         # Bubble Tea model
│   │   ├── update.go        # Message handling
│   │   ├── view.go          # Rendering
│   │   ├── keys.go          # Key bindings
│   │   └── styles.go        # Lip Gloss styles
│   ├── api/
│   │   ├── client.go        # TMDB HTTP client
│   │   └── types.go         # API response structs
│   └── config/
│       └── config.go        # Environment config
├── initialRequirement.md
└── README.md
```

## License

MIT
