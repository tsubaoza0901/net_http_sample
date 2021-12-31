package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	var (
		targetURL string    = "http://go:6060/" // コンテナ間通信用に、ホストは「localhost」ではなく、「go（server側のdocker-compose.ymlのservice名）」を指定
		method    string    = "GET"
		body      io.Reader = nil
	)

	log.Println("Start Http Request")

	var count = 0
	for i := 0; i < 30; i++ {
		req, err := http.NewRequest(method, targetURL, body)
		if err != nil {
			log.Println(err.Error())
		}
		if count == 0 {
			dump, _ := httputil.DumpRequestOut(req, true)
			fmt.Printf("dump req: %s\n", dump)
		}

		client := new(http.Client)
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer res.Body.Close()
		if count == 0 {
			dumpResp, err := httputil.DumpResponse(res, true)
			if err != nil {
				log.Println(err.Error())
			}
			fmt.Printf("dump res: %s\n", dumpResp)

			// body, _ := ioutil.ReadAll(res.Body)
			// fmt.Printf("res.Body: %s\n", string(body))
		}
		count = count + 1
	}
}
