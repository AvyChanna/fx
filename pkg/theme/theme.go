package theme

import (
	"github.com/charmbracelet/lipgloss"
)

type Theme struct {
	Cursor    Color
	Syntax    Color
	Preview   Color
	StatusBar Color
	Search    Color
	Key       func(i, len int) Color
	String    Color
	Null      Color
	Boolean   Color
	Number    Color
}
type Color func(s string) string

var (
	defaultCursor    = lipgloss.NewStyle().Reverse(true).Render
	defaultPreview   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("8")).Render
	defaultStatusBar = lipgloss.NewStyle().Background(lipgloss.Color("7")).Foreground(lipgloss.Color("0")).Render
	defaultSearch    = lipgloss.NewStyle().Background(lipgloss.Color("11")).Foreground(lipgloss.Color("16")).Render
)

var BWTheme = Theme{
	Cursor:    defaultCursor,
	Syntax:    noColor,
	Preview:   noColor,
	StatusBar: noColor,
	Search:    defaultSearch,
	Key:       func(_, _ int) Color { return noColor },
	String:    noColor,
	Null:      noColor,
	Boolean:   noColor,
	Number:    noColor,
}

var ColorTheme = Theme{
	Cursor:    defaultCursor,
	Syntax:    noColor,
	Preview:   defaultPreview,
	StatusBar: defaultStatusBar,
	Search:    defaultSearch,
	Key:       func(_, _ int) Color { return fg("#0072cf") },
	String:    fg("#6BCB77"),
	Null:      fg("#666666"),
	Boolean:   fg("#FF6B6B"),
	Number:    fg("#FFD93D"),
}

func noColor(s string) string {
	return s
}

func fg(color string) Color {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Render
}
