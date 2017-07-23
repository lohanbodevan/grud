package api

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"os"
)

type Repository interface {
	DB(name string) *mgo.Database
	Close()
}

type Api struct {
	Repository Repository
}

func (app *Api) ConfigureRoutes(router *mux.Router) {
	router.HandleFunc("/tvseries", app.GetTVSeriesHandler).Methods("get")
	router.HandleFunc("/tvseries", app.CreateTVSerieHandler).Methods("post")
	router.HandleFunc("/tvseries/{code}", app.UpdateTVSerieHandler).Methods("put")
	router.HandleFunc("/tvseries/{code}", app.DeleteTVSerieHandler).Methods("delete")
}

func InitServer() {
	repository := DatabseInit()
	defer repository.Close()

	app := Api{
		Repository: repository,
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

	return session
}
