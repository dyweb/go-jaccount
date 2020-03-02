# Go Client for jAccount

A Go Client for jAccount.

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

oauthClient := config.Client(oauth2.NoContext, token)

// jAccount API client
client = jaccount.NewClient(oauthClient)

// Get the profile of the user
profile, err := client.GetProfile(context.Background())
```

## License

MIT
