/*
Package reqholder provides a data structure to keep and handle status from http requests.
*/
package reqholder

import "fmt"

const RESP_OVER_LIMIT = "Error"

type ReqHolder struct {
	Counter int
	Limit   int
}

// count request
func (rh *ReqHolder) Count() {
	rh.Counter++
}

// validate request counter and return result
func (rh *ReqHolder) Result() string {
	result := fmt.Sprintf("%d", rh.Counter)
	// response "Error" if over the limit
	if rh.Counter > rh.Limit {
		result = RESP_OVER_LIMIT
	}
	return result
}

// reset request counter
func (rh *ReqHolder) Reset() {
	rh.Counter = 0
}
