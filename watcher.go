package main

import "fmt"

func main() {
	c, err := DialDefault()
	if err != nil {
		fmt.Print(err)
	}

	txpoolStatus, err := c.TxPoolInspect()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(txpoolStatus)
}
