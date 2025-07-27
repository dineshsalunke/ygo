package ygo

type UpdateBlocks struct {
	clients map[ClientID][]Block
}
type Update struct {
	blocks     *UpdateBlocks
	delete_set IdSet
}

func decode_block(id *ID, decoder UpdateDecoder) (Block, error) {
	info, err := decoder.ReadVarUint()
	if err != nil {
		return nil, err
	}
	switch info {
	case uint64(BlockKindSkip):
		length, err := decoder.ReadVarUint()
		if err != nil {
			return nil, err
		}
		val := newSkip(id, length)
		return val, nil
	case uint64(BlockKindGC):
		length, err := decoder.ReadLength()
		if err != nil {
			return nil, err
		}
		val := newGC(id, length)
		return val, nil
	default:
		item, err := decode_item(decoder, byte(info))
		if err != nil {
			return nil, err
		}
		return item, nil
	}
}

func decode_update(decoder UpdateDecoder) (*Update, error) {
	clients_length, err := decoder.ReadVarUint()
	if err != nil {
		return nil, err
	}

	clients := make(map[ClientID][]Block, clients_length)
	update_blocks := &UpdateBlocks{
		clients: clients,
	}

	for range clients_length {
		blocks_length, err := decoder.ReadVarUint()
		if err != nil {
			return nil, err
		}
		client, err := decoder.ReadClient()
		if err != nil {
			return nil, err
		}
		clock, err := decoder.ReadVarUint()
		if err != nil {
			return nil, err
		}

		blocks := make([]Block, 0, blocks_length)

		for range blocks_length {
			id := newID(client, clock)
			block, err := decode_block(id, decoder)
			if err != nil {
				return nil, err
			}
			clock += block.Length()
			blocks = append(blocks, block)
		}
	}

	delete_set, err := decode_id_set(decoder)
	if err != nil {
		return nil, err
	}

	return &Update{blocks: update_blocks, delete_set: delete_set}, nil
}

func DecodeUpdate(decoder UpdateDecoder) (*Update, error) {
	return decode_update(decoder)
}
