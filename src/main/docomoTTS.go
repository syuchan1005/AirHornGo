package main

import (
	"net/http"
	"fmt"
	"net/http/httputil"
	"strings"
)

var baseUrl string = "https://api.apigw.smt.docomo.ne.jp/aiTalk/v1/textToSpeech"

type docomoTTS struct {
	apiKey string
}

func (docomo *docomoTTS) sendText(ssml string) (responseBody []byte) {
	fmt.Println("send")
	request, _ := http.NewRequest("POST",
		baseUrl + "?APIKEY=" + docomo.apiKey,
		strings.NewReader("<?xml version=\"1.0\" encoding=\"utf-8\" ?>" + ssml))
	header := request.Header
	header.Set("Content-Type", "application/ssml+xml")
	header.Set("Accept", "audio/L16")
	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Can't Connection")
		responseBody = []byte{}
	} else {
		responseBody, _ = httputil.DumpResponse(response, true)
	}
	return
}