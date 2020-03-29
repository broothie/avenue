package router

import (
	"net/http"
	"strings"
)

func (r route) match(request *http.Request) (bool, *http.Request) {
	matchChan := make(chan bool, 3)

	var req *http.Request
	go func() {
		var match bool
		match, req = r.matchPath(request)
		matchChan <- match
	}()

	go func() {
		matchChan <- r.matchQuery(request)
	}()

	go func() {
		matchChan <- r.matchHeaders(request)
	}()

	for i := 0; i < 3; i++ {
		if !<-matchChan {
			return false, nil
		}
	}

	return true, req
}

func (r route) matchPath(request *http.Request) (bool, *http.Request) {
	reqSegments := strings.Split(request.URL.Path, "/")[1:] // [1:] to skip the first `""` from the leading `/`
	rteSegments := r.segments
	if len(reqSegments) != len(rteSegments) {
		return false, nil
	}

	for i, reqSegment := range reqSegments {
		rteSegment := r.segments[i]

		// If not a named segment, simply check for match
		if !strings.HasPrefix(rteSegment, ":") {
			if reqSegment != rteSegment {
				return false, nil
			} else {
				continue
			}
		}

		// Named segment can't have empty value
		if reqSegment == "" {
			return false, nil
		}

		request = setPathVarOnRequest(request, strings.TrimPrefix(rteSegment, ":"), reqSegment)
	}

	return true, request
}

var filler struct{}

func (r route) matchQuery(request *http.Request) bool {
	return matchValuesAgainstStringMap(request.URL.Query(), r.queries)
}

func (r route) matchHeaders(request *http.Request) bool {
	return matchValuesAgainstStringMap(request.Header, r.headers)
}

func matchValuesAgainstStringMap(values map[string][]string, stringMap map[string]string) bool {
	if len(stringMap) == 0 {
		return true
	}

	reqSet := make(map[string]map[string]struct{})
	for key, values := range values {
		reqSet[key] = make(map[string]struct{})

		for _, value := range values {
			reqSet[key][value] = filler
		}
	}

	for rteKey, rteValue := range stringMap {
		reqValueSet, keyExists := reqSet[rteKey]
		if !keyExists {
			return false
		}

		if rteValue == "" {
			continue
		}

		if _, valueExists := reqValueSet[rteValue]; !valueExists {
			return false
		}
	}

	return true
}
