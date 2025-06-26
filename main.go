package main

import (
	"github.com/syumai/workers"
	_ "github.com/v2fly/v2ray-core/v5/main/distro/all"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		msg := "Myservers"
		w.Write([]byte(msg))
	})
	workers.Serve(nil) // use http.DefaultServeMux
}
