package route

import (
	"log"
	"net/http"
	"planify/model"
	"time"
)

type Sign struct{
}

var SigninHandler http.Handler = &Sign{}

func (s *Sign) ServeHTTP(w http.ResponseWriter, r *http.Request){
    switch r.Method{
        case http.MethodPost:
            post(w, r)
         default:
            s.Get(w, r)
    }

}

func (s *Sign) Get(w http.ResponseWriter, r *http.Request){
    CreatePage(nil, w, "view/page.html", "view/sign.tmpl")
}

func post(w http.ResponseWriter, r *http.Request){
    email := r.FormValue("email")
    password := r.FormValue("password")
    
    cookie := http.Cookie{
        Name: "auth-token",
        Value: "token",
        HttpOnly: true,
        Secure: true,
        Expires: time.Now().Add(time.Minute * 10),
        SameSite: http.SameSiteStrictMode,
        Path: "/",
    }
    http.SetCookie(w, &cookie)
    
    w.Write([]byte("success"))
    log.Println(email, password)
}

func (s *Sign) Put(w http.ResponseWriter, r *http.Request){

    //^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*\W)(?!.* ).{8,16}$
    var user model.User
    if err := ReadJsonBody(r.Body, user); err != nil{
        return
    }
    if user.Confirmation != user.Password{
        return
    }
    user.Create()
}
