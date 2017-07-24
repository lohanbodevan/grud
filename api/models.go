package api

import (
	"errors"
	log "github.com/Sirupsen/logrus"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ActorActress struct {
	Name string `json:"name"`
}

type TVSerie struct {
	Code        string         `json:"code,omitempty"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Casting     []ActorActress `json:"casting"`
	Stars       int            `json:"stars"`
}

func (serie TVSerie) Validate() error {
	if len(serie.Title) == 0 {
		log.Errorf("API - Validate - Title is required")
		return errors.New(VALIDATION_ERROR)
	}

	if len(serie.Description) == 0 {
		log.Errorf("API - Validate - Description is required")
		return errors.New(VALIDATION_ERROR)
	}

	if len(serie.Casting) == 0 {
		log.Errorf("API - Validate - At least one Casting item is required")
		return errors.New(VALIDATION_ERROR)
	}

	if serie.Stars == 0 {
		log.Errorf("API - Validate - Star should be gratter then zero")
		return errors.New(VALIDATION_ERROR)
	}

	return nil
}

func (user User) Validate() error {
	if len(user.Email) == 0 {
		log.Errorf("API - Validate - Email is required")
		return errors.New(VALIDATION_ERROR)
	}

	if len(user.Password) == 0 {
		log.Errorf("API - Validate - Password is required")
		return errors.New(VALIDATION_ERROR)
	}

	return nil
}
