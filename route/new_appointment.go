package route

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type NewAppointment struct{
    EmployeeId string `json:"employee"`
    Service []string `json:"service"`
    Date string `json:"date"`

}

var NewAppointmentHandler http.Handler = &NewAppointment{}

func (a *NewAppointment) ServeHTTP(w http.ResponseWriter, r *http.Request){
    switch r.Method{
        case http.MethodPut:
            a.Put(w, r)
        case http.MethodPost:
            a.Post(w, r)
        default:
            a.Get(w, r)
    }

}

func (n *NewAppointment) Get(w http.ResponseWriter, r *http.Request){

    if err := CreatePage(nil, w, "view/page.html", "view/new_appointment.tmpl"); err != nil{
        return
    }
}

func (n *NewAppointment) Post(w http.ResponseWriter, r *http.Request){
    
    var appointment NewAppointment
    if err := ReadJsonBody(r.Body, &appointment); err != nil{
        log.Printf("error reading the body: %s", err)
        return
    }
    log.Println(appointment)

}

func (n *NewAppointment) Put(w http.ResponseWriter, r *http.Request){
    log.Println(r.PathValue("id"))
    ReadJsonBody(r.Body, &n)
    log.Println(n)
    var hours []string
    for  v := range 16{
        test := time.Date(2025, time.February, 19, 6, 30, 0, 0, time.UTC)
        hours = append(hours, test.Add(time.Minute * (30 * time.Duration(v))).Format(time.TimeOnly))
    }
    temp, err := template.New("test").Parse(`{{range .}}<button type="button" name="time" class="btn" onclick="onTimePick(this)">{{.}}</button>{{end}}`)
    if err != nil{
        log.Printf("error parsing the template: %s", err)
        return
    }
    if err := temp.Execute(w, hours); err != nil{
        log.Printf("error executiong template: %s", err)
    }
}
