package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"os"
)

type Repository struct {
	Session *mgo.Session
}

type Api struct {
	Repository *Repository
}

func (app *Api) ConfigureRoutes(router *mux.Router) {
	router.HandleFunc("/tvseries", app.GetTVSeriesHandler).Methods("get")
	router.HandleFunc("/tvseries", app.CreateTVSerieHandler).Methods("post")
}

func InitServer() {
	repository := DatabseInit()
	defer repository.Session.Close()

	app := Api{
		Repository: &repository,
	}

	mux := mux.NewRouter()
	app.ConfigureRoutes(mux)

	server := negroni.New(negroni.NewRecovery())
	server.UseHandler(mux)

	serverAddr := ":" + os.Getenv("PORT")
	server.Run(serverAddr)
}

func DatabseInit() Repository {
	session, err := mgo.Dial(os.Getenv("DB_HOST"))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	repo := Repository{
		Session: session,
	}
	return repo
}
