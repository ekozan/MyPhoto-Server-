package api

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"myapp/internal/auth"
	"myapp/internal/device"
)

// @Summary Create Device
// @Description Create a new device
// @Tags devices
// @Accept  json
// @Produce  json
// @Param device body device.Device true "Device"
// @Success 201 {object} device.Device
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /devices [post]
func CreateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())

	var dev device.Device
	err := json.NewDecoder(r.Body).Decode(&dev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dev.UserID = userID

	createdDevice, err := device.CreateDevice(dev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdDevice)
}

// @Summary Get Device
// @Description Get a device by ID
// @Tags devices
// @Produce  json
// @Param id path string true "Device ID"
// @Success 200 {object} device.Device
// @Failure 404 {string} string "Device not found"
// @Failure 500 {string} string "Internal server error"
// @Router /devices/{id} [get]
func GetDeviceHandler(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	vars := mux.Vars(r)
	id := vars["id"]

	dev, err := device.GetDevice(id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(dev)
}

// @Summary Update Device
// @Description Update a device by ID
// @Tags devices
// @Accept  json
// @Produce  json
// @Param id path string true "Device ID"
// @Param device body device.Device true "Device"
// @Success 200 {string} string "Device updated"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /devices/{id} [put]
func UpdateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	vars := mux.Vars(r)
	id := vars["id"]

	var dev device.Device
	err := json.NewDecoder(r.Body).Decode(&dev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dev.ID = id
	dev.UserID = userID

	err = device.UpdateDevice(dev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete Device
// @Description Delete a device by ID
// @Tags devices
// @Param id path string true "Device ID"
// @Success 204 {string} string "Device deleted"
// @Failure 404 {string} string "Device not found"
// @Failure 500 {string} string "Internal server error"
// @Router /devices/{id} [delete]
func DeleteDeviceHandler(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	vars := mux.Vars(r)
	id := vars["id"]
	err := device.DeleteDevice(id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// @Summary List Devices
// @Description List all devices
// @Tags devices
// @Produce  json
// @Success 200 {array} device.Device
// @Failure 500 {string} string "Internal server error"
// @Router /devices [get]
func ListDevicesHandler(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	devices, err := device.ListDevices(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(devices)
}
