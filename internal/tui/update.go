package tui

import (
	"context"
	"time"

	"github.com/Utkarsh932006/SMDB/internal/api"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the Bubble Tea program, starting the spinner and loading initial data.
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.fetchCategory(m.category),
	)
}

// Update handles incoming messages and events, returning the updated model and any commands.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case moviesLoadedMsg:
		m.loading[msg.category] = false
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}
		m.err = nil
		m.movies[msg.category] = msg.movies
		return m, nil

	case tea.KeyMsg:
		return m.handleKeyPress(msg)
	}

	return m, nil
}

func (m Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Global Quit
	if key.Matches(msg, m.keys.Quit) {
		return m, tea.Quit
	}

	// Help Toggle
	if key.Matches(msg, m.keys.Help) {
		m.showHelp = !m.showHelp
		return m, nil
	}

	// When help is visible, any key closes help overlay
	if m.showHelp {
		if key.Matches(msg, m.keys.Back) || key.Matches(msg, m.keys.Enter) {
			m.showHelp = false
		}
		return m, nil
	}

	// Direct category switching (1-4) works in both list and detail views
	switch {
	case key.Matches(msg, m.keys.Cat1):
		return m.switchCategory(CategoryNowPlaying)
	case key.Matches(msg, m.keys.Cat2):
		return m.switchCategory(CategoryPopular)
	case key.Matches(msg, m.keys.Cat3):
		return m.switchCategory(CategoryTopRated)
	case key.Matches(msg, m.keys.Cat4):
		return m.switchCategory(CategoryUpcoming)
	case key.Matches(msg, m.keys.Tab):
		return m.nextCategory()
	case key.Matches(msg, m.keys.PrevTab):
		return m.prevCategory()
	}

	// Detail View navigation
	if m.view == viewDetail {
		if key.Matches(msg, m.keys.Back) {
			m.view = viewList
			return m, nil
		}
		return m, nil
	}

	// List View navigation
	if m.view == viewList {
		// Retry on Enter when an error is present
		if m.err != nil {
			if key.Matches(msg, m.keys.Enter) {
				return m, m.fetchCategory(m.category)
			}
			return m, nil
		}

		movies := m.CurrentMovies()
		currIdx := m.cursor[m.category]

		switch {
		case key.Matches(msg, m.keys.Up):
			if currIdx > 0 {
				m.cursor[m.category] = currIdx - 1
				m.adjustScroll()
			}
			return m, nil

		case key.Matches(msg, m.keys.Down):
			if currIdx < len(movies)-1 {
				m.cursor[m.category] = currIdx + 1
				m.adjustScroll()
			}
			return m, nil

		case key.Matches(msg, m.keys.Enter):
			if len(movies) > 0 {
				m.view = viewDetail
			}
			return m, nil
		}
	}

	return m, nil
}

func (m *Model) adjustScroll() {
	curr := m.cursor[m.category]
	offset := m.scrollOffset[m.category]
	pageSize := m.visibleListHeight()

	if pageSize <= 0 {
		pageSize = 10
	}

	if curr < offset {
		m.scrollOffset[m.category] = curr
	} else if curr >= offset+pageSize {
		m.scrollOffset[m.category] = curr - pageSize + 1
	}
}

func (m *Model) visibleListHeight() int {
	// Account for Header (app title + tabs + margins ≈ 4 lines)
	// and Footer (status/keys ≈ 3 lines)
	h := m.height - 7
	if h < 5 {
		return 5
	}
	return h
}

func (m Model) switchCategory(cat Category) (tea.Model, tea.Cmd) {
	if m.category == cat && m.view == viewList {
		return m, nil
	}
	m.category = cat
	m.view = viewList

	// If already cached, switch instantly without refetching
	if _, ok := m.movies[cat]; ok {
		m.err = nil
		return m, nil
	}

	return m, m.fetchCategory(cat)
}

func (m Model) nextCategory() (tea.Model, tea.Cmd) {
	next := (int(m.category) + 1) % len(Categories)
	return m.switchCategory(Category(next))
}

func (m Model) prevCategory() (tea.Model, tea.Cmd) {
	prev := (int(m.category) - 1 + len(Categories)) % len(Categories)
	return m.switchCategory(Category(prev))
}

func (m *Model) fetchCategory(cat Category) tea.Cmd {
	m.loading[cat] = true
	m.err = nil
	client := m.client

	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		var resp *api.MovieListResponse
		var err error

		switch cat {
		case CategoryNowPlaying:
			resp, err = client.GetNowPlaying(ctx)
		case CategoryPopular:
			resp, err = client.GetPopular(ctx)
		case CategoryTopRated:
			resp, err = client.GetTopRated(ctx)
		case CategoryUpcoming:
			resp, err = client.GetUpcoming(ctx)
		}

		if err != nil {
			return moviesLoadedMsg{category: cat, err: err}
		}

		return moviesLoadedMsg{category: cat, movies: resp.Results}
	}
}
