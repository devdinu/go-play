package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func modifyResponse(resp *http.Response) error {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	if err != nil {
		return err
	}

	b = bytes.Replace(b, []byte("server"), []byte("new gojek server"), -1)
	body := ioutil.NopCloser(bytes.NewReader(b))

	resp.Body = body
	resp.ContentLength = int64(len(b))
	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))

	return nil
}

func main() {
	target, err := url.Parse("https://www.google.com")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ModifyResponse = modifyResponse

	http.Handle("/", proxy)
	http.ListenAndServe(":8080", nil)
}
