package route

import "net/http"

type Search struct{

}

var SearchHandler http.Handler = &Search{}

func (s *Search) ServeHTTP(w http.ResponseWriter, r *http.Request){

    switch r.Method{
        case http.MethodPost:
            s.Post(w, r)
        default:
            CreatePage(nil, w, "view/page.html", "view/search.tmpl")
    }
}

func (s *Search) Post(w http.ResponseWriter, r *http.Request){


}
