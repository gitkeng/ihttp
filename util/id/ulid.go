package id

import (
	"encoding/hex"
	"fmt"
	"github.com/segmentio/ksuid"
	"strings"
)

func UUID() string {
	id := ksuid.New()
	return id.String()
}

func Inspect(uuid string) (string, error) {
	id, err := ksuid.Parse(uuid)
	if err != nil {
		return "", err
	}
	const inspectFormat = `
REPRESENTATION:
  String: %v
     Raw: %v
COMPONENTS:
       Time: %v
  Timestamp: %v
    Payload: %v
`
	return fmt.Sprintf(inspectFormat,
		id.String(),
		strings.ToUpper(hex.EncodeToString(id.Bytes())),
		id.Time(),
		id.Timestamp(),
		strings.ToUpper(hex.EncodeToString(id.Payload())),
	), nil
}
