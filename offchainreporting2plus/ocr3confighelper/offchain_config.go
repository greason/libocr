package ocr3confighelper

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

type OffchainConfig struct {
	ExpirationWindow uint32          `json:"expirationWindow"` // window in which a report can be verified in seconds
	BaseUSDFee       decimal.Decimal `json:"baseUSDFee"`       // Base USD fee, 10c base fee to verify the report
}

func DecodeOffchainConfig(b []byte) (o OffchainConfig, err error) {
	err = json.Unmarshal(b, &o)
	if err != nil {
		return o, fmt.Errorf("failed to decode offchain config: must be valid JSON (got: 0x%x); %w", b, err)
	}
	return
}

func (c OffchainConfig) Encode() ([]byte, error) {
	return json.Marshal(c)
}
