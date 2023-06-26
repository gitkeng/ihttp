package httputil

import "net/url"

func QueryString(queries map[string]string) string {
	params := url.Values{}
	for key,value := range queries{
		params.Add(key,value)
	}
	return params.Encode()
}
