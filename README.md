# Go Client for jAccount

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/dyweb/go-jaccount)
[![License](https://img.shields.io/github/license/dyweb/go-jaccount)](https://github.com/dyweb/go-jaccount/blob/master/LICENSE)

go-jaccount is a Go Client for jAccount API.

## Installation

```shell
go get github.com/dyweb/go-jaccount
```

## Example

```go
// OAuth 2.0 configuration
var config = &oauth2.Config{
    ClientID:     os.Getenv("clientid"),
    ClientSecret: os.Getenv("secretkey"),
    Endpoint:     jaccount.Endpoint,
    RedirectURL:  "http://localhost:8000/callback",
    Scopes:       []string{"essential"},
}

var client *jaccount.Client

c := config.Client(oauth2.NoContext, token)

// jAccount API client
client = jaccount.NewClient(c)

// Get the profile of the user
profile, err := client.Profile.Get(context.Background())
```

## References

- [google/go-github](https://github.com/google/go-github)

## License

MIT
