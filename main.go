package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"
)

var err error
const uploadDir = "./uploads"

func main() {
    err = os.Mkdir(uploadDir, os.ModePerm)
    handErr(err)

    http.HandleFunc("POST /submit", formSubHandler)
    port := "8080" 
    fmt.Println("server is running in localhost:" + port)

    err = http.ListenAndServe(":" + port, nil)
    /* to also host on the lan
    err = http.ListenAndServe("0.0.0.0:" + port, nil)
    */
    

    handErr(err)
}

func formSubHandler(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.Method)

    err = os.Mkdir(uploadDir+"/"+strings.ToLower(r.FormValue("teamName")), os.ModePerm)

    for _, subdir := range []string{"code", "reverse"}{
         os.Mkdir(uploadDir+"/"+strings.ToLower(r.FormValue("teamName")+"/"+subdir), os.ModePerm)
    }

    if err != nil {
	http.Error(w, "your name is crazy", http.StatusBadRequest)
		return
    }

    fmt.Fprintf(w, "thank you " + r.FormValue("teamName"))
}

func handErr(err error){
    if err!=nil {
        fmt.Println(err)
    }
}
