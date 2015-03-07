package main

import (
  "encoding/json"
  "net/http"
)

type Profile struct {
  FirstName  	string
  LastName		string 
  Hobbies 		[]string
}

func main() {
  http.HandleFunc("/profiles", handleProfile)
  http.ListenAndServe(":3000", nil)
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
  profile := Profile{"Christopher", "Bartling", []string {"horse riding", "programming"}}

  js, err := json.Marshal(profile)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json") 
  w.Write(js) 
}
