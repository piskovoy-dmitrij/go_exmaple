package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    "time"
)

type Status struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func main () {
    router := httprouter.New()
    router.GET("/status", getStatus)
    router.GET("/keys/:key", getKey)
    router.POST("keys/:key", addKey)
    router.DELETE("keys/:key", deleteKey)
    http.ListenAndServe(":8080", router)
}

func getStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rec := Status{Message: "Ok", Timestamp: time.Now().Unix()}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if json.NewEncoder(w).Encode(rec) != nil {
		w.WriteHeader(500)
	}
}

func getKey(key string) string {

}