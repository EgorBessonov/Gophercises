package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(rw http.ResponseWriter, req *http.Request) {
		path := req.URL.Path
		if fullPath, fpExists := pathsToUrls[path]; fpExists {
			http.Redirect(rw, req, fullPath, http.StatusFound)
			return
		}
		fallback.ServeHTTP(rw, req)
	}
}

type pathAndURL struct {
	path string
	URL  string
}

func createMap(pathAndURLs []pathAndURL) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, el := range pathAndURLs {
		pathsToUrls[el.path] = el.URL
	}
	return pathsToUrls
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(ymlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathsAndURLs []pathAndURL
	err := yaml.Unmarshal(ymlBytes, &pathsAndURLs)
	if err != nil {
		return nil, err
	}

	pathsToUrls := createMap(pathsAndURLs)
	return MapHandler(pathsToUrls, fallback), nil
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathsAndURLs []pathAndURL
	err := json.Unmarshal(jsonBytes, &pathsAndURLs)
	if err != nil {
		return nil, err
	}
	pathsToUrls := createMap(pathsAndURLs)
	return MapHandler(pathsToUrls, fallback), nil
}
