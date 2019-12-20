package main

import (
	"net/http"
	"net/http/httptest"
	"reqholder"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	t.Run("client1 under limit request", func(t *testing.T) {
		for i := 1; i <= REQUEST_LIMIT; i++ {
			request, _ := http.NewRequest(http.MethodGet, "/", nil)
			request.Header.Add("X-Forwarded-For", "127.0.0.1")
			response := httptest.NewRecorder()

			handler(response, request)

			respText := response.Body.String()

			exceptReqCountText := strconv.Itoa(i)
			assert.Equal(t, respText, exceptReqCountText)
		}
	})

	t.Run("client1 over limit request", func(t *testing.T) {
		for i := 1; i <= 1; i++ {
			request, _ := http.NewRequest(http.MethodGet, "/", nil)
			request.Header.Add("X-Forwarded-For", "127.0.0.1")
			response := httptest.NewRecorder()

			handler(response, request)

			respText := response.Body.String()

			exceptReqCountText := reqholder.RESP_OVER_LIMIT
			assert.Equal(t, respText, exceptReqCountText)
		}
	})

	t.Run("client2 under limit request", func(t *testing.T) {
		for i := 1; i <= REQUEST_LIMIT; i++ {
			request, _ := http.NewRequest(http.MethodGet, "/", nil)
			request.Header.Add("X-Forwarded-For", "::1")
			response := httptest.NewRecorder()

			handler(response, request)

			respText := response.Body.String()

			exceptReqCountText := strconv.Itoa(i)
			assert.Equal(t, respText, exceptReqCountText)
		}
	})

	t.Run("client2 over limit request", func(t *testing.T) {
		for i := 1; i <= 1; i++ {
			request, _ := http.NewRequest(http.MethodGet, "/", nil)
			request.Header.Add("X-Forwarded-For", "::1")
			response := httptest.NewRecorder()

			handler(response, request)

			respText := response.Body.String()

			exceptReqCountText := reqholder.RESP_OVER_LIMIT
			assert.Equal(t, respText, exceptReqCountText)
		}
	})
}
