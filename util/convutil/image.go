package convutil

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

type ImageFormat string

const (
	PNG  ImageFormat = "PNG"
	JPEG ImageFormat = "JPEG"
	GIF  ImageFormat = "GIF"
)

var supportMimeTypes = map[ImageFormat]string{
	PNG:  "data:image/png;base64,",
	JPEG: "data:image/jpeg;base64,",
	GIF:  "data:image/gif;base64,",
}

func Base64Image2Byte(base64Image string) ([]byte, error) {
	// Remove the prefix that tells us the mime type
	// of the image file
	base64Image = strings.TrimSpace(base64Image)
	var withoutPrefix string
	for _, prefix := range supportMimeTypes {
		if prefix == strings.ToLower(base64Image[0:len(prefix)]) {
			withoutPrefix = strings.TrimSpace(base64Image[len(prefix):])
			break
		}
	}
	// Decode the base64 string into bytes
	b, err := base64.StdEncoding.DecodeString(withoutPrefix)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func Byte2Base64Image(b []byte) (string, error) {
	// Determine the content type of the image file
	mimeType := http.DetectContentType(b)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	builder := strings.Builder{}

	switch mimeType {
	case "image/jpeg":
		builder.WriteString(supportMimeTypes[JPEG])
	case "image/png":
		builder.WriteString(supportMimeTypes[PNG])
	case "image/gif":
		builder.WriteString(supportMimeTypes[GIF])
	default:
		return "", fmt.Errorf("unsupported image type: %s", mimeType)
	}

	// Append the base64 encoded output
	builder.WriteString(base64.StdEncoding.EncodeToString(b))
	return builder.String(), nil
}
