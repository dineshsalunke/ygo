package ygo

import "fmt"

type ItemContentDecoderFactory func(decoder UpdateDecoder, info Kind) (ItemContent, error)

var Decoders map[Kind]ItemContentDecoderFactory = make(map[Kind]ItemContentDecoderFactory, 0)

type ItemContent interface {
	IsCountable() bool
	Length() uint64
	GetRef() Kind
	Splice(offset uint64) ItemContent
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
