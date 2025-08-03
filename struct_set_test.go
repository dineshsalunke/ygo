package ygo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExcludeSet(t *testing.T) {
	ss := newStructSet(4)
	ss.addRange(1, []Block{
		&Item{
			block: &block{
				id:     newID(1, 10),
				length: 1,
			},
			content: newItemContentAny([]any{1}),
		},
		&Item{
			block: &block{
				id:     newID(1, 11),
				length: 4,
			},
			content: newItemContentAny([]any{1, 2, 3, 4}),
		},
		&Item{
			block: &block{
				id:     newID(1, 15),
				length: 5,
			},
			content: newItemContentAny([]any{1, 2, 3, 4, 5}),
		},
		&Item{
			block: &block{
				id:     newID(1, 20),
				length: 1,
			},
			content: newItemContentAny([]any{1}),
		},
	})
	idSet := newIdSet()
	idSet.add(1, 11, 8)
	err := ss.excludeIdSet(idSet)
	assert.Nil(t, err)
	assert.Equal(t, []Block{
		&Item{
			block: &block{
				id:     newID(1, 10),
				length: 1,
			},
			content: newItemContentAny([]any{1}),
		},
		&Skip{
			block: &block{
				id:     newID(1, 11),
				length: 8,
			},
		},
		&Item{
			block: &block{
				id:     newID(1, 20),
				length: 1,
			},
			content: newItemContentAny([]any{1}),
		},
	}, ss.clients[1].refs)
}
