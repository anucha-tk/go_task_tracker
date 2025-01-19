package style

import (
	"github.com/charmbracelet/lipgloss"
)

func ErrorStyle() lipgloss.Style {
	return lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("1"))
}

func SuccessStyle() lipgloss.Style {
	return lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#04B575"))
}
