package common

import (
	"math/big"
	pb "github.com/xlab-si/emmy/comm/pro"
)

type ProtocolType uint8

const (
	Sigma ProtocolType = iota + 1 // SigmaProtocol
	ZKP //ZeroKnowledgeProof
	ZKPOK //ZeroKnowledgeProofOfKnowledge
)

type ECGroupElement struct {
	X *big.Int
	Y *big.Int
}

func ToECGroupElement(el *pb.ECGroupElement) *ECGroupElement {
	x := ECGroupElement{X: new(big.Int).SetBytes(el.X), Y: new(big.Int).SetBytes(el.Y)}
	return &x
}

func ToPbECGroupElement(el *ECGroupElement) *pb.ECGroupElement {
	x := pb.ECGroupElement{X: el.X.Bytes(), Y: el.Y.Bytes()}
	return &x
}