package types

import "github.com/ethereum/go-ethereum/common"

type UToC struct {
	DN       []byte         `json:"dn"        gencodec:"required"`
	ET       []byte         `json:"et"        gencodec:"required"`
	CAS      []*common.Address	`json:"cas"        gencodec:"required"`
}

func (u *UToC) Dn() []byte            { return common.CopyBytes(u.DN) }
func (u *UToC) Et() []byte            { return common.CopyBytes(u.ET) }



