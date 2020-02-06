package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	scf "github.com/tencentyun/scf-go-lib/cloudevents/scf"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

type DefineEvent struct {
	// test event define
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

type ConfigStruct struct {
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
}

func hello(ctx context.Context, event DefineEvent) (scf.APIGatewayProxyResponse, error) {
	fmt.Println("this is number 22:", event.Key1)
	fmt.Println("what are you doing:", event.Key2)
	f, err := os.Open("./public/index.html")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	text, err := ioutil.ReadAll(f)
	config := new(scf.APIGatewayProxyResponse)
	config.Body = string(text)
	config.IsBase64Encoded = false
	config.StatusCode = http.StatusOK
	header := make(http.Header)
	header.Set("Content-Type", "text/html; charset=utf8")
	h := make(map[string]string)
	for k, v := range header {
		if len(v) == 1 {
			h[k] = v[0]
		}
	}
	config.Headers = h
	return *config, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(hello)
}
