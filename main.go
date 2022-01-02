package main

import (
	"PhoneBook/other/subhandle"
	"fmt"
	"net/http"
	"time"
)

func main() {

	err := subhandle.ReadCSVFile(subhandle.CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = subhandle.CreateIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         subhandle.PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	fmt.Println("Server ready to serves")

	mux.Handle("/list", http.HandlerFunc(subhandle.ListHandler))
	mux.Handle("/insert/", http.HandlerFunc(subhandle.InsertHandler))
	mux.Handle("/insert", http.HandlerFunc(subhandle.InsertHandler))
	mux.Handle("/search", http.HandlerFunc(subhandle.SearchHandler))
	mux.Handle("/search/", http.HandlerFunc(subhandle.SearchHandler))
	mux.Handle("/delete/", http.HandlerFunc(subhandle.DeleteHandler))
	mux.Handle("/status", http.HandlerFunc(subhandle.StatusHandler))
	mux.Handle("/", http.HandlerFunc(subhandle.DefaultHandler))

	fmt.Println("Ready to serve at", subhandle.PORT)
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
