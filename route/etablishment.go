package route

import "net/http"

type Etablishment struct{

}

var EtablishmentHandler http.Handler = &Etablishment{}

func (e *Etablishment) ServeHTTP(w http.ResponseWriter, r *http.Request){

    switch r.Method{
        default:
            e.Get(w, r)
    }
}

func (e *Etablishment) Get(w http.ResponseWriter, r *http.Request){

}
