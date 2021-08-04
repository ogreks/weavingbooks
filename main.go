package main

import (
    "books/src/member"
    "encoding/json"
    "fmt"
    "os"
    "time"
)

func main() {
    var user member.User = member.User{
        Id: 1,
        Account: "vip1",
        Password: "123456",
        Name: "admin",
        Locknums: 0,
        Locker: false,
        Createtime: time.Now().Unix(),
    }
    jsonString, err := json.Marshal(user)
    fmt.Println(string(jsonString))
    if err != nil {
        panic(err)
    }
	file, err := os.OpenFile(".env", os.O_APPEND,0666)
    if err != nil {
        panic(err)
    }
	defer file.Close()
	n, err := file.WriteString(string(jsonString) + "\n")
    if err != nil {
        panic(err)
    }
    fmt.Println(n)
}
