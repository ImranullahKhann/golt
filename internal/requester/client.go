package requester

import (
	"fmt"
	"time"
	"net/http"
	"io"
)

func Makereq(url string) {
	printf := fmt.Printf;

	start := time.Now()
	
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// time until the headers and the start of the body are recieved; Time to First Byte (TTFB)
	ttfbDuration := time.Since(start)

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// time taken to recieve and read the entire response body
	duration := time.Since(start)

	printf("Response Status: %v \n", resp.Status)
	printf("Time to First Byte: %v \n", ttfbDuration)
	printf("Response Duration: %v \n", duration)
}
