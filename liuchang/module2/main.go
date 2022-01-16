package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func handle(w http.ResponseWriter, r *http.Request) {
	//测试
	w.Write([]byte("<h1> hello world </h1>"))

	//step1
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}

	//step2
	os.Setenv("Version", "1.0")
	version := os.Getenv("Version")
	w.Header().Set("Version", version)

	//step3
	//clientIP := r.RemoteAddr
	clientIP := getCurrentIP(r)
	httpCode := http.StatusOK
	fmt.Printf("clientIP is %s, httpCode is %s \n", clientIP, httpCode)

}

//step4
func healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("200"))
}

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-REAL_IP")

	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}

	return ip
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handle)
	mux.HandleFunc("healthz", healthz)

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal("not start module2 ,%s \n ", err.Error())
	}
}
