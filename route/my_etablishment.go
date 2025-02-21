package route

import "net/http"

type MyEtablishmentRoute struct{

}

var MyEtablishmentHandler http.Handler = &MyEtablishmentRoute{}

func (me *MyEtablishmentRoute) ServeHTTP(w http.ResponseWriter, r *http.Request){

    switch r.Method{
        default:
            me.Get(w, r)
    }
}

func (me *MyEtablishmentRoute) Get(w http.ResponseWriter, r *http.Request){
    if err := CreatePage(nil, w, "view/page.html", "view/my_etablishment.tmpl"); err != nil{
        return
    }
}
