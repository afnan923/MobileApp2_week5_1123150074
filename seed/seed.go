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
			Name:        "Nasi Goreng Spesial",
			Price:       25000,
			Category:    "Makanan",
			Stock:       50,
			Description: "Nasi goreng gurih dengan topping telur, ayam, dan kerupuk.",
			ImageURL:    "https://images.unsplash.com/photo-1603133872878-684f208fb84b?q=80&w=600&h=400&fit=crop",
		},
		{
			Name:        "Es Teh Manis",
			Price:       5000,
			Category:    "Minuman",
			Stock:       100,
			Description: "Minuman teh segar dengan es batu dan gula asli.",
			ImageURL:    "https://images.unsplash.com/photo-1556679343-c7306c1976bc?q=80&w=600&h=400&fit=crop",
		},
		{
			Name:        "Nasi Megono",
			Price:       15000,
			Category:    "Makanan",
			Stock:       40,
			Description: "Nasi khas dengan cacahan nangka muda dan parutan kelapa berbumbu.",
			ImageURL:    "https://images.unsplash.com/photo-1512058560366-cd24295982cd?q=80&w=600&h=400&fit=crop",
		},
		{
			Name:        "Soto Tauco",
			Price:       20000,
			Category:    "Makanan",
			Stock:       30,
			Description: "Soto daging/ayam dengan kuah tauco kental nan sedap.",
			ImageURL:    "https://images.unsplash.com/photo-1547592166-23ac45744acd?q=80&w=600&h=400&fit=crop",
		},
		{
			Name:        "Telur Asin",
			Price:       6000,
			Category:    "Makanan",
			Stock:       200,
			Description: "Telur asin khas yang masir, gurih, dan berminyak.",
			ImageURL:    "https://images.unsplash.com/photo-1587486912202-3e44b1d3f675?q=80&w=600&h=400&fit=crop",
		},
	}


	for _, p := range products {
		config.DB.Create(&p)
	}

	log.Printf("Seed berhasil: %d produk kuliner dengan gambar baru telah ditambahkan!", len(products))
}