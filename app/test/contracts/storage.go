package contracts

import (
	"github.com/lankaiyun/kaiyunchain/contractapi"
)

type StorageContract struct {
	contractapi.Contract
}

func (s *StorageContract) SetValue(k, v []byte) {
	s.Contract.SetValue(k, v)
}

func (s *StorageContract) GetValue(k []byte) []byte {
	v, _ := s.Contract.GetValue(k)
	return v
}
