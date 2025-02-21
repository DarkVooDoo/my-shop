package model

import "context"


func (u *User) Create() error{
    conn := GetDBPoolConn()

    conn.QueryRowContext(context.Background(), "INSERT INTO Users(firstname, lastname, email) VALUES($1,$2,$3)", u.Firstname, u.Lastname, u.Email)

    return nil
}
