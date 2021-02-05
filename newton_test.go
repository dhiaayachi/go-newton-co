package newton

import (
	"testing"
)

func TestBalance(t *testing.T) {

	n := New(ClientId, ClientSecret)

	_, err := n.Balances("BTC")

	if err != nil {
		t.Error("test failed: " + err.Error())
	}

}

func TestAction(t *testing.T) {

	n := New(ClientId, ClientSecret)

	_, err := n.Actions("", 0, 0, 0, 0)

	if err != nil {
		t.Error("test failed: " + err.Error())
	}

}

func TestOrderHistory(t *testing.T) {

	n := New(ClientId, ClientSecret)

	_, err := n.OrdersHistory(0, 0, 0, 0, "", "")

	if err != nil {
		t.Error("test failed: " + err.Error())
	}

}

func TestOpenOrders(t *testing.T) {

	n := New(ClientId, ClientSecret)

	_, err := n.OpenOrders(0, 0, "", "")

	if err != nil {
		t.Error("test failed: " + err.Error())
	}

}

func TestNewOrder(t *testing.T) {

	n := New(ClientId, ClientSecret)

	_, err := n.NewOrder("LIMIT", "IOC", "BUY", "BTC_QCAD", 10000.1, 0.001)

	if err != nil {
		t.Error("test failed: " + err.Error())
	}

}
