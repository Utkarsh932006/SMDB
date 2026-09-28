package tui

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// View renders the current user interface according to the Model state.
func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Initializing SMDB..."
	}

	if m.showHelp {
		return m.helpView()
	}

	header := m.headerView()
	var body string

	switch {
	case m.err != nil:
		body = m.errorView()
	case m.loading[m.category]:
		body = m.loadingView()
	case m.view == viewDetail:
		body = m.detailView()
	default:
		body = m.listView()
	}

	footer := m.footerView()

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

func (m Model) headerView() string {
	title := m.styles.AppTitle.Render("🎬 SMDB — Simple Movie DataBrowser")

	var tabs []string
	for i, cat := range Categories {
		tabLabel := fmt.Sprintf("[%d] %s", i+1, CategoryNames[cat])
		if cat == m.category {
			tabs = append(tabs, m.styles.TabActive.Render(tabLabel))
		} else {
			tabs = append(tabs, m.styles.TabInactive.Render(tabLabel))
		}
	}

	tabRow := lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
	tabBar := m.styles.TabBar.Width(m.width - 2).Render(tabRow)

	return lipgloss.JoinVertical(lipgloss.Left, title, tabBar)
}

func (m Model) listView() string {
	movies := m.CurrentMovies()
	if len(movies) == 0 {
		return m.styles.Loading.Render("No movies found in this category.")
	}

	start := m.scrollOffset[m.category]
	pageSize := m.visibleListHeight()
	end := start + pageSize
	if end > len(movies) {
		end = len(movies)
	}

	// Calculate space for title
	// Total line = prefix (3) + title (flexible) + rating (12) + date (14) + padding
	fixedWidth := 3 + 12 + 14 + 4
	titleWidth := m.width - fixedWidth
	if titleWidth < 15 {
		titleWidth = 15
	}

	var rows []string
	for i := start; i < end; i++ {
		movie := movies[i]
		isSelected := i == m.cursor[m.category]

		// Selection cursor
		cursorStr := "  "
		if isSelected {
			cursorStr = "▸ "
		}

		// Truncate title
		titleStr := truncateString(movie.Title, titleWidth)
		// Pad title to uniform width for column alignment
		titlePadded := lipgloss.NewStyle().Width(titleWidth).Render(titleStr)

		// Visual stars rating
		starsStr := renderStars(movie.VoteAverage)

		// Release date
		dateStr := movie.ReleaseDate
		if dateStr == "" {
			dateStr = "N/A"
		}
		dateStr = fmt.Sprintf("(%s)", dateStr)

		// Apply styles
		c := m.styles.ItemCursor.Render(cursorStr)
		var t string
		if isSelected {
			t = m.styles.ItemTitleSel.Render(titlePadded)
		} else {
			t = m.styles.ItemTitle.Render(titlePadded)
		}
		r := m.styles.ItemRating.Render(fmt.Sprintf("%-11s", starsStr))
		d := m.styles.ItemDate.Render(fmt.Sprintf("%-13s", dateStr))

		rowContent := fmt.Sprintf("%s%s  %s  %s", c, t, r, d)

		if isSelected {
			rows = append(rows, m.styles.ListItemSelected.Width(m.width-4).Render(rowContent))
		} else {
			rows = append(rows, m.styles.ListItem.Width(m.width-4).Render(rowContent))
		}
	}

	// Pad remaining vertical lines if needed so layout stays stable
	for len(rows) < pageSize {
		rows = append(rows, "")
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (m Model) detailView() string {
	movie := m.SelectedMovie()
	if movie == nil {
		return m.styles.Loading.Render("No movie selected.")
	}

	backHint := m.styles.DetailBackHint.Render("← Press Esc or Backspace to return to list")
	title := m.styles.DetailTitle.Render(movie.Title)

	stars := renderStars(movie.VoteAverage)
	rating := fmt.Sprintf("%s (%s / %d votes)", stars, fmt.Sprintf("%.1f/10", movie.VoteAverage), movie.VoteCount)

	releaseDate := movie.ReleaseDate
	if releaseDate == "" {
		releaseDate = "Unknown"
	}

	lang := strings.ToUpper(movie.OriginalLanguage)
	if lang == "" {
		lang = "Unknown"
	}

	pop := fmt.Sprintf("%.2f", movie.Popularity)

	contentWidth := m.width - 8
	if contentWidth < 30 {
		contentWidth = 30
	}

	metaLines := []string{
		fmt.Sprintf("%s %s", m.styles.DetailMetaLabel.Render("⭐ Rating:        "), m.styles.DetailMetaValue.Render(rating)),
		fmt.Sprintf("%s %s", m.styles.DetailMetaLabel.Render("📅 Release Date:  "), m.styles.DetailMetaValue.Render(releaseDate)),
		fmt.Sprintf("%s %s", m.styles.DetailMetaLabel.Render("🌐 Language:      "), m.styles.DetailMetaValue.Render(lang)),
		fmt.Sprintf("%s %s", m.styles.DetailMetaLabel.Render("🔥 Popularity:    "), m.styles.DetailMetaValue.Render(pop)),
	}

	metaBlock := strings.Join(metaLines, "\n")

	overviewText := movie.Overview
	if overviewText == "" {
		overviewText = "No synopsis available."
	}

	synopsisHeader := m.styles.DetailMetaLabel.Render("📖 Overview / Synopsis:")
	synopsisBody := m.styles.DetailOverview.Width(contentWidth).Render(overviewText)

	detailContent := lipgloss.JoinVertical(
		lipgloss.Left,
		backHint,
		title,
		metaBlock,
		"",
		synopsisHeader,
		synopsisBody,
	)

	return m.styles.DetailContainer.Width(m.width - 4).Render(detailContent)
}

func (m Model) loadingView() string {
	catName := CategoryNames[m.category]
	msg := fmt.Sprintf("%s Fetching %s movies from TMDB...", m.spinner.View(), catName)
	return m.styles.Loading.Render(msg)
}

func (m Model) errorView() string {
	errTitle := lipgloss.NewStyle().Bold(true).Render("❌ Error fetching movies")
	errMsg := fmt.Sprintf("Failed to load %s:\n%s", CategoryNames[m.category], m.err.Error())
	hint := lipgloss.NewStyle().Italic(true).Render("\n• Press Enter to retry\n• Press 1-4 or Tab to switch categories\n• Press q to quit")

	content := lipgloss.JoinVertical(lipgloss.Left, errTitle, errMsg, hint)
	return m.styles.Error.Width(m.width - 4).Render(content)
}

func (m Model) footerView() string {
	var leftHint string
	if m.view == viewDetail {
		leftHint = m.styles.FooterDim.Render("Esc: Back • Tab/1-4: Switch Category • ?: Help • q: Quit")
	} else if m.err != nil {
		leftHint = m.styles.FooterDim.Render("Enter: Retry • Tab/1-4: Switch Category • ?: Help • q: Quit")
	} else {
		leftHint = m.styles.FooterDim.Render("↑/k, ↓/j: Navigate • Enter: Details • Tab/1-4: Category • ?: Help • q: Quit")
	}

	// Right item counter
	var rightCounter string
	movies := m.CurrentMovies()
	if len(movies) > 0 && m.view == viewList {
		rightCounter = m.styles.FooterDim.Render(fmt.Sprintf("%d/%d", m.cursor[m.category]+1, len(movies)))
	}

	space := m.width - lipgloss.Width(leftHint) - lipgloss.Width(rightCounter) - 4
	if space < 1 {
		space = 1
	}

	bar := leftHint + strings.Repeat(" ", space) + rightCounter
	return m.styles.FooterBar.Width(m.width - 2).Render(bar)
}

func (m Model) helpView() string {
	title := m.styles.HelpTitle.Render("🎬 SMDB — Keyboard Shortcuts")

	shortcuts := [][]string{
		{"↑ / k", "Move cursor up"},
		{"↓ / j", "Move cursor down"},
		{"Enter", "View movie details / Retry on error"},
		{"Esc / Backspace", "Return to movie list"},
		{"Tab", "Next category"},
		{"Shift+Tab", "Previous category"},
		{"1", "Now Playing category"},
		{"2", "Popular category"},
		{"3", "Top Rated category"},
		{"4", "Upcoming category"},
		{"?", "Toggle this help dialog"},
		{"q / Ctrl+C", "Quit SMDB"},
	}

	var lines []string
	lines = append(lines, title, "")

	for _, s := range shortcuts {
		k := m.styles.HelpKey.Width(18).Render(s[0])
		d := m.styles.HelpDesc.Render(s[1])
		lines = append(lines, fmt.Sprintf("%s %s", k, d))
	}

	lines = append(lines, "", m.styles.FooterDim.Render("Press ? or Esc to close help"))

	content := strings.Join(lines, "\n")
	return m.styles.HelpBox.Width(m.width - 6).Render(content)
}

func renderStars(vote float64) string {
	starsCount := int(math.Round(vote / 2.0))
	if starsCount > 5 {
		starsCount = 5
	}
	if starsCount < 0 {
		starsCount = 0
	}

	filled := strings.Repeat("★", starsCount)
	empty := strings.Repeat("☆", 5-starsCount)
	return fmt.Sprintf("%s%s %4.1f", filled, empty, vote)
}

func truncateString(s string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return string(runes[:maxLen])
	}
	return string(runes[:maxLen-1]) + "…"
}
