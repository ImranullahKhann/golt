# golt
A high-concurrency HTTP load generator built in Go to measure service latency distribution.

### Usage
- The utility accepts five parameters, of which only one is required (url) for the program to run:
    - url: The url of the endpoint to send requests to
    - n: The number of requests to make per second (default: 1000) 
    - c: The number of go routines to concurrently use to make requests (default: 100)
    - t: The amount of time to send requests for in seconds (default: 10)
    - p: The percentile latency to return (default: 90)
- Example Usage:
```bash
golt -url=http://localhost:8080/users -n=2000 -c=400 -t=30 -p=95
# Output
# Responses: 60000
# P95 Latency: 4193
```
