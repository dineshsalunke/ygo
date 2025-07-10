package sync

import (
	"fmt"

	"github.com/dineshsalunke/ygo/pkg/core"
	"github.com/dineshsalunke/ygo/pkg/encoding"
	"github.com/dineshsalunke/ygo/pkg/encoding/binary"
)

const (
	MessageSyncStep1 = 0
	MessageSyncStep2 = 1
	MessageUpdate    = 2
)

func WriteSyncStep1(dec *binary.Decoder, enc *binary.Encoder, doc *core.Doc, txOrigin any) error {
	if err := enc.WriteVarUint(MessageSyncStep1); err != nil {
		return err
	}
	buf, err := encoding.EncodeStateVector(doc)
	if err != nil {
		return err
	}
	return enc.WriteVarUintBytes(buf)
}

func WriteSyncStep2(enc *binary.Encoder, doc *core.Doc, sv []byte) error {
	if err := enc.WriteVarUint(MessageSyncStep2); err != nil {
		return err
	}
	buf, err := encoding.EncodeStateAsUpdate(doc, sv)
	if err != nil {
		return err
	}
	return enc.WriteVarUintBytes(buf)
}

func ReadSyncStep1(dec *binary.Decoder, enc *binary.Encoder, doc *core.Doc) error {
	sv, err := dec.ReadVarUintBytes()
	if err != nil {
		return err
	}
	return WriteSyncStep2(enc, doc, sv)
}

func ReadSyncStep2(dec *binary.Decoder, doc *core.Doc, txOrigin any) error {
	return encoding.ApplyUpdateV1(doc, dec.Bytes(), txOrigin)
}

func ReadUpdate(dec *binary.Decoder, doc *core.Doc, txOrigin any) error {
	return ReadSyncStep2(dec, doc, txOrigin)
}

func WriteUpdate(enc *binary.Encoder, update []byte) error {
	if err := enc.WriteVarUint(MessageUpdate); err != nil {
		return err
	}
	if err := enc.WriteVarUintBytes(update); err != nil {
		return err
	}
	return nil
}

func ReadSyncMessage(dec *binary.Decoder, enc *binary.Encoder, doc *core.Doc, txOrigin any) error {
	messageType, err := dec.ReadVarUint()
	if err != nil {
		return err
	}

	switch messageType {
	case MessageSyncStep1:
		return ReadSyncStep1(dec, enc, doc)
	case MessageSyncStep2:
		return ReadSyncStep2(dec, doc, txOrigin)
	case MessageUpdate:
		return ReadUpdate(dec, doc, txOrigin)
	default:
		return fmt.Errorf("unknow message type %v", messageType)
	}
}
