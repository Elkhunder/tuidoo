package main

import (
	"log"
	"tuidoo/internal/app"
)

func main() {
	log.Println("🌱 Seeding database...")
	app.SeedDatabase()
}
