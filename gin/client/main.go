package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		resp, err := http.PostForm("http://127.0.0.1:3000/post_context", url.Values{"host_client_context": {"192.168.56.12"}, "host_client_status": {"server_on"}})
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("httpPost receive context is err:%v", err)
		}
		fmt.Println(string(body))
		time.Sleep(100 * time.Millisecond)
	}
}