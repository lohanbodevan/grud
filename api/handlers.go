package api

import (
	"encoding/json"
	"net/http"
)

func (a *Api) GetTVSeriesHandler(w http.ResponseWriter, r *http.Request) {
	series, err := GetTVSeries(a.Repository)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var response []byte

	if len(series) > 0 {
		response, err = json.Marshal(series)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(response)
}

func (a *Api) CreateTVSerieHandler(w http.ResponseWriter, r *http.Request) {
	var payload TVSerie
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = CreateTVSerie(payload, a.Repository)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
