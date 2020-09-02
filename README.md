你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
