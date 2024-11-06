package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("./web/public"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(
		":3000",
		http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			resp.Header().Add("Cache-Control", "no-cache")
			if strings.HasSuffix(req.URL.Path, ".wasm") {
				resp.Header().Set("content-type", "application/wasm")
			}
			fs.ServeHTTP(resp, req)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
}
