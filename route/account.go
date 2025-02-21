package route

import "net/http"

type Account struct{

}

var AccountHandler http.Handler = &Account{}

func (a *Account) ServeHTTP(w http.ResponseWriter, r *http.Request){

    switch r.Method{
        default:
            a.Get(w, r)
    }
}

func (a *Account) Get(w http.ResponseWriter, r *http.Request){

    CreatePage(nil, w, "view/page.html", "view/account.tmpl")
}
