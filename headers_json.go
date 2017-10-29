package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/levigross/grequests"
)

// Makes sure a url starts with http:// or https://
// It prepends http:// to it if it doesn't
func guaranteeProtocol(url string) string {
	hasProtocol := strings.HasPrefix(url, "http://") ||
		strings.HasPrefix(url, "https://")

	if !hasProtocol {
		url = "http://" + url
	}
	return url
}

// Tries to fetch a URL, and returns the response's headers
// as a map
func getHeaders(url string) map[string]string {

	resp, err := grequests.Get(guaranteeProtocol(url), nil)
	headers := make(map[string]string)
	if err != nil {
		headers["error"] = "There was a problem fetchign your url"
	}
	for key, val := range resp.Header {
		headers[key] = strings.Join(val, "")
	}
	return headers
}

// This is the main handler for this micro-service.
// If called empty, it returns some informative text.
// If passed a "url" querystring, it tries to fetch that URL
// and returns its response headers as JSON
func mainHandler(ctx echo.Context) error {
	url := ctx.QueryParam("url")
	fmt.Println(url)
	if url == "" {
		welcome := "This micro-service does a GET request to a url " +
			"and returns the response's headers as JSON. " +
			"\nThe request should look like: " +
			"GET thisservice.com/www.yoururl.com"
		return ctx.String(http.StatusOK, welcome)
	}
	return ctx.JSONPretty(http.StatusOK, getHeaders(url), "  ")
}

func main() {
	app := echo.New()
	app.GET("/", mainHandler)
	app.Start(":8080")
}
