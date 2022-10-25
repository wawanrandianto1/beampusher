package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	instanceId = "dc1dfd43-1815-4b60-9ea0-b59937f51dd7"
	secretKey  = "abcde"
)

type Notif struct {
}

func (n *Notif) publishToInterest() (error, map[string]interface{}) {
	url := "https://" + instanceId + ".pushnotifications.pusher.com/publish_api/v1/instances/" + instanceId + "/publishes/interests"
	body := []byte(`{
		"title": "Post title",
		"body": "Post description",
		"userId": 1
	}`)

	resp, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err, map[string]interface{}{}
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	return nil, res
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	http.ListenAndServe(":3001", r)
}
