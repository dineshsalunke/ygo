package ygo

type UpdateDecoderV1 struct {
	*BinaryDecoder
}

func newUpdateDecoderV1(buf []byte) *UpdateDecoderV1 {
	return &UpdateDecoderV1{
		BinaryDecoder: newDecoder(buf),
	}
}

func NewUpdateDecoderV1(buf []byte) *UpdateDecoderV1 {
	return newUpdateDecoderV1(buf)
}

func (dec *UpdateDecoderV1) ResetDsCurrVal() {
	// This is a noop
}

func (dec *UpdateDecoderV1) ReadDsClock() (uint64, error) {
	b, err := dec.ReadVarUint()
	return b, err
}

func (dec *UpdateDecoderV1) ReadDsLength() (uint64, error) {
	b, err := dec.ReadVarUint()
	return b, err
}

func (dec *UpdateDecoderV1) ReadLeftID() (*ID, error) {
	client, err := dec.ReadVarUint()
	if err != nil {
		return nil, err
	}
	clock, err := dec.ReadVarUint()
	if err != nil {
		return nil, err
	}
	return newID(client, clock), nil
}

func (dec *UpdateDecoderV1) ReadRightID() (*ID, error) {
	client, err := dec.ReadVarUint()
	if err != nil {
		return nil, err
	}
	clock, err := dec.ReadVarUint()
	if err != nil {
		return nil, err
	}
	return newID(client, clock), nil
}

func (dec *UpdateDecoderV1) ReadClient() (uint64, error) {
	client_id, err := dec.ReadVarUint()
	return client_id, err
}

func (dec *UpdateDecoderV1) ReadInfo() (byte, error) {
	return dec.ReadByte()
}

func (dec *UpdateDecoderV1) ReadParentInfo() (bool, error) {
	client, err := dec.ReadVarUint()
	return client == 1, err
}

func (dec *UpdateDecoderV1) ReadTypeRef() (byte, error) {
	return dec.ReadByte()
}

func (dec *UpdateDecoderV1) ReadLength() (uint64, error) {
	b, err := dec.ReadVarUint()
	return b, err
}

func (dec *UpdateDecoderV1) ReadAny() (any, error) {
	return dec.BinaryDecoder.ReadAny()
}

func (dec *UpdateDecoderV1) ReadJson() (any, error) {
	return dec.BinaryDecoder.ReadAny()
}

func (dec *UpdateDecoderV1) ReadKey() (string, error) {
	return dec.ReadVarString()
}
