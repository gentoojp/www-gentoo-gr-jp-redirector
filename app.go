package main

import (
	"net/http"
	"os"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.gentoo.jp"+r.URL.Path, http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", RedirectHandler)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
