package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func create() {

  url := "https://hng-internship-task2-gprm.onrender.com/api"
  method := "POST"

  payload := strings.NewReader(`{
    "name": "Sayyu Dantata"
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