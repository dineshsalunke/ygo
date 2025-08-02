package ygo

func readStructSet(decoder UpdateDecoder, tx *Transaction) (*StructSet, error) {
	numOfUpdates, err := decoder.ReadVarUint()
	if err != nil {
		return nil, err
	}
	ss := newStructSet(numOfUpdates)

	for range numOfUpdates {
		numOfStructs, err := decoder.ReadVarUint()
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

		id := newID(client, clock)
		refs := make([]Block, numOfStructs)
		for i := range numOfStructs {
			info, err := decoder.ReadInfo()
			if err != nil {
				return nil, err
			}
			switch Kind(info & 31) {
			case BlockKindGC:
				length, err := decoder.ReadLength()
				if err != nil {
					return nil, err
				}
				block := newGC(id, length)
				refs[i] = block
				clock = clock + length
			case BlockKindSkip:
				length, err := decoder.ReadVarUint()
				if err != nil {
					return nil, err
				}
				block := newSkip(id, length)
				refs[i] = block
				clock = clock + length
			default:
				block, err := decodeItem(id, decoder, Kind(info), tx.doc)
				if err != nil {
					return nil, err
				}
				refs[i] = block
				clock = clock + block.Length()
			}

		}
		ss.addRange(client, refs)
	}

	return ss, nil
}

func applyUpdate(decoder UpdateDecoder, tx *Transaction) error {
	// Read remote updates
	_, err := readStructSet(decoder, tx)
	if err != nil {
		return err
	}

	// TODO: find all local updates for remote clients
	// TODOD: remove all the overlapping updates of same remote and local clients

	// TODO: Integrate remote updates, and return udpates for which deps could not be resolved
	// TODO: Check if we have pending updates to be merged, if any then merge else assign the missing deps update to pending updates

	// TODO: Read DeleteSet
	// TODO: Apply DeleteSet, return the ones which couldn't be applied
	// TODO: Check for pending DeleteSet and apply it, else if we have any DeleteSet which couldn't be applied from earlier step then make a note of them

	// TODO: check if something couldn't be applied due to missing deps, if any then retry applying update
	return nil
}

func ApplyUpdateV2(doc *Doc, update []byte) error {
	tx := newTransaction(doc, TransactionWithLocal(false))
	decoder := newUpdateDecoderV1(update)

	if err := applyUpdate(decoder, tx); err != nil {
		return err
	}

	return tx.commitTransaction()
}
