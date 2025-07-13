package ygo

type DocOption func(doc *Doc)

func WithGUID(guid string) DocOption {
	return func(doc *Doc) {
		doc.guid = guid
	}
}

func WithGC(gc bool) DocOption {
	return func(doc *Doc) {
		doc.gc = gc
	}
}

func WithGCFilter(predicate gcFilter) DocOption {
	return func(doc *Doc) {
		doc.gcFilter = predicate
	}
}

func WithCollectionId(id string) DocOption {
	return func(doc *Doc) {
		doc.collectionId = &id
	}
}

func WithShouldLoad(b bool) DocOption {
	return func(doc *Doc) {
		doc.shouldLoad = b
	}
}

func WithAutoLoad(b bool) DocOption {
	return func(doc *Doc) {
		doc.autoLoad = b
	}
}

func WithMeta(b any) DocOption {
	return func(doc *Doc) {
		doc.meta = b
	}
}
