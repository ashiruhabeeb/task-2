package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func delete() {

  url := "https://hng-internship-task2-gprm.onrender.com/api/eb14793c-bcfe-4505-b557-14b4e5b210ea"
  method := "DELETE"

  payload := strings.NewReader(``)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}