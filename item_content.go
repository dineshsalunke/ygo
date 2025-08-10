package ygo

import "fmt"

type ItemContentDecoderFactory func(decoder UpdateDecoder, info Kind) (ItemContent, error)

var Decoders map[Kind]ItemContentDecoderFactory = make(map[Kind]ItemContentDecoderFactory, 0)

type ItemContent interface {
	Content() []any
	Length() uint64
	IsCountable() bool
	Copy() ItemContent
	Splice(offset uint64) ItemContent
	MergeWith(right ItemContent) bool
	Integrate(tx *Transaction, item Block) error
	Delete(tx *Transaction) error
	GC(tx *Transaction) error
	Write(encoder UpdateEncoder, offset uint64, offsetEnd uint64) error
	GetRef() Kind
}

func decodeItemContent(decoder UpdateDecoder, info Kind) (ItemContent, error) {
	item_decoder, ok := Decoders[Kind(info&31)]
	if !ok {
		return nil, fmt.Errorf("unregistered item content kind : %d", info)
	}
	content, err := item_decoder(decoder, info)
	if err != nil {
		return nil, err
	}
	return content, nil
}
