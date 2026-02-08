package main

import (
	"fmt"
	"log"
	"os"

	"tuidoo/internal/app"
)

func main() {
	// Check for subcommands
	if len(os.Args) > 1 {
		command := os.Args[1]

		switch command {
		case "seed":
			app.SeedDatabase()
			return
		case "reset":
			app.ResetDatabase()
			return
		case "clean":
			app.CleanDatabase()
			return
		case "help", "-h", "--help":
			printHelp()
			return
		case "version", "-v", "--version":
			printVersion()
			return
		default:
			fmt.Printf("Unknown command: %s\n\n", command)
			printHelp()
			os.Exit(1)
		}
	}

	// No command provided - run TUI app
	if err := app.RunTUI(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}

func printHelp() {
	fmt.Println(`TUIDOO - Terminal UI Todo Application

Usage:
  tuidoo                Run the TUI application (default)
  tuidoo seed           Seed the database with sample data
  tuidoo reset          Reset and reseed the database
  tuidoo clean          Clean all data from the database
  tuidoo help           Show this help message
  tuidoo version        Show version information

Options:
  -h, --help            Show help
  -v, --version         Show version

Examples:
  tuidoo                # Start the TUI
  tuidoo seed           # Add sample data
  tuidoo reset          # Fresh start with sample data`)
}

func printVersion() {
	fmt.Println("TUIDOO version 0.1.0-alpha")
}
