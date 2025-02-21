package route

import (
	"log"
	"net/http"
)

type Store struct{
    Id string
}


var StoreHandler http.Handler = &Store{}

func (s *Store) ServeHTTP(w http.ResponseWriter, r *http.Request){

    switch r.Method{
        default:
            s.Get(w, r)
    }
}

func (s *Store) Get(w http.ResponseWriter, r *http.Request){

    log.Println(r.PathValue("id"))
    CreatePage(nil, w, "view/page.html", "view/etablissement.tmpl")
}

