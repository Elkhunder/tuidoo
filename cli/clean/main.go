package main

import (
	"log"
	"tuidoo/internal/app"
)

func main() {
	log.Println("🧹 Cleaning database...")
	app.CleanDatabase()
}
