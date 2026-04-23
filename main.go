package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/afnan923/MobileApp2_week5_1123150074/config"
	"github.com/afnan923/MobileApp2_week5_1123150074/routes"
	"github.com/afnan923/MobileApp2_week5_1123150074/pkg/logger"
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment variable sistem")
	}
	logger.Init()
	config.InitFirebase()
	
	config.InitDatabase()

	router := routes.SetupRouter()
	
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	logger.L.Info("server starting",
		"url", "http://0.0.0.0:"+port,
		"health", "http://0.0.0.0:"+port+"/v1/health",
	)
	
	if err := router.Run(":" + port); err != nil {
		logger.L.Error("server gagal berjalan", "error", err)
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}