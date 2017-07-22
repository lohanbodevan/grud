package api

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func GetTVSeries(r *Repository) ([]TVSerie, error) {
	collection := r.Session.DB(os.Getenv("DB_NAME")).C("tv_series")

	series := []TVSerie{}
	err := collection.Find(nil).Iter().All(&series)
	if err != nil {
		log.Errorf("API - GetTVSeries - Fail to select: %s", err)
		return nil, err
	}

	log.Infof("API - GetTVSeries - Get from db: %s", series)
	return series, nil
}

func CreateTVSerie(serie TVSerie, r *Repository) error {
	collection := r.Session.DB(os.Getenv("DB_NAME")).C("tv_series")

	err := collection.Insert(&serie)
	if err != nil {
		log.Errorf("API - CreateTVSerie - Fail to insert: %s", err)
		return err
	}

	log.Infof("API - CreateTVSerie - TV Serie created: %s", serie)
	return nil
}
