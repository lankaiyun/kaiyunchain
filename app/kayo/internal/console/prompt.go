package console

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func MainPrompt(acc string) int {
	color.Blue("Welcome to the Kayo console!")
	fmt.Printf("Account: %s\n", acc)
	return ScanChoose()
}

func ScanChoose() int {
	arr := []string{
		"Quit: leave console",
		"Balance: check your account balance",
		"Transaction: conduct a transfer transaction",
		"TxPool: look at the transactions in the pool",
		"Mine: start mining",
		"Blockchain: look at block info",
		"Mpt: look at state",
	}
	prompt := promptui.Select{
		Label: "Select Function",
		Items: arr,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return -1
	}

	for i := 0; i < len(arr); i++ {
		if result == arr[i] {
			return i
		}
	}
	return -1
}
