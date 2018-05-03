package main

import (
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte(`
    <html>
      <head>
        <title>chat</title>
      </head>
      <body>
        Let's Chat!!
      </body>
    </html>
    `))
  })

  // web server start
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
