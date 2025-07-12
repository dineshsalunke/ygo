package ygo

type SharedTypeIntegrator interface {
	Integrate(doc *Doc, item any) error
}

type SharedType interface {
	SharedTypeIntegrator
}
