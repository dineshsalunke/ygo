package ygo

type UpdateDecoderV1 struct {
	*BinaryDecoder
}

func (dec *UpdateDecoderV1) ResetDsCurrVal() {
	// This is a noop
}

func (dec *UpdateDecoderV1) ReadDsClock() (uint32, error) {
	b, err := dec.ReadVarUint()
	return uint32(b), err
}

func (dec *UpdateDecoderV1) ReadDsLength() (uint32, error) {
	b, err := dec.ReadVarUint()
	return uint32(b), err
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
	return newID(client, uint32(clock)), nil
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
	return newID(client, uint32(clock)), nil
}

func (dec *UpdateDecoderV1) ReadClient() (ClientID, error) {
	return dec.ReadVarUint()
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

func (dec *UpdateDecoderV1) ReadLength() (uint32, error) {
	b, err := dec.ReadVarUint()
	return uint32(b), err
}

func (dec *UpdateDecoderV1) ReadAny() (any, error) {
	return dec.ReadAny()
}

func (dec *UpdateDecoderV1) ReadJson() (any, error) {
	return dec.ReadAny()
}

func (dec *UpdateDecoderV1) ReadKey() (string, error) {
	return dec.ReadVarString()
}
