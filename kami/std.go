package main

import (
	"io"
	"log"
	"net/http"

	"gopkg.in/redis.v2"
)

var redisDB *redis.Client

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/secret/message", requireKey(secretMessageHandler))
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func requireKey(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("key") != "12345" { // HL
			http.Error(w, "bad key", http.StatusForbidden)
			return
		}
		h(w, r)
	}
}

func requireKeyRedis(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		userID, err := redisDB.Get("auth:" + key).Result() // HL
		if key == "" || err != nil {                       // HL
			http.Error(w, "bad key", http.StatusForbidden)
			return
		}
		log.Println("user", userID, "viewed message")
		h(w, r)
	}
}

func requireKeyOrSession(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		// set key from db if we have a session
		if session := r.FormValue("session"); session != "" { // HL
			var err error
			if key, err = redisDB.Get("session:" + session).Result(); err != nil { // HL
				http.Error(w, "bad session", http.StatusForbidden)
				return
			}
		}
		userID, err := redisDB.Get("auth:" + key).Result()
		if key == "" || err != nil {
			http.Error(w, "bad key", http.StatusForbidden)
			return
		}
		log.Println("user", userID, "viewed message")
		h(w, r)
	}
}

func logRequest()

func secretMessageHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "42")
}
