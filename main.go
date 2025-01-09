package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	encoder "github.com/adeelkhan/base64encoder/encoder"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Base64Encode(w http.ResponseWriter, r *http.Request) {
	var p encoder.RequestEncode
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	encodedString := encoder.EncodeString(p.PlainText)
	bytes, err := json.Marshal(encoder.ResponseEncode{Base64EncodedString: encodedString})
	if err != nil {
		errorBytes, _ := json.Marshal(encoder.ResponseError{Message: "Error occured"})
		fmt.Fprint(w, string(errorBytes))
		return
	}
	fmt.Fprint(w, string(bytes))
}

func Base64Decode(w http.ResponseWriter, r *http.Request) {
	var p encoder.RequestDecode
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	plainText, err := encoder.DecodeString(p.Base64EncodedString)
	if err != nil {
		errorBytes, _ := json.Marshal(encoder.ResponseError{Message: "Error occured"})
		fmt.Fprint(w, string(errorBytes))
		return
	}
	bytes, err := json.Marshal(encoder.ResponseDecode{PlainText: plainText})
	if err != nil {
		errorBytes, _ := json.Marshal(encoder.ResponseError{Message: "Error occured"})
		fmt.Fprint(w, string(errorBytes))
		return
	}
	fmt.Fprint(w, string(bytes))
}

func main() {
	// server setup
	r := mux.NewRouter()

	r.HandleFunc("/encode", Base64Encode)
	r.HandleFunc("/decode", Base64Decode)

	serverAddress := "localhost:8000"

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	srv := &http.Server{
		Handler:      handler,
		Addr:         serverAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Server starting at...%s", serverAddress)
	log.Fatal(srv.ListenAndServe(), handler)
}
