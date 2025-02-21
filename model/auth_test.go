package model

import (
	"log"
	"regexp"
	"testing"
	"time"
)

func TestRegex(t *testing.T){

    //var dates []string
    test := time.Date(2025, time.February, 19, 12, 45, 0, 0, time.UTC)
    for v := range 6{
        ti := test.Add(time.Minute * (30 * time.Duration(v)))
        log.Println(ti.Format(time.TimeOnly))
    }
}

func TestPassword(t *testing.T){
    r := regexp.MustCompile(`(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}`)
    log.Println(r.Match([]byte("Test0423!")))
}
