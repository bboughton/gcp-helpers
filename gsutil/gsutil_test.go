package gsutil_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/storage"
	"github.com/bboughton/gcp-helpers/gsutil"
	"github.com/stretchr/testify/assert"
)

var testdataStorageBucket = os.Getenv("GSUTIL_TEST_BUCKET")

func TestReadFile_error(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{"invalid scheme", "file://bad-scheme"},
		{"invalid url", "://a"},
		{"invalid bucket", "gs://a"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Given
			ctx := context.Background()

			// When
			bytes, err := gsutil.ReadFile(ctx, tc.url)

			// Then
			assert.Error(t, err)
			assert.Nil(t, bytes)
		})
	}
}

func TestReadFile(t *testing.T) {
	if testdataStorageBucket == "" {
		t.Skip("to run TestReadFile set env var 'GSUTIL_TEST_BUCKET'")
	}
	// Given
	path := fmt.Sprintf("%s-%d", t.Name(), time.Now().Unix())
	expected := []byte("a")
	url, done := genTestFile(t, path, expected)
	defer done()
	ctx := context.Background()

	// When
	actual, err := gsutil.ReadFile(ctx, url)

	// Then
	assert.NoError(t, err)
	assert.Equalf(t, expected, actual, "unable to retrieve test file, %s", url)
}

func genTestFile(t *testing.T, path string, data []byte) (string, func()) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		t.Fatal(err)
	}
	object := client.Bucket(testdataStorageBucket).Object(path)
	wc := object.NewWriter(ctx)
	defer wc.Close()
	_, err = wc.Write(data)
	if err != nil {
		client.Close()
		t.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		client.Close()
		t.Fatal(err)
	}
	return fmt.Sprintf("gs://%s/%s", testdataStorageBucket, path), func() {
		defer client.Close()
		err := object.Delete(ctx)
		if err != nil {
			t.Logf("object '%s' was not able to be deleted", path)
		}
	}
}
