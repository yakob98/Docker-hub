package main

import (
        "fmt"
        "net/http"
)
// func main() {
//     fmt.Println("Hello, World!")
// }
func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World%s!", r.URL.Path[1:])
}
