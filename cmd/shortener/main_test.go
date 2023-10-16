package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_webhookPost(t *testing.T) {

	type want struct {
		code        int
		contentType string
	}

	tests := []struct {
		name string
		body string
		want want
	}{
		{
			name: "400 bad request",
			want: want{
				code:        400,
				contentType: "",
			},
		},
		{
			name: "200 simple",
			body: "http://yandex.ru",
			want: want{
				code:        201,
				contentType: "text/plain",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			reqPost := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			recPost := httptest.NewRecorder()
			webhookPost(recPost, reqPost)

			res := recPost.Result()
			assert.Equal(t, tt.want.code, res.StatusCode)
			defer res.Body.Close()

			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))

		})
	}
}
