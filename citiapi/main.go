package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "os"
  "time"
)

type Result struct {
  ExecutionTime string
  StationBeanList []struct {
    Id, AvailableDocks, TotalDocks int
    StationName, StatusValue, LastCommunicationTime string
  }
}

type Event struct {
  ExecutionTime string `json:"executionTime"`
  Id int `json:"id"`
  AvailableDocks int
  TotalDocks int
  StationName string
  StatusValue string
  LastCommunicationTime string
}

func main() {
  client := &http.Client{Timeout: 10 * time.Second}
  var res Result

  r, err := client.Get("https://feeds.citibikenyc.com/stations/stations.json")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  defer r.Body.Close()

  json.NewDecoder(r.Body).Decode(&res)
  
  for _, v := range res.StationBeanList {
    x := Event{ res.ExecutionTime, v.Id, v.AvailableDocks, v.TotalDocks, v.StationName, v.StatusValue, v.LastCommunicationTime }
    fmt.Println(x)
  }

  //fmt.Println(res)
}



