package uuid

import (
	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
)

func UUIDv4() string{
	id := uuid.NewV4()
	return id.String()
}

func UUIDv4Raw()[uuid.Size]byte{
	return uuid.NewV4()
}


// NewUUID return new UUID as string
func NewUUID() string {
	id := ksuid.New()
	return id.String()
}

func InspectUUID(uuid string) string{
	id, err := ksuid.Parse(uuid)
	if err != nil {
		return "Invalid UUID"
	}
	return id.String()
}