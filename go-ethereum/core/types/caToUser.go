package types

import "github.com/ethereum/go-ethereum/common"

type CToU struct {
	DN       []byte         `json:"dn"        gencodec:"required"`
	ET       []byte         `json:"et"        gencodec:"required"`
	CA	 *common.Address	`json:"cas"        gencodec:"required"`
	Signatures []byte 		`json:"sig"        gencodec:"required"`
}

func (c *CToU) Sig() []byte            { return common.CopyBytes(c.Signatures) }