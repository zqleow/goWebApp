package main

import (
    "fmt"
    "math/rand"
    "net/url"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", content_render)
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
    http.ListenAndServe(":8080", nil)
}

func content_render(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Pre-Course Assignment</h1>")
    fmt.Fprintf(w, "<title>Go Web App</title>")
    fmt.Fprintf(w, "<img src='assets/image.jpeg'>")
    dynamicText:= [5]string{"Logic will get you from A to B. Imagination will take you everywhere.", "There are 10 kinds of people. Those who know binary and those who don't.", "There are two ways of constructing a software design. One way is to make it so simple that there are obviously no deficiencies and the other is to make it so complicated that there are no obvious deficiencies.", 
   "It's not that I'm so smart, it's just that I stay with problems longer.",
   "It is pitch dark. You are likely to be eaten by a grue."}
   randomIndex := rand.Intn(len(dynamicText))
   pick := dynamicText[randomIndex]
   fmt.Fprintf(w, "<h2>%s</h2>", pick)
   URL, error := url.Parse("https://github.com/zqleow/goWebApp")
   if error != nil {
        log.Fatal("Error accessing repository" ,error)
   }
   fmt.Fprintf(w, "<body><p>Repository:%v\r\n<p><body>",URL)
}
