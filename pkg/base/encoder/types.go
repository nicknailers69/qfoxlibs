package encoder

import (
	"math/big"
	"reflect"
)

type Uint *big.Int
type Uint256 *big.Int
type Int256 *big.Int
type BigNumber *big.Int
type Hex []byte
type String string

type BaseTypes interface {
	ToString() string
	ToHex() Hex
	ToHexString() string
	ToBigNumber() *BigNumber
	FromBigNumber() interface{}
	FromRawBytes() interface{}
}

type baseTypes struct {
	BaseTypes
	topic interface{}
}

func (b *baseTypes) typeOfTopic() string {
	return reflect.TypeOf(b.topic).String()
}

func (b *baseTypes) ToString() string {
	if b.typeOfTopic() == "String" {
		return string(b.topic.(String))
	}
	if b.typeOfTopic() == "string" {
		return string(b.topic.(String))
	}

	return ""

}
