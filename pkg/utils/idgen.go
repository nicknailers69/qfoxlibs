package utils

import (
	"encoding/base64"
	"math/big"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

const idLen int = 48

type ID interface {
	Raw() []byte
	String() string
	Base64() string
	Big() *big.Int
	CTBytes() []byte
	SSBytes() []byte
	PubKey() []byte
}

type id struct {
	RawBytes []byte
	CT       []byte
	SS1      []byte
}

func NewID() (ID, error) {

	rnd := NewRand()
	algName := "Kyber512"

	key := oqs.KeyEncapsulation{}
	defer key.Clean()

	err := key.Init(algName, rnd)
	if err != nil {
		return nil, err
	}

	raw, _ := key.GenerateKeyPair()
	ciphertext, sharedSecret, err := key.EncapSecret(raw)
	if err != nil {
		return nil, err
	}

	return &id{RawBytes: raw, CT: ciphertext, SS1: sharedSecret}, nil

}

func (i *id) Raw() []byte {

	return i.RawBytes

}

func (i *id) String() string {

	return i.Big().String()

}

func (i *id) Base64() string {

	return base64.StdEncoding.EncodeToString(i.Big().Bytes())

}

func (i *id) Big() *big.Int {

	b := new(big.Int).SetBytes(i.RawBytes)

	return b

}

func (i *id) CTBytes() []byte {

	return i.CT

}

func (i *id) SSBytes() []byte {

	return i.SS1

}

func (i *id) PubKey() []byte {

	return i.Big().Bytes()

}
