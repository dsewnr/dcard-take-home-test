# Take Home Test for Web Backend Developer

## How it works
Set environment variable and run server.
```
$ export GO111MODULE=on
$ npm start
```
Open browser with http://localhost:8080/

## Test
```
# test server
$ go test -v 

# test request handler
$ go test -v reqholder 
```

## Set request limit and time slot
```
const (
	REQUEST_LIMIT = 60
	TIME_SLOT     = 60 // in second
)
```


## About ReqHolder structure

Thus we need to keep each request status by IP address. So we use `ReqHolder` structure from `reqholder` package, and define a map to keep its status.
```
var reqHolders map[string]*reqholder.ReqHolder
```
Count requests
```
reqHolders[ip].Count()
```
Validate and response
```
reqHolders[ip].Result()
```
Reset counter every n seconds
```
reqHolders[ip].Reset()
```
