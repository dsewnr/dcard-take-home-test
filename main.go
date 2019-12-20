package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"reqholder"
)

const (
	REQUEST_LIMIT = 60
	TIME_SLOT     = 60 // in second
)

var reqHolders map[string]*reqholder.ReqHolder

func init() {
	// reqHolder = reqholder.ReqHolder{Counter: 0, Limit: REQUEST_LIMIT}
	reqHolders = make(map[string]*reqholder.ReqHolder)
}

func GetClientIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	headers := []string{"X-Real-Ip", "X-Forwarded-For"}
	for _, h := range headers {
		ipFromHeader := r.Header.Get(h)
		if ipFromHeader != "" {
			ip = ipFromHeader
			break
		}
	}
	return ip
}

func handler(w http.ResponseWriter, r *http.Request) {
	// skip favicon request
	if r.RequestURI == "/favicon.ico" {
		return
	}

	// get client IP
	ip := GetClientIP(r)

	// create reqholder if the ip is not in map
	if reqHolders[ip] == nil {
		reqHolders[ip] = &reqholder.ReqHolder{Counter: 0, Limit: REQUEST_LIMIT}
	}

	// increase request counter
	reqHolders[ip].Count()

	fmt.Fprintf(w, "%s", reqHolders[ip].Result())
}

func main() {
	uptime := 0 // server uptime

	// time ticker
	go func() {
		for {
			currentSecond := uptime % TIME_SLOT
			// check all IPs from map
			for k, v := range reqHolders {
				log.Printf("[%s] uptime = %d, currentSecond = %d, result = %s\n", k, uptime, currentSecond, v.Result())
				// reset request counter every TIME_SLOT second(s)
				if currentSecond == 0 {
					v.Reset()
					log.Printf("Counter reset!\n")
				}
			}
			time.Sleep(time.Second)
			uptime++
		}
	}()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
