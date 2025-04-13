package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	resp,err := http.Get("https://data-asg.goldprice.org/dbXRates/USD")
  if err!= nil{
    fmt.Println("Error:",err)
    return
  }
  defer resp.Body.Close()

  body,err :=ioutil.ReadAll(resp.Body)
  if err!= nil{
    fmt.Println("Read Error:",err)
    return
  }

  var result struct{
      Items [] struct{
        XAUPrice float64 `json:"xauPrice"`
      }`json:"items"`
  }

  err = json.Unmarshal(body ,&result)
  if err!= nil{
    fmt.Println("JSON parse error:",err)
    return
  }

  fmt.Println("Price:",result.Items[0].XAUPrice)
}


//https://data-asg.goldprice.org/dbXRates/USD

/*{
  "ts": 1744524928510,
  "tsj": 1744524925206,
  "date": "Apr 13th 2025, 02:15:25 am NY",
  "items": [
    {
      "curr": "USD",
      "xauPrice": 3238.825,
      "xagPrice": 32.2775,
      "chgXau": 77.285,
      "chgXag": 1.4926,
      "pcXau": 2.4445,
      "pcXag": 4.8485,
      "xauClose": 3161.54,
      "xagClose": 30.7849
    }
  ]
}*/