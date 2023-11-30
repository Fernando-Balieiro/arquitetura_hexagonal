package handler

import "encoding/json"

func jsonError(msg string) []byte {
  err := struct {
    Message string `json:"message"`
  }{}

  r, errr := json.Marshal(err)
  if errr != nil {
    return []byte(errr.Error())
  }
}
