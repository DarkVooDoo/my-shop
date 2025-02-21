package route

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Landpage struct{
    Name string
}

var LandpageHandler http.Handler = &Landpage{}

func (l *Landpage) ServeHTTP(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/"{
        http.NotFound(w, r)

        return
    }
    switch r.Method{
        default:
            l.Get(w, r)
    }
}

func (l *Landpage) Get(w http.ResponseWriter, r *http.Request){
    var test []Landpage
    test = append(test, Landpage{Name: "Hello"}, Landpage{Name: "Moises"})
    CreatePage(test, w, "view/page.html", "view/landpage.tmpl", "view/component/AppointmentCard.tmpl", "view/component/EtablishmentCard.tmpl")
}

func CreatePage(data any, w http.ResponseWriter, pattern ...string)error{
    temp, err := template.ParseFiles(pattern...)
    if err != nil{
        log.Printf("error loading template: %s", err)
        return errors.New("error parsing template")
    }
    w.Header().Add("Content-Encoding", "gzip")
    gz, err := gzip.NewWriterLevel(w, gzip.BestCompression)
    if err != nil{
        log.Printf("error compressing the file: %s", err)
        return nil
    }
    defer gz.Close()
    if err = temp.Execute(gz, data); err != nil{
        log.Printf("error executiong template: %s", err)
        return errors.New("error executing the template")
    }

    return nil
}

func ReadJsonBody(body io.ReadCloser, j interface{})error{
    dec := json.NewDecoder(body)
    if err := dec.Decode(j); err != nil{
        log.Printf("error decoding the json: %s", err)
        return errors.New("error decoding the json")
    }
    return nil
}
