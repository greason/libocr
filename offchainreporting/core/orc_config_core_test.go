package core

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ocrConfigHelper "github.com/smartcontractkit/libocr/offchainreporting/confighelper"
	"github.com/smartcontractkit/libocr/offchainreporting/internal/test"
	ocrTypes "github.com/smartcontractkit/libocr/offchainreporting/types"
	"testing"
	"time"
)

func AproOffChainAggregatorConfig(numberNodes int, target int) test.OffChainAggregatorConfig {
	if numberNodes <= 4 {
		fmt.Printf("insufficient number of nodes (%d) supplied for OCR, need at least 5", numberNodes)
	}
	s := []int{1}
	// First node's stage already inputted as a 1 in line above, so numberNodes-1.
	for i := 0; i < numberNodes-1; i++ {
		s = append(s, 2)
	}
	// chainTypeSlowUpdates
	// limits.MinDeltaC(10min) <= cfg.DeltaC
	// limits.MinDeltaStage(10s) <= cfg.DeltaStage
	// cfg.DeltaGrace < cfg.DeltaRound
	// cfg.DeltaRound < cfg.DeltaProgress
	// 0 < cfg.RMax && cfg.RMax < 255
	// len(cfg.S) < 1000

	var AlphaPPB = uint64(10000000)
	var DeltaC = time.Hour * 24
	var DeltaProgress = time.Second * 35
	var DeltaRound = time.Second * 30

	switch target {
	case CoreBtc:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	}

	return test.OffChainAggregatorConfig{
		AlphaPPB:         AlphaPPB, // 10 ^9
		DeltaC:           DeltaC,
		DeltaGrace:       time.Second * 12,
		DeltaProgress:    DeltaProgress,
		DeltaStage:       time.Second * 60,
		DeltaResend:      time.Second * 17,
		DeltaRound:       DeltaRound,
		RMax:             6,
		S:                s,
		N:                numberNodes,
		F:                1,
		OracleIdentities: []ocrConfigHelper.OracleIdentityExtra{},
	}
}

const (
	CoreUsdt = iota
	CoreUsdc
	CoreBtc
	CoreEth
)

func GetNodeConfigs(target int) []test.NodeOCRConfig {
	nodeConfigs := make(map[int][]test.NodeOCRConfig)

	// mode-main
	{
		nodeConfigsCoreBtc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xCeA72bBF5577F9BC674dFD9a128D0FF2dbD30d70",
				SignAddress:     "0xc3c3Bf3e145b6326925bb23d612Bf7941eD15D33",
				ConfigPubKey:    "69a02a3df1a92c6bc4cf666472ce13d3a538edd50df8e3dfe285bf5787f5be0f",
				OffChainPubKey:  "4e63adc14895af73674c498ce91b558b2878427bf3ae6aa61425953b088a0fbf",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "4bf103bca4cbc1dc6f905737578828b505b5f2562a1bc15576dd431c5b06bcec",
			}, {
				Id:              2,
				TransmitAddress: "0xc355159B16962609eB041Ab749E8BA337dC498f9",
				SignAddress:     "0x38643f74E1B7Cb7e2a4e03F53D1D61Bf08C1b2b3",
				ConfigPubKey:    "be5cbcb185f91d49a579d35e27c4ebe4e38a06e8ca912286e1d7a996b0d42866",
				OffChainPubKey:  "77b1da8dbb8a3cc2bac43d862874924a6b45cb74ccfbae349e8b35c8953db901",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "2f1e3fd8d25bc893472c4997c2cf7826d30ddaf2d7229d172d0550270e879d7c",
			}, {
				Id:              3,
				TransmitAddress: "0x4AbcA6C08236776308e008dE2ECAd7C0867d4b08",
				SignAddress:     "0x66B6cD6293f4EbFFE06D4Cb97E6531fABD10F772",
				ConfigPubKey:    "c9e3e3ec742e50c6f39d42b0a70de8c0371eed077e4ec35db0d94654f3173043",
				OffChainPubKey:  "766258f9d2c417d62c45c9c4d5266255c824b071cf8f7d6ebfcfa0e7a464aa57",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "59399627e153422510fc2d04422470af29e4ffef228486c5ed2847faeba5517f",
			}, {
				Id:              4,
				TransmitAddress: "0xaEC1B644e3217370E7973953aa15DEA76e15288B",
				SignAddress:     "0x2D706C456F7Ba733A0714D618E1bCf6A7D371BC3",
				ConfigPubKey:    "9d56443ff3533380c3f95bc210cefd8b7babef4fdcde8a6ee7dfc0426c261b6f",
				OffChainPubKey:  "66f2a113329326463f5b59a3e1332dd267ce4d0fe57ebd7f0d68876b5839b19e",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "04826b1b816745e9fcc0d7ec86663f4d141b5e2e76d38f28d90231b46c84c484",
			}, {
				Id:              5,
				TransmitAddress: "0x41fCC88c8Dc99732E81A18B3ed037e6B040026F4",
				SignAddress:     "0xf51Df33Ef5dAA3f639BC061673f97e773DAc9960",
				ConfigPubKey:    "299714216374487d9183e4835255fa5a57c2f3489e263bc24dd71e062d1b961f",
				OffChainPubKey:  "afb79741ff1e9310c84282d7ca5b523a316c9ebcc79f91afdee088cfc4984a67",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "d70e505e6cd7606393a8c339baadd29869bca483a6400defc4c4e05fda5cde96",
			}, {
				Id:              6,
				TransmitAddress: "0xC183015730d6CB67137f2D4E7e41f04275C29989",
				SignAddress:     "0xA0ef06551755ea7e5d854fD0b73044752f5b757C",
				ConfigPubKey:    "8b786e01dd0cd27d903b9e39869c6f27472db514aa9dc83d7ac84aa73629f713",
				OffChainPubKey:  "6e7d044e5513fe61d4fec5b0e9ce284359095cdc85733ae79a74a8d65956d443",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "4bb8d4b0b46fcf2faa3bfa7327e48c4d0387f2de638ee1f3fa6294d8e3e02c8b",
			}, {
				Id:              7,
				TransmitAddress: "0x1A7EEe3C1C6ff6dC44Ed7425A9F6b7A7C0E017ef",
				SignAddress:     "0x8BBae17f7CeC01070ce886E8B9404399a6060070",
				ConfigPubKey:    "ca14f9a09a17a4a9992e2e127bc461e8f65a4e8117237ca241b2747395304d7f",
				OffChainPubKey:  "24d6d0f31d5476e7bdd5f89e4f8b8bc932df342dbb237c6c34ccbba33188fc3c",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "f77a1e898b6123662b07e5886a86d2087b10963491eece73264fa40d82f8ed80",
			}, {
				Id:              8,
				TransmitAddress: "0x4a2005Fc56fd40Db346d4a6713D27765D7e1d44a",
				SignAddress:     "0xAaAB4554058D2D960145E1DA30206ae633FC6cc4",
				ConfigPubKey:    "96b9bad8f6326248d2ff9804377fd71ea9429dc7eaaf80fde2d3f4344c10c01c",
				OffChainPubKey:  "3561c3d135a2f34886bb3afbfe8254c96bf0bfdde693c1abe15b44e27442af18",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "7527cf389d90e28d65caaa927190635152e12b3029248d0962766d85dcaeb743",
			}, {
				Id:              9,
				TransmitAddress: "0xfE20d7AfE1edE8d064d8679747fcf00f9FBEd486",
				SignAddress:     "0xC9c9f2cea2e5f4E1bD8fB5a5aC634Be1a9276526",
				ConfigPubKey:    "db13e328e66d384f5f2dcb15b00c75a973aec7c71c156584305e90112e3dd979",
				OffChainPubKey:  "6939189feed3dfeac719f92e7c4c49d8f8baa4049f7ed492ebe8bd8b8b209785",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "777c1ab5090b2dbabea6d40ab0cacd24d5a280799d190fa0bb2bd5ca891410eb",
			},
		}
		nodeConfigs[CoreBtc] = nodeConfigsCoreBtc
	}

	return nodeConfigs[target]
}

func GetOffChainAggregatorConfig(target int) test.OffChainAggregatorConfig {
	nodeConfigs := GetNodeConfigs(target)
	ocrConfig := AproOffChainAggregatorConfig(len(nodeConfigs), target)
	for _, nodeConfig := range nodeConfigs {
		// Need to convert the key representations
		var onChainSigningAddress [20]byte
		var configPublicKey [32]byte
		offchainSigningAddress, err := hex.DecodeString(nodeConfig.OffChainPubKey)
		if err != nil {
			panic(err)
		}
		decodeConfigKey, err := hex.DecodeString(nodeConfig.ConfigPubKey)
		if err != nil {
			panic(err)
		}

		// https://stackoverflow.com/questions/8032170/how-to-assign-string-to-bytes-array
		copy(onChainSigningAddress[:], common.HexToAddress(nodeConfig.SignAddress).Bytes())
		copy(configPublicKey[:], decodeConfigKey)

		oracleIdentity := ocrConfigHelper.OracleIdentity{
			TransmitAddress:       common.HexToAddress(nodeConfig.TransmitAddress),
			OnChainSigningAddress: onChainSigningAddress,
			PeerID:                nodeConfig.PeerID,
			OffchainPublicKey:     offchainSigningAddress,
		}
		oracleIdentityExtra := ocrConfigHelper.OracleIdentityExtra{
			OracleIdentity:                  oracleIdentity,
			SharedSecretEncryptionPublicKey: ocrTypes.SharedSecretEncryptionPublicKey(configPublicKey),
		}

		ocrConfig.OracleIdentities = append(ocrConfig.OracleIdentities, oracleIdentityExtra)
	}
	return ocrConfig
}

func TestEncodeOCRConfig(t *testing.T) {
	ocrConfig := GetOffChainAggregatorConfig(CoreBtc)
	signers, transmitters, threshold, encodedConfigVersion, encodedConfig, err := ocrConfigHelper.ContractSetConfigArgs(
		ocrConfig.DeltaProgress,
		ocrConfig.DeltaResend,
		ocrConfig.DeltaRound,
		ocrConfig.DeltaGrace,
		ocrConfig.DeltaC,
		ocrConfig.AlphaPPB,
		ocrConfig.DeltaStage,
		ocrConfig.RMax,
		ocrConfig.S,
		ocrConfig.OracleIdentities,
		ocrConfig.F,
	)
	fmt.Printf("signers: %v, transmitters: %v, threshold: %v, encodedConfigVersion: %v, encodedConfig: %v, err: %v",
		signers, transmitters, threshold, encodedConfigVersion, encodedConfig, err)
	fmt.Printf("\nencodedConfig: %v", hexutil.Encode(encodedConfig))
}
