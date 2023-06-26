package stringutil

import "encoding/json"

func JsonPretty[T any](obj T) string {
	jsonRaw, err := json.MarshalIndent(&obj, " ", " ")
	if err != nil {
		return ""
	}
	return string(jsonRaw)
}

func Json[T any](obj T) string {
	jsonRaw, err := json.Marshal(&obj)
	if err != nil {
		return ""
	}
	return string(jsonRaw)
}

func Json2Object[T any](jsonStr string, obj T) error {
	return json.Unmarshal([]byte(jsonStr), &obj)
}
