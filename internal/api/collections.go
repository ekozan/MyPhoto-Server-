package api

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"myapp/internal/auth"
	"myapp/internal/collection"

// @Summary Create Collection
// @Description Create a new collection
// @Tags collections
// @Accept  json
// @Produce  json
// @Param collection body collection.Collection true "Collection"
// @Success 201 {object} collection.Collection
// @Failure 400 {string} string "Bad request"
// @Failure 403 {string} string "Collection limit reached"
// @Failure 500 {string} string "Internal server error"
// @Router /collections [post]
func CreateCollectionHandler(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())

	collections, err := collection.ListCollections(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(collections) >= 10 {
		http.Error(w, "Collection limit reached", http.StatusForbidden)
		return
	}

	var coll collection.Collection
	err = json.NewDecoder(r.Body).Decode(&coll)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	coll.UserID = userID

	createdCollection, err := collection.CreateCollection(&coll)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdCollection)
}

// @Summary Get Collection
// @Description Get a collection by ID
// @Tags collections
// @Produce  json
// @Param id path string true "Collection ID"
// @Success 200 {object} collection.Collection
// @Failure 404 {string} string "Collection not found"
// @Failure 500 {string} string "Internal server error"
// @Router /collections/{id} [get]
func GetCollectionHandler(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	vars := mux.Vars(r)
	id := vars["id"]
	coll, err := collection.GetCollection(id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(coll)
}

// @Summary Update Collection
// @Description Update a collection by ID
// @Tags collections
// @Accept  json
// @Produce  json
// @Param id path string true "Collection ID"
// @Param collection body collection.Collection true "Collection"
// @Success 200 {string} string "Collection updated"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /collections/{id} [put]
func UpdateCollectionHandler(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	vars := mux.Vars(r)
	id := vars["id"]

	var coll collection.Collection
	err := json.NewDecoder(r.Body).Decode(&coll)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	coll.ID = id
	coll.UserID = userID

	err = collection.UpdateCollection(&coll)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete Collection
// @Description Delete a collection by ID
// @Tags collections
// @Param id path string true "Collection ID"
// @Success 204 {string} string "Collection deleted"
// @Failure 404 {string} string "Collection not found"
// @Failure 500 {string} string "Internal server error"
// @Router /collections/{id} [delete]
func DeleteCollectionHandler(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	vars := mux.Vars(r)
	id := vars["id"]
	err := collection.DeleteCollection(id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// @Summary List Collections
// @Description List all collections
// @Tags collections
// @Produce  json
// @Success 200 {array} collection.Collection
// @Failure 500 {string} string "Internal server error"
// @Router /collections [get]
func ListCollectionsHandler(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	collections, err := collection.ListCollections(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(collections)
}
