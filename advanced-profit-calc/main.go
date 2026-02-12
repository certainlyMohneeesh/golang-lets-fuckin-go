package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// 1. Updated Struct with new fields
type FinancialReport struct {
	Timestamp        time.Time `json:"timestamp"`
	Revenue          float64   `json:"revenue"`
	Expenses         float64   `json:"expenses"`
	TaxRate          float64   `json:"tax_rate"`
	EBT              float64   `json:"ebt"`
	Profit           float64   `json:"profit"`
	TaxRetentionRate float64   `json:"tax_retention_rate"` // New
	PretaxMargin     float64   `json:"pretax_margin"`      // New
	NetProfitMargin  float64   `json:"net_profit_margin"`  // New
}

const historyFile = "history.json"

func main() {
	// Load existing history at startup
	history := loadHistory()

	fmt.Println("╔══════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                          Advanced Profit Calculator                          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════════════════╝")

	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. New Calculation")
		fmt.Println("2. View History")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			// Clear buffer if user types text
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			// Run calculation and get the populated struct
			report := runCalculation()

			// Append to history slice
			history = append(history, report)

			// Save to file
			saveHistory(history)
			fmt.Println("\nSuccess! Results saved to history.json")

		case 2:
			printHistory(history)
			printProfitChart(history)

		case 3:
			fmt.Println("Goodbye! See you again later.")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func runCalculation() FinancialReport {
	// Get Valid Inputs
	revenue := getValidFloat("Enter Revenue: ")
	expenses := getValidFloat("Enter Expenses: ")
	taxRateInput := getValidFloat("Enter Tax Rate (%): ")

	// Logic
	taxRate := taxRateInput / 100
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate)

	// Avoid division by zero for ratios
	var taxRetention, pretaxMargin, netMargin float64
	if ebt != 0 {
		taxRetention = profit / ebt
	}
	if revenue != 0 {
		pretaxMargin = ebt / revenue
		netMargin = profit / revenue
	}

	// Loading animation (User's original touch)
	fmt.Print("Calculating profit")
	for i := 0; i < 3; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Print("\r" + strings.Repeat(" ", 30) + "\r") // Clear line

	// Display Results
	fmt.Println("\n--- Financial Report ---")
	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", ebt)
	fmt.Printf("Net Profit:                %.2f\n", profit)
	fmt.Printf("Tax Retention Rate:        %.2f\n", taxRetention)
	fmt.Printf("Pretax Margin:             %.2f\n", pretaxMargin)
	fmt.Printf("Net Profit Margin:         %.2f\n", netMargin)

	// Return struct to be saved
	return FinancialReport{
		Timestamp:        time.Now(),
		Revenue:          revenue,
		Expenses:         expenses,
		TaxRate:          taxRateInput,
		EBT:              ebt,
		Profit:           profit,
		TaxRetentionRate: taxRetention,
		PretaxMargin:     pretaxMargin,
		NetProfitMargin:  netMargin,
	}
}

// Helper to handle input loops
func getValidFloat(prompt string) float64 {
	var input float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&input)
		if err == nil && input >= 0 {
			return input
		}
		var discard string
		fmt.Scanln(&discard)
		fmt.Println("Invalid input. Please enter a positive number.")
	}
}

// Save logic
func saveHistory(data []FinancialReport) {
	file, _ := json.MarshalIndent(data, "", "  ") // "  " makes the JSON readable
	os.WriteFile(historyFile, file, 0644)
}

// Load logic
func loadHistory() []FinancialReport {
	var data []FinancialReport
	file, err := os.ReadFile(historyFile)
	if err != nil {
		return []FinancialReport{}
	}
	json.Unmarshal(file, &data)
	return data
}

// Updated Print History to show new columns
func printHistory(data []FinancialReport) {
	if len(data) == 0 {
		fmt.Println("No history found.")
		return
	}

	// Header
	fmt.Printf("\n%-16s | %-10s | %-10s | %-10s | %-10s | %-10s\n", "Date", "Revenue", "Expenses", "Profit", "PreTax %", "Net %")
	fmt.Println(strings.Repeat("-", 70))

	for _, item := range data {
		dateStr := item.Timestamp.Format("2006-01-02 15:04")
		// Displaying the new margins in the history table
		fmt.Printf("%-16s | %-10.2f | %-10.2f | %-10.2f | %-10.2f | %-10.2f\n",
			dateStr, item.Revenue, item.Expenses, item.Profit, item.PretaxMargin, item.NetProfitMargin)
	}
}

func printProfitChart(data []FinancialReport) {
	if len(data) == 0 {
		return
	}

	// 1. Find the maximum profit to scale the chart correctly
	var maxProfit float64
	for _, item := range data {
		if item.Profit > maxProfit {
			maxProfit = item.Profit
		}
	}

	fmt.Println("\n--- Profit Trend (Visual) ---")
	// 2. Loop through history and print bars
	for _, item := range data {
		// Calculate bar length relative to the max profit (max width 50 chars)
		barLength := int((item.Profit / maxProfit) * 50)
		if barLength < 0 {
			barLength = 0
		} // Handle negative profit gracefully

		// "█" is a special block character that looks like a bar
		bar := strings.Repeat("█", barLength)

		fmt.Printf("%-12s | %s $%.0f\n",
			item.Timestamp.Format("Jan 02"), // Short date
			bar,
			item.Profit,
		)
	}
}
