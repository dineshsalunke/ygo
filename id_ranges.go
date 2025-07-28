package ygo

import "slices"

type IdRanges struct {
	sorted    bool
	last_used bool
	ids       []*IdRange
}

func newIdRanges(ids []*IdRange) *IdRanges {
	return &IdRanges{
		sorted:    false,
		last_used: false,
		ids:       ids,
	}
}

func (self *IdRanges) addIdRange(clock, length uint64) {
	var last *IdRange
	if len(self.ids) > 0 {
		last = self.ids[len(self.ids)-1]
	}
	if last != nil && last.clock+last.length == clock {
		if self.last_used {
			self.ids[len(self.ids)-1] = newIdRange(last.clock, last.length+length)
			last.length += length
			self.last_used = false
		} else {
			last.length += length
		}
	} else {
		self.sorted = false
		self.ids = append(self.ids, newIdRange(clock, length))
	}
}

func (irs *IdRanges) squash() []*IdRange {
	irs.last_used = true
	if !irs.sorted {
		irs.sorted = true
		slices.SortFunc(irs.ids, func(a, b *IdRange) int {
			return int(a.clock) - int(b.clock)
		})

		j := 1
		for i := 1; i < len(irs.ids); i++ {
			left := irs.ids[j-1]
			right := irs.ids[i]
			if left.end() >= right.start() {
				r := right.end() - left.start()
				if left.length < r {
					irs.ids[j-1] = &IdRange{clock: left.clock, length: r}
				}
			} else if left.end() == 0 {
				irs.ids[j-1] = right
			} else {
				if j < i {
					irs.ids[j] = right
				}
				j++
			}
		}
		if irs.ids[j-1].length == 0 {
			irs.ids = irs.ids[:j-1]
		} else {
			irs.ids = irs.ids[:j]
		}
	}

	return irs.ids
}
