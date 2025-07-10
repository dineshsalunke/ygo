package encoding

import (
	"github.com/dineshsalunke/ygo/pkg/core"
	"github.com/dineshsalunke/ygo/pkg/encoding/binary"
)

func applyUpdate(decoder UpdateDecoder, doc *core.Doc, update []byte, txOrigin any) error {
	panic("not implemented")
}

func ApplyUpdateV1(doc *core.Doc, update []byte, txorigin any) error {
	dec := &UpdateDecoderV1{}
	return applyUpdate(dec, doc, update, txorigin)
}

func ApplyUpdateV2(doc *core.Doc, update []byte, txorigin any) error {
	dec := &UpdateDecoderV2{}
	return applyUpdate(dec, doc, update, txorigin)
}

func encodeStateVector(doc *core.Doc, dsencoder DsEncoder) ([]byte, error) {
	panic("not implemented")
}

func EncodeStateVectorV1(doc *core.Doc) ([]byte, error) {
	encoder := &DsEncoderV1{}
	return encodeStateVector(doc, encoder)
}

func EncodeStateVectorV2(doc *core.Doc) ([]byte, error) {
	encoder := &DsEncoderV2{}
	return encodeStateVector(doc, encoder)
}

func encodeStateAsUpdate(enc UpdateEncoder, doc *core.Doc, encodedtargetStateVector []byte) ([]byte, error) {
	panic("not implemented")
}

func EncodeStateAsUpdateV1(doc *core.Doc, targetStateVector []byte) ([]byte, error) {
	enc := &UpdateEncoderV1{}
	return encodeStateAsUpdate(enc, doc, targetStateVector)
}

func EncodeStateAsUpdateV2(doc *core.Doc, targetStateVector []byte) ([]byte, error) {
	enc := &UpdateEncoderV2{}
	return encodeStateAsUpdate(enc, doc, targetStateVector)
}

func EncodeStateAsUpdate(doc *core.Doc, targetStateVector []byte) ([]byte, error) {
	return EncodeStateAsUpdateV2(doc, targetStateVector)
}

func readStateVector(dec *binary.Decoder) (map[uint64]uint64, error) {
	length, err := dec.ReadVarUint()
	if err != nil {
		return nil, err
	}
	ss := make(map[uint64]uint64, length)

	for range length {
		client, err := dec.ReadVarUint()
		if err != nil {
			return nil, err
		}
		clock, err := dec.ReadVarUint()
		if err != nil {
			return nil, err
		}
		ss[client] = clock
	}

	return ss, nil
}

func decodeStateVector(buff []byte) (map[uint64]uint64, error) {
	return readStateVector(binary.NewDecoder(buff))
}
