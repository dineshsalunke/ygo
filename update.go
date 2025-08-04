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

func mergeUpdates(updates [][]byte) ([]byte, error) {
	panic("not implemented")
}

func applyUpdate(decoder UpdateDecoder, tx *Transaction) error {
	retry := false
	store := tx.doc.store
	// Read remote updates
	remoteUpdates, err := readStructSet(decoder, tx)
	if err != nil {
		return err
	}

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

	if err := remoteUpdates.excludeIdSet(localState); err != nil {
		return err
	}

	restStructs, err := store.IntegrateStructs(tx)
	if err != nil {
		return err
	}
	if store.pendingStructs != nil {
		for client, clock := range store.pendingStructs.missingState {
			_, remoteHasClient := remoteUpdates.clients[client]
			clientClockLength := store.GetClientClockLength(client)
			if remoteHasClient || clock < clientClockLength {
				retry = true
				break
			}
		}

		if restStructs != nil {
			for client, clock := range restStructs.missingState {
				mclock, has := store.pendingStructs.missingState[client]
				if has || mclock > clock {
					store.pendingStructs.missingState[client] = clock
				}
			}
			update, err := mergeUpdates([][]byte{
				store.pendingStructs.update,
				restStructs.update,
			})
			if err != nil {
				return err
			}
			store.pendingStructs.update = update
		}
	} else {
		store.pendingStructs = restStructs
	}

	restIdSetUpdate, err := store.ApplyIdSet(decoder, tx)
	if err != nil {
		return err
	}
	if store.pendingIdSetUpdate != nil {
		pendingDsDecoder := newUpdateDecoderV1(store.pendingIdSetUpdate)
		// We only encode pending deletes, so lets get rid of the first 0
		_, err := pendingDsDecoder.ReadVarUint()
		if err != nil {
			return err
		}
		pendingIdSetUpdate, err := store.ApplyIdSet(pendingDsDecoder, tx)
		if err != nil {
			return err
		}

		if restIdSetUpdate != nil && pendingIdSetUpdate != nil {
			pendingDsUpdate, err := mergeUpdates([][]byte{restIdSetUpdate, pendingIdSetUpdate})
			if err != nil {
				return err
			}
			store.pendingIdSetUpdate = pendingDsUpdate
		} else {
			store.pendingIdSetUpdate = pendingIdSetUpdate
			if restIdSetUpdate != nil {
				store.pendingIdSetUpdate = restIdSetUpdate
			}
		}
	} else {
		store.pendingIdSetUpdate = restIdSetUpdate
	}

	if retry {
		update := store.pendingStructs.update
		store.pendingStructs = nil
		if err := ApplyUpdateV2(tx.doc, update); err != nil {
			return err
		}
	}
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
