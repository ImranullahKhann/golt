package requester

import (
	"time"
	"net/http"
)

func MakeReq(url string) (int32, int16){
	start := time.Now()
	
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	statusCode := int16(resp.StatusCode)
	defer resp.Body.Close()

	// time until the headers and the start of the body are recieved; Time to First Byte (TTFB) in microseconds
	ttfbDuration := int32(time.Since(start).Nanoseconds() / 1000)

	return ttfbDuration, statusCode
}
