package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pusher/pusher-http-go/v5"
	"github.com/rs/cors"
)

// const (
// 	instanceId = "dc1dfd43-1815-4b60-9ea0-b59937f51dd7"
// 	secretKey  = "abcde"
// )

// type Notif struct {
// }

// func (n *Notif) publishToInterest() (error, map[string]interface{}) {
// 	url := "https://" + instanceId + ".pushnotifications.pusher.com/publish_api/v1/instances/" + instanceId + "/publishes/interests"
// 	body := []byte(`{
// 		"title": "Post title",
// 		"body": "Post description",
// 		"userId": 1
// 	}`)

// 	resp, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
// 	if err != nil {
// 		return err, map[string]interface{}{}
// 	}

// 	var res map[string]interface{}
// 	json.NewDecoder(resp.Body).Decode(&res)
// 	return nil, res
// }

type Respon struct {
	Message string `json:"message"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Post("/pusher/auth", func(w http.ResponseWriter, r *http.Request) {
		pusherClient := pusher.Client{
			AppID:   "1494118",
			Key:     "e48e401e9da913009f95",
			Secret:  "8bf1185ce7e6c74b65dd",
			Cluster: "ap1",
		}

		log.Printf("received method: %s", r.Method)
		params, _ := io.ReadAll(r.Body) // expects params to contain socket_id & channel_name

		log.Printf("params: %v", string(params))

		// userData := map[string]interface{}{"id": "1234"} // get user id from auth service
		response, err := pusherClient.AuthorizePrivateChannel(params)
		if err != nil {
			log.Printf("error pusher: %v", err)
			http.Error(w, http.StatusText(500), 500)
			w.Write([]byte("auth private channel :" + err.Error()))
			return
		}

		log.Printf("response: %s", response)

		responseOk, err := json.Marshal(response)
		if err != nil {
			log.Printf("error: %v", err)
			http.Error(w, http.StatusText(500), 500)
			w.Write([]byte("json convert : " + err.Error()))
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseOk)
	})

	handler := cors.AllowAll().Handler(r)
	http.ListenAndServe(":5000", handler)
}
