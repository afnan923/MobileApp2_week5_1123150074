package main

import (
	"log"

	"github.com/afnan923/MobileApp2_week5_1123150074/config"
	"github.com/afnan923/MobileApp2_week5_1123150074/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.InitDatabase()

	products := []models.Product{
		{
			Name:        "Paket Pancing Pemula",
			Price:       75000,
			Category:    "Paket",
			Stock:       20,
			Description: "Paket lengkap untuk pemula: joran, reel, senar, dan umpan dasar.",
			ImageURL:    "https://images.unsplash.com/photo-1578269174936-2709b6aeb913?q=80&w=600&h=400&fit=crop",
		},
		{
			Name:        "Paket Pancing Profesional",
			Price:       250000,
			Category:    "Paket",
			Stock:       10,
			Description: "Paket pancing kualitas tinggi untuk mancing di laut atau sungai besar.",
			ImageURL:    "https://images.unsplash.com/photo-1504274066651-8d31a536b11a?q=80&w=600&h=400&fit=crop",
		},
		{
			Name:        "Paket Pancing Harian",
			Price:       50000,
			Category:    "Paket",
			Stock:       30,
			Description: "Cocok untuk mancing santai di kolam atau danau.",
			ImageURL:    "https://images.unsplash.com/photo-1525104698733-04c4cfa4f5b4?q=80&w=600&h=400&fit=crop",
		},
		{
			Name:        "Paket Pancing Malam",
			Price:       90000,
			Category:    "Paket",
			Stock:       15,
			Description: "Dilengkapi lampu, pelampung, dan perlengkapan khusus mancing malam.",
			ImageURL:    "https://images.unsplash.com/photo-1507525428034-b723cf961d3e?q=80&w=600&h=400&fit=crop",
		},
	}

	for _, p := range products {
		config.DB.Where("name = ?", p.Name).Assign(p).FirstOrCreate(&p)
	}

	log.Printf("Seed berhasil: %d paket pancing ditambahkan!", len(products))
}