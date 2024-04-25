package test

import (
	ocrConfigHelper "github.com/smartcontractkit/libocr/offchainreporting/confighelper"
	"time"
)

type NodeOCRConfig struct {
	Id              int
	TransmitAddress string
	SignAddress     string
	ConfigPubKey    string
	OffChainPubKey  string
	PeerID          string
	OffChainKeyId   string
}

type OffChainAggregatorConfig struct {
	DeltaProgress    time.Duration // The duration in which a leader must achieve progress or be replaced
	DeltaResend      time.Duration // The interval at which nodes resend NEWEPOCH messages
	DeltaRound       time.Duration // The duration after which a new round is started
	DeltaGrace       time.Duration // The duration of the grace period during which delayed oracles can still submit observations
	DeltaC           time.Duration // Limits how often updates are transmitted to the contract as long as the median isn’t changing by more then AlphaPPB
	AlphaPPB         uint64        // Allows larger changes of the median to be reported immediately, bypassing DeltaC
	DeltaStage       time.Duration // Used to stagger stages of the transmission protocol. Multiple Ethereum blocks must be mineable in this period
	RMax             uint8         // The maximum number of rounds in an epoch
	S                []int         // Transmission Schedule
	F                int           // The allowed number of "bad" oracles
	N                int           // The number of oracles
	OracleIdentities []ocrConfigHelper.OracleIdentityExtra
}
