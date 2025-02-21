package route

import (
	"log"
	"net/http"
)

type Appointment struct{
    Id string
    Employee string `json:"employee"`
    Service []string `json:"service"`
    Date string `json:"date"`
}

var AppointmentHandler http.Handler = &Appointment{}
var dayHours = make([]int, 4)

func (a *Appointment) ServeHTTP(w http.ResponseWriter, r *http.Request){

    switch r.Method{
        default:
            a.Get(w, r)
    }
}

func (a *Appointment) Get(w http.ResponseWriter, r *http.Request){

    switch r.URL.Query().Get("type"){
        case "oldest":
            //Get The oldest appointments
            log.Println("oldest")
        default:
            //Get the latest appointments
            log.Println("latest")
    }
    if err := CreatePage(dayHours, w, "view/page.html", "view/appointment.tmpl", "view/component/AppointmentCard.tmpl");err != nil{
        log.Printf("error creating the page: %s", err)
        return
    }
    
}


