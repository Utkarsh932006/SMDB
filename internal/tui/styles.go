package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Styles holds all Lip Gloss style definitions for SMDB.
type Styles struct {
	// App Header
	AppTitle lipgloss.Style

	// Tabs
	TabActive   lipgloss.Style
	TabInactive lipgloss.Style
	TabBar      lipgloss.Style

	// List
	ListItem         lipgloss.Style
	ListItemSelected lipgloss.Style
	ItemCursor       lipgloss.Style
	ItemTitle        lipgloss.Style
	ItemTitleSel     lipgloss.Style
	ItemDate         lipgloss.Style
	ItemRating       lipgloss.Style

	// Detail View
	DetailContainer lipgloss.Style
	DetailTitle     lipgloss.Style
	DetailMetaLabel lipgloss.Style
	DetailMetaValue lipgloss.Style
	DetailOverview  lipgloss.Style
	DetailBackHint  lipgloss.Style

	// Help View
	HelpBox   lipgloss.Style
	HelpTitle lipgloss.Style
	HelpKey   lipgloss.Style
	HelpDesc  lipgloss.Style

	// Error & Loading
	Spinner lipgloss.Style
	Loading lipgloss.Style
	Error   lipgloss.Style

	// Footer
	FooterBar lipgloss.Style
	FooterKey lipgloss.Style
	FooterDim lipgloss.Style
}

// DefaultStyles returns standard styles using adaptive colors suitable for both
// light and dark terminal backgrounds.
func DefaultStyles() Styles {
	// Adaptive color definitions
	primary := lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7D56F4"}
	accent := lipgloss.AdaptiveColor{Light: "#0284C7", Dark: "#38BDF8"}
	gold := lipgloss.AdaptiveColor{Light: "#D97706", Dark: "#FBBF24"}
	fg := lipgloss.AdaptiveColor{Light: "#1F2937", Dark: "#F3F4F6"}
	fgDim := lipgloss.AdaptiveColor{Light: "#6B7280", Dark: "#9CA3AF"}
	border := lipgloss.AdaptiveColor{Light: "#D1D5DB", Dark: "#4B5563"}
	errColor := lipgloss.AdaptiveColor{Light: "#DC2626", Dark: "#EF4444"}
	selectedBg := lipgloss.AdaptiveColor{Light: "#E5E7EB", Dark: "#2D3748"}

	return Styles{
		AppTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(primary).
			MarginBottom(1),

		TabActive: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(primary).
			Padding(0, 2),

		TabInactive: lipgloss.NewStyle().
			Foreground(fgDim).
			Padding(0, 2),

		TabBar: lipgloss.NewStyle().
			BorderBottom(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottomForeground(border).
			MarginBottom(1),

		ListItem: lipgloss.NewStyle().
			Padding(0, 1),

		ListItemSelected: lipgloss.NewStyle().
			Background(selectedBg).
			Padding(0, 1).
			Bold(true),

		ItemCursor: lipgloss.NewStyle().
			Foreground(primary).
			Bold(true),

		ItemTitle: lipgloss.NewStyle().
			Foreground(fg),

		ItemTitleSel: lipgloss.NewStyle().
			Foreground(primary).
			Bold(true),

		ItemDate: lipgloss.NewStyle().
			Foreground(fgDim),

		ItemRating: lipgloss.NewStyle().
			Foreground(gold).
			Bold(true),

		DetailContainer: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(accent).
			Padding(1, 2).
			Margin(0, 1),

		DetailTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(primary).
			MarginBottom(1),

		DetailMetaLabel: lipgloss.NewStyle().
			Bold(true).
			Foreground(fgDim),

		DetailMetaValue: lipgloss.NewStyle().
			Foreground(fg),

		DetailOverview: lipgloss.NewStyle().
			Foreground(fg).
			MarginTop(1),

		DetailBackHint: lipgloss.NewStyle().
			Foreground(fgDim).
			Italic(true).
			MarginBottom(1),

		HelpBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primary).
			Padding(1, 2),

		HelpTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(primary).
			MarginBottom(1),

		HelpKey: lipgloss.NewStyle().
			Bold(true).
			Foreground(accent),

		HelpDesc: lipgloss.NewStyle().
			Foreground(fgDim),

		Spinner: lipgloss.NewStyle().
			Foreground(primary),

		Loading: lipgloss.NewStyle().
			Foreground(fgDim).
			Italic(true).
			Padding(1, 2),

		Error: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(errColor).
			Foreground(errColor).
			Padding(1, 2).
			Margin(1, 0),

		FooterBar: lipgloss.NewStyle().
			BorderTop(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderTopForeground(border).
			Padding(0, 1).
			MarginTop(1),

		FooterKey: lipgloss.NewStyle().
			Foreground(primary).
			Bold(true),

		FooterDim: lipgloss.NewStyle().
			Foreground(fgDim),
	}
}
