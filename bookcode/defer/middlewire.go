package main

import (
	"log"
	"net/http"
	"fmt"
	"time"
)

type MyHandler struct {}

func (m MyHandler) ServeHTTP(w http.ResponseWriter,  r *http.Request) {
	before:= time.Now()
	defer func(){
		after := time.Now()
		fmt.Println(after.After(before))
	}()
	fmt.Fprintf(w,"hello world")
}

func main(){
	handler :=MyHandler{}
	server:= http.Server{
		Addr:"127.0.0.1:8080",
		Handler:recoverHandler(handler),
	}
	server.ListenAndServe()
}

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}