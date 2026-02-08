package main

import (
	"log"
	"tuidoo/internal/app"
)

func main() {
	log.Println("🔄 Resetting database...")
	app.ResetDatabase()
}
