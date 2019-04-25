// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*txdataMarshaling)(nil)

func (t txdata) MarshalJSON() ([]byte, error) {
	type txdata struct {
		Code         uint8           `json:"code"     gencodec:"required"` // differ txs --Agzs 09.17

		DN string				`json:"dn"     gencodec:"required"`
	    ET string					`json:"et"     gencodec:"required"`
		CAS string      `json:"cas"     gencodec:"required"`
		Signatures string		`json:"sig"     gencodec:"required"`

		AccountNonce hexutil.Uint64  `json:"nonce"    gencodec:"required"`
		Price        *hexutil.Big    `json:"gasPrice" gencodec:"required"`
		GasLimit     hexutil.Uint64  `json:"gas"      gencodec:"required"`
		Recipient    *common.Address `json:"to"       rlp:"nil"`
		Amount       *hexutil.Big    `json:"value"    gencodec:"required"`
		Payload      hexutil.Bytes   `json:"input"    gencodec:"required"`
		V            *hexutil.Big    `json:"v" gencodec:"required"`
		R            *hexutil.Big    `json:"r" gencodec:"required"`
		S            *hexutil.Big    `json:"s" gencodec:"required"`
		Hash         *common.Hash    `json:"hash" rlp:"-"`
	}
	var enc txdata

	enc.DN = t.DN
	enc.ET = t.ET
	enc.CAS = t.CAS
	enc.Signatures = t.Signatures
	enc.AccountNonce = hexutil.Uint64(t.AccountNonce)
	enc.Price = (*hexutil.Big)(t.Price)
	enc.GasLimit = hexutil.Uint64(t.GasLimit)
	enc.Recipient = t.Recipient
	enc.Amount = (*hexutil.Big)(t.Amount)
	enc.Payload = t.Payload
	enc.V = (*hexutil.Big)(t.V)
	enc.R = (*hexutil.Big)(t.R)
	enc.S = (*hexutil.Big)(t.S)
	enc.Hash = t.Hash
	return json.Marshal(&enc)
}

func (t *txdata) UnmarshalJSON(input []byte) error {
	type txdata struct {
		Code         uint8           `json:"code"     gencodec:"required"` // differ txs --Agzs 09.18
		DN string					`json:"dn"     gencodec:"required"`
	    ET string					`json:"et"     gencodec:"required"`
		CAS string     `json:"cas"     gencodec:"required"`
		Signatures string		`json:"sig"     gencodec:"required"`
		

		AccountNonce *hexutil.Uint64 `json:"nonce"    gencodec:"required"`
		Price        *hexutil.Big    `json:"gasPrice" gencodec:"required"`
		GasLimit     *hexutil.Uint64 `json:"gas"      gencodec:"required"`
		Recipient    *common.Address `json:"to"       rlp:"nil"`
		Amount       *hexutil.Big    `json:"value"    gencodec:"required"`
		Payload      *hexutil.Bytes  `json:"input"    gencodec:"required"`
		V            *hexutil.Big    `json:"v" gencodec:"required"`
		R            *hexutil.Big    `json:"r" gencodec:"required"`
		S            *hexutil.Big    `json:"s" gencodec:"required"`
		Hash         *common.Hash    `json:"hash" rlp:"-"`
	}
	var dec txdata
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.AccountNonce == nil {
		return errors.New("missing required field 'nonce' for txdata")
	}
	t.AccountNonce = uint64(*dec.AccountNonce)
	if dec.Price == nil {
		return errors.New("missing required field 'gasPrice' for txdata")
	}
	t.Price = (*big.Int)(dec.Price)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gas' for txdata")
	}
	t.GasLimit = uint64(*dec.GasLimit)
	if dec.Recipient != nil {
		t.Recipient = dec.Recipient
	}
	if dec.Amount == nil {
		return errors.New("missing required field 'value' for txdata")
	}
	t.Amount = (*big.Int)(dec.Amount)
	if dec.Payload == nil {
		return errors.New("missing required field 'input' for txdata")
	}
	t.Payload = *dec.Payload

	if len(dec.DN) == 0{
		return errors.New("missing required field 'DN' for txdata")
	}
	t.DN = dec.DN
	if len(dec.ET) == 0{
		return errors.New("missing required field 'ET' for txdata")
	}
	t.ET = dec.ET
	if len(dec.CAS) == 0{
		return errors.New("missing required field 'CAS' for txdata")
	}
	t.CAS = dec.CAS
	if len(dec.Signatures) == 0{
		return errors.New("missing required field 'Signatures' for txdata")
	}
	t.Signatures= dec.Signatures


	if dec.V == nil {
		return errors.New("missing required field 'v' for txdata")
	}
	t.V = (*big.Int)(dec.V)
	if dec.R == nil {
		return errors.New("missing required field 'r' for txdata")
	}
	t.R = (*big.Int)(dec.R)
	if dec.S == nil {
		return errors.New("missing required field 's' for txdata")
	}
	t.S = (*big.Int)(dec.S)
	if dec.Hash != nil {
		t.Hash = dec.Hash
	}
	
	

	return nil
}
