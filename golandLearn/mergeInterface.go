package main

import (
	"net/http"
	"sync"
	"time"
)

func mergeInterface(timeout int) []*http.Response {
	result := make([]*http.Response, 3)
	urls := []string{"url1", "url2", "url3"}
	begin := time.Second.Milliseconds()
	a := sync.WaitGroup{}
	for k, v := range urls {
		go func(v string) {
			var r *http.Response
			defer func(r1 *http.Response) {
				a.Done()
				result[k] = r
			}(r)
			r, _ = http.Get(v)
		}(v)
	}
	go func() {
		for {
			end := time.Second.Milliseconds()
			if end-begin > int64(timeout) {
				a.Done()
				a.Done()
				a.Done()
				a.Done()
			}
		}
	}()
	a.Add(4)
	a.Wait()
	return result
}
