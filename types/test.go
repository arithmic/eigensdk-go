package types

import "github.com/arithmic/eigensdk-go/crypto/bls"

type TestOperator struct {
	OperatorId     OperatorId
	StakePerQuorum map[QuorumNum]StakeAmount
	BlsKeypair     *bls.KeyPair
}
