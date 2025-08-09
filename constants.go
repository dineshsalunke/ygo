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
	BlockKindItemString  Kind = 4
	BlockKindItemEmbed   Kind = 5
	BlockKindItemFormat  Kind = 6
	BlockKindItemType    Kind = 7
	BlockKindItemAny     Kind = 8
	BlockKindItemDoc     Kind = 9
	BlockKindSkip        Kind = 10
	BlockKindItemMove    Kind = 11
)

const (
	SkipTypeRef byte = 10
	GCTypeRef   byte = 0
)

const (
	ItemFlagLinked    uint16 = 0b0001_0000_0000
	ItemFlagMarked    uint16 = 0b0000_1000
	ItemFlagDeleted   uint16 = 0b0000_0100
	ItemFlagCountable uint16 = 0b0000_0010
	ItemFlagKeep      uint16 = 0b0000_0001
)
