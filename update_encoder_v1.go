package ygo

type IdSetEncoderV1 struct {
	*BinaryEncoder
}

func (enc *IdSetEncoderV1) ResetDsCurrVal() {
	panic("not implemented")
}

func (enc *IdSetEncoderV1) WriteDsClock(length uint64) error {
	return enc.WriteVarUint(length)
}

func (enc *IdSetEncoderV1) WriteDsLength(length uint64) error {
	return enc.WriteVarUint(length)
}

func newIdSetEncoderV1() *IdSetEncoderV1 {
	return &IdSetEncoderV1{
		BinaryEncoder: newEncoder(),
	}
}

type UpdateEncoderV1 struct {
	*IdSetEncoderV1
}

func newUpdateEncoderV1() *UpdateEncoderV1 {
	return &UpdateEncoderV1{
		IdSetEncoderV1: newIdSetEncoderV1(),
	}
}

func (enc *UpdateEncoderV1) WriteLeftID(id *ID) error {
	if err := enc.WriteVarUint(id.clock); err != nil {
		return err
	}
	if err := enc.WriteVarUint(id.client); err != nil {
		return err
	}
	return nil
}

func (enc *UpdateEncoderV1) WriteRightID(id *ID) error {
	if err := enc.WriteVarUint(id.clock); err != nil {
		return err
	}
	if err := enc.WriteVarUint(id.client); err != nil {
		return err
	}
	return nil
}

func (enc *UpdateEncoderV1) WriteClient(clientId uint64) error {
	return enc.WriteVarUint(clientId)
}

func (enc *UpdateEncoderV1) WriteInfo(info byte) error {
	return enc.WriteByte(info)
}

func (enc *UpdateEncoderV1) WriteParentInfo(hasParentInfo bool) error {
	b := uint64(0)
	if hasParentInfo {
		b = 1
	}
	return enc.WriteVarUint(b)
}

func (enc *UpdateEncoderV1) WriteTypeRef(ref byte) error {
	return enc.WriteByte(ref)
}

func (enc *UpdateEncoderV1) WriteLength(length uint64) error {
	return enc.WriteVarUint(length)
}

func (enc *UpdateEncoderV1) WriteAny(val any) error {
	return enc.BinaryEncoder.WriteAny(val)
}

func (enc *UpdateEncoderV1) WriteJson(val any) error {
	return enc.BinaryEncoder.WriteAny(val)
}

func (enc *UpdateEncoderV1) WriteKey(key string) error {
	return enc.WriteVarString(key)
}

func (enc *UpdateEncoderV1) Bytes() []byte {
	return enc.BinaryEncoder.Bytes()
}
