package components

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tearingItUp786/nekot/util"
)

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(util.ListItemPaddingLeft)
	selectedItemStyle = lipgloss.
				NewStyle().
				PaddingLeft(util.ListRightShiftedItemPadding)
	activeItemStyle = itemStyle.Copy()
)

type SessionListItem struct {
	Id       int
	Text     string
	IsActive bool
}

type SessionsList struct {
	list list.Model
}

func (i SessionListItem) FilterValue() string { return "" }

type sessionItemDelegate struct{}

func (d sessionItemDelegate) Height() int  { return 1 }
func (d sessionItemDelegate) Spacing() int { return 0 }
func (d sessionItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	var cmds []tea.Cmd

	return tea.Batch(cmds...)
}

func (d sessionItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(SessionListItem)
	if !ok {
		log.Println("not okay")
		return
	}

	str := fmt.Sprintf("%s", i.Text)
	str = util.TrimListItem(str, m.Width())

	fn := itemStyle.Render
	selectedRender := selectedItemStyle.Render

	if i.IsActive {
		fn = activeItemStyle.Render
		selectedRender = activeItemStyle.Render
	}

	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedRender("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func (l *SessionsList) GetSelectedItem() (SessionListItem, bool) {
	item, ok := l.list.SelectedItem().(SessionListItem)
	return item, ok
}

func (l *SessionsList) SetItems(items []list.Item) {
	l.list.SetItems(items)
}

func (l *SessionsList) SetSize(w, h int) {
	l.list.SetWidth(w)
	l.list.SetHeight(h)
}

func (l SessionsList) GetWidth() int {
	return l.list.Width()
}

func (l SessionsList) Update(msg tea.Msg) (SessionsList, tea.Cmd) {
	var cmd tea.Cmd
	l.list, cmd = l.list.Update(msg)
	return l, cmd
}

func NewSessionsList(items []list.Item, w, h int, colors util.SchemeColors) SessionsList {
	l := list.New(items, sessionItemDelegate{}, w, h)

	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.DisableQuitKeybindings()

	l.Paginator.ActiveDot = lipgloss.NewStyle().Foreground(colors.HighlightColor).Render("■")
	l.Paginator.InactiveDot = lipgloss.NewStyle().Foreground(colors.DefaultTextColor).Render("•")
	selectedItemStyle = selectedItemStyle.Copy().Foreground(colors.AccentColor)
	activeItemStyle = activeItemStyle.Copy().Foreground(colors.HighlightColor)
	itemStyle = itemStyle.Copy().Foreground(colors.DefaultTextColor)

	return SessionsList{
		list: l,
	}
}

func (l *SessionsList) EditListView(paneHeight int) string {
	l.list.SetHeight(paneHeight)
	return lipgloss.
		NewStyle().
		MaxHeight(paneHeight).
		PaddingLeft(util.DefaultElementsPadding).
		Render(l.list.View())
}
