package tui

import (
	"strings"
	"testing"

	"github.com/Utkarsh932006/SMDB/internal/api"
	tea "github.com/charmbracelet/bubbletea"
)

func TestRenderStars(t *testing.T) {
	tests := []struct {
		vote     float64
		contains string
	}{
		{0.0, "☆☆☆☆☆"},
		{5.0, "★★★☆☆"},
		{8.0, "★★★★☆"},
		{10.0, "★★★★★"},
	}

	for _, tt := range tests {
		res := renderStars(tt.vote)
		if !strings.Contains(res, tt.contains) {
			t.Errorf("renderStars(%f) = %q, expected to contain %q", tt.vote, res, tt.contains)
		}
	}
}

func TestTruncateString(t *testing.T) {
	tests := []struct {
		input    string
		maxLen   int
		expected string
	}{
		{"Short", 10, "Short"},
		{"ExactLength", 11, "ExactLength"},
		{"A very long title that should truncate", 15, "A very long ti…"},
		{"Hi", 0, ""},
	}

	for _, tt := range tests {
		res := truncateString(tt.input, tt.maxLen)
		if res != tt.expected {
			t.Errorf("truncateString(%q, %d) = %q, expected %q", tt.input, tt.maxLen, res, tt.expected)
		}
	}
}

func TestModelNavigation(t *testing.T) {
	client := api.NewClient("test-token")
	model := NewModel(client)

	// Simulate window size
	newM, _ := model.Update(tea.WindowSizeMsg{Width: 100, Height: 40})
	model = newM.(Model)

	// Inject sample movies into CategoryNowPlaying
	sampleMovies := []api.Movie{
		{ID: 1, Title: "Movie 1", VoteAverage: 8.5},
		{ID: 2, Title: "Movie 2", VoteAverage: 7.0},
		{ID: 3, Title: "Movie 3", VoteAverage: 6.5},
	}
	model.movies[CategoryNowPlaying] = sampleMovies

	if model.cursor[CategoryNowPlaying] != 0 {
		t.Fatalf("expected initial cursor 0, got %d", model.cursor[CategoryNowPlaying])
	}

	// Move down
	newM, _ = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
	model = newM.(Model)
	if model.cursor[CategoryNowPlaying] != 1 {
		t.Errorf("expected cursor 1 after 'j', got %d", model.cursor[CategoryNowPlaying])
	}

	// Move down again
	newM, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	model = newM.(Model)
	if model.cursor[CategoryNowPlaying] != 2 {
		t.Errorf("expected cursor 2 after Down arrow, got %d", model.cursor[CategoryNowPlaying])
	}

	// Move down at boundary
	newM, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
	model = newM.(Model)
	if model.cursor[CategoryNowPlaying] != 2 {
		t.Errorf("expected cursor to stay at 2 at bottom boundary, got %d", model.cursor[CategoryNowPlaying])
	}

	// Move up
	newM, _ = model.Update(tea.KeyMsg{Type: tea.KeyUp})
	model = newM.(Model)
	if model.cursor[CategoryNowPlaying] != 1 {
		t.Errorf("expected cursor 1 after Up arrow, got %d", model.cursor[CategoryNowPlaying])
	}

	// Open detail view with Enter
	newM, _ = model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	model = newM.(Model)
	if model.view != viewDetail {
		t.Fatalf("expected view to be viewDetail after Enter, got %d", model.view)
	}

	// Ensure view renders movie title in details
	viewStr := model.View()
	if !strings.Contains(viewStr, "Movie 2") {
		t.Errorf("expected detail view to contain 'Movie 2', got:\n%s", viewStr)
	}

	// Return to list with Esc
	newM, _ = model.Update(tea.KeyMsg{Type: tea.KeyEsc})
	model = newM.(Model)
	if model.view != viewList {
		t.Fatalf("expected view to be viewList after Esc, got %d", model.view)
	}

	// Switch category with '2' (Popular)
	newM, _ = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'2'}})
	model = newM.(Model)
	if model.category != CategoryPopular {
		t.Errorf("expected category Popular after pressing '2', got %v", model.category)
	}
}

func TestModelHelpToggle(t *testing.T) {
	client := api.NewClient("test-token")
	model := NewModel(client)
	newM, _ := model.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	model = newM.(Model)

	// Toggle help
	newM, _ = model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'?'}})
	model = newM.(Model)
	if !model.showHelp {
		t.Fatal("expected showHelp to be true after pressing '?'")
	}

	rendered := model.View()
	if !strings.Contains(rendered, "Keyboard Shortcuts") {
		t.Errorf("expected help view to render 'Keyboard Shortcuts', got:\n%s", rendered)
	}

	// Close help
	newM, _ = model.Update(tea.KeyMsg{Type: tea.KeyEsc})
	model = newM.(Model)
	if model.showHelp {
		t.Fatal("expected showHelp to be false after pressing Esc")
	}
}
