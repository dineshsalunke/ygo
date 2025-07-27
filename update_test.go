package ygo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeUpdate(t *testing.T) {
	/*
	   Generated with:
	   	```js
	   	var Y = require('yjs');

	   	var doc = new Y.Doc()
	   	var map = doc.getMap()
	   	map.set('keyB', 'valueB')

	   	// Merge changes from remote
	   	var update = Y.encodeStateAsUpdate(doc)
	   	```
	*/

	update := []byte{
		1, 1, 176, 249, 159, 198, 7, 0, 40, 1, 0, 4, 107, 101, 121, 66, 1, 119, 6, 118, 97,
		108, 117, 101, 66, 0,
	}

	id := newID(2026372272, 0)
	decoder := newUpdateDecoderV1(update)
	u, err := decode_update(decoder)
	assert.Nil(t, err)
	block, ok := u.blocks.clients[id.client]
	assert.True(t, ok)
	assert.NotNil(t, block)
	expected := make([]Block, 1)
	content := make([]any, 1)
	content[0] = "valueB"
	expected[0] = &Item{
		block:      newBlock(id, 1),
		parent_sub: "keyB",
		content:    newItemContentAny(content),
	}
	assert.Equal(t, expected, block)
}
