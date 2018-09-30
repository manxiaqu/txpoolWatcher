package main

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	c, err := DialDefault()
	if err != nil {
		fmt.Print(err)
	}

	txpoolStatus, err := c.TxPoolContent()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(txpoolStatus)
}
