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
		return
	}

	var response []byte
	if len(series) > 0 {
		response, err = json.Marshal(series)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(response)
}

func (a *Api) CreateTVSerieHandler(w http.ResponseWriter, r *http.Request) {
	var payload TVSerie
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = CreateTVSerie(payload, a.Repository)
	if err != nil && err.Error() == VALIDATION_ERROR {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *Api) UpdateTVSerieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["code"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var payload TVSerie
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = UpdateTVSerie(vars["code"], payload, a.Repository)
	if err != nil && err.Error() == NOT_FOUND {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil && err.Error() == VALIDATION_ERROR {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func (a *Api) DeleteTVSerieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["code"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := DeleteTVSerie(vars["code"], a.Repository)
	if err != nil && err.Error() == NOT_FOUND {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
