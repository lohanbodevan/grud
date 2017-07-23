package api

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func GetTVSeries(r Repository) ([]TVSerie, error) {
	collection := r.DB(os.Getenv("DB_NAME")).C("tv_series")

	series := []TVSerie{}
	err := collection.Find(nil).Iter().All(&series)
	if err != nil {
		log.Errorf("API - GetTVSeries - Fail to select: %s", err)
		return nil, err
	}

	log.Infof("API - GetTVSeries - Get from db: %s", series)
	return series, nil
}

func CreateTVSerie(serie TVSerie, r Repository) error {
	err := serie.Validate()
	if err != nil {
		return err
	}

	serie.Code = generateUUID()

	collection := r.DB(os.Getenv("DB_NAME")).C("tv_series")
	err = collection.Insert(&serie)
	if err != nil {
		log.Errorf("API - CreateTVSerie - Fail to insert: %s", err)
		return err
	}

	log.Infof("API - CreateTVSerie - TV Serie created: %s", serie)
	return nil
}

func UpdateTVSerie(code interface{}, serie TVSerie, r Repository) error {
	err := serie.Validate()
	if err != nil {
		return err
	}

	codeValue, ok := code.(string)
	if ok == false {
		return errors.New("Code should be a string")
	}
	serie.Code = codeValue

	collection := r.DB(os.Getenv("DB_NAME")).C("tv_series")
	err = collection.Update(bson.M{"code": code}, &serie)
	if err != nil {
		log.Errorf("API - UpdateTVSerie - Fail to update: %s", err)
		return err
	}

	log.Infof("API - UpdateTVSerie - TV Serie updated: %s", serie)
	return nil
}

func DeleteTVSerie(code interface{}, r Repository) error {
	collection := r.DB(os.Getenv("DB_NAME")).C("tv_series")

	err := collection.Remove(bson.M{"code": code})
	if err != nil {
		log.Errorf("API - DeleteTVSerie - Fail to delete ID: %s Error: %s", code, err)
		return err
	}

	log.Infof("API - DeleteTVSerie - TV Serie deleted: %s", code)
	return nil
}

func generateUUID() string {
	return uuid.NewV4().String()
}
