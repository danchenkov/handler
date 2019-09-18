package main

import (
    "bytes"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    var query bytes.Buffer
    for k, v := range r.Form {
        query.WriteString(fmt.Sprintf("%s=%s&", k, strings.Join(v, "")))
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    newURL, err := url.Parse("http://localhost:5005/sample-form.html")
    if err != nil {
        log.Fatal(err)
    }
    newURL.RawQuery = strings.TrimSuffix(query.String(), "&")
    http.Redirect(w, r, newURL.String(), 301)
}

func main() {
    http.HandleFunc("/handle", handler)      // setting router rule
    err := http.ListenAndServe(":5004", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
