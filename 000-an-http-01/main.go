package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/GoesToEleven")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}

	fmt.Println("Header :", resp.Header.Get("Content-Type"))

	/*
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatalf("error: %s", err)
		}
	*/

	type RcvdData struct {
		Name         string
		Public_Repos int
	}

	var rd RcvdData
	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&rd); err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(rd)
	fmt.Printf("%#v\n", rd)

	// fmt.Println("Body :", resp.Body)
	// fmt.Println("Close :", resp.Close)
	// fmt.Println("ContentLength :", resp.ContentLength)
	// fmt.Println("Header :", resp.Header)
	// fmt.Println("Proto :", resp.Proto)
	// fmt.Println("ProtoMajor :", resp.ProtoMajor)
	// fmt.Println("ProtoMinor :", resp.ProtoMinor)
	// fmt.Println("Request :", resp.Request)
	// fmt.Println("Status :", resp.Status)
	// fmt.Println("TLS :", resp.TLS)
	// fmt.Println("Trailer :", resp.Trailer)
	// fmt.Println("TransferEncoding :", resp.TransferEncoding)
	// fmt.Println("Uncompressed :", resp.Uncompressed)
	// fmt.Println("Cookies :", resp.Cookies())
	// fmt.Println("Request.Method :", resp.Request.Method)
	// fmt.Println("Request.Proto :", resp.Request.Proto)
	// fmt.Println("Request.URL :", resp.Request.URL)

	/*
		JSON <-> GO
		true/false <-> true/false
		string <-> string
		null <-> nil
		number <-> many numeric types
		arry <-> []any
		object <-> map[string]any, struct

		to/from
		[]byte	marshal & unmarshal
		wire	encode & decode
	*/
}
