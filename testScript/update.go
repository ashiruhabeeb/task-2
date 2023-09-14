package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func update() {

  url := "https://hng-internship-task2-gprm.onrender.com/api?id=940440d7-915c-4d37-9f21-c02d423796c3&name=Habeeb%20Ashiru"
  method := "PUT"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

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