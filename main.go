package main

import (
	"log"
	"net/http"
	"planify/model"
	"planify/route"
	"time"
)

const (
    Addr = ":8000"
)

type Test struct{
    Data interface{}
}

func (t Test) ServeHTTP(w http.ResponseWriter, r *http.Request){
}

func main(){
    if err := model.InitDB(); err != nil{
        log.Fatalf("error db init\n %v", err)
    }   
    mux := http.NewServeMux()
    mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
    mux.HandleFunc("/", route.LandpageHandler.ServeHTTP)
    mux.HandleFunc("/connexion",  route.SigninHandler.ServeHTTP)
    mux.HandleFunc("/compte", route.AccountHandler.ServeHTTP)
    mux.HandleFunc("/etablissement", route.EtablishmentHandler.ServeHTTP)
    mux.HandleFunc("/recherche", route.SearchHandler.ServeHTTP)
    mux.HandleFunc("/rendez-vous", route.AppointmentHandler.ServeHTTP)
    mux.HandleFunc("/etablissement/{id}", route.StoreHandler.ServeHTTP)
    mux.HandleFunc("/etablissement/{id}/rendez-vous", route.NewAppointmentHandler.ServeHTTP)
    mux.HandleFunc("/planning", route.PlanningHandler.ServeHTTP)
    mux.HandleFunc("/pro", route.MyEtablishmentHandler.ServeHTTP)

    server := http.Server{
        Addr: Addr,
        ReadTimeout: time.Second * 5,
        Handler: mux,
        
    }

    if err := server.ListenAndServe(); err != nil{
        log.Fatalf("error spinning server: %s", Addr)
    }
}








