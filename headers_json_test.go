package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGuaranteeProtocolWorks(t *testing.T) {
	withProtocol := guaranteeProtocol("google.com")
	if withProtocol != "http://google.com" {
		t.Error("Expected http://google.com")
	}
}

func TestGetHeadersDisregardsOkURLs(t *testing.T) {
	okURL := "https://google.com"
	if guaranteeProtocol(okURL) != okURL {
		t.Error("Modifying OK url!")
	}
}

func TestGetHeadersHandlesErrors(t *testing.T) {
	url := "bad url"
	errResp := getHeaders(&url)
	if _, ok := errResp["error"]; !ok {
		t.Error("getHeaders is not handling errors properly")
	}
}

func TestRequestNoArgs(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, mainHandler(c)) {
		assert.Equal(t, rec.Code, 200)
		assert.NotEmpty(t, rec.Body.String())
	}
}

func TestRequestWithURL(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/?url=www.google.com", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, mainHandler(c)) {
		assert.Equal(t, rec.Code, 200)
		var jsonResp map[string]string
		json.Unmarshal(rec.Body.Bytes(), &jsonResp)
		assert.NotEqual(t, len(jsonResp), 0)
	}
}
