package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main_page(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("http://localhost:8081/api/hello")
    if err != nil {
        fmt.Println(err)
        http.Error(w, "error", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
	
    fmt.Fprintf(w, "Content from service B", body)
}

func main() {
    http.HandleFunc("/api/main", main_page)
    http.ListenAndServe(":8080", nil)
}
