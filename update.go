package ygo

import "fmt"

func readBlockSet(decoder UpdateDecoder, tx *Transaction) (*BlockSet, error) {
	numOfClients, err := decoder.ReadVarUint()
	if err != nil {
		return nil, fmt.Errorf("failed to read number of clients in update")
	}
	blockSet := newBlockSet(numOfClients)

	for ci := range numOfClients {
		numOfBlocks, err := decoder.ReadVarUint()
		if err != nil {
			return nil, fmt.Errorf("failed to read number of blocks for client at index %d", ci)
		}

		client, err := decoder.ReadClient()
		if err != nil {
			return nil, fmt.Errorf("failed to read client id at index %d", ci)
		}

		clock, err := decoder.ReadVarUint()
		if err != nil {
			return nil, fmt.Errorf("failed to read clock for client %d at index %d", client, ci)
		}

		id := newID(client, clock)
		refs := make([]Block, numOfBlocks)
		for bi := range numOfBlocks {
			info, err := decoder.ReadInfo()
			if err != nil {
				return nil, fmt.Errorf("failed to read block info for client %d at index %d", client, bi)
			}
			switch Kind(info & 31) {
			case BlockKindGC:
				length, err := decoder.ReadLength()
				if err != nil {
					return nil, fmt.Errorf("failed to read gc length at index %d for client %d", bi, client)
				}
				block := newGC(id, length)
				refs[bi] = block
				clock = clock + length
			case BlockKindSkip:
				length, err := decoder.ReadVarUint()
				if err != nil {
					return nil, fmt.Errorf("failed to read skip length at index %d for client %d", bi, client)
				}
				block := newSkip(id, length)
				refs[bi] = block
				clock = clock + length
			default:
				block, err := decodeItem(id, decoder, Kind(info), tx.doc)
				if err != nil {
					return nil, fmt.Errorf("failed to decode item block for client %d at index at %d with info %d: %w", client, bi, info&31, err)
				}
				refs[bi] = block
				clock = clock + block.Length()
			}

		}
		blockSet.addRange(client, refs)
	}

	return blockSet, nil
}

func mergeUpdatesV1(updates [][]byte) ([]byte, error) {
	panic("not implemented")
}

func applyUpdate(decoder UpdateDecoder, tx *Transaction, txOrigin any) error {
	retry := false
	store := tx.doc.store
	// Read remote updates
	remoteBlockSet, err := readBlockSet(decoder, tx)
	if err != nil {
		return err
	}

	localState := newIdSet()
	for remoteClientId := range remoteBlockSet.clients {
		localBlocks, has := store.clients[remoteClientId]
		if has {
			// Assume all the local updates are contigous and insert the remote updates at last
			lastBlock := localBlocks[len(localBlocks)-1]
			localState.add(remoteClientId, 0, lastBlock.ClockEnd())
			idRanges, ok := store.skips.clients[remoteClientId]
			if ok {
				for _, idRange := range idRanges.getIdRanges() {
					localState.delete(remoteClientId, idRange.clock, idRange.length)
				}
			}
		}
	}

	if err := remoteBlockSet.excludeIdSet(localState); err != nil {
		return err
	}

	restStructs, err := store.IntegrateStructs(tx, remoteBlockSet)
	if err != nil {
		return err
	}
	if store.pendingStructs != nil {
		for client, clock := range store.pendingStructs.missingState {
			_, remoteHasClient := remoteBlockSet.clients[client]
			clientClockLength := store.GetClientClockEnd(client)
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
			update, err := mergeUpdatesV1([][]byte{
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
			pendingDsUpdate, err := mergeUpdatesV1([][]byte{restIdSetUpdate, pendingIdSetUpdate})
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
		if err := ApplyUpdateV1(tx.doc, update, nil); err != nil {
			return err
		}
	}
	return nil
}

func ApplyUpdateV1(doc *Doc, update []byte, txOrigin any) error {
	_, err := Transact(doc, func(tx *Transaction) (any, error) {
		decoder := newUpdateDecoderV1(update)

		if err := applyUpdate(decoder, tx, txOrigin); err != nil {
			return nil, err
		}
		return nil, nil
	}, txOrigin, false)
	return err
}

func diffUpdateV1(update []byte, otherUpdate []byte) ([]byte, error) {
	_, err := decodeStateVector(update)
	if err != nil {
		return nil, err
	}
	panic("not implemented")
}

func decodeStateVector(buf []byte) (StateVector, error) {
	return readStateVector(newIdSetDecoderV1(buf))
}

func readStateVector(decoder IdSetDecoder) (StateVector, error) {
	ssLength, err := decoder.ReadVarUint()
	if err != nil {
		return nil, err
	}
	sv := make(StateVector, ssLength)
	for range ssLength {
		client, err := decoder.ReadVarUint()
		if err != nil {
			return nil, err
		}
		clock, err := decoder.ReadVarUint()
		if err != nil {
			return nil, err
		}
		sv[client] = clock
	}
	return sv, nil
}

func writeStateAsUpdate(encoder UpdateEncoder, doc *Doc, targetVector StateVector) error {
	panic("not implemented")
}

func encodeStateAsUpate(doc *Doc, encodedTargetStateVector []byte, encoder UpdateEncoder) ([]byte, error) {
	decoder := newIdSetDecoderV1(encodedTargetStateVector)
	targetStateVector, err := readStateVector(decoder)
	if err != nil {
		return nil, err
	}
	if err := writeStateAsUpdate(encoder, doc, targetStateVector); err != nil {
		return nil, err
	}
	updates := make([][]byte, 1)
	updates = append(updates, encoder.Bytes())
	if doc.store.pendingIdSetUpdate != nil {
		updates = append(updates, doc.store.pendingIdSetUpdate)
	}
	if doc.store.pendingStructs != nil {
		diffedUpdates, err := diffUpdateV1(doc.store.pendingIdSetUpdate, encodedTargetStateVector)
		if err != nil {
			return nil, err
		}
		updates = append(updates, diffedUpdates)
	}

	if len(updates) > 1 {
		return mergeUpdatesV1(updates)
	}

	return updates[0], nil
}
