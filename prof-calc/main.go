package main

import (
	"encoding/json"
	"fmt"

	// "math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// --- Constants & Styles ---

const (
	historyFile = "financial_data.json"
	colorAccent = "#04B575" // Green for profit
	colorText   = "#FAFAFA" // White
	colorSub    = "#7D7D7D" // Grey
)

var (
	appStyle   = lipgloss.NewStyle().Padding(1, 2)
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorAccent)).
			Bold(true).
			Padding(0, 1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(colorAccent))

	infoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(colorSub))
)

// --- Data Models ---

type FinancialReport struct {
	Timestamp       time.Time `json:"timestamp"`
	Revenue         float64   `json:"revenue"`
	Expenses        float64   `json:"expenses"`
	Profit          float64   `json:"profit"`
	NetProfitMargin float64   `json:"net_profit_margin"`
}

type Profile struct {
	Name     string            `json:"name"`
	Currency string            `json:"currency"` // "$" or "₹"
	History  []FinancialReport `json:"history"`
}

type AppData struct {
	Profiles       []Profile `json:"profiles"`
	CurrentProfile int       `json:"current_profile"`
}

// --- Bubble Tea Model ---

type state int

const (
	stateMenu state = iota
	stateNewCalc
	stateHistory
	stateProfileSelect
	stateNewProfile
)

type model struct {
	state           state
	data            AppData
	inputs          []textinput.Model // 0: Revenue, 1: Expenses, 2: Tax
	focusIndex      int
	cursor          int // Main menu cursor
	table           table.Model
	newProfileInput textinput.Model
	width, height   int
}

// --- Initialization ---

func initialModel() model {
	// Initialize Inputs for Calculation
	inputs := make([]textinput.Model, 3)
	labels := []string{"Revenue", "Expenses", "Tax Rate (%)"}
	for i := range inputs {
		t := textinput.New()
		t.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(colorAccent))
		t.Placeholder = labels[i]
		t.CharLimit = 20
		inputs[i] = t
	}

	// Initialize New Profile Input
	np := textinput.New()
	np.Placeholder = "Enter Profile Name (e.g. SaaS, Freelance)"
	np.CharLimit = 30

	// Load Data
	data := loadData()
	if len(data.Profiles) == 0 {
		// Default profile if none exists
		data.Profiles = append(data.Profiles, Profile{Name: "Default", Currency: "$", History: []FinancialReport{}})
	}

	return model{
		state:           stateMenu,
		data:            data,
		inputs:          inputs,
		newProfileInput: np,
		table:           createTable(data.Profiles[data.CurrentProfile]),
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

// --- Update Loop ---

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.table.SetWidth(msg.Width - 10)

	case tea.KeyMsg:
		switch m.state {

		// 1. MAIN MENU
		case stateMenu:
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < 4 { // 5 menu items
					m.cursor++
				}
			case "enter":
				switch m.cursor {
				case 0: // New Calc
					m.state = stateNewCalc
					m.focusIndex = 0
					m.inputs[0].Focus()
					return m, textinput.Blink
				case 1: // History
					m.state = stateHistory
					m.table = createTable(m.getCurrentProfile())
				case 2: // Switch Profile
					m.state = stateProfileSelect
					m.cursor = 0 // Reset cursor for sub-menu
				case 3: // Toggle Currency
					curr := m.getCurrentProfile().Currency
					if curr == "$" {
						m.data.Profiles[m.data.CurrentProfile].Currency = "₹"
					} else {
						m.data.Profiles[m.data.CurrentProfile].Currency = "$"
					}
					saveData(m.data)
				case 4: // Exit
					return m, tea.Quit
				}
			}

		// 2. NEW CALCULATION FORM
		case stateNewCalc:
			switch msg.String() {
			case "esc":
				m.state = stateMenu
				return m, nil
			case "tab", "shift+tab", "enter", "up", "down":
				s := msg.String()

				if s == "enter" && m.focusIndex == len(m.inputs)-1 {
					// Calculate & Save
					rev, _ := strconv.ParseFloat(m.inputs[0].Value(), 64)
					exp, _ := strconv.ParseFloat(m.inputs[1].Value(), 64)
					tax, _ := strconv.ParseFloat(m.inputs[2].Value(), 64)

					ebt := rev - exp
					profit := ebt * (1 - (tax / 100))
					margin := 0.0
					if rev > 0 {
						margin = profit / rev
					}

					report := FinancialReport{
						Timestamp:       time.Now(),
						Revenue:         rev,
						Expenses:        exp,
						Profit:          profit,
						NetProfitMargin: margin,
					}

					// Append to current profile
					m.data.Profiles[m.data.CurrentProfile].History = append(m.data.Profiles[m.data.CurrentProfile].History, report)
					saveData(m.data)

					// Reset inputs
					for i := range m.inputs {
						m.inputs[i].Reset()
					}
					m.state = stateMenu
					return m, nil
				}

				// Cycle inputs
				if s == "up" || s == "shift+tab" {
					m.focusIndex--
				} else {
					m.focusIndex++
				}

				if m.focusIndex > len(m.inputs)-1 {
					m.focusIndex = 0
				} else if m.focusIndex < 0 {
					m.focusIndex = len(m.inputs) - 1
				}

				cmds := make([]tea.Cmd, len(m.inputs))
				for i := 0; i <= len(m.inputs)-1; i++ {
					if i == m.focusIndex {
						cmds[i] = m.inputs[i].Focus()
						continue
					}
					m.inputs[i].Blur()
				}
				return m, tea.Batch(cmds...)
			}
			// Handle typing
			cmd = m.updateInputs(msg)
			return m, cmd

		// 3. HISTORY VIEW
		case stateHistory:
			if msg.String() == "esc" || msg.String() == "q" {
				m.state = stateMenu
				return m, nil
			}
			m.table, cmd = m.table.Update(msg)
			return m, cmd

		// 4. PROFILE SELECT
		case stateProfileSelect:
			switch msg.String() {
			case "esc":
				m.state = stateMenu
				return m, nil
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.data.Profiles) { // +1 for "Create New"
					m.cursor++
				}
			case "enter":
				if m.cursor == len(m.data.Profiles) {
					// Create New Profile
					m.state = stateNewProfile
					m.newProfileInput.Focus()
					return m, textinput.Blink
				}
				// Select Profile
				m.data.CurrentProfile = m.cursor
				saveData(m.data)
				m.state = stateMenu
			}

		// 5. NEW PROFILE
		case stateNewProfile:
			switch msg.String() {
			case "esc":
				m.state = stateProfileSelect
				return m, nil
			case "enter":
				name := m.newProfileInput.Value()
				if name != "" {
					newProfile := Profile{
						Name:     name,
						Currency: "$", // Default
						History:  []FinancialReport{},
					}
					m.data.Profiles = append(m.data.Profiles, newProfile)
					m.data.CurrentProfile = len(m.data.Profiles) - 1
					saveData(m.data)
					m.newProfileInput.Reset()
					m.state = stateMenu
				}
				return m, nil
			}
			m.newProfileInput, cmd = m.newProfileInput.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return tea.Batch(cmds...)
}

// --- View Rendering ---

func (m model) View() string {
	s := strings.Builder{}

	// Global Header
	s.WriteString(titleStyle.Render("PROFIT CALCULATOR PRO"))
	s.WriteString("\n\n")

	// Current Context
	profile := m.getCurrentProfile()
	s.WriteString(fmt.Sprintf("Profile: %s | Currency: %s\n", profile.Name, profile.Currency))
	s.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color(colorSub)).Render(strings.Repeat("-", 40)) + "\n\n")

	switch m.state {
	case stateMenu:
		choices := []string{
			"1. New Calculation",
			"2. View Analytics & History",
			"3. Switch Profile",
			"4. Toggle Currency ($/₹)",
			"5. Exit",
		}
		for i, choice := range choices {
			cursor := "  "
			if m.cursor == i {
				cursor = "> "
				choice = lipgloss.NewStyle().Foreground(lipgloss.Color(colorAccent)).Bold(true).Render(choice)
			}
			s.WriteString(fmt.Sprintf("%s%s\n", cursor, choice))
		}

		// Forecasting Widget
		s.WriteString("\n" + getForecasting(profile) + "\n")

	case stateNewCalc:
		s.WriteString("Enter Financial Data:\n\n")
		for i, input := range m.inputs {
			s.WriteString(input.View())
			if i < len(m.inputs)-1 {
				s.WriteString("\n\n")
			}
		}
		s.WriteString("\n\n(Press Enter to Submit, Esc to Cancel)\n")

	case stateHistory:
		s.WriteString(m.table.View())
		s.WriteString("\n\n(Press Q to Back)\n")

	case stateProfileSelect:
		s.WriteString("Select Profile:\n\n")
		for i, p := range m.data.Profiles {
			cursor := "  "
			name := p.Name
			if m.cursor == i {
				cursor = "> "
				name = lipgloss.NewStyle().Foreground(lipgloss.Color(colorAccent)).Bold(true).Render(name)
			}
			s.WriteString(fmt.Sprintf("%s%s\n", cursor, name))
		}
		// Option to create new
		cursor := "  "
		text := "+ Create New Profile"
		if m.cursor == len(m.data.Profiles) {
			cursor = "> "
			text = lipgloss.NewStyle().Foreground(lipgloss.Color(colorAccent)).Bold(true).Render(text)
		}
		s.WriteString(fmt.Sprintf("%s%s\n", cursor, text))

	case stateNewProfile:
		s.WriteString("Create New Profile:\n\n")
		s.WriteString(m.newProfileInput.View())
		s.WriteString("\n\n(Enter to Save, Esc to Cancel)\n")
	}

	return appStyle.Render(s.String())
}

// --- Helpers ---

func (m model) getCurrentProfile() Profile {
	if len(m.data.Profiles) == 0 {
		return Profile{Name: "None", Currency: "$"}
	}
	if m.data.CurrentProfile >= len(m.data.Profiles) {
		return m.data.Profiles[0]
	}
	return m.data.Profiles[m.data.CurrentProfile]
}

func getForecasting(p Profile) string {
	if len(p.History) < 3 {
		return infoStyle.Render("ℹ️  Forecast: Need at least 3 entries to project revenue.")
	}

	// Simple Moving Average of last 3 entries
	start := len(p.History) - 3
	sumProfit := 0.0
	for i := start; i < len(p.History); i++ {
		sumProfit += p.History[i].Profit
	}
	avg := sumProfit / 3.0

	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.Color(colorAccent)).
		Padding(0, 1).
		Render(fmt.Sprintf("🔮 Forecast: Projected Next Month Profit: %s%.2f", p.Currency, avg))
}

func createTable(p Profile) table.Model {
	columns := []table.Column{
		{Title: "Date", Width: 12},
		{Title: "Revenue", Width: 10},
		{Title: "Profit", Width: 10},
		{Title: "Margin", Width: 8},
	}

	rows := []table.Row{}
	for _, item := range p.History {
		rows = append(rows, table.Row{
			item.Timestamp.Format("2006-01-02"),
			fmt.Sprintf("%.2f", item.Revenue),
			fmt.Sprintf("%.2f", item.Profit),
			fmt.Sprintf("%.0f%%", item.NetProfitMargin*100),
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}

// --- Persistence ---

func saveData(data AppData) {
	file, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(historyFile, file, 0644)
}

func loadData() AppData {
	var data AppData
	file, err := os.ReadFile(historyFile)
	if err != nil {
		return AppData{CurrentProfile: 0}
	}
	json.Unmarshal(file, &data)
	return data
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
