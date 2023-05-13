package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, f := context.WithTimeout(context.Background(), 6*time.Second)
	defer f()
	r, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:9000/api/v1", nil)

	client := http.Client{}
	rsp, err := client.Do(r)
	if err != nil {
		// To check for any timeout or context cancel we should
		// always compare the 'context.Err()' value of the context
		// object sent in request instead of error retured
		// by 'Client.Do' method.
		if ctx.Err() == context.DeadlineExceeded {
			log.Println(err, "REQUEST TIMEOUT")
			return
		}
		if ctx.Err() == context.Canceled {
			log.Println(err, "DEADLINE EXCEEDED")
			return
		}
		log.Println(err)
	}
	log.Printf("%+v", rsp)
}
