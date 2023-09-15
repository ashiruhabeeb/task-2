package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func update() {

  url := "https://hng-internship-task2-gprm.onrender.com/api/eb14793c-bcfe-4505-b557-14b4e5b210ea"
  method := "PUT"

  payload := strings.NewReader(`{
    "name": "Ibrahim Mahama"
}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")

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