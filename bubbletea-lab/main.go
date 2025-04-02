package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	name  string
	count int
}

func (model) Init() tea.Cmd {
	return nil
}

type customMsg string

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case customMsg:
		m.name = string(msg)

	case tea.KeyMsg:
		switch msg.String() {

		case "c":
			return m, func() tea.Msg {
				<-time.After(2 * time.Second)
				return customMsg(fmt.Sprintf("%s (command)", m.name))
			}

		case "r":
			return m, func() tea.Msg {
				n := strings.Split(m.name, " ")
				return customMsg(n[0])
			}
		}

	}

	m.count++

	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("\n%s with count %d and pressed key", m.name, m.count)
}

type item struct {
	key         string
	description string
	model       tea.Model
}

type menu struct {
	items  []item
	active int
}

func (menu) Init() tea.Cmd {
	return nil
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msgtype := msg.(type) {

	case tea.KeyMsg:
		switch msgtype.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "b":
			m.active = -1
			return m, nil

		default:
			if m.active == -1 {
				for i, item := range m.items {
					if item.key == msgtype.String() {
						m.active = i
					}
				}

				return m, nil
			}

			i := m.items[m.active]
			ni, cmd := i.model.Update(msg)
			m.items[m.active].model = ni

			return m, cmd
		}

	default:
		if m.active != -1 {
			i := m.items[m.active]
			ni, cmd := i.model.Update(msg)
			m.items[m.active].model = ni
			return m, cmd
		}
	}

	return m, nil
}

func (m menu) View() string {
	var b strings.Builder

	b.WriteString("default header\n")

	if m.active != -1 {
		b.WriteString(m.items[m.active].model.View())
	} else {
		for _, i := range m.items {
			b.WriteString(fmt.Sprintf("\n(%s) %s", i.key, i.description))
		}
	}

	b.WriteString("\n\ndefault footer")

	return b.String()
}

func main() {
	first := model{
		name: "first",
	}
	second := model{
		name: "second",
	}
	third := model{
		name: "third",
	}

	m := menu{
		active: -1,
		items: []item{
			{
				key:         "f",
				description: "this is the first item",
				model:       first,
			},
			{
				key:         "s",
				description: "this is the second item",
				model:       second,
			},
			{
				key:         "t",
				description: "i'm the third and I was just added in the items list and everything is still working",
				model:       third,
			},
		},
	}

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
