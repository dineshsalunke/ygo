package ygo

type IdRange struct{}
type IdRanges []*IdRange

type IdSet map[ClientID]IdRanges

func decode_id_set(decoder UpdateDecoder) (IdSet, error) {
	panic("not implemented")
}
