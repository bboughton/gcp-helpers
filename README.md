# GCP Helpers

This repository contains Google Cloud Platform (GCP) helper packages
for Go.

## bqurl

This package may be used to parse a BigQuery URL string into an URL
struct. The primary use case for thie package is to allow a BigQuery
table URL to be passed into a program as a single configuration.

```go
func main() {
	u, err := bqurl.Parse(os.Getenv("BQ_URL"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Project ID:", u.ProjectID)
	fmt.Println("Dataset ID:", u.DatasetID)
	fmt.Println("Table ID:", u.TableID)
}
```
