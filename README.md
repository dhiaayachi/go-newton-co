# go-newton-co
go api for http://newton.co/ bitcoin exchance 


To create a client:

```go
n := New(<ClientId>, <ClientSecret>)

	_, err := n.Balances("BTC")

	if err != nil {
		t.Error("test failed: " + err.Error())
	}
  
```

