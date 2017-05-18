package main

import (
    "fmt"
    "log"

    Metrika "github.com/xboston/yametrikago"
)

func main() {

    log.Println("Start")

    var (
        clientId, clientSecret, username, password, token, code string
        err                                                     error
    )

    //
    clientId = "bd59b41747a247b1ad96c90131dc6568"
    clientSecret = ""
    username = ""
    password = "2583e49576854531bbc5d7e8669276f5"

    //
    code = ""

    // при указании токена остальные параметры не обязательны
    token = ""

    metrika := Metrika.NewMetrika(clientId, clientSecret, username, password, token, code)
    metrika.SetDebug(true)

    err = metrika.Authorize()
    PanicIfErr(err)

    counterList, err := metrika.GetCounterList()
    PanicIfErr(err)

    Debug(counterList)
}
