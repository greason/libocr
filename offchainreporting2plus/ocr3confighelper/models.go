package ocr3confighelper

type NodeOCR2MercuryConfig struct {
	Id                int
	ConfigPublicKey   string
	OnchainPublicKey  string
	OffchainPublicKey string

	PeerID        string
	OffChainKeyId string
	CSAPublicKey  string
}
