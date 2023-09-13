package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func UpdateByName() {

  url := "https://hnginternship-task2.onrender.com/api/?name=Adekunle%20Alao&newName=Issac%20Olanrewaju"
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