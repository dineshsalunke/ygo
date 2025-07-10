package encoding

import "github.com/dineshsalunke/ygo/pkg/core"

func EncodeStateVectorV1(doc *core.Doc) ([]byte, error) {
	panic("not implemented")
}

func EncodeStateVector(doc *core.Doc) ([]byte, error) {
	return EncodeStateVectorV1(doc)
}

func EncodeStateAsUpdateV2(doc *core.Doc, encodedTargetStateVector []byte) ([]byte, error) {
	panic("not implemented")
}

func EncodeStateAsUpdate(doc *core.Doc, encodedTargetStateVector []byte) ([]byte, error) {
	return EncodeStateAsUpdateV2(doc, encodedTargetStateVector)
}

func ApplyUpdateV1(doc *core.Doc, update []byte, txorigin any) error {
	panic("not implemented")
}

func ApplyUpdate(doc *core.Doc, update []byte, txorigin any) error {
	return ApplyUpdateV1(doc, update, txorigin)
}
