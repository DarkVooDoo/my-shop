package route

import "net/http"

type Appt struct{
    Hour int
    Minute int
    Length int

}

type PlanningPayload struct{
    Hours []int
    Schedule []Appt
}

var Test []Appt = []Appt{Appt{Hour:9, Minute: 30}, Appt{Hour: 4, Minute: 10}}

var PlanningHandler http.Handler = &PlanningPayload{Hours: make([]int, 24)}

func (p *PlanningPayload) ServeHTTP(w http.ResponseWriter, r *http.Request){

    switch r.Method{
        default:
            p.Get(w, r)
    }
}

func (p *PlanningPayload) Get(w http.ResponseWriter, r *http.Request){
    p.Schedule = Test
    if err := CreatePage(p, w, "view/page.html", "view/planning.tmpl"); err != nil{
        return
    }
}
