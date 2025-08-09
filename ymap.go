package ygo

type YMap struct {
	*BaseSharedType

	// Preliminary content before integration
	prelimContent map[string]any
}

func newYMap(prelim map[string]any) *YMap {
	return &YMap{
		BaseSharedType: newBaseSharedType(),
		prelimContent:  prelim,
	}
}

// NewYMap creates a new YMap instance
func NewYMap() *YMap {
	return newYMap(make(map[string]any))
}

func NewYMapWithEntries(entries map[string]any) *YMap {
	return newYMap(entries)
}

func (ymap *YMap) Write(encoder UpdateEncoder) error {
	return encoder.WriteTypeRef(1)
}
