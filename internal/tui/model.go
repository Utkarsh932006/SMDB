package tui

import (
	"github.com/Utkarsh932006/SMDB/internal/api"
	"github.com/charmbracelet/bubbles/spinner"
)

// Category represents a movie category supported by the TMDB API.
type Category int

const (
	CategoryNowPlaying Category = iota
	CategoryPopular
	CategoryTopRated
	CategoryUpcoming
)

// Categories lists all supported movie categories in order.
var Categories = []Category{
	CategoryNowPlaying,
	CategoryPopular,
	CategoryTopRated,
	CategoryUpcoming,
}

// CategoryNames maps category enums to human-readable names.
var CategoryNames = map[Category]string{
	CategoryNowPlaying: "Now Playing",
	CategoryPopular:    "Popular",
	CategoryTopRated:   "Top Rated",
	CategoryUpcoming:   "Upcoming",
}

// viewState represents which screen is currently visible.
type viewState int

const (
	viewList viewState = iota
	viewDetail
)

// moviesLoadedMsg is sent when a category's movies have been fetched from the TMDB API.
type moviesLoadedMsg struct {
	category Category
	movies   []api.Movie
	err      error
}

// Model represents the main Bubble Tea application state.
type Model struct {
	client       *api.Client
	category     Category
	movies       map[Category][]api.Movie
	cursor       map[Category]int
	scrollOffset map[Category]int
	loading      map[Category]bool
	err          error

	spinner  spinner.Model
	view     viewState
	showHelp bool

	width  int
	height int

	keys   KeyMap
	styles Styles
}

// NewModel creates an initialized Model ready for Bubble Tea.
func NewModel(client *api.Client) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	styles := DefaultStyles()
	s.Style = styles.Spinner

	m := Model{
		client:       client,
		category:     CategoryNowPlaying,
		movies:       make(map[Category][]api.Movie),
		cursor:       make(map[Category]int),
		scrollOffset: make(map[Category]int),
		loading:      make(map[Category]bool),
		spinner:      s,
		view:         viewList,
		keys:         DefaultKeyMap(),
		styles:       styles,
	}

	return m
}

// CurrentMovies returns the movies for the currently active category.
func (m *Model) CurrentMovies() []api.Movie {
	return m.movies[m.category]
}

// SelectedMovie returns a pointer to the currently highlighted movie, or nil if none.
func (m *Model) SelectedMovie() *api.Movie {
	movies := m.CurrentMovies()
	idx := m.cursor[m.category]
	if len(movies) == 0 || idx < 0 || idx >= len(movies) {
		return nil
	}
	return &movies[idx]
}
