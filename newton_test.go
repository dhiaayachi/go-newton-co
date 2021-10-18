package newton

import (
	"testing"
	"time"
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

	_, err := n.Actions(DEPOSIT, 1, 0, time.Date(2020, 01, 01, 00, 00, 00, 00, time.Local).Unix(), time.Date(2020, 01, 02, 00, 00, 00, 00, time.Local).Unix())

	if err != nil {
		t.Error("test failed: " + err.Error())
	}
}

func TestOrderHistory(t *testing.T) {
	ClientId, ClientSecret := getSecrets()
	n := New(ClientId, ClientSecret)

	_, err := n.OrdersHistory(1, 0, time.Date(2020, 01, 01, 00, 00, 00, 00, time.Local).Unix(), time.Date(2020, 01, 01, 01, 00, 00, 00, time.Local).Unix(), "", "")

	if err != nil {
		t.Error("test failed: " + err.Error())
	}
}

func TestOpenOrders(t *testing.T) {
	ClientId, ClientSecret := getSecrets()
	n := New(ClientId, ClientSecret)

	_, err := n.OpenOrders(1, 0, "", "")

	if err != nil {
		t.Error("test failed: " + err.Error())
	}
}
