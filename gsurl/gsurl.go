package gsurl

import (
	"errors"
	"net/url"
	"strings"
)

// parseURL exists to allow us to use 'url' as a request param
var parseURL = url.Parse

// Parse parses a Google Cloud Storage string into a URL struct. The expected
// format of the string is gs://[bucket-name]/[object-path]. If the provided
// URL is formatted incorrectly an error will be returned.
func Parse(url string) (*URL, error) {
	u, err := parseURL(url)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "gs" {
		return nil, errors.New("Invalid protocal specified, the only protocal that is permitted is 'gs'.")
	}

	bucket, object := u.Host, strings.TrimLeft(u.Path, "/")

	if bucket == "" {
		return nil, errors.New("Bucket name is required")
	}

	if object == "" {
		return nil, errors.New("Object name is required")
	}

	return &URL{
		Bucket: bucket,
		Object: object,
	}, nil
}

// URL contains the information needed to identify the location of an object
// located in Google Cloud Storage.
type URL struct {
	// Bucket is the name of the Google Cloud Storage bucket where the object
	// is located.
	Bucket string

	// Object is the name and or path of the object stored in the bucket. It
	// should not start with a foward slash.
	Object string
}
