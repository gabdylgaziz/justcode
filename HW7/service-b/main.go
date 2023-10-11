package main

import (
    "fmt"
    "net/http"
	"io/ioutil"
)

func main_page(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("http://localhost:8080/api/main")
    if err != nil {
        fmt.Println(err)
        http.Error(w, "error", http.StatusInternalServerError)
        return
    }
	body, _ := ioutil.ReadAll(resp.Body)
	
    fmt.Fprintf(w, "Success! Code: %s", body)
}

func main() {
    http.HandleFunc("/api/hello", main_page)
    http.ListenAndServe(":8081", nil)
}
