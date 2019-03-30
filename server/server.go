package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
	"uaas/img"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root")
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	width, _ := strconv.ParseUint(vars["width"], 10, 32)
	height, _ := strconv.ParseUint(vars["height"], 10, 32)

	image := img.ProcessImage(uint(width), uint(height))

	_, err := w.Write(image.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}

func New() {
	r := mux.NewRouter()

	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/{width}/{height}", ImageHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	fmt.Println("vim-go")
}
