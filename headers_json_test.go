package main

import "testing"

func TestGetHeadersInjectsHttp(t *testing.T) {
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
