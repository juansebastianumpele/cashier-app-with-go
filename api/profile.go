package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	// View with response image `img-avatar.png` from path `assets/images`
	var image, _ = os.ReadFile("/assets/images/img-avatar.png")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/png")
	w.Write(image)

	// TODO: answer here
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	// Update image `img-avatar.png` from path `assets/images`
	uploadedFile, _, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	readImage, _ := ioutil.ReadAll(uploadedFile)
	ioutil.WriteFile("/assets/images/img-avatar.png", readImage, 0644)

	// TODO: answer here

	api.dashboardView(w, r)
	defer uploadedFile.Close()
}
