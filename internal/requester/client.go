package requester

import (
	"time"
	"net/http"
)

func Makereq(url string) (int32, int16, error){
	start := time.Now()
	
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	statusCode := int16(resp.StatusCode)
	defer resp.Body.Close()

	// time until the headers and the start of the body are recieved; Time to First Byte (TTFB) in microseconds
	ttfbDuration := int32(time.Since(start).Nanoseconds() / 1000)

	// _, err = io.ReadAll(resp.Body)
	// if err != nil {
		// panic(err)
	// }
// 
	// // time taken to recieve and read the entire response body
	// duration := time.Since(start)

	return ttfbDuration, statusCode, nil
}
