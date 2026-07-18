package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/oddman/mandrill/internal/tui/styles"
)

var mandrillASCII = `
8b    d8    db    88b 88 8888b.  88""Yb 88 88     88
88b  d88   dPYb   88Yb88  8I  Yb 88__dP 88 88     88
88YbdP88  dP__Yb  88 Y88  8I  dY 88"Yb  88 88  .o 88  .o
88 YY 88 dP""""Yb 88  Y8 8888Y"  88  Yb 88 88ood8 88ood8

  NOT affiliated with Mailchimp's Mandrill email service
`

type Model struct {
	theme      styles.Theme
	idle       bool
	tickCount  int
	toolsCount int
	registryOK bool
	gitClean   bool
}

type tickMsg struct{}

func InitialModel(themeName string) Model {
	theme := styles.GetTheme(themeName)
	return Model{
		theme:      theme,
		idle:       false,
		toolsCount: 0,
		registryOK: true,
		gitClean:   true,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.idle = false
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tickMsg:
		m.tickCount++
		if m.tickCount > 30 {
			m.idle = true
		}
		return m, tea.Tick(1000000000, func(t time.Time) tea.Msg { return tickMsg{} })
	}
	return m, nil
}

func (m Model) View() string {
	if m.idle {
		return m.screensaverView()
	}
	return m.mainView()
}

func (m Model) screensaverView() string {
	status := fmt.Sprintf("  tools installed: %d | registry: up to date | git: main (clean)", m.toolsCount)
	return fmt.Sprintf("%s\n%s\n\n  press any key to return\n", mandrillASCII, status)
}

func (m Model) mainView() string {
	return fmt.Sprintf("%s\n  mandrill v0.1.0\n  press 'q' to quit\n", mandrillASCII)
}

func Run(theme string) error {
	p := tea.NewProgram(InitialModel(theme), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
