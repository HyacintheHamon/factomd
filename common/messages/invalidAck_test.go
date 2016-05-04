// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages_test

import (
	"testing"

	"github.com/FactomProject/factomd/common/constants"
	. "github.com/FactomProject/factomd/common/messages"
	"github.com/FactomProject/factomd/common/primitives"
)

func TestMarshalUnmarshalInvalidAck(t *testing.T) {
	msg := newInvalidAck()

	hex, err := msg.MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Marshalled - %x", hex)

	msg2, err := UnmarshalMessage(hex)
	if err != nil {
		t.Error(err)
	}
	str := msg2.String()
	t.Logf("str - %v", str)

	if msg2.Type() != constants.INVALID_ACK_MSG {
		t.Error("Invalid message type unmarshalled")
	}

	hex2, err := msg2.(*InvalidAck).MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	if len(hex) != len(hex2) {
		t.Error("Hexes aren't of identical length")
	}
	for i := range hex {
		if hex[i] != hex2[i] {
			t.Error("Hexes do not match")
		}
	}

	if msg.IsSameAs(msg2.(*InvalidAck)) != true {
		t.Errorf("InvalidAck messages are not identical")
	}
}

func TestSignAndVerifyInvalidAck(t *testing.T) {
	msg := newSignedInvalidAck()
	hex, err := msg.MarshalBinary()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Marshalled - %x", hex)

	t.Logf("Sig - %x", *msg.Signature.GetSignature())
	if len(*msg.Signature.GetSignature()) == 0 {
		t.Error("Signature not present")
	}

	valid, err := msg.VerifySignature()
	if err != nil {
		t.Error(err)
	}
	if valid == false {
		t.Error("Signature is not valid")
	}

	msg2, err := UnmarshalMessage(hex)
	if err != nil {
		t.Error(err)
	}

	if msg2.Type() != constants.INVALID_ACK_MSG {
		t.Error("Invalid message type unmarshalled")
	}
	eomProper := msg2.(*InvalidAck)

	valid, err = eomProper.VerifySignature()
	if err != nil {
		t.Error(err)
	}
	if valid == false {
		t.Error("Signature 2 is not valid")
	}

}

func newInvalidAck() *InvalidAck {
	msg := new(InvalidAck)
	msg.Timestamp.SetTimeNow()

	return msg
}

func newSignedInvalidAck() *InvalidAck {
	ack := newInvalidAck()

	key, err := primitives.NewPrivateKeyFromHex("07c0d52cb74f4ca3106d80c4a70488426886bccc6ebc10c6bafb37bf8a65f4c38cee85c62a9e48039d4ac294da97943c2001be1539809ea5f54721f0c5477a0a")
	if err != nil {
		panic(err)
	}
	err = ack.Sign(&key)
	if err != nil {
		panic(err)
	}

	return ack
}