package main

import (
	"encoding/json"
	"github.com/Parth-PDF/Visimage-App/dao"
	"net/http"
)

// Image is ...
type Image struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	ImageTag string `json:"imageTag"`
}

// UploadRequest is ...
type UploadRequest struct {
	DataURI string `json:"datauri"`
}

// DeleteRequest is ...
type DeleteRequest struct {
	DeleteID string `json:"id"`
}

// ImagesHandler is ...
func ImagesHandler(imageDao *dao.ImageDao) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		images, err := imageDao.GetImages()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error talking to db"))
			return
		}

		// convert db response to http response for images
		imagesResp := []Image{}
		for _, image := range images {
			imagesResp = append(imagesResp, Image{ID: image.ID, UserID: image.UserID, ImageTag: image.ImageTag})
		}

		payload, _ := json.Marshal(imagesResp)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(payload))
	}
}

// UploadHandler is ...
func UploadHandler(imageDao *dao.ImageDao) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var uploadReq UploadRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&uploadReq)
		if err != nil {
			// w.WriteHeader() Fill this out as http status invalid argument
			w.Write([]byte("Unable to parse request"))
			return
		}

		err = imageDao.PostImage(uploadReq.DataURI)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error uploading image to db"))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// DeleteHandler is ...
func DeleteHandler(imageDao *dao.ImageDao) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var deleteReq DeleteRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&deleteReq)

		if err != nil {
			// w.WriteHeader() Fill this out as http status invalid argument
			w.Write([]byte("Unable to parse request"))
			return
		}

		err = imageDao.DeleteImage(deleteReq.DeleteID)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error uploading image to db"))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
