package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	w := NewWallet()
	fmt.Println(w.Address.Hex())
}
