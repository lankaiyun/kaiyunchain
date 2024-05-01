package contractapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lankaiyun/kaiyunchain/rpc/pb"
	"github.com/mitchellh/mapstructure"
	"math/big"
	"reflect"

	"github.com/lankaiyun/kaiyunchain/common"
)

type ContractInterface interface {
	IAmContract()
}

type Contract struct {
	BelongBlock     *big.Int
	Deployer        common.Address
	Hash            common.Hash
	Caller          string
	CallerPass      string
	ContractAddress string
}

func (c Contract) IAmContract() {}

func (c Contract) SetValue(k, v []byte) error {
	rpcClient := GetRpcClient()
	r, _ := rpcClient.Set(context.Background(), &pb.SetReq{Key: k, Value: v, Account: c.Caller, Password: c.CallerPass, ContractAddress: c.ContractAddress})
	if r.Result == "1" {
		return nil
	} else {
		return fmt.Errorf("%s", r.Result)
	}
}

func (c Contract) GetValue(k []byte) ([]byte, error) {
	rpcClient := GetRpcClient()
	r, _ := rpcClient.Get(context.Background(), &pb.GetReq{Key: k, ContractAddress: c.ContractAddress})
	if r.Result == "1" {
		return r.Value, nil
	} else {
		return nil, fmt.Errorf("%s", r.Result)
	}
}

func (c Contract) GetContractInfo() ContractInfo {
	return ContractInfo{}
}

func (c Contract) GetGlobalInfo() GlobalInfo {
	return GlobalInfo{}
}

func Deploy(contract interface{}, account, password string) (string, error) {
	contractBs, err := json.Marshal(contract)
	if err != nil {
		panic(err)
	}
	rpcClient := GetRpcClient()
	r, _ := rpcClient.Deploy(context.Background(), &pb.DeployReq{Contract: contractBs, Account: account, Password: password})
	if r.Result == "1" {
		return r.ContractAddress, nil
	} else {
		return "", fmt.Errorf("%s", r.Result)
	}
}

func GetContract(contract ContractInterface, contractAddress string) error {
	rpcClient := GetRpcClient()
	r, _ := rpcClient.GetContract(context.Background(), &pb.GetContractReq{ContractAddress: contractAddress})
	if r.Result == "1" {
		var intf interface{}
		err := json.Unmarshal(r.Contract, &intf)
		if err != nil {
			panic(err)
		}

		err = mapstructure.Decode(intf, contract)
		if err != nil {
			panic(err)
		}
		return nil
	} else {
		return fmt.Errorf("%s", r.Result)
	}
}

func Save(contract ContractInterface, contractAddress, account, password string) error {
	rpcClient := GetRpcClient()
	r, _ := rpcClient.GetContract(context.Background(), &pb.GetContractReq{ContractAddress: contractAddress})
	var temp *ContractInterface
	var intf interface{}
	err := json.Unmarshal(r.Contract, &intf)
	if err != nil {
		panic(err)
	}

	err = mapstructure.Decode(intf, temp)
	if err != nil {
		panic(err)
	}
	if temp == &contract {
		contractBs, err := json.Marshal(contract)
		if err != nil {
			panic(err)
		}
		rpcClient.Call(context.Background(), &pb.CallReq{ContractAddress: contractAddress, Contract: contractBs, Account: account, Password: password})
		if r.Result == "1" {
			return nil
		} else {
			return fmt.Errorf("%s", r.Result)
		}
	} else {
		return errors.New("contract something wrong")
	}
	return nil
}

func Call(obj interface{}, name string, args []interface{}) (interface{}, error) {
	objValue := reflect.ValueOf(obj)
	method := objValue.MethodByName(name)

	if !method.IsValid() {
		return nil, fmt.Errorf("Method '%s' not found", name)
	}

	methodType := method.Type()
	numIn := methodType.NumIn()
	if len(args) != numIn {
		return nil, fmt.Errorf("Incorrect number of arguments. Expected: %d, Got: %d", numIn, len(args))
	}

	for i := 0; i < numIn; i++ {
		argType := methodType.In(i)
		if reflect.TypeOf(args[i]).Kind() != argType.Kind() {
			return nil, fmt.Errorf("Argument at position %d has incorrect type. Expected: %s, Got: %s", i, argType.Kind(), reflect.TypeOf(args[i]).Kind())
		}
	}

	var input []reflect.Value
	for _, arg := range args {
		input = append(input, reflect.ValueOf(arg))
	}

	result := method.Call(input)
	if len(result) > 0 {
		return result[0].Interface(), nil
	}
	return nil, nil
}
