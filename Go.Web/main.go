package main  

import (  
    "fmt"  
    "log"  
    "net/http"
)

func IndexAction(w http.ResponseWriter, r *http.Request) {  
    fmt.Fprintf(w, "Hello World")  
}

func main() {  
    http.HandleFunc("/", IndexAction)  
    err := http.ListenAndServe(":80", nil)  
    if err != nil {  
        log.Fatal("ListenAndServe: ", err)  
    }
}