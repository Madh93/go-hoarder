# go-hoarder

[![Go Version](https://img.shields.io/badge/Go-1.22%2B-blue)](https://go.dev/doc/install)
[![Latest tag](https://img.shields.io/github/v/tag/Madh93/go-hoarder?label=go%20module)](https://github.com/Madh93/go-hoarder/tags)
[![Go Reference](https://pkg.go.dev/badge/github.com/Madh93/go-hoarder.svg)](https://pkg.go.dev/github.com/Madh93/go-hoarder)
[![License](https://img.shields.io/badge/License-MIT-brightgreen)](LICENSE)

[Go](https://go.dev/) client library for [Hoarder](https://hoarder.app).

The `go-hoarder` client is auto-generated using the
[`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen) tool, which allows convert [Hoarder OpenAPI](https://github.com/hoarder-app/hoarder/blob/v0.20.0/packages/open-api/hoarder-openapi-spec.json) specification to Go code.

## Requirements

- [Go 1.22](https://golang.org/dl/) or higher.
- A valid API key from [Hoarder](https://docs.hoarder.app/screenshots#settings).

## Installation

To install `go-hoarder`, use `go get`:

```sh
go get github.com/Madh93/go-hoarder
```

## Usage

Here is a basic example of how to use the `go-hoarder` library:

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/Madh93/go-hoarder"
)

func main() {
    // Basic configuration
    apiUrl := "https://<YOUR_HOARDER_HOSTNAME>/api/v1" // Replace this with your API URL
    apiKey := "<YOUR_HOARDER_API_KEY>"                 // Replace this with your actual token

    // Set up Bearer authentication
    auth := func(ctx context.Context, req *http.Request) error {
        req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
        return nil
    }

    // Create the Hoarder client
    client, err := hoarder.NewClient(apiUrl, hoarder.WithRequestEditorFn(auth))
    if err != nil {
        log.Fatalf("Error creating client: %v", err)
    }

    log.Printf("Hello world from %s", client.Server)
}
```

For more code examples, check out the [examples](examples) directory.

## Documentation

For detailed usage and API documentation, refer to the [GoDoc](https://pkg.go.dev/github.com/Madh93/go-hoarder).

Additionally, it's recommended to check out the latest [Hoarder API documentation](https://docs.hoarder.app/api) for more information.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any bug fixes or enhancements.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a Pull Request.

## License

This project is licensed under the [MIT license](LICENSE).
