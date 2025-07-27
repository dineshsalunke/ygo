package ygo

const (
	HasOriginFlag      byte = 0b10000000
	HasRightOriginFlag byte = 0b01000000
	HasParentSubFlag   byte = 0b00100000
)

const (
	BlockKindGC          Kind = 0
	BlockKindItemDeleted Kind = 1
	BlockKindItemJson    Kind = 2
	BlockKindItemBinary  Kind = 3
	blockkindItemString  Kind = 4
	BlockKindItemEmbed   Kind = 5
	BlockKindItemFormat  Kind = 6
	BlockKindItemType    Kind = 7
	BlockKindAny         Kind = 8
	BlockKindDoc         Kind = 9
	BlockKindSkip        Kind = 10
	BlockKindItemMove    Kind = 11
)
