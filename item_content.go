package ygo

import "fmt"

type item_content_decoder func(decoder UpdateDecoder, info Kind) (ItemContent, error)

var decoders map[Kind]item_content_decoder = make(map[Kind]item_content_decoder, 0)

type ItemContent interface {
	IsCountable() bool
	Length() uint64
	GetRef() Kind
}

func decode_item_content(decoder UpdateDecoder, info Kind) (ItemContent, error) {
	item_decoder, ok := decoders[Kind(info&31)]
	if !ok {
		return nil, fmt.Errorf("unregistered item content kind : %d", info)
	}
	content, err := item_decoder(decoder, info)
	if err != nil {
		return nil, err
	}
	return content, nil
}
