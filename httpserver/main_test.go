package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handleHeaderFunc(t *testing.T) {

	tests := []struct {
		name   string
		key    string
		value  string
		expect string
	}{
		{name: "success", key: "X-Native-Cloud", value: "20210926", expect: "20210926"},
		{name: "fail", key: "X-Name", value: "qyq", expect: "qyq"},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/header", nil)
			req.Header.Set(v.key, v.value)
			w := httptest.NewRecorder()

			http.DefaultServeMux.ServeHTTP(w, req)
			result := w.Body.Bytes()

			assert.Equal(t, v.expect, w.Header().Get(v.key))
			assert.Equal(t, "ok", string(result))
		})
	}
}

func Test_handleVersionFunc(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	w := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(w, req)
	result := w.Body.Bytes()

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, VERSION, string(result))
}

func Test_handlehealthzFunc(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(w, req)
	result := w.Body.Bytes()

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", string(result))
}
