package main

import (
	"fmt"
	"github.com/lankaiyun/kaiyunchain/app/test/contracts"
	"github.com/lankaiyun/kaiyunchain/contractapi"
)

func Deploy(contract contractapi.ContractInterface, account string, password string) (string, error) {
	return contractapi.Deploy(contract, account, password)
}

type Hello struct {
	contractapi.Contract
}

func main() {
	//contractAddr, err := Deploy(&contracts.StorageContract{}, "0x760Bb7e3ae87f5a030560E6512435Eee48588800", "pass1234")
	//if err != nil {
	//	fmt.Println("err-->", err)
	//} else {
	//	fmt.Println("--->", contractAddr)
	//}
	contract := &contracts.StorageContract{}
	err := contractapi.GetContract(contract, "0x918d9E3CaA3606234482db3EA34656830b7839De")
	if err != nil {
		fmt.Println("err-->", err)
	}
	contract.Caller = "0x760Bb7e3ae87f5a030560E6512435Eee48588800"
	contract.CallerPass = "pass1234"
	contract.ContractAddress = "0x918d9E3CaA3606234482db3EA34656830b7839De"
	// contract.SetValue([]byte("hello"), []byte("world"))
	fmt.Println(string(contract.GetValue([]byte("hello"))))
}
