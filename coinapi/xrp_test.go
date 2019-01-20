package coinapi

import (
	"testing"
)

func TestGetBalance(t *testing.T) {
	data := GetBalance("rKu9QEsBax7h4fv2qJnfXPVyqL98rMqBvG")
	t.Log(data)
}

func TestNewAddress(t *testing.T) {
	data := NewAddress("rKu9QEsBax7h4fv2qJnfXPVyqL98rMqBvG")
	t.Log(data)
}
