package ygo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdRangessquash(t *testing.T) {
	irs := newIdRanges([]*IdRange{
		{clock: 0, length: 2},
		{clock: 3, length: 2},
		{clock: 6, length: 2},
	})
	irs.squash()
	assert.Len(t, irs.ids, 2)
	assert.Equal(t, []*IdRange{
		{clock: 0, length: 5},
		{clock: 6, length: 2},
	}, irs.ids)
	//
}
