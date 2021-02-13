package urlshort

import (
	"encoding/json"
	"net/http"
)

// JSONHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//	[
//		{
//			"path": "/some-path",
//			"url": "https://www.some-url.com/demo"
//		},
//		{
//			"path": "/some-other-path",
//			"url": "https://www.some-other-url.com/demo"
//		}
//	]
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(jsonByte []byte, fallback http.Handler) (http.HandlerFunc, error) {
	jsonPaths, err := parseJSON(jsonByte)
	if err != nil {
		return nil, err
	}
	redirectMap := buildRedirectMap(jsonPaths)
	return MapHandler(redirectMap, fallback), nil
}

func parseJSON(jsonByte []byte) ([]redirect, error) {
	var redirects []redirect
	err := json.Unmarshal(jsonByte, &redirects)
	return redirects, err
}
