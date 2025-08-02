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
	store := tx.doc.store
	// Read remote updates
	remoteUpdates, err := readStructSet(decoder, tx)
	if err != nil {
		return err
	}

	// TODO: find all local updates for remote clients
	localState := newIdSet()
	for remoteClientId := range remoteUpdates.clients {
		localBlocks, has := store.clients[remoteClientId]
		if has {
			// Assume all the local updates are contigous and insert the remote updates at last
			lastBlock := localBlocks[len(localBlocks)-1]
			localState.add(remoteClientId, 0, lastBlock.ClockLength())
			idRanges, ok := store.skips.clients[remoteClientId]
			if ok {
				for _, idRange := range idRanges.getIdRanges() {
					localState.delete(remoteClientId, idRange.clock, idRange.length)
				}
			}
		}
	}

	// TODO: remove all the overlapping updates of same remote and local clients

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
