package bqurl_test

import (
	"strings"
	"testing"

	"github.com/bboughton/gcp-helpers/bqurl"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name      string
		url       string
		projectID string
		datasetID string
		tableID   string
		err       bool
	}{
		{
			name:      "basic url",
			url:       "a-1234:b.c",
			projectID: "a-1234",
			datasetID: "b",
			tableID:   "c",
			err:       false,
		},
		{
			name:      "empty url not allowed",
			url:       "",
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
		{
			name:      "only table id is provided",
			url:       "a",
			projectID: "",
			datasetID: "",
			tableID:   "a",
			err:       false,
		},
		{
			name:      "only dataset and table id's are provided",
			url:       "a.b",
			projectID: "",
			datasetID: "a",
			tableID:   "b",
			err:       false,
		},
		{
			name:      "project id is too short",
			url:       strings.Repeat("a", 5) + ":b.c",
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
		{
			name:      "project id is too long",
			url:       strings.Repeat("a", 31) + ":b.c",
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
		{
			name:      "project id may only contain lowercase letters",
			url:       "aaaaaA:b.c",
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
		{
			name:      "project id may not contain spaces",
			url:       "aa aaa:b.c",
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
		{
			name:      "project id may not special characters",
			url:       "aaa!aa:b.c",
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
		{
			name:      "project id may not start with a digit",
			url:       "1aaaaa:b.c",
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
		{
			name:      "project id may not start with a hyphen",
			url:       "-aaaaa:b.c",
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
		{
			name:      "project id may not end with a hyphen",
			url:       "aaaaa-:b.c",
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
		{
			name:      "dataset id is too long",
			url:       "aaaaaa:" + strings.Repeat("b", 1025) + ".c",
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
		{
			name:      "dataset id may contain uppercase letters",
			url:       "aaaaaa:B.c",
			projectID: "aaaaaa",
			datasetID: "B",
			tableID:   "c",
			err:       false,
		},
		{
			name:      "dataset id may contain digits",
			url:       "aaaaaa:1.c",
			projectID: "aaaaaa",
			datasetID: "1",
			tableID:   "c",
			err:       false,
		},
		{
			name:      "dataset id may contain underscores",
			url:       "aaaaaa:_.c",
			projectID: "aaaaaa",
			datasetID: "_",
			tableID:   "c",
			err:       false,
		},
		{
			name:      "dataset id may not contain special characters",
			url:       "aaaaaa:b!.c",
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
		{
			name:      "table id is too short",
			url:       "aaaaaa:b.",
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
		{
			name:      "table id is too long",
			url:       "aaaaaa:b." + strings.Repeat("c", 1025),
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
		{
			name:      "table id may contain uppercase letters",
			url:       "aaaaaa:b.C",
			projectID: "aaaaaa",
			datasetID: "b",
			tableID:   "C",
			err:       false,
		},
		{
			name:      "table id may contain digits",
			url:       "aaaaaa:b.1",
			projectID: "aaaaaa",
			datasetID: "b",
			tableID:   "1",
			err:       false,
		},
		{
			name:      "table id may contain underscores",
			url:       "aaaaaa:b._",
			projectID: "aaaaaa",
			datasetID: "b",
			tableID:   "_",
			err:       false,
		},
		{
			name:      "table id may not contain special characters",
			url:       "aaaaaa:b.c!",
			projectID: "",
			datasetID: "",
			tableID:   "",
			err:       true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Given

			// When
			u, err := bqurl.Parse(tc.url)

			// Then
			if u != nil {
				assert.Equal(t, tc.projectID, u.ProjectID)
				assert.Equal(t, tc.datasetID, u.DatasetID)
				assert.Equal(t, tc.tableID, u.TableID)
			}
			assert.Equal(t, tc.err, err != nil)
		})
	}
}
