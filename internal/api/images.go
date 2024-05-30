package api

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"myapp/internal/image"
)

// @Summary Upload Image
// @Description Upload an image
// @Tags images
// @Accept  mpfd
// @Produce  json
// @Param file formData file true "Image file"
// @Param uid formData string true "User ID"
// @Success 201 {string} string "Image uploaded successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /images/upload [post]
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	image.UploadImage(w, r)
}

// @Summary Download Image
// @Description Download an image
// @Tags images
// @Produce  octet-stream
// @Param id path string true "Image ID"
// @Success 200 {file} file "Image file"
// @Failure 404 {string} string "Image not found"
// @Failure 500 {string} string "Internal server error"
// @Router /images/{id} [get]
func DownloadImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	image.DownloadImage(w, r, id)
}

// @Summary Delete Image
// @Description Delete an image
// @Tags images
// @Param id path string true "Image ID"
// @Success 204 {string} string "Image deleted"
// @Failure 404 {string} string "Image not found"
// @Failure 500 {string} string "Internal server error"
// @Router /images/{id} [delete]
func DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	image.DeleteImage(w, r, id)
}

// @Summary List Images
// @Description List all images
// @Tags images
// @Produce  json
// @Success 200 {array} image.ImageMetadata
// @Failure 500 {string} string "Internal server error"
// @Router /images [get]
func ListImagesHandler(w http.ResponseWriter, r *http.Request) {
	metadataList := image.ListMetadata()
	json.NewEncoder(w).Encode(metadataList)
}
