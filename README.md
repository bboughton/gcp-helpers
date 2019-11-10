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

## gsurl

This package may be used to parse Google Cloud Storage URL strings
into an URL struct. The primary use case for this package is to all
GCS URL to be passed into a program as a single configuration.

```go
func main() {
	u, err := gsurl.Parse(os.Getenv("GCS_URL"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bucket:", u.Bucket)
	fmt.Println("Object:", u.Object)
}
```

## gsutil

Package gsutil provides utility helpers for working with Google Cloud
Storage. These helpers are meant to simplify some common use cases. To
keep these helpers simple authentication credentials are retrieved
from the standard credential resolution chain used by the storage SDK
'cloud.google.com/go/storage'.

```go
func main() {
	b, err := gsutil.ReadFile("gs://bucket-name/path/to/object")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
```
