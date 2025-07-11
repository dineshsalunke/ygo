package core

type Integrator interface{}

type Writer interface{}

type IntegrateWriter interface {
	Integrator
	Writer
	Length() uint64
	Clock() uint64
}
