package reqholder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReqHolder(t *testing.T) {
	limit := 60
	// init instance
	reqHolder := ReqHolder{Counter: 0, Limit: limit}
	assert.Equal(t, reqHolder.Result(), "0", "they should be equal")
	// increase request counter
	reqHolder.Count()
	assert.Equal(t, reqHolder.Result(), "1", "they should be equal")
	// reset request counter
	reqHolder.Reset()
	assert.Equal(t, reqHolder.Result(), "0", "they should be equal")
	// let request counter to be over limit
	for i := 0; i < limit+1; i++ {
		reqHolder.Count()
	}
	assert.Equal(t, reqHolder.Result(), "Error", "they should be equal")
	// reset request counter
	reqHolder.Reset()
	assert.Equal(t, reqHolder.Result(), "0", "they should be equal")
}
