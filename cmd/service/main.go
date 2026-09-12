package main

import (
 "net/http"
 "log"
)

func main() {
 http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK); _, _ = w.Write([]byte(`{"status":"ok"}`)) })
 log.Fatal(http.ListenAndServe(":8080", nil))
}
