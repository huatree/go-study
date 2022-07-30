package infra

import (
  "io/ioutil"
  "net/http"
)

type Retriever struct{}

func (Retriever) Get(url string) string {
  res, err := http.Get(url)
  if err != nil {
    panic(err)
  }

  defer res.Body.Close()
  bytes, _ := ioutil.ReadAll((res.Body))
  return string(bytes)
}
