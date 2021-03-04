package newton

import (
	"os"
	"testing"
)

func getSecrets() (string, string) {
	return os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET")
}

func TestBalance(t *testing.T) {
	ClientId, ClientSecret := getSecrets()
	n := New(ClientId, ClientSecret)

	_, err := n.Balances("BTC")

	if err != nil {
		t.Error("test failed: " + err.Error())
	}
}

func TestAction(t *testing.T) {
	ClientId, ClientSecret := getSecrets()
	n := New(ClientId, ClientSecret)

	_, err := n.Actions("", 0, 0, 0, 0)

	if err != nil {
		t.Error("test failed: " + err.Error())
	}
}

func TestOrderHistory(t *testing.T) {
	ClientId, ClientSecret := getSecrets()
	n := New(ClientId, ClientSecret)

	_, err := n.OrdersHistory(0, 0, 0, 0, "", "")

	if err != nil {
		t.Error("test failed: " + err.Error())
	}
}

func TestOpenOrders(t *testing.T) {
	ClientId, ClientSecret := getSecrets()
	n := New(ClientId, ClientSecret)

	_, err := n.OpenOrders(0, 0, "", "")

	if err != nil {
		t.Error("test failed: " + err.Error())
	}
}
