package main

import (
	"github.com/codegangsta/negroni"
	"github.com/intervention-engine/fhir/server"
	"github.com/intervention-engine/ie/controllers"
	"github.com/intervention-engine/ie/middleware"
	"os"
)

func main() {
	// Check for a linked MongoDB container if we are running in Docker
	mongoHost := os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	if mongoHost == "" {
		mongoHost = "localhost"
	}

	s := server.NewServer(mongoHost)

	s.AddMiddleware("QueryCreate", negroni.HandlerFunc(middleware.QueryExecutionHandler))

	s.AddMiddleware("PatientCreate", negroni.HandlerFunc(middleware.FactHandler))
	s.AddMiddleware("PatientUpdate", negroni.HandlerFunc(middleware.FactHandler))
	s.AddMiddleware("PatientDelete", negroni.HandlerFunc(middleware.FactHandler))

	s.AddMiddleware("ConditionCreate", negroni.HandlerFunc(middleware.FactHandler))
	s.AddMiddleware("ConditionUpdate", negroni.HandlerFunc(middleware.FactHandler))
	s.AddMiddleware("ConditionDelete", negroni.HandlerFunc(middleware.FactHandler))

	s.AddMiddleware("EncounterCreate", negroni.HandlerFunc(middleware.FactHandler))
	s.AddMiddleware("EncounterUpdate", negroni.HandlerFunc(middleware.FactHandler))
	s.AddMiddleware("EncounterDelete", negroni.HandlerFunc(middleware.FactHandler))

	s.AddMiddleware("ObservationCreate", negroni.HandlerFunc(middleware.FactHandler))
	s.AddMiddleware("ObservationUpdate", negroni.HandlerFunc(middleware.FactHandler))
	s.AddMiddleware("ObservationDelete", negroni.HandlerFunc(middleware.FactHandler))

	s.Router.HandleFunc("/QueryConditionTotal/{id}", controllers.ConditionTotalHandler)
	s.Router.HandleFunc("/QueryEncounterTotal/{id}", controllers.EncounterTotalHandler)
	s.Router.HandleFunc("/QueryList/{id}", controllers.PatientListHandler)

	filterBase := s.Router.Path("/Filter").Subrouter()
	filterBase.Methods("GET").Handler(negroni.New(negroni.HandlerFunc(controllers.FilterIndexHandler)))
	filterBase.Methods("POST").Handler(negroni.New(negroni.HandlerFunc(controllers.FilterCreateHandler)))

	filter := s.Router.Path("/Filter/{id}").Subrouter()
	filter.Methods("GET").Handler(negroni.New(negroni.HandlerFunc(controllers.FilterShowHandler)))
	filter.Methods("PUT").Handler(negroni.New(negroni.HandlerFunc(controllers.FilterUpdateHandler)))
	filter.Methods("DELETE").Handler(negroni.New(negroni.HandlerFunc(controllers.FilterDeleteHandler)))

	s.Run()
}
