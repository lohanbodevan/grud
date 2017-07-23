package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	NOT_FOUND        = "not found"
	VALIDATION_ERROR = "Validation error"
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
	if err.Error() == VALIDATION_ERROR {
		w.WriteHeader(http.StatusBadRequest)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (a *Api) UpdateTVSerieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["code"] == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	var payload TVSerie
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = UpdateTVSerie(vars["code"], payload, a.Repository)
	if err.Error() == NOT_FOUND {
		w.WriteHeader(http.StatusNotFound)
	} else if err.Error() == VALIDATION_ERROR {
		w.WriteHeader(http.StatusBadRequest)
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (a *Api) DeleteTVSerieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["code"] == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := DeleteTVSerie(vars["code"], a.Repository)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
