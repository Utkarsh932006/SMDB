# SMDB — Initial Requirements

> **S**imple Movie DataBrowser — a TUI for browsing movies via TMDB

## 1. Project Overview

Build an interactive Terminal User Interface (TUI) that fetches movie data from
[The Movie Database (TMDB) API](https://developer.themoviedb.org/) and presents
it in a browsable, keyboard-driven interface inside the terminal.

This is a TUI adaptation of the [roadmap.sh TMDB CLI project](https://roadmap.sh/projects/tmdb-cli).

### Why a TUI instead of a CLI?

| CLI (original)                     | TUI (this project)                        |
| ---------------------------------- | ----------------------------------------- |
| One-shot output, then exits        | Persistent, interactive session            |
| Flags select category              | Tab/key to switch categories live          |
| Plain text list                    | Styled panels, colors, scrollable lists    |
| No navigation                     | Keyboard-first: arrows, vim keys, search   |

---

## 2. Tech Stack

| Component        | Choice                                                      |
| ---------------- | ------------------------------------------------------------ |
| **Language**     | Go                                                           |
| **TUI Framework**| [Bubble Tea](https://github.com/charmbracelet/bubbletea) (Elm-architecture TUI framework) |
| **Styling**      | [Lip Gloss](https://github.com/charmbracelet/lipgloss)       |
| **HTTP Client**  | `net/http` (stdlib)                                          |
| **JSON Parsing** | `encoding/json` (stdlib)                                     |
| **API**          | [TMDB API v3](https://developer.themoviedb.org/reference/intro/getting-started) |

---

## 3. Core Features (MVP)

### 3.1 Movie Categories

The app must support browsing **four** movie categories, matching the original
project scope:

| Category       | TMDB Endpoint                           |
| -------------- | --------------------------------------- |
| Now Playing    | `GET /movie/now_playing`                |
| Popular        | `GET /movie/popular`                    |
| Top Rated      | `GET /movie/top_rated`                  |
| Upcoming       | `GET /movie/upcoming`                   |

### 3.2 Movie List View

- Display a scrollable list of movies for the selected category.
- Each list item shows at minimum:
  - **Title**
  - **Release Date**
  - **Rating** (vote_average) — rendered as a visual indicator (e.g. `★★★★☆`)
- Highlight the currently selected movie.

### 3.3 Movie Detail View

When a movie is selected (Enter), show a detail panel/page with:

- Title
- Overview / Synopsis
- Release Date
- Rating (vote_average / vote_count)
- Original Language
- Popularity score

### 3.4 Category Switching

- Use **Tab** / **Shift+Tab** or number keys (`1`–`4`) to switch between
  categories.
- A visible tab bar or header indicates the active category.

### 3.5 Keyboard Navigation

| Key              | Action                              |
| ---------------- | ----------------------------------- |
| `↑` / `k`        | Move selection up                   |
| `↓` / `j`        | Move selection down                 |
| `Enter`          | Open detail view for selected movie |
| `Esc` / `Backspace` | Go back from detail to list      |
| `Tab` / `1-4`    | Switch category                     |
| `q` / `Ctrl+C`   | Quit                                |
| `?`              | Toggle help overlay                 |

---

## 4. API Integration

### 4.1 Authentication

- TMDB uses a **Bearer token** (API Read Access Token) or an **API key** query
  param.
- The app reads the token from the environment variable `TMDB_API_TOKEN`.
- If the variable is missing, display a clear error message with setup
  instructions and exit gracefully.

### 4.2 Request Handling

- All API calls are `GET` requests returning JSON.
- Parse responses into Go structs.
- Implement a simple **loading state** — show a spinner or "Loading…" indicator
  while fetching.

### 4.3 Error Handling

Handle the following gracefully (display in-TUI error, don't crash):

- Missing / invalid API token
- Network unreachable
- API rate limiting (HTTP 429)
- Malformed JSON response
- Non-200 HTTP status codes

---

## 5. UI / UX Guidelines

1. **Layout**: Fixed header (app title + category tabs) → scrollable list body →
   fixed footer (keybinding hints).
2. **Colors**: Use Lip Gloss adaptive colors so the TUI looks good on both dark
   and light terminal backgrounds.
3. **Resize**: Handle terminal resize — the layout must reflow on terminal resize.
   Bubble Tea handles `SIGWINCH` via `WindowSizeMsg` automatically.
4. **Minimal chrome**: No excessive borders or decoration. Use subtle
   box-drawing characters only where they aid readability.
5. **Responsive text**: Truncate long titles with `…` rather than wrapping
   mid-word.

---

## 6. Project Structure (Proposed)

```
SMDB/
├── main.go                 # Entry point, initialise Bubble Tea program
├── go.mod
├── go.sum
├── internal/
│   ├── tui/
│   │   ├── model.go        # Top-level Bubble Tea model
│   │   ├── update.go       # Msg handling / state transitions
│   │   ├── view.go         # Rendering logic
│   │   ├── keys.go         # Key bindings
│   │   └── styles.go       # Lip Gloss styles
│   ├── api/
│   │   ├── client.go       # TMDB HTTP client
│   │   └── types.go        # API response structs
│   └── config/
│       └── config.go       # Env var loading, validation
├── initialRequirement.md
└── README.md
```

---

## 7. Stretch Goals (Post-MVP)

These are **not** required for the initial version but are natural extensions:

| Feature                 | Description                                           |
| ----------------------- | ----------------------------------------------------- |
| **Search**              | `/` to open a search bar, query `GET /search/movie`   |
| **Pagination**          | Load next page of results on scroll-to-bottom         |
| **Caching**             | Cache API responses in-memory with TTL to reduce calls |
| **Genre filter**        | Filter the current list by genre tags                 |
| **Movie poster art**    | Render low-res poster via terminal sixel / kitty protocol (very stretch) |

---

## 8. Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- A free [TMDB account](https://www.themoviedb.org/signup) and an
  [API Read Access Token](https://www.themoviedb.org/settings/api)
- A terminal emulator with 256-color or TrueColor support (most modern
  terminals)

---

## 9. Success Criteria

The MVP is considered **done** when:

- [ ] The app compiles and runs with `go run .`
- [ ] All four movie categories are browsable
- [ ] Movie detail view works for any selected movie
- [ ] Category switching is instant (no full re-render flicker)
- [ ] Errors (no token, no network) are shown in the TUI, not panics
- [ ] Keyboard navigation matches the table in §3.5
- [ ] The app exits cleanly on `q` / `Ctrl+C`
- [ ] A `README.md` explains setup and usage
