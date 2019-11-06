package gsurl_test

import (
	"testing"

	"github.com/bboughton/gcp-helpers/gsurl"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name  string
		url   string
		gsurl *gsurl.URL
		err   bool
	}{
		{
			name: "Basic URL",
			url:  "gs://aaa/b",
			gsurl: &gsurl.URL{
				Bucket: "aaa",
				Object: "b",
			},
			err: false,
		},
		{
			name:  "Protocal required",
			url:   "//a/b",
			gsurl: nil,
			err:   true,
		},
		{
			name:  "Protocal must be gs",
			url:   "http://a/b",
			gsurl: nil,
			err:   true,
		},
		{
			name:  "Bucket name is required",
			url:   "gs:///b",
			gsurl: nil,
			err:   true,
		},
		{
			name:  "Object name is required",
			url:   "gs://a",
			gsurl: nil,
			err:   true,
		},
		{
			name:  "Object name is required, check trailing slash",
			url:   "gs://a/",
			gsurl: nil,
			err:   true,
		},
		{
			name: "Object name doesn't have leading slash",
			url:  "gs://bucket/object",
			gsurl: &gsurl.URL{
				Bucket: "bucket",
				Object: "object",
			},
			err: false,
		},
		{
			name:  "Invalid URL",
			url:   "%",
			gsurl: nil,
			err:   true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Given

			// When
			u, err := gsurl.Parse(tc.url)

			// Then
			if u != nil {
				assert.EqualValues(t, tc.gsurl, u)
			}
			assert.Equal(t, tc.err, err != nil)
		})
	}
}
