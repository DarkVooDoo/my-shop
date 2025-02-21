package model

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct{
    Id string `json:"id"`
    Email string `json:"email"`
    Picture string `json:"picture"`
    EntrepriseId string
    jwt.RegisteredClaims
}

type User struct{
    Email string `json:"email"`
    Password string `json:"password"`
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
    Confirmation string `json:"confirmation"`
}

const (
    ACCESS_TOKEN_EXPIRE = time.Hour * 24 * 30 
)


func Signin()error{

    return nil
}

func CreateAccessToken(id string, email string)(string, error){
    claim := UserClaim{
        id,
        email,
        "",
        "",
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(ACCESS_TOKEN_EXPIRE)),
            NotBefore: jwt.NewNumericDate(time.Now()),
            IssuedAt: jwt.NewNumericDate(time.Now()),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
    ss, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_KEY")))
    if err != nil{
        log.Printf("error signing the token: %s", err)
        return "", errors.New("error signing the token")
    }
    return ss, nil

}

func VerifyAccessToken(token string)error{
    tk, err := jwt.ParseWithClaims(token, &UserClaim{}, func(t *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("ACCESS_TOKEN_KEY")), nil
    })

    if err != nil{
        return errors.New("error parsing token")
    }else if _, ok := tk.Claims.(*UserClaim); ok {
        return nil
    } else {
        return errors.New("unknown claims type, cannot proceed")
    }
}
