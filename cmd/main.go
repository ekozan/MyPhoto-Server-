package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"myapp/internal/api"
	"myapp/internal/auth"
	"myapp/internal/config"
	"myapp/internal/db"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to the database
	db.Connect()

	r := mux.NewRouter()

	// Authentication
	r.HandleFunc("/api/auth/login", api.LoginHandler).Methods("POST")

	// Image management
	r.HandleFunc("/api/images/upload", api.UploadImageHandler).Methods("POST")
	r.HandleFunc("/api/images/{id}", api.DownloadImageHandler).Methods("GET")
	r.HandleFunc("/api/images/{id}", api.DeleteImageHandler).Methods("DELETE")
	r.HandleFunc("/api/images", api.ListImagesHandler).Methods("GET")

	// Collection management
	r.HandleFunc("/api/collections", api.CreateCollectionHandler).Methods("POST")
	r.HandleFunc("/api/collections/{id}", api.GetCollectionHandler).Methods("GET")
	r.HandleFunc("/api/collections/{id}", api.UpdateCollectionHandler).Methods("PUT")
	r.HandleFunc("/api/collections/{id}", api.DeleteCollectionHandler).Methods("DELETE")
	r.HandleFunc("/api/collections", api.ListCollectionsHandler).Methods("GET")

	// Device management
	r.HandleFunc("/api/devices", api.CreateDeviceHandler).Methods("POST")
	r.HandleFunc("/api/devices/{id}", api.GetDeviceHandler).Methods("GET")
	r.HandleFunc("/api/devices/{id}", api.UpdateDeviceHandler).Methods("PUT")
	r.HandleFunc("/api/devices/{id}", api.DeleteDeviceHandler).Methods("DELETE")
	r.HandleFunc("/api/devices", api.ListDevicesHandler).Methods("GET")

	// Use the auth middleware for all routes
	r.Use(auth.Middleware)

	port := viper.GetString("server.port")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
