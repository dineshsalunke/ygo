package ygo

type SharedTypeIntegrator interface {
	Integrate(doc *Doc, item SharedStruct) error
}

type SharedType interface {
	SharedTypeIntegrator
}
