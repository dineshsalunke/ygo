package ygo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExcludeSet(t *testing.T) {
	ss := newBlockSet(4)
	ss.addRange(1, []Block{
		newItem(newID(1, 10), WithContent(newItemContentAny([]any{1}))),
		newItem(newID(1, 11), WithContent(newItemContentAny([]any{1, 2, 3, 4}))),
		newItem(newID(1, 15), WithContent(newItemContentAny([]any{1, 2, 3, 4, 5}))),
		newItem(newID(1, 20), WithContent(newItemContentAny([]any{1}))),
	})
	idSet := newIdSet()
	idSet.add(1, 11, 8)
	err := ss.excludeIdSet(idSet)
	assert.Nil(t, err)
	assert.Equal(t, []Block{
		newItem(newID(1, 10), WithContent(newItemContentAny([]any{1}))),
		newSkip(newID(1, 11), 8),
		newItem(newID(1, 20), WithContent(newItemContentAny([]any{1}))),
	}, ss.clients[1].refs)
}
