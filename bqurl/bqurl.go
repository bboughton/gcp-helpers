package bqurl

import (
	"errors"
	"strings"
)

// Parse parses a BigQuery url string into a URL struct. The expected format of
// the string is [project-id]:[dataset-id].[table-id] where the table id is the
// only required section. If the provided url string is formatted incorrectly
// or if the individual parts are not valid an error will be returned.
func Parse(u string) (*URL, error) {
	if u == "" {
		return nil, errors.New("url must not be blank")
	}

	var p, d, t string
	if i := strings.Index(u, ":"); i > -1 {
		p, u = u[:i], u[i+1:]
	}
	if i := strings.Index(u, "."); i > -1 {
		d, u = u[:i], u[i+1:]
	}
	t = u

	if err := validateProjectID(p); err != nil {
		return nil, err
	}
	if err := validateDatasetID(d); err != nil {
		return nil, err
	}
	if err := validateTableID(t); err != nil {
		return nil, err
	}

	return &URL{
		ProjectID: p,
		DatasetID: d,
		TableID:   t,
	}, nil
}

// URL contains the information needed to identify the location of a BigQuery
// table.
type URL struct {
	// ProjectID is a Google Cloud Platform Project ID. It must be between 6
	// and 30 characters. It may consist of lowercase letters, digits, or
	// hyphens. It must start with a letter and it may not end with a hyphen.
	ProjectID string

	// DatasetID is a BigQuery Dataset ID. It must be no longer then 1024
	// characters. It may consist of lowercase letters, uppercase letters,
	// digits, or underscores.
	DatasetID string

	// TableID is a BigQuery Table ID. It must be no longer then 1024
	// characters. It may consist of lowercase letters, uppercase letters,
	// digits, or underscores.
	TableID string
}

// validateProjectID validates the project id against a set of rules and
// returns an error if there is a violation. As a special case if the project
// id is empty no error will be returned since it is valid for the project id
// to not be supplied.
//
// Rules:
//   - Must be 6 to 30 characters.
//   - Must only contain lowercase letters, digits, or hyphens.
//   - Must start with a letter.
//   - Must not end with a hyphen.
func validateProjectID(p string) error {
	// special case: when string is empty return nil error because we allow for
	// no project id to be set
	if p == "" {
		return nil
	}
	if len(p) < 6 {
		return errors.New("project_id must be at least 6 characters")
	}
	if len(p) > 30 {
		return errors.New("project_id must be no more then 30 characters")
	}

	runes := []rune(p)
	if runes[0] < 'a' || runes[0] > 'z' {
		return errors.New("project_id must start with a lowercase letter")
	}
	if runes[len(runes)-1] == '-' {
		return errors.New("project_id must not end with a hyphen")
	}
	for _, r := range runes {
		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= '0' && r <= '9' {
			continue
		}
		if r == '-' {
			continue
		}
		return errors.New("project_id may only contain lowercase letters, digits, or hyphens")
	}
	return nil
}

// validateDatasetID validates the dataset id against a set of rules and
// returns an error if there is a violation. As a special case if the dataset
// id is empty no error will be returned since it is valid for the dataset id
// to not be supplied.
//
// Rules:
//   - Must less then 1024 characters.
//   - Must only contain lowercase letters, uppercase letters, digits, or underscores.
func validateDatasetID(d string) error {
	if d == "" {
		return nil
	}
	if len(d) > 1024 {
		return errors.New("dataset_id must be no more then 1024 characters")
	}

	runes := []rune(d)
	for _, r := range runes {
		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= 'A' && r <= 'Z' {
			continue
		}
		if r >= '0' && r <= '9' {
			continue
		}
		if r == '_' {
			continue
		}
		return errors.New("dataset_id may only contain lowercase letters, uppercase letters, digits, or underscores")
	}
	return nil
}

// validateTableID validates the table id against a set of rules and
// returns an error if there is a violation.
//
// Rules:
//   - Must be between 1 and 1024 characters.
//   - Must only contain lowercase letters, uppercase letters, digits, or underscores.
func validateTableID(t string) error {
	if len(t) < 1 {
		return errors.New("table_id must be at least 1 character")
	}
	if len(t) > 1024 {
		return errors.New("table_id must be no more then 1024 characters")
	}

	runes := []rune(t)
	for _, r := range runes {
		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= 'A' && r <= 'Z' {
			continue
		}
		if r >= '0' && r <= '9' {
			continue
		}
		if r == '_' {
			continue
		}
		return errors.New("table_id may only contain lowercase letters, uppercase letters, digits, or underscores")
	}
	return nil
}
