package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/powerman/check"
)

func TestCountService(tt *testing.T) {
	t := check.T(tt)

	h := countHandler{}
	r := httptest.NewRequest("GET", "/count", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	res := w.Result()
	t.Nil(res.Body.Close())
	t.Equal(res.StatusCode, http.StatusOK)
	rawbody, err := ioutil.ReadAll(res.Body)
	t.Nil(err)
	body := string(rawbody)
	t.Equal(body, "hello world")
}
