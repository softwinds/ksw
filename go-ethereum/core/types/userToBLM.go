package types

import "github.com/ethereum/go-ethereum/common"

type UToBLM struct {
	DN       []byte         `json:"dn"        gencodec:"required"`
	ET       []byte         `json:"et"        gencodec:"required"`
	SigMap   map[common.Address][]byte	`json:"sigMap"        gencodec:"required"`
}

func (u *UToBLM) Dn() []byte            { return common.CopyBytes(u.DN) }
func (u *UToBLM) Et() []byte            { return common.CopyBytes(u.ET) }
