package core

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ocrConfigHelper "github.com/smartcontractkit/libocr/offchainreporting/confighelper"
	"github.com/smartcontractkit/libocr/offchainreporting/internal/test"
	ocrTypes "github.com/smartcontractkit/libocr/offchainreporting/types"
	"os"
	"path/filepath"
	"strings"
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
	case CoreCore:
		// 1% / 3600s
		AlphaPPB = uint64(10000000)
		DeltaC = time.Hour * 1
	case CoreBtc:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case CoreUsdt:
		// 0.1% / 86400s
		AlphaPPB = uint64(1000000)
		DeltaC = time.Hour * 24
	case CoreUsdc:
		// 0.1% / 86400s
		AlphaPPB = uint64(1000000)
		DeltaC = time.Hour * 24

	case CoreSolvBtcMbtc:
		// 0.5%/3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case CoreSolvBtc:
		// 0.5%/3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1

	case CoreMBtcBtcER:
		// 0.2%/10天
		AlphaPPB = uint64(2000000)
		DeltaC = time.Hour * 24 * 10
	case CoreMBtcUsd:
		// 0.5%/3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case CoreSolvBtcBtcbER:
		// 0.5%/3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1

	case CoreSolvBtc_temp:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case CoreStCore:
		// 1% / 3600s
		AlphaPPB = uint64(10000000)
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
		F:                (numberNodes - 1) / 3,
		OracleIdentities: []ocrConfigHelper.OracleIdentityExtra{},
	}
}

const (
	CoreUsdt = iota
	CoreUsdc
	CoreBtc
	CoreCore
	CoreSolvBtcMbtc
	CoreSolvBtc

	CoreMBtcBtcER
	CoreMBtcUsd
	CoreSolvBtcBtcbER

	// unused
	CoreSolvBtc_temp
	CoreStCore
)

func GetNodeConfigs(target int) []test.NodeOCRConfig {
	nodeConfigs := make(map[int][]test.NodeOCRConfig)

	// core-main
	{
		nodeConfigsCoreUsdt := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xCeA72bBF5577F9BC674dFD9a128D0FF2dbD30d70",
				SignAddress:     "0xe9918A3ae893bd0A6C3241E094de169Acec82F88",
				ConfigPubKey:    "d00d003d6468e023ac1213ee099d12738b2f2ad439f7c91bcd593f280cbe3d7f",
				OffChainPubKey:  "cc77a16b27cc23d370f8b7b6bf0bafbf546d8be85b49c582cb1b88680e6284f3",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "1b23497a001352daefef03fedf2cadb41353c1a5dc6cb96882f65b956cb84ef8",
			}, {
				Id:              2,
				TransmitAddress: "0xc355159B16962609eB041Ab749E8BA337dC498f9",
				SignAddress:     "0xc3f46b31615794208d3A84141b4196F8d0353701",
				ConfigPubKey:    "f856f240405e51279c5cdce0d12997b0fbfd8e798e33ff1270a9b3814d8fa678",
				OffChainPubKey:  "e0fcc315c7915a2e16cef1c4fcef7e38c97b6b272509116d293b4d9f70cafe40",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "60fa64c7c6e0671a447eee196ff85c8b5036ae1b4af777cb372b686be9990a60",
			}, {
				Id:              3,
				TransmitAddress: "0x4AbcA6C08236776308e008dE2ECAd7C0867d4b08",
				SignAddress:     "0x1183E0eD3B8E3B722D653B5AAb4da6eE05F6E442",
				ConfigPubKey:    "7dfe0c7f1581f9c83b2e70941d5c6b2d129a7a0fb00bff403c620fb4869a5f59",
				OffChainPubKey:  "7a8af90efce269d6cd53bcd28d613e7f0cc98b01a4c433a0b7fc1c0a2e935ba1",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "475f3651af401e96e4bcb93cbe3089fec1648dc23901363802b010ef5b659f7a",
			}, {
				Id:              4,
				TransmitAddress: "0xaEC1B644e3217370E7973953aa15DEA76e15288B",
				SignAddress:     "0x9e185919f6A98EbB7ee07aE395501A5f3b98D7b6",
				ConfigPubKey:    "aa091f19921ae650aabb337c296a1d563006e695c959f8be43b839e339a7d559",
				OffChainPubKey:  "7c730b44801bdbffabe281bb5cd57572cd970609e0e11925b13fd9332548fe2e",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "76331b5d37758948f6afd16e09db86c9782bc9e359b9e34386b833aa89dae0e8",
			}, {
				Id:              5,
				TransmitAddress: "0x41fCC88c8Dc99732E81A18B3ed037e6B040026F4",
				SignAddress:     "0xF62315b4a7129f2024006f52d0362472A1bB36F1",
				ConfigPubKey:    "27860fba292dd3ccbe138dc02eb8e66a1e94e584f5111274f3397758ee2e346e",
				OffChainPubKey:  "6b0d3a9600a16c8710f920196e26a8589ab07eb4bdc019aedd72060d4ccd1780",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "626f15e0205c94b75568f1295ab31d019e38ce26b379fa6483ae421419c51b31",
			}, {
				Id:              6,
				TransmitAddress: "0xC183015730d6CB67137f2D4E7e41f04275C29989",
				SignAddress:     "0x50d6F8f6D6CD1E11259CBcdB864733684a547FE4",
				ConfigPubKey:    "5b9b5e2d4fe9a5e128c375fa17505a0876a3a03574d328cc96f0a664c0b3b130",
				OffChainPubKey:  "74d39b6dccfada9d52f810c52528478329b07e0a150a771837118f404458d6c4",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "50ff1b1eac949b7b22779ff90871990428ead778a5feb670f509dc2533b14c6f",
			}, {
				Id:              7,
				TransmitAddress: "0x1A7EEe3C1C6ff6dC44Ed7425A9F6b7A7C0E017ef",
				SignAddress:     "0xE228c58EB3eE2e3a42ecE246F6A5C38D6fb068eD",
				ConfigPubKey:    "01d3efecead8cdf0a9ab53d06474b51c7397bd79e030b2c6e56a26ceeb530057",
				OffChainPubKey:  "46e03df829e4e6371b03c90e558966805207a3b3d1571e3552e97fbf29b06131",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "930217db1bafd1243fee895df496ea74ce43f300cf4a84c720cd88ffb1838922",
			}, {
				Id:              8,
				TransmitAddress: "0x4a2005Fc56fd40Db346d4a6713D27765D7e1d44a",
				SignAddress:     "0xa39Ce1CB66553404b118Fb8Bf5bCc683A9879E8F",
				ConfigPubKey:    "a5d8cccd3705a2faffef336452861d6a247429ccaaeafced36c1788f0364fa11",
				OffChainPubKey:  "d927b1783e0d638e8d1a9c22184ae9eb0501e6dfb6b20720f80c47dcc326d3d6",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "b7b58cb397b2b9b2191e328a0c199052572f52bdfa4ef0aa806956b017256dd3",
			}, {
				Id:              9,
				TransmitAddress: "0xfE20d7AfE1edE8d064d8679747fcf00f9FBEd486",
				SignAddress:     "0xD57ca8939E2163c5ba16580557c8e307303fa711",
				ConfigPubKey:    "916a9baa5887241c28d137076e45ff4d74a2dd0ed5b3bfaf1433223e3ed3fb1e",
				OffChainPubKey:  "3c2c67c6ad81e5b2a0819ac2edbb1cf45d68389545bc3e99044dddf1d8578f60",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "e66a81afc1fce7bac1d3f952e765cb814821309255d4ddee53638ab7ccacbd8e",
			},
		}
		nodeConfigs[CoreUsdt] = nodeConfigsCoreUsdt
	}

	{
		nodeConfigsCoreUsdc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xCeA72bBF5577F9BC674dFD9a128D0FF2dbD30d70",
				SignAddress:     "0x84C119Bb4ef0e4f66489e6D9Cf04473CA45F3931",
				ConfigPubKey:    "70ffb3dca8c86a1748e370a07df2796b404ede6079f19eb91976f8499ff6fa45",
				OffChainPubKey:  "ffb1c17f4b0740dd1d0c38c666d844a83faa67faa96ad6bbe6b5090a51891373",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "64eedcbfcf641600417415a1fb9bf816d02ff52c5ac19673908ee986a807c7d3",
			}, {
				Id:              2,
				TransmitAddress: "0xc355159B16962609eB041Ab749E8BA337dC498f9",
				SignAddress:     "0x3710972261cD8Bfa52E6fc46cEc927e6AE3C4AB8",
				ConfigPubKey:    "e2420ac8b71ae5c852301c4fac72d74f7d65dafc7dde6f3a18511a574e9bbb1a",
				OffChainPubKey:  "fcf892caf7aa4b0df713b14f953bd39946b309ba366adbd2d784205a1e745b6d",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "1d419fb5697446b981e0bba200ebd804de7c60e63b014abaa0667b4bd32f4813",
			}, {
				Id:              3,
				TransmitAddress: "0x4AbcA6C08236776308e008dE2ECAd7C0867d4b08",
				SignAddress:     "0x14D97fB404f453D81FA9a6Ed902169678c6aCAf4",
				ConfigPubKey:    "787fd51cdf5c6f1ff658c782ed15bae762d97dcd19f235d5df5c413d360d0046",
				OffChainPubKey:  "4ae08d0ab089602cc51a869bc60d3c9d9baf6bd9de5e115773f85cbd89bbc2c0",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "2531342b50c50c32fc903806377b2c22ef111523893dfc6452cb9ab7b307ac63",
			}, {
				Id:              4,
				TransmitAddress: "0xaEC1B644e3217370E7973953aa15DEA76e15288B",
				SignAddress:     "0x872f4340fa1CfC3906845933947a8d057C8e1DEe",
				ConfigPubKey:    "570d401b2f6ff8af33818e01868a01344eac70232a5541fae0777502c27e2955",
				OffChainPubKey:  "d9584c5be5d70933030f05bceef5f53574fa611ea9733c58e624a7788e7fad73",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "b5fa21382c442d91bdc5ab907b3b118c4b84bd3de5b560399127e6b254b8911a",
			}, {
				Id:              5,
				TransmitAddress: "0x41fCC88c8Dc99732E81A18B3ed037e6B040026F4",
				SignAddress:     "0x4fb65140e0e9bdc4A401f5d5AD968fdbb61D4A8a",
				ConfigPubKey:    "3a88d63423ba19e859b42e62f02d4f2df07b2b31418315a0fae5d4822678dd46",
				OffChainPubKey:  "6203b062cd0b01a68ea8357aaa1b5e77add843399b30bf8f7d27f3fca65c79af",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "2465c894cbeaa008d57f9c3bf2ea148c6d0f000f60a0f63329dbac63fc3a029f",
			}, {
				Id:              6,
				TransmitAddress: "0xC183015730d6CB67137f2D4E7e41f04275C29989",
				SignAddress:     "0x72F58B5e19F249401e2EB08125d9BF9619C126D0",
				ConfigPubKey:    "208681ad09cdaf995f850ac2de53d953777c19be405618b01286d914055bed67",
				OffChainPubKey:  "fc3412d410a34ef8183e0e437222e2c762c7f7cb19314738f4b33b717df229d1",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "2126ba186ff4102e4f2fbfc73c0841153af03b7421f6162df279867f51a4c7b5",
			}, {
				Id:              7,
				TransmitAddress: "0x1A7EEe3C1C6ff6dC44Ed7425A9F6b7A7C0E017ef",
				SignAddress:     "0x7e91663d0eB534CB0037afB366BA8e608ae32c62",
				ConfigPubKey:    "66edf57e1915ed329d695ea34bd189dddb5eddcef6782f344c30959e20bb9f38",
				OffChainPubKey:  "2ca8694393d49987cade2f98f55b954ad39fb57f070a9c6906f003459e967dc2",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "9b82fbeb0ff1eb955cf8cba6960a74984281a4656a953033e9452f8fe3cdf03a",
			}, {
				Id:              8,
				TransmitAddress: "0x4a2005Fc56fd40Db346d4a6713D27765D7e1d44a",
				SignAddress:     "0x77A9ACb501dEb0100bb73eaF212c3905a211Fa57",
				ConfigPubKey:    "22f7af2dc7eb0daec74107b0070ba2583e5ae8a8a3986a444d8a1460191e1404",
				OffChainPubKey:  "c9c29e2ffa118def8ccfd4ae8b8ec650bf1ea4a00b82290a42da18677cbb3f14",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "3885a6bf25a4d991ff945f2c39a2fe0b3cc31ee25ad21c0b3a91eab1f4859c63",
			}, {
				Id:              9,
				TransmitAddress: "0xfE20d7AfE1edE8d064d8679747fcf00f9FBEd486",
				SignAddress:     "0x8F6c3f94C3b015aA161CDC344Aa3821aF97F5Fd6",
				ConfigPubKey:    "db3688f040e9d83fb3183f07546b4e1079ed5b60e623d15d2e8ee5f09a1d7e52",
				OffChainPubKey:  "be07d1567a36cf60a21990f9794bc8d9aa9a891eef52e811cab4f2e0c9fee44c",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "1f9dea93dcf1a581e0eb3b9ebdc3e1d63b4ea85035520540f31fd88b1cbfb31b",
			},
		}
		nodeConfigs[CoreUsdc] = nodeConfigsCoreUsdc
	}

	{
		nodeConfigsCoreBtc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xCeA72bBF5577F9BC674dFD9a128D0FF2dbD30d70",
				SignAddress:     "0x15535Ba93638f983F3A1a98Fe9d6A7d854f07444",
				ConfigPubKey:    "dd713ac803f731f2fb1e44fda9f094edb392a08983562dd7bd8d091235b23e69",
				OffChainPubKey:  "8916e4e1a41c724139e46580db93ab3ed31d8f2ed7123b2ea14919358f558b86",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "ebec6d7fdde23712bcfd3aee4eeaaa1d4260c719b542659ddfa6aebc20559667",
			}, {
				Id:              2,
				TransmitAddress: "0xc355159B16962609eB041Ab749E8BA337dC498f9",
				SignAddress:     "0xCEA8a49457675fF6f06DFD9e8Ad5cD32dFFa4C69",
				ConfigPubKey:    "aa92b160cd32a91d6eb4d6684835bacd22b948b25fa2c7f60c48f6c3ec8c0e43",
				OffChainPubKey:  "443a2087ef2dcac00b21d4ea551893819b28eed29f1a5d83a1d7f2e8d07617c8",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "ede89d8980c8bc6b9d83235104268b582db33a8c45a4fc81fb3c3f8f9342cebd",
			}, {
				Id:              3,
				TransmitAddress: "0x4AbcA6C08236776308e008dE2ECAd7C0867d4b08",
				SignAddress:     "0xd07BE30BB170dd9C3b253126e88287dC9bA09D29",
				ConfigPubKey:    "c6fb935957b5475ecda218291839b9ef94e0a54edae159710df5227cb510a830",
				OffChainPubKey:  "ed19bfc7b3f4f3efea2d959461c00d04826b5d326eebc1dff5d065e5f79495ee",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "9448e3e839119ba8cbac107b4149fa9a1a8d5db5ed6e7bd2316225904533a1ff",
			}, {
				Id:              4,
				TransmitAddress: "0xaEC1B644e3217370E7973953aa15DEA76e15288B",
				SignAddress:     "0xEcDe82433793bcef60E852D961f5ADcFDa341Ee6",
				ConfigPubKey:    "a38755e4ec51a36b8b3f474f44df925abb54766438ab49cd13e409ef351f3846",
				OffChainPubKey:  "a9b675470c872feb382db1e9140fb1b2083af6709b0a8913896c8bdc6940ce70",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "eea7ba1f73638b0e685480b5fb572e489bf584ff834544d1b3a517c34fa5d2ab",
			}, {
				Id:              5,
				TransmitAddress: "0x41fCC88c8Dc99732E81A18B3ed037e6B040026F4",
				SignAddress:     "0x177BF47be91310bb7c1F380F6253bD07c069D7f7",
				ConfigPubKey:    "0e9ac12c49e3f4c2718a71b8c5059f0783fff0bd592ebab57b3c077a3164df1f",
				OffChainPubKey:  "6284987f61f8dbc68036dace1f5eaa919ca32f26a8096f07f19e201804ac0924",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "cc5ca88e398174c39c0fe525f23661e3ae5d00a159177188f266da785d5aa310",
			}, {
				Id:              6,
				TransmitAddress: "0xC183015730d6CB67137f2D4E7e41f04275C29989",
				SignAddress:     "0xF032AACA42C074bB592E5a3599E27537BE97BBBa",
				ConfigPubKey:    "486d56f0408bd1e01a71d1378b044e4f3b12eaf3df2927349a8b556c58213528",
				OffChainPubKey:  "4b6ede9712a0b384261872422d561195696847824de0bfb1734e90d95f4cac1e",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "47a2fabd252b8c5a23497a6aaa577040a1335d3b47fee489ecae5b26f3c5d672",
			}, {
				Id:              7,
				TransmitAddress: "0x1A7EEe3C1C6ff6dC44Ed7425A9F6b7A7C0E017ef",
				SignAddress:     "0xf12896BC135d48Ea14E5eeC49B9dDec55e18f541",
				ConfigPubKey:    "904d35ab5b66dbad1a98348519d17712211485ddc31ed2d5ab1eb2f6c76fea70",
				OffChainPubKey:  "7114a2957bbd8c423913e52c7e05762fc2043ece688101e0c5be702843427096",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "b99b560bfbd277ad6e6a83c0536bacb1638cfac451cbccad326d62aeeab8df2c",
			}, {
				Id:              8,
				TransmitAddress: "0x4a2005Fc56fd40Db346d4a6713D27765D7e1d44a",
				SignAddress:     "0x6f3322718697203a745C87529Feb4C41A4618B87",
				ConfigPubKey:    "9a3301df2cb1b1771294ac608a8f3ea163c7d16e65c9f3410a08e3599a1bbf58",
				OffChainPubKey:  "b69740f2c1910cb5f0bae3054440f5a667dd2462e4e49f9985d042a422c7ffef",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "55fa21ce6580a9865ec8e1a8a39c065fef751d035fad6b352f9fee49ccecaa8d",
			}, {
				Id:              9,
				TransmitAddress: "0xfE20d7AfE1edE8d064d8679747fcf00f9FBEd486",
				SignAddress:     "0x70d7733832212aD056414CfbE63435a2997FEbC8",
				ConfigPubKey:    "123662d5c49fb2c39d67714b169331eb1cff1c37fc079da59a9b5bb2118de912",
				OffChainPubKey:  "3e416635528714edfe6b202e8e94296da0ee759441a411de32a7eeaa6143c490",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "c2e8f355d84d4c6cf3708083d1780a8bd5258d5701f453a0df7b78cee980c54d",
			},
		}
		nodeConfigs[CoreBtc] = nodeConfigsCoreBtc
	}

	{
		nodeConfigsCoreCore := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xCeA72bBF5577F9BC674dFD9a128D0FF2dbD30d70",
				SignAddress:     "0xfbD7D512b009C5f0A47A14a3581986939E77ebfF",
				ConfigPubKey:    "09b989ad4c5cf4abd861c5baddcc31e31d520f45762a072f77aa5cc981e3e536",
				OffChainPubKey:  "f27193db2f55ef2ebe9161c7475ea9ed2df037b26946d4161a8ba6a7e6682cc7",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "f0cf48b073312a0039336c29545b4e01f2beac448ae24297cb3ec1456c42d692",
			}, {
				Id:              2,
				TransmitAddress: "0xc355159B16962609eB041Ab749E8BA337dC498f9",
				SignAddress:     "0xA5D1218c59B95A3FF3fa8249Ff9Ce4B13d09A3B3",
				ConfigPubKey:    "62fd0288cad1a6e25ed677153dfaaccc3d90401ea8db215ab3d6c1d0bf3c1c25",
				OffChainPubKey:  "604b3c7204dafe3a2e1a63ac9f06c03f55259a7cf8e54c9681b859110a12a028",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "161cf1e29da9a60dde28d7f4b476b52e74b9798a94aa381379d222af4785d4dc",
			}, {
				Id:              3,
				TransmitAddress: "0x4AbcA6C08236776308e008dE2ECAd7C0867d4b08",
				SignAddress:     "0x145Ab977A31BEff7EE99Ca9051A6962839727b20",
				ConfigPubKey:    "35ec501933f126f935c915b1d4595cf7634390d09d13d37f082fac1e3c859d0d",
				OffChainPubKey:  "58ec96e3f86d2af8850748c84a801c9c7041f36f1b3c26f51210bad937043df9",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "3d2e5541d4febeb57ade52c4eb9b453738da7c8e9cbe8ee8c483a8f8fddf5b54",
			}, {
				Id:              4,
				TransmitAddress: "0xaEC1B644e3217370E7973953aa15DEA76e15288B",
				SignAddress:     "0x2118Aa02780BA2D41c5C0761CD1FB6a12A7d231a",
				ConfigPubKey:    "fd423c4e1b6c32e91f888367eac65bae608d30c928b58820306cbb9645927a4c",
				OffChainPubKey:  "0c6a75f67d4456a87d5588d5178bba49c34b548338d69be2d9853e56f6d51793",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "c2817d4a70efe878f9c80e13b0b3905416049d0a93a735eb2643779c8a7e4050",
			}, {
				Id:              5,
				TransmitAddress: "0x41fCC88c8Dc99732E81A18B3ed037e6B040026F4",
				SignAddress:     "0xca44EEA1047D64579Cf24cAE801A33AD28233237",
				ConfigPubKey:    "2f8aa980fda9e6ac3b1a10fe78f6561dc01a465d88ad0496258652833326601a",
				OffChainPubKey:  "0b2905eef33f6186e42877da5ef3ab086c3827a9de67760df2d8e6783fd1f9af",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "ec7e017a6dd9736f691d1caf6e7aadf5823c4a1892cd863a71d71b907afba736",
			}, {
				Id:              6,
				TransmitAddress: "0xC183015730d6CB67137f2D4E7e41f04275C29989",
				SignAddress:     "0xE955902fdb754e5B908d93807D4412F2f908D150",
				ConfigPubKey:    "0de45acb15df527cffcd0af362d736cd9e55beaf4004424f8e6a37fef1cc616d",
				OffChainPubKey:  "0cbc232bdff2b77e29b99147848afff2b440eb6f81fac44fc777c36a1e074de4",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "ce219f5962e750a4340358b80bcd050ecf9abb72e47fcaa96ac2023e7ef83975",
			}, {
				Id:              7,
				TransmitAddress: "0x1A7EEe3C1C6ff6dC44Ed7425A9F6b7A7C0E017ef",
				SignAddress:     "0x6e97F014B4CD4a7E68eA664e4a342dbdD2cCBC27",
				ConfigPubKey:    "e763f62a45d674371a8fe8abf882f3d14e2f5ff86503425cc5ff25606b01fc7b",
				OffChainPubKey:  "071445e594f503dc662d2532453ca32681dc067fc5eeb2d277a44de58873e569",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "d8b73112385f1b2a6ce3f9307bb39ea0e5e8a68b9d6c1b343a781a474f4dfc58",
			}, {
				Id:              8,
				TransmitAddress: "0x4a2005Fc56fd40Db346d4a6713D27765D7e1d44a",
				SignAddress:     "0x5fAC21EA6BE042102aD985dD6642F05586118064",
				ConfigPubKey:    "8c41da77ce4f5c4a879353428145f5a78f13a520027f802f59427c66788df406",
				OffChainPubKey:  "a136597cb95906d0c51f16af0a16f0451e23b621f88c3393093304d1cb71bc01",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "7e10c6926a1dbe44ca7cf7544ee426af580945ae73647b78b98d0f6fd0ed2f1a",
			}, {
				Id:              9,
				TransmitAddress: "0xfE20d7AfE1edE8d064d8679747fcf00f9FBEd486",
				SignAddress:     "0x7b87f22f8812dA3267C0D223EE6C1e11E9e77A65",
				ConfigPubKey:    "218cc8b35481c421262165a232c494e8108e995efcc7c0ed7a5b0999d1c71118",
				OffChainPubKey:  "81635bb7b68c198eda4ffd0cbd7a2701281dfbe34a836f6aec00a3f21242eedc",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "5cf470cfa88a96b61f696007dd3526104e1b050f40972b8fd64b98e7f8de6d9c",
			},
		}
		nodeConfigs[CoreCore] = nodeConfigsCoreCore
	}

	{
		nodeConfigsCoreSolvBtc := []test.NodeOCRConfig{
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
		nodeConfigs[CoreSolvBtc_temp] = nodeConfigsCoreSolvBtc
	}

	{
		nodeConfigsCoreStCore := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xCeA72bBF5577F9BC674dFD9a128D0FF2dbD30d70",
				SignAddress:     "0xd0621ECDD74b2d599431739b41Db1eB32832c265",
				ConfigPubKey:    "4fb38babe8473835d1c715a595a039f23fad463a51d63cb8facf2ca9c3ce7709",
				OffChainPubKey:  "9582baf731842533467fa303c6288452fefe046f2e6346164b10636b2b387c0f",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "0ad85b4b414d97e3890929d582ad771668eb6ea718bea10eba68ee6a6b6cdec2",
			}, {
				Id:              2,
				TransmitAddress: "0xc355159B16962609eB041Ab749E8BA337dC498f9",
				SignAddress:     "0xE4ADf2Df7e9dEA5cE1f8a4d3D37F1BA844AEEc87",
				ConfigPubKey:    "d9da6419e968ca1c2ffafe514c7151d0777f16c28727faacb1aee4f6f8279b67",
				OffChainPubKey:  "415188ebc18bac93ace2c36beeba26753807a030b95e0d671178e4a9e33029e4",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "1b612f1d6a4efa0c357740370ba82595824c1057546191dc1a5939905ce9887e",
			}, {
				Id:              3,
				TransmitAddress: "0x4AbcA6C08236776308e008dE2ECAd7C0867d4b08",
				SignAddress:     "0x2e04042479C0ff6c142CdFE4E4d39a85c14CE119",
				ConfigPubKey:    "ce4529196620c13418d9f2f5c13a09fd3d1744bc49027df45cb043b76a7f560d",
				OffChainPubKey:  "4e1cd253c1b773aed9b791afe2f23212d33d038986453affbb20abf4ebeeca36",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "9e1e3527557f54b31310b03842e100566cb802744d3aa1ae4783537dcd016e93",
			}, {
				Id:              4,
				TransmitAddress: "0xaEC1B644e3217370E7973953aa15DEA76e15288B",
				SignAddress:     "0xAeaE28582Abd2926888c9d3D195ECA73C36D9570",
				ConfigPubKey:    "098bd314476e81816ad817d01546f9c2ffb5fc8e60bf674631dd7e5e6832fc0b",
				OffChainPubKey:  "f6e1deb6e986a6281da9ab54ce1bf88e20d23c978c653ff4b80038d0fca29928",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "b81c235e92166c892605fd838c6f2a54de7cdcedf634fcc7e510f6034d1c3b06",
			}, {
				Id:              5,
				TransmitAddress: "0x41fCC88c8Dc99732E81A18B3ed037e6B040026F4",
				SignAddress:     "0xD230b061FA19618608251BbA118C222c5f9D3970",
				ConfigPubKey:    "74e346a32177e04eac99c2ac5c71540f9aaeb522b3188bb677d7202bfdb5b63a",
				OffChainPubKey:  "b7c4434b2f25d740b976cbfcf86a03a0d7b359de57fa501dff277716de922901",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "152f8cd35abbe829eb5e2017d47e62a2378ebb26ce83ebca9032d293fc288332",
			}, {
				Id:              6,
				TransmitAddress: "0xC183015730d6CB67137f2D4E7e41f04275C29989",
				SignAddress:     "0x681bD73380a83BC3C4ffd8B5fbA5659f507C7478",
				ConfigPubKey:    "d6049143b6e6759a385c51bf61c1d7a2f491956440e9df7977ad04a6e4a39f39",
				OffChainPubKey:  "c42ed7bc781349e85c09314597479220ab5267c40e9a2d2ed8cfe22a5f8070a1",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "d4e5679a37eda426919d8cd978aa360b91915523e291542fcd432df1a7046109",
			}, {
				Id:              7,
				TransmitAddress: "0x1A7EEe3C1C6ff6dC44Ed7425A9F6b7A7C0E017ef",
				SignAddress:     "0x3c514E66B6D1184507049dE3894E804b36091337",
				ConfigPubKey:    "8986e4851bb338f307bf7ef5a62845ec1fb1f604eb93fb0bb66301bb5d072f3b",
				OffChainPubKey:  "a1b1b4e5afcea0afe6ff9d3218f59a3b662edadd19e1543d277a75ff6a53a2e0",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "c2689046752cfbc225f5b017a6878dd6eb9e12600ee5fe72c358745896423aa5",
			}, {
				Id:              8,
				TransmitAddress: "0x4a2005Fc56fd40Db346d4a6713D27765D7e1d44a",
				SignAddress:     "0x97560fB873571B9b42F39CfAC5ED658772be4665",
				ConfigPubKey:    "125fd223bad4651e80be90d9a52da843fa79004af4dd61c656fef8b75ceb2963",
				OffChainPubKey:  "2b7f7021bd5ba215cdb41e3b50f79c1b9cef1b6c20d1f261c7f4d00f2a456044",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "58ddbb51492ff9b53e09672e688de26cb5d4d3531b4a9d3451be2efc3bed84e6",
			}, {
				Id:              9,
				TransmitAddress: "0xfE20d7AfE1edE8d064d8679747fcf00f9FBEd486",
				SignAddress:     "0x6080FeE90F7C795Fd21a55bCc16416638ea05a4B",
				ConfigPubKey:    "0526ef0f8118751e4d09ef857d3c9104ec6e23fe6d73c3a9301b452607c10b7e",
				OffChainPubKey:  "d7326bf6e7d9427436c230449e955fa492115ba5000fca0d056dbcc7d96c02d4",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "24baa484b25640a4b8df6cf576aadff1d3a6e4669e4e2c1d548954ae70da6594",
			},
		}
		nodeConfigs[CoreStCore] = nodeConfigsCoreStCore
	}

	return nodeConfigs[target]
}

func GetOffChainAggregatorConfig(target int, publicKeyPath string) test.OffChainAggregatorConfig {
	nodeConfigs := GetNodeConfigs(target)

	if len(publicKeyPath) > 0 {
		content, _ := os.ReadFile(publicKeyPath)
		var ocrConfigs []test.NodeOCRConfig
		err := json.Unmarshal(content, &ocrConfigs)
		if err != nil {
			return test.OffChainAggregatorConfig{}
		}
		ocrConfigs = ocrConfigs[2:]
		nodeConfigs = ocrConfigs
		fmt.Printf("nodeConfigs: %v", nodeConfigs)
	}
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
	readPublicKeyFromFIle := true
	publicKeyPath := ""
	if readPublicKeyFromFIle {
		publicKeyFileName := "publicKeys_mbtc_btc.json"
		publicKeyPath = filepath.Join("/Users/greason/Documents/workspace_bitlayer/chainlink/apro.configs/coreMain/publicKeys/",
			publicKeyFileName)
	}

	ocrConfig := GetOffChainAggregatorConfig(CoreMBtcBtcER, publicKeyPath)
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
	fmt.Printf("\nerr: %v\n", err)
	fmt.Printf("\nsigners: %v\n", PrintList(signers))
	fmt.Printf("transmitters: %v\n", PrintList(transmitters))
	fmt.Printf("threshold: %v\n", threshold)
	fmt.Printf("encodedConfigVersion: %v\n", encodedConfigVersion)
	fmt.Printf("encodedConfig: %v", hexutil.Encode(encodedConfig))
}

func PrintList(content []common.Address) string {
	txt := strings.Replace(fmt.Sprintf("%v", content), " ", "\",\"", -1)
	txt = strings.Replace(txt, "[", "[\"", -1)
	txt = strings.Replace(txt, "]", "\"]", -1)
	return txt
}
