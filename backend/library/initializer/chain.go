package initializer

import "fmt"

type InitFunc interface {
	Init() error
}

type InitChain struct {
	Initializers []InitFunc
}

func NewInitChain() *InitChain {
	return &InitChain{}
}

func (chain *InitChain) Add(initializers ...InitFunc) {
	chain.Initializers = append(chain.Initializers, initializers...)
}

func (chain *InitChain) Execute() {
	for _, initializer := range chain.Initializers {
		err := initializer.Init()
		if err != nil {
			panic(fmt.Sprintf("initializer failed: %v", err))
		}
	}
}
