package sessions

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/tearingItUp786/chatgpt-tui/util"
)

func (m *Model) _settingsContainer() lipgloss.Style {
	borderColor := util.NormalTabBorderColor

	if m.isFocused {
		borderColor = util.ActiveTabBorderColor
	}

	container := lipgloss.NewStyle().
		AlignVertical(lipgloss.Top).
		Border(lipgloss.ThickBorder(), true).
		BorderForeground(borderColor)

	return container
}

func listHeader(str ...string) string {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		MarginLeft(2).
		Render(str...)
}

func listItem(heading string, value string, isActive bool) string {
	headingColor := util.Pink100
	color := "#bbb"
	if isActive {
		colorValue := util.Pink200
		color = colorValue
		headingColor = colorValue
	}
	headingEl := lipgloss.NewStyle().
		PaddingLeft(2).
		Foreground(lipgloss.Color(headingColor)).
		Render
	spanEl := lipgloss.NewStyle().
		Foreground(lipgloss.Color(color)).
		Render

	return headingEl(" "+heading, spanEl(value))
}

func (m Model) normalListView() string {
	sessionListItems := []string{}
	for _, session := range m.AllSessions {
		isCurrentSession := m.CurrentSessionID == session.ID
		sessionListItems = append(
			sessionListItems,
			listItem(fmt.Sprint(session.ID), session.SessionName, isCurrentSession),
		)
	}

	return lipgloss.NewStyle().
		// TODO: figure out how to get height from the settings model
		Height(m.terminalHeight - 22).
		MaxHeight(m.terminalHeight - 22).
		Render(strings.Join(sessionListItems, "\n"))
}
