package core

type NeuralChainBlock interface {
}

type neuralChainBlock struct{}

type neuronLinkHash uint64

type neuronLink map[neuronLinkHash]map[int32]interface{}
