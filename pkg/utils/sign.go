package utils

import (
	"encoding/json"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

type Signature []byte
type SigPubKey []byte

type message []byte

type SignedMessage struct {
	msg message
	sig Signature
	pub SigPubKey
	Raw []byte
}

func (m message) Sign() ([]byte, error) {

	sigName := "Falcon-512"
	signer := oqs.Signature{}
	defer signer.Clean()

	if err := signer.Init(sigName, nil); err != nil {
		return nil, err
	}
	pubKey, _ := signer.GenerateKeyPair()
	pk := SigPubKey(pubKey)
	sig, _ := signer.Sign(m)
	sm := &SignedMessage{
		m,
		Signature(sig),
		pk, nil,
	}
	sm.Raw, _ = json.Marshal(sm)
	return sm.Raw, nil
}

func (sm SignedMessage) Verify(raw []byte) error {

	var smsg SignedMessage
	json.Unmarshal(raw, &smsg)
	v := oqs.Signature{}
	defer v.Clean()
	if err := v.Init("Falcon-512", nil); err != nil {
		return err
	}
	_, err := v.Verify(smsg.msg, smsg.sig, smsg.pub)
	if err != nil {
		return err
	}
	return nil

}
