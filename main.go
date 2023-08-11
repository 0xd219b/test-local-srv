package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"
)

func main() {
	srvURL := make(chan string)
	go RunSrv(srvURL)

	reqURL := <-srvURL
	// replace 127.0.0.1 to 0.0.0.0
	u, err := url.Parse(reqURL)
	if err != nil {
		panic(err)
	}

	l := u.Scheme + "://0.0.0.0:" + u.Port()
	fmt.Println(l)
	r, err := http.NewRequest("GET", l, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}

func RunSrv(srvURL chan string) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}))
	defer srv.Close()
	srvURL <- srv.URL
	time.Sleep(60 * time.Second)
}
