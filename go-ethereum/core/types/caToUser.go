package types

import "github.com/ethereum/go-ethereum/common"

type CToU struct {
	Signatures []byte 		`json:"sig"        gencodec:"required"`
}

func (c *CToU) Sig() []byte            { return common.CopyBytes(c.Signatures) }