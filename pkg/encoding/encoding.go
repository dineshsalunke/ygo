package encoding

import (
	"slices"

	"github.com/dineshsalunke/ygo/pkg/core"
	"github.com/dineshsalunke/ygo/pkg/encoding/binary"
)

func writeStateVector(enc DsEncoder, sv map[uint64]uint64) error {
	if err := enc.Encoder().WriteVarUint(uint64(len(sv))); err != nil {
		return err
	}
	clients := make([]uint64, len(sv))
	for client := range sv {
		clients = append(clients, client)
	}
	slices.Sort(clients)
	slices.Reverse(clients)

	for _, client := range clients {
		clock := sv[client]
		if err := enc.Encoder().WriteVarUint(client); err != nil {
			return err
		}
		if err := enc.Encoder().WriteVarUint(clock); err != nil {
			return err
		}
	}
	return nil
}

func writeDocumentStateVector(enc DsEncoder, doc *core.Doc) error {
	return writeStateVector(enc, doc.Store().GetStateVector())
}

func EncodeStateVectorV1(doc *core.Doc) ([]byte, error) {
	encoder := &DsEncoderV1{
		encoder: binary.NewEncoder([]byte{}),
	}
	if err := writeDocumentStateVector(encoder, doc); err != nil {
		return nil, err
	}
	return encoder.Bytes(), nil
}

func EncodeStateVector(doc *core.Doc) ([]byte, error) {
	return EncodeStateVectorV1(doc)
}

func EncodeStateAsUpdateV1(doc *core.Doc, encodedTargetStateVector []byte) ([]byte, error) {
	panic("not implemented")
}

func EncodeStateAsUpdate(doc *core.Doc, encodedTargetStateVector []byte) ([]byte, error) {
	return EncodeStateAsUpdateV1(doc, encodedTargetStateVector)
}

func ApplyUpdateV1(doc *core.Doc, update []byte, txorigin any) error {
	panic("not implemented")
}

func ApplyUpdate(doc *core.Doc, update []byte, txorigin any) error {
	return ApplyUpdateV1(doc, update, txorigin)
}
