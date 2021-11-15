# go-newton-co
go api for http://newton.co/ cryptocurrency exchance 


## Usage:

For every API enpoint, there is an appropriate struct in the `query` package
that has all body and query parameters as part of the struct's fields.

Once the appropriate query is selected, use `Do` function of the `newton` package with
the query provided.

Example of calling Newton's `/balances` endpoint:

```go
newton := New(<ClientId>, <ClientSecret>)

q := &query.Balances{Asset: "BTC"}
response, err := newton.Do(q)
  
```

For more examples, please reference `newton_test.go`. 

## Testing

Newton's mock endpoint does not take authentication headers into account. To test the API with authentication enabled,
take a look at `auth_test.sh` where the following environment variables are set before running `go test`:
- TEST_AUTH - should be set to "true"
- CLIENT_ID - your real client id from Newton
- CLIENT_SECRET - your real client secret from Newton

Running the script with the correct client id and client secret, will make
a request to Newton's production balances endpoint to ensure that authentication is functioning correctly.
On the other hand, running `go test` without these environment variables set will test every endpoint of the mock server.
