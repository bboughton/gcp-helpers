// Package gsutil provides utility helpers for working with Google Cloud
// Storage. These helpers are meant to simplify some common use cases. To keep
// these helpers simple authentication credentials are retrieved from the
// standard credential resolution chain used by the storage SDK
// 'cloud.google.com/go/storage'.
package gsutil

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"github.com/bboughton/gcp-helpers/gsurl"
	"google.golang.org/api/option"
)

// ReadFile reads the contents of a file located in GCloud Storage returning
// a byte array with its contents or an error. The url is expected to have the
// following format "gs://the-bucket-name/path/to/the/file.ext". Users of this
// func should take care to only load objects of a known size as errors can
// arise if the requested object is too large.
func ReadFile(ctx context.Context, url string) ([]byte, error) {
	u, err := gsurl.Parse(url)
	if err != nil {
		return nil, &InvalidURLError{URL: url, Err: err}
	}

	client, err := storage.NewClient(ctx, option.WithScopes(storage.ScopeReadOnly))
	if err != nil {
		return nil, &ReadError{URL: url, Err: err}
	}
	defer client.Close()

	rc, err := client.Bucket(u.Bucket).Object(u.Object).NewReader(ctx)
	if err != nil {
		return nil, &ReadError{URL: url, Err: err}
	}
	defer rc.Close()

	bytes, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, &ReadError{URL: url, Err: err}
	}

	return bytes, nil
}

// ReadError is an error type that is returned when there is an error reading
// an object from Google Cloud Storage.
type ReadError struct {
	// URL is the Google Cloud Storage url related to the error.
	URL string

	// Err is the underlying error that caused the read failure.
	Err error
}

// Error returns a human readable error message.
func (e *ReadError) Error() string { return "unable to read file " + e.URL }

// InvalidURLError is an error type that is returned if the provided Google
// Cloud Storage url is invalid.
type InvalidURLError struct {
	// URL is the Google Cloud Storage url related to the error.
	URL string

	// Err is the underlying error detailing how the url is invalid.
	Err error
}

// Error returns a human readable error message.
func (e *InvalidURLError) Error() string { return "invalid url " + e.URL + " " + e.Err.Error() }
