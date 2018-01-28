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
	return func(w http.ResponseWriter, r *http.Request) {
		if path, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, path, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []struct {
		path string `yaml:"path"`
		URL  string `yaml:"url"`
	}
	if err := yaml.Unmarshal(yml, &pathUrls); err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		for _, pu := range pathUrls {
			if pu.path == r.URL.Path {
				http.Redirect(w, r, pu.URL, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}

func JSONHandler(myJson []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []struct {
		Path string `json:"path"`
		URL  string `json:"url"`
	}
	if err := json.Unmarshal(myJson, &pathUrls); err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		for _, pu := range pathUrls {
			if pu.Path == r.URL.Path {
				http.Redirect(w, r, pu.URL, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}
