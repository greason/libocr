package bsc

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
	case BSCBtc:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case BSCEth:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case BSCUsdt:
		// 0.1% / 86400s
		AlphaPPB = uint64(1000000)
		DeltaC = time.Hour * 24
	case BSCUsdc:
		// 0.1% / 86400s
		AlphaPPB = uint64(1000000)
		DeltaC = time.Hour * 24
	case BSCBNB:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case BSCstBTC_BTC:
		// 10% / 10天
		AlphaPPB = uint64(100000000)
		DeltaC = time.Hour * 240

	case BSCFBTC:
		// 0.5%/3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case BSC_TONBTC:
		// 10%/10天
		AlphaPPB = uint64(100000000)
		DeltaC = time.Hour * 240
	case BSC_TONDOGS:
		// 10%/10天
		AlphaPPB = uint64(100000000)
		DeltaC = time.Hour * 240
	case BSC_TONNOT:
		// 10%/10天
		AlphaPPB = uint64(100000000)
		DeltaC = time.Hour * 240
	case BSC_TONTON:
		// 10%/10天
		AlphaPPB = uint64(100000000)
		DeltaC = time.Hour * 240

	case BSC_ordi:
		// 1%/86400s
		AlphaPPB = uint64(10000000)
		DeltaC = time.Hour * 24
	case BSC_sats:
		// 1%/86400s
		AlphaPPB = uint64(10000000)
		DeltaC = time.Hour * 24
	case BSC_bitcoin_puppets_btc:
		// 2% / 14400s
		AlphaPPB = uint64(20000000)
		DeltaC = time.Hour * 4
	case BSC_nodeMonkey_btc:
		// 2% / 14400s
		AlphaPPB = uint64(20000000)
		DeltaC = time.Hour * 4
	case BSC_ordinalsMaxiBiz_btc:
		// 2% / 14400s
		AlphaPPB = uint64(20000000)
		DeltaC = time.Hour * 4
	case BSC_quantumCats_btc:
		// 2% / 14400s
		AlphaPPB = uint64(20000000)
		DeltaC = time.Hour * 4

	case BSC_FDUSD:
		// 0.1%/86400s
		AlphaPPB = uint64(1000000)
		DeltaC = time.Hour * 24
	case BSC_LISTA_USD:
		// 0.5%/3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case BSC_FB_USD:
		// 0.5%/3600s
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
		F:                (numberNodes - 1) / 3,
		OracleIdentities: []ocrConfigHelper.OracleIdentityExtra{},
	}
}

const (
	BSCUsdt = iota
	BSCUsdc
	BSCBtc
	BSCEth
	BSCBNB
	BSCstBTC_BTC

	BSCFBTC
	BSC_TONBTC
	BSC_TONDOGS
	BSC_TONNOT
	BSC_TONTON

	BSC_ordi
	BSC_sats
	BSC_bitcoin_puppets_btc
	BSC_nodeMonkey_btc
	BSC_ordinalsMaxiBiz_btc
	BSC_quantumCats_btc

	BSC_FDUSD

	BSC_LISTA_USD
	BSC_FB_USD
)

func GetNodeConfigs(target int) []test.NodeOCRConfig {
	nodeConfigs := make(map[int][]test.NodeOCRConfig)

	// bsc-main
	{
		nodeConfigsBSCBtc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xEf3397302D05b2EC482FA52F53b1366D617EFBDd",
				SignAddress:     "0x642886FF7E78e1ED24a0B1DE797fd1A523259264",
				ConfigPubKey:    "3b836443fd8d98d46e6c3c13d065e5222f70c7031883428e528116d5dd4b7f45",
				OffChainPubKey:  "629bf2e1f42bbfcd6898c19e806ddd35addcaf2423c6150be6acaf85975ba837",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "b6acd607f32996684b4c420f7d945ffb929a5bff41fba0df7794aa838e509fac",
			}, {
				Id:              2,
				TransmitAddress: "0x958FD347C4e336fB2ca58EdDb31fC51cf3d8B525",
				SignAddress:     "0x7D05DC390a650A903110eE7dA4a386D7F031644e",
				ConfigPubKey:    "6609dd2fb24f7f321ec293a2d700578741d6202fca9af2dd759ec22aa9af4a06",
				OffChainPubKey:  "7941ad0d6e28d4c08676bc833158b469092cb9133ee3855a83c244fb25123bcf",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "a228f4ba25f6017caf041e019595f27cf75330aded9edbf293f5bbc96e2cc88b",
			}, {
				Id:              3,
				TransmitAddress: "0xDE2fe1Bdc26714d757881023b4Cb18e6607F3816",
				SignAddress:     "0xCa0C82c86D6d0a16A4B8DCBa09c71f160223998F",
				ConfigPubKey:    "49a140cb025a14365fa715f0d7207ac66907166563bdba37b57bfb84a8a0e44c",
				OffChainPubKey:  "ac6cb647262658a8e6da884fc77bba874da0c4a4ffa23fb6854996a3d7c328ee",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "42945c416536c88fa142f7fb65a8d25b19cc00dc2375e6daeaadda1e79134e7e",
			}, {
				Id:              4,
				TransmitAddress: "0x17d49f2f839C331Ae8ca6CebEC67A58AB7b2Fff2",
				SignAddress:     "0xd96f062bfed32ECAb8228051DCb6C0b705183772",
				ConfigPubKey:    "afc99779266e1dd45e4042bb91f3d4027cf8f2ef2e23758db60ca60ca5d07a4e",
				OffChainPubKey:  "3d1517d9b4ad6d2e3d2585d5a91ad220f4853be176967730f1e4defc7ce30e50",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "b9a1c4885a554b397f38ce23ef7107a2a0b3384c27572a85b4e65960c601f67a",
			}, {
				Id:              5,
				TransmitAddress: "0x6B8f6f0B13293f79Fa9660F0826e5318955C3B8F",
				SignAddress:     "0xcBEd26EF9d06f8Ff00fBF4dDC3E6b6D6dc6E731a",
				ConfigPubKey:    "fa9d4f8d52564e0de9f97e76060683978dc3854d16dd09dc07e42ef0735c9135",
				OffChainPubKey:  "19f9103e84d6faea04cdea50164ff031b9b1fd45d97170940f579202dda419b0",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "d9ac99fd450fa05646c97dca2e90ab48a7f15d25a82a028fe610bd467db7b4b1",
			}, {
				Id:              6,
				TransmitAddress: "0xA59bd950bd3f3C03a4d8d7ea7Ea4Efd7eb069979",
				SignAddress:     "0x8c5558E31c3524BDCc7Ad41efcE85E5289B71C9A",
				ConfigPubKey:    "8923e1adabe4f89185524d395e083a225d3cb1c462d2bbde7b943c0551ba6070",
				OffChainPubKey:  "b8a7a697ba2ad00ddd1ceb4b77cf317afa124af90a6896070b9460b16653b51f",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "075fcd26d6a55826c50c59459e7ec9d1fa9315520450b8d02040f7bba31b6e02",
			}, {
				Id:              7,
				TransmitAddress: "0x3BBFb8dE877FF256881BfaCaB503Cbc50b447270",
				SignAddress:     "0xC71Bdaff7973b77Df4f642166fAE01E7b9B420E1",
				ConfigPubKey:    "dc9bd768cbdd853cd6ad19e1dd136cf1c6bb49189568d4bb0f13365a03806b20",
				OffChainPubKey:  "a30a5241e09266c4aa3d97442c2e629ebfc68946c1cf089482ccbe5d98ff736b",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "f6cc4d7db0e925a0c2f0835f03163c0a48c167e79cb731cacf108b4a3e9155da",
			}, {
				Id:              8,
				TransmitAddress: "0x98C67459A03d25ae98781fDDc8e69B05b5d188C8",
				SignAddress:     "0x3674fa953c8e7891B58A765c6716a48bf446C16a",
				ConfigPubKey:    "b5970ec4d98b0ffce8e8c7f1827bf5c351c458394a747fe47f49c025e783294d",
				OffChainPubKey:  "f418fbda9bff453a99515077fb5487fd12d122dbe3ef3d09f4bd01fffc6d1d07",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "07a67ed3550d8f398da05219209acd3d7c7abb706e19a1e8376eef4cde897ca2",
			}, {
				Id:              9,
				TransmitAddress: "0x8cE54324473A2ee30956F2d7473aBBB7455078B1",
				SignAddress:     "0x65f0fCAdDAcCFfaD2AD033406c38Eb817fBeFb90",
				ConfigPubKey:    "ec99ca588384b207d430c7169eed59c57f54f7b48100c7416074bfba5840963f",
				OffChainPubKey:  "e12ada8d74ae2c05515dc610a915e6e67a2a5335bb06b1021285da1ab709507c",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "742edf0230029f3da7cec7734b69a24106dc3c9af29f757701e38425eea7a0f4",
			},
		}
		nodeConfigs[BSCBtc] = nodeConfigsBSCBtc
	}

	{
		nodeConfigsBSCEth := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xEf3397302D05b2EC482FA52F53b1366D617EFBDd",
				SignAddress:     "0xf09d1D0928392Ac2377f1d208BDb47d0d73F919e",
				ConfigPubKey:    "23a6c0e337497377bfb6f4afcc4f44bf9e2e146275fd20d456d81a9862cb6510",
				OffChainPubKey:  "442d0b9d18183a8e782e97e58ee4335706dd40534e5d52cd7c144831e303e4f1",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "10284eb639449fb46c74d15ffcc574626fe8c8f2093e7624a42ac50b49f7a75d",
			}, {
				Id:              2,
				TransmitAddress: "0x958FD347C4e336fB2ca58EdDb31fC51cf3d8B525",
				SignAddress:     "0x554ae8e566449EB028cF5534ca2614366a3CdA6c",
				ConfigPubKey:    "593d1d9aa2592591a6c412b9737f3a25de218826e097367d3f2bb3b8d9f81551",
				OffChainPubKey:  "3f78fecad7b1378755eb6e40c85557e6e4da43c85dc783c4fc4e0130876afe5e",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "338227c7184a65028d328261f8ed5ad1961d7099167f8fa90bdde8717e3d69c9",
			}, {
				Id:              3,
				TransmitAddress: "0xDE2fe1Bdc26714d757881023b4Cb18e6607F3816",
				SignAddress:     "0x881364ec5622E8A8d422C5AE35BB99A77B84C698",
				ConfigPubKey:    "22e5ea3fffc3d740addbbae392f266a6527109e9e19e307fb89064aa2e5cfb71",
				OffChainPubKey:  "10ee915dcdf166db927d941066656e169535862e0c99e85ae3af8a03e2827779",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "0d1db91f401b8b2112be8266066d6867ab22d0e5d49adf05c5c62dc6e3dcbbde",
			}, {
				Id:              4,
				TransmitAddress: "0x17d49f2f839C331Ae8ca6CebEC67A58AB7b2Fff2",
				SignAddress:     "0x63e247c93f55999A8f90fAa6691B92A7c84C3e2D",
				ConfigPubKey:    "bd0099dfd08a15768fb0d88bc4b6e205e3b6fa6ea68e73e9e16fd8cccb963b66",
				OffChainPubKey:  "a0fa1fd8e832051405a21fe2af207b71b027ffe9880d00c910e95d966757ea8d",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "20d4ad1058f653a4a948c83d730056983fac3659094a0aacfe14043a1748bed7",
			}, {
				Id:              5,
				TransmitAddress: "0x6B8f6f0B13293f79Fa9660F0826e5318955C3B8F",
				SignAddress:     "0xab2A16d68d44C5B7376ee096337EBaA76f49ff79",
				ConfigPubKey:    "c47aa7f80eba813568e041019a52b10b3064749cd9417e5272e658a5c3523b25",
				OffChainPubKey:  "ebe0ea11003f93cc1f8cb379b0780d1860c5e47345d5afbf7f45dd1da5775056",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "d20579c9b0dae046b7b7a92dc116eedd443396695dee844af93529ba920e6e93",
			}, {
				Id:              6,
				TransmitAddress: "0xA59bd950bd3f3C03a4d8d7ea7Ea4Efd7eb069979",
				SignAddress:     "0x5e148A5fA8c47BD4b6950De49914aEA04C0bC957",
				ConfigPubKey:    "a4cc70f058acdfad20fe989e730171f1fc9e8ef9cce783fab054012f819e0203",
				OffChainPubKey:  "928c027e97219d404ba5a0384510ea1f4776447844ebc518981778df74822bf6",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "de14c452904318bc92f9b2887ce0f3ac0428b037ab04d8de84aeb22f228ee276",
			}, {
				Id:              7,
				TransmitAddress: "0x3BBFb8dE877FF256881BfaCaB503Cbc50b447270",
				SignAddress:     "0x93492Da4eF1F92c3E37902DA00c547E995F84008",
				ConfigPubKey:    "9f13603564b3c1dbc84c2e891aab617a99d30a236b1b28a90c160432e5b89e63",
				OffChainPubKey:  "732eace241ae4c9ca41fa3a32b671cee8241475d10a10758ae2faa3ab207c347",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "030f128e14b0ef223c47cc2b3f9c555bc36554af415114cc36a6a4050ccff365",
			}, {
				Id:              8,
				TransmitAddress: "0x98C67459A03d25ae98781fDDc8e69B05b5d188C8",
				SignAddress:     "0x310aD03dd6f5e9f3F215Bd56A7c8BBffE242f587",
				ConfigPubKey:    "c03a58db5382ba0789d37426312dc5c46e7216998e8c7aac43ea4b26f7944a37",
				OffChainPubKey:  "0a78215a92d58c33b7b59cbddae56b65b3a756a69c26dae21d8dd0ce6cec719f",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "66a30d82953f313a8c6cf4f9a7652636f90dff75c2dae24c71b5c744632a0bb7",
			}, {
				Id:              9,
				TransmitAddress: "0x8cE54324473A2ee30956F2d7473aBBB7455078B1",
				SignAddress:     "0x7F53E7B6F63BbB227409ED48A0620C82b48564A5",
				ConfigPubKey:    "46c4290798ba5b25b51f08d7aeaae8c497dbe9cf794f6adaa77d9dab40d93a08",
				OffChainPubKey:  "4eaae4e8132722542940340a4db9c66077e084fb1b99306baed3143c028a4532",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "33c00cfa4176719d7408d1700d8428178258373e756995b06413f234b1d1b46e",
			},
		}
		nodeConfigs[BSCEth] = nodeConfigsBSCEth
	}

	{
		nodeConfigsBSCUsdc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xEf3397302D05b2EC482FA52F53b1366D617EFBDd",
				SignAddress:     "0xA9aE11cF90E236C848F86a162DBc3570C230ce52",
				ConfigPubKey:    "00b20a4a5c2b7bb480c743aaac0f57ca4ed5c1fa348f04e48510fd0994b64207",
				OffChainPubKey:  "0e4ad30dc83ee0f41fd942e5101422523b7d531872b228db587fc1bc7d84c760",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "5eb623c6bf55589ca6fb3170e0752dc8fb3055273b76c28be50a923737471491",
			}, {
				Id:              2,
				TransmitAddress: "0x958FD347C4e336fB2ca58EdDb31fC51cf3d8B525",
				SignAddress:     "0xCAAA60e81D2e84ae652C147200209FE29380527a",
				ConfigPubKey:    "9e0bd5d4e1bdaa9fe85d9b99b1fbc379e20d8e9243b991000ce22f7cf3cbd91f",
				OffChainPubKey:  "2dc6f534b6801d148116d0568df221e57d28549e7392df7a4985810068e50d2d",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "5ff76a8dad6dd5183809d735f5b41fa4b2d337939aedddda694a6d2c0d845cc1",
			}, {
				Id:              3,
				TransmitAddress: "0xDE2fe1Bdc26714d757881023b4Cb18e6607F3816",
				SignAddress:     "0x82c2Ce043C2Eb3D1EB1a0cE6E6c966617a951D81",
				ConfigPubKey:    "5cf0f05c7f12ad7af8d8a4a516aa729ddf1326f4554bc4382b7eaf4123b7b04d",
				OffChainPubKey:  "adba4b9132c7a771743bde169b84ed5011f29efca98537de698d0c4a2e61e648",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "ce09aa503cb3688e678620f93e2010d7652703199261fcb9f3e6e59229f80e43",
			}, {
				Id:              4,
				TransmitAddress: "0x17d49f2f839C331Ae8ca6CebEC67A58AB7b2Fff2",
				SignAddress:     "0x445c87c9fB63AC6BFa183f411b587f71b9d9b5Fb",
				ConfigPubKey:    "67f4e1260ccd976893f675455f1a90cebe6eb196b8fd3f3532e3af6137433051",
				OffChainPubKey:  "f7ea200d32990ab254fb121823500fd4b18a26d0ffa65de42674d18d71835084",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "2aaa8edc6ce3484ff86122bd7146891dfd025c421736ece368aee2f9976a4c37",
			}, {
				Id:              5,
				TransmitAddress: "0x6B8f6f0B13293f79Fa9660F0826e5318955C3B8F",
				SignAddress:     "0x650363825D3c8F7D55c482873d242181815CD090",
				ConfigPubKey:    "2c485451db70c94ad108c721782e8c85d0099edcf05f5140f820517733878a4a",
				OffChainPubKey:  "9674132e9b4abfbd1ef6b5a784809b20ea70f484e41b22b9e91268655eeaa613",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "f195e547b9ca47afd1a5dbdc34f371abb64d1164647a0c1657ed8424b9a3fadd",
			}, {
				Id:              6,
				TransmitAddress: "0xA59bd950bd3f3C03a4d8d7ea7Ea4Efd7eb069979",
				SignAddress:     "0x873799ed9520Aef9Dc9E04d624b6a414ee445c21",
				ConfigPubKey:    "34645e6d6270e3408ada7694174ce96fb7fc67dd0118a30c02360d0746d6b421",
				OffChainPubKey:  "4ddf24ef13e0ddc18dd8e694bb4d7eaaad6d580b125212d2716090dd897d576d",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "6e303f368bff39a86f4f4354b5d78e38f868357360490eef28e43730fb6332c7",
			}, {
				Id:              7,
				TransmitAddress: "0x3BBFb8dE877FF256881BfaCaB503Cbc50b447270",
				SignAddress:     "0x5520005C9F6B0cdfA28a910270cB616c44528bf0",
				ConfigPubKey:    "03e72a93e3fcfe0b11f900c94986c4c80255b644a8d5f6ae0828041a344d2c02",
				OffChainPubKey:  "cb4f7d0111165f04f59993244c52dd276c4cdab66784eed467d4c25b6a5b7491",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "7a92f20766ebd57a818691f0097f07ef2a52af76b281bb7fc2c6c8944d4f7436",
			}, {
				Id:              8,
				TransmitAddress: "0x98C67459A03d25ae98781fDDc8e69B05b5d188C8",
				SignAddress:     "0x078EC78d108f029D07610475181C952602b1D6ff",
				ConfigPubKey:    "5636c906b63acf8cad05f23785ca79250ce9160d29ebe14dbcb1a162dc6c340f",
				OffChainPubKey:  "3032e4041f27e811e5c4aa013d30d40261fbe487cda815edf2e2e961cd72c391",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "b07581b1e8aaebfbadc5e872e0e3915a7e35d505e98b2bcf6168bd228e1f71f0",
			}, {
				Id:              9,
				TransmitAddress: "0x8cE54324473A2ee30956F2d7473aBBB7455078B1",
				SignAddress:     "0x11A673210c96448Ca15e0fd75d48b6B27a00a9ff",
				ConfigPubKey:    "ac75f33bca94e656810450add66f9205ccf06faebb600c46259a5f8677c3c615",
				OffChainPubKey:  "e0a5cf67cf28d451c9e8f0b1701583e08199366cb6dba8bb4512f18247567f5d",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "9e31c7f62587fdd3166842eb9097e9b2da84f56ac68925ab15370ae3b35cff89",
			},
		}
		nodeConfigs[BSCUsdc] = nodeConfigsBSCUsdc
	}

	{
		nodeConfigsBSCUsdt := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xEf3397302D05b2EC482FA52F53b1366D617EFBDd",
				SignAddress:     "0x1Ab100428daaA0Ca5eAb1a7A6F11902e83d3D07F",
				ConfigPubKey:    "c1ebeb8a5b235f53591cb167c2b774c9295f1fa3fb751efe9b99ccc58a684919",
				OffChainPubKey:  "c6920f81dfc5cfb6716a8a90af80879b4e471c9ebf4db9cd116156514fde05f5",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "0e7645b640a1ff4bf5b1466a7c16b87973fbfbf2fd369c657c3d7d9d1c255458",
			}, {
				Id:              2,
				TransmitAddress: "0x958FD347C4e336fB2ca58EdDb31fC51cf3d8B525",
				SignAddress:     "0x314f45186DEf0858a8f9Fe4A93DfBdBe9d14917F",
				ConfigPubKey:    "0c12ebe9fd927bb4a9a5c4f705e496d593e44faeeaf4926f035fa050543e8d15",
				OffChainPubKey:  "4aff9ffe330c0817bea267547a55246bbc46e4e3dbf89f768421a8b8424c9865",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "4cbf6c27f4d3883c21481ced3a4a4f24147ef8fdd23e839309228ff34f296590",
			}, {
				Id:              3,
				TransmitAddress: "0xDE2fe1Bdc26714d757881023b4Cb18e6607F3816",
				SignAddress:     "0xC73b96D861894d1F597Df8429dc9039b46f2416b",
				ConfigPubKey:    "366d8fd10ad6410910dd1898e81f1a014f00ae20add0c763fe27e56baf1c6940",
				OffChainPubKey:  "c1c954a1717db403b7809d8158fabbda54a5cd64b3c1eebf2abcc461d5560f10",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "0a940e292b556bd29a515b9e2efb5c931bbbab64e3d85395b3f327f20754195f",
			}, {
				Id:              4,
				TransmitAddress: "0x17d49f2f839C331Ae8ca6CebEC67A58AB7b2Fff2",
				SignAddress:     "0x639BC71090786ae94555d974F7133257509B782B",
				ConfigPubKey:    "c958532b9e8fc973212d73f57c50c22b575badd15ba2107f324dd85d845f571a",
				OffChainPubKey:  "f8a2bfd9119965fcc9799f10f28662aa53ba2e21d5ad9bab1ec7c5472a2255f2",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "35710032d177a4df8a732ec70e33dcb705cd6e4890bd92d218822be6c854469a",
			}, {
				Id:              5,
				TransmitAddress: "0x6B8f6f0B13293f79Fa9660F0826e5318955C3B8F",
				SignAddress:     "0x850AF4719E9F983F7D97994591AB6159e760471D",
				ConfigPubKey:    "9c943448e92d6a4df769d75bd6e397bf9b172b46d655948f38c67b3d7740453c",
				OffChainPubKey:  "c10ff8d564df9bc46c846b1760316e5708bc9a7ddb31134fd0ecc04b074b1c6a",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "17b458e030a550e29b50a8129293ed27371c11195b66f7260cafa968578a5a2d",
			}, {
				Id:              6,
				TransmitAddress: "0xA59bd950bd3f3C03a4d8d7ea7Ea4Efd7eb069979",
				SignAddress:     "0xB81a5a9f214010cA224E107e01C87C2443b3A9C7",
				ConfigPubKey:    "e256c9f03e43bf4fde66743361356e079eed9829b655a0bca70590e0eb672559",
				OffChainPubKey:  "5e0686895f961bba82c13eedd0e56129d0cc8aff194c8ac8e71ee36eb564d619",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "d432b9e9872c9114ad465b48139d98bfe7a81ac10023251b411f7ffa55eee936",
			}, {
				Id:              7,
				TransmitAddress: "0x3BBFb8dE877FF256881BfaCaB503Cbc50b447270",
				SignAddress:     "0x74235854D48cf4b4696Ef5FC124044A86c9c2481",
				ConfigPubKey:    "9762a59158cb2651b51599a7497c630ae46d187e9cbab3c7045616a3f6d3b85d",
				OffChainPubKey:  "ac8c3822d7c4aa8915bafb92555cfbf4f04609724d08c1f67e0f680e72f4f725",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "0f23d434cb76bdac2f652c62a3714c19a8adc17fe0ca15e1a17c26b1f3db99dd",
			}, {
				Id:              8,
				TransmitAddress: "0x98C67459A03d25ae98781fDDc8e69B05b5d188C8",
				SignAddress:     "0xdA9bF1b347E630Ea6B449c6CE1245224D8b70Df9",
				ConfigPubKey:    "f7adaf47164eeb67db38703501696ed983405b2fe2c80258c714c9ed59f6884f",
				OffChainPubKey:  "6ad6fc4a151b98abdc7a5d339f0472dc53524f0418bd04a4717793cbd2985a64",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "1de1e28e76ecfe38d13cc46fa224bc6b2bebe79c27c80be89004ff5a440b61a5",
			}, {
				Id:              9,
				TransmitAddress: "0x8cE54324473A2ee30956F2d7473aBBB7455078B1",
				SignAddress:     "0x63b6D1B82830F24581E6cd6862143a3b86a6795F",
				ConfigPubKey:    "cd75bf4579fe1a98e49cee14af6c3681974cfac463df112006878721c6da1378",
				OffChainPubKey:  "cf1eacb23ef7531ffbbd3449caf292aeab9a76064686f00a21eac2a237c51606",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "2c3d29263ccc33c7f98bf1d47ea12b3b487ea6ef284f7c1c7f9ca758f715ed7b",
			},
		}
		nodeConfigs[BSCUsdt] = nodeConfigsBSCUsdt
	}

	{
		nodeConfigsBSCBNB := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xEf3397302D05b2EC482FA52F53b1366D617EFBDd",
				SignAddress:     "0x3B0Bd9F82225f72EE68375969afb916d7974F6ee",
				ConfigPubKey:    "e4b8c6b7afd81424b2405082c17ab2e442c13bf2db00b7d0073511c156408d06",
				OffChainPubKey:  "0d45e86206d14aacc618ffcf0a669be2b17276dde3d09d205667a10ff0e61500",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "7f828778fb50a4a4af60174179cfd3b7c79d2f02b62571aa98971704612f86c6",
			}, {
				Id:              2,
				TransmitAddress: "0x958FD347C4e336fB2ca58EdDb31fC51cf3d8B525",
				SignAddress:     "0xeC88126ABF61987446154394D1af3BCB0CB0Bd0C",
				ConfigPubKey:    "3885186645cdb63c710027075b9aeb660cf9cfa2bf3fa0953be8e5e0e9f83420",
				OffChainPubKey:  "241b758f9bf9839dcb2d8aa69cba6c445a63cb549316d49d31461b968d495809",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "a7dae5f5014cc130e61ff374efb0875eded401d736d27bc28e626529de9bf061",
			}, {
				Id:              3,
				TransmitAddress: "0xDE2fe1Bdc26714d757881023b4Cb18e6607F3816",
				SignAddress:     "0x7DCbc67B107Cf82972b9dABf152C8fa429Bd181B",
				ConfigPubKey:    "a21e03aff17e47a49e4d4a1ec3ad6c8c74ef8cfce99d54507bf02b76df292b60",
				OffChainPubKey:  "8e2774dce23c29148bfce144b9dda34c5bfd8ee00863b5c9e3b1d154b2bf0a07",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "affebb33dce014f5221fb75b06e543952ff771392bc1fc5d20310f4e574875fc",
			}, {
				Id:              4,
				TransmitAddress: "0x17d49f2f839C331Ae8ca6CebEC67A58AB7b2Fff2",
				SignAddress:     "0x28E629d19e5281d5CD8c471F20dC637df30a56a2",
				ConfigPubKey:    "9370eb997b2dc275933a859ab0b91938bc1b6e42df64711ccca959145df01e7b",
				OffChainPubKey:  "c8a91aac99e43b344498cdd1ae5ac6229fa5141e634a045672d153a114f6bc71",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "efdde2f29c7e9f7d10eaf2fab8aa1878ead47b25a028035631cac231b54ecc9f",
			}, {
				Id:              5,
				TransmitAddress: "0x6B8f6f0B13293f79Fa9660F0826e5318955C3B8F",
				SignAddress:     "0x6dD3Fe84f73219A0a99Acb5aa67e05Ac98e04a1E",
				ConfigPubKey:    "15c87959b54a2d43f1ccf36106f88c922e897931d131a23f9ccd96d7a6c63565",
				OffChainPubKey:  "1ab81e7ae896228d102190a32657058fed1518e5020866b651a901d5a0ac6bfd",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "11f9dc283107923ff46978a3d2e364ade2eb12f46ac9f1216d5e8fec28bc00a2",
			}, {
				Id:              6,
				TransmitAddress: "0xA59bd950bd3f3C03a4d8d7ea7Ea4Efd7eb069979",
				SignAddress:     "0xD902428d6e4Df9413f521e41a65eDC9905a4588e",
				ConfigPubKey:    "e43426657270c5b6c1d4c49d4d2a4b74c1ea2034945df0dd5f7d4f19efbaf336",
				OffChainPubKey:  "0ffd1b1f81c8506808b1f8387004909ba1e4fd6e0e9288e3bbc8217f87edc6b3",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "24853caffe890b6c16ae041b07bb5a08e6cac22f41d3cfc152858b3301ebe477",
			}, {
				Id:              7,
				TransmitAddress: "0x3BBFb8dE877FF256881BfaCaB503Cbc50b447270",
				SignAddress:     "0xcB55a2c4B3c4E5200C7ed00fBbC44dB8ff28E3A6",
				ConfigPubKey:    "ef0e958a6b6b6d72585b4e0acd65e73e7d6a3e37f18f8c5b1b6387eb488bff03",
				OffChainPubKey:  "0b6cbe1da35aedd1880571cb203ebda3cc4a45e23dba121775857d5292fb740f",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "be326fc2cfe787b874ef45287aeff759c185dd6bd32919552da9306936adec12",
			}, {
				Id:              8,
				TransmitAddress: "0x98C67459A03d25ae98781fDDc8e69B05b5d188C8",
				SignAddress:     "0x4791B97439C46f3DF949673a53827DE9B556e750",
				ConfigPubKey:    "b022d654d58b9bb94f344437aafc92a71168f4e83cfd8b1ffb60d13b54af281d",
				OffChainPubKey:  "b9eb2d750b013c18968ed3b86ea683b7e5e5856f022b1c3db758a390449c5ed0",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "485076ad20bcb6c343aa0a1123ad6335c2c496d572ceee6ce90daa0a01a62fbf",
			}, {
				Id:              9,
				TransmitAddress: "0x8cE54324473A2ee30956F2d7473aBBB7455078B1",
				SignAddress:     "0x46915A359A70aAbb862D5fF3A04148034f2458Cf",
				ConfigPubKey:    "047182cccefe00d4a75666638580d03cefe86c600fdc0e0fc024885754267622",
				OffChainPubKey:  "623f64a90b92af339397c3126a20f141b8b6c484cd31e717f7ac88420b813307",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "ccfc4e70d088dc8925994d63578692c5d1cf4c4ef6e934ba8304e23cbbc5fef9",
			},
		}
		nodeConfigs[BSCBNB] = nodeConfigsBSCBNB
	}

	{
		nodeConfigsBSCstBTC_BTC := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xEf3397302D05b2EC482FA52F53b1366D617EFBDd",
				SignAddress:     "0x7D1E62cCAa85830c70295B4E06ccF698F9dc95A3",
				ConfigPubKey:    "0b9e783dc177a03b9fc68f3e78c7f13d356b5a3777dbd065e95adab97c6f4e26",
				OffChainPubKey:  "92f0e689794bf6ac1ce03c2e4f4ed1d578d90e6db2380f378588ff77c9eab59f",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "b888399df85d25f1f6591059748ec393fb8eacd351d1f295eab458467bfafb7c",
			}, {
				Id:              2,
				TransmitAddress: "0x958FD347C4e336fB2ca58EdDb31fC51cf3d8B525",
				SignAddress:     "0x28523A7639fdf599A0ee2FB8175E689De78aca29",
				ConfigPubKey:    "6c96483c164e47f07cb203fee1418642fd41f9d73d80902cd9e70edb54a8a669",
				OffChainPubKey:  "af6c56ee9fd1a06bd982bfd9cdb2ecd9c3706ce74a2be2e338a97f80c3b21c44",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "d681721dd231b71d04bcca27e5ca6986b53a190f3ea7f9ea5ce4c300de509e3f",
			}, {
				Id:              3,
				TransmitAddress: "0xDE2fe1Bdc26714d757881023b4Cb18e6607F3816",
				SignAddress:     "0xF36a02Ed88D899feF79bCa2d1C3A780a0CBFF792",
				ConfigPubKey:    "3499088efdcb09101ee729509b12b7ad96a9443ee7edc758627dc2ef4fb39610",
				OffChainPubKey:  "f1f383233b29366e1be7d8ad12d23c75802b7976c12b477b364f57f791621814",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "5ad3b205b45e61d1da205a2ba7d3a9393ea632108ff1494258bff28de1977654",
			}, {
				Id:              4,
				TransmitAddress: "0x17d49f2f839C331Ae8ca6CebEC67A58AB7b2Fff2",
				SignAddress:     "0x2d75d5D4Ae3D77A6559729e23F4e4DcAAa958281",
				ConfigPubKey:    "b0954cad8e6c04d41e0145d88521b859b1b4993810e6ccb58c55240c85e41b08",
				OffChainPubKey:  "b9af9a890508ab3edf8e6d2d6436cb97f636b331593f10ef9da91df442707275",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "0d079774bfe547d1e1f43af093b0e1a4c75190310cd562877caca484006d9d5d",
			}, {
				Id:              5,
				TransmitAddress: "0x6B8f6f0B13293f79Fa9660F0826e5318955C3B8F",
				SignAddress:     "0xDde5d8Ce1218393667Bd924195EC2988CA19767A",
				ConfigPubKey:    "f2ca9bd0592348a2fd2d1ae3409f24b2adc54149e89825ff99bea4de8e72fa6b",
				OffChainPubKey:  "a71640e2b5f3cdab3521a680b9413de7c46a634beca18918c1c0415eaeb3650d",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "1a1a4471b5b99b08f61eb37482d109408e4d0071cea4edda367f776d95f9b3e7",
			}, {
				Id:              6,
				TransmitAddress: "0xA59bd950bd3f3C03a4d8d7ea7Ea4Efd7eb069979",
				SignAddress:     "0x4157f74DCd12ecAeb84ba53417Df7554cc0339Eb",
				ConfigPubKey:    "f57d48dd670d4befd4d6c51751b8051019c5d497efb0756a79b776589d02bc78",
				OffChainPubKey:  "4c04084c2089fca555aa6787df3bc4d6e61d128b40825a842ec85e6529a1b746",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "86726464471a1cbfe1811571916be6b096714d62f274e1aedd8a553b8a32bc43",
			}, {
				Id:              7,
				TransmitAddress: "0x3BBFb8dE877FF256881BfaCaB503Cbc50b447270",
				SignAddress:     "0x95c642D7a9074bCb6761D987602E563465397020",
				ConfigPubKey:    "9a9c781be58d8106ea631f839b91a6030af6738c28f1a4908a8d08d2fc144505",
				OffChainPubKey:  "deabea892332b24bae9aeea72a0b88bed7eaa925eb6fa1c36a3099739428dc2a",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "c26fb9dd63cabf536eb51358fd6c3f6183e8fbd9b2067ef2bab984d5d1464cc1",
			}, {
				Id:              8,
				TransmitAddress: "0x98C67459A03d25ae98781fDDc8e69B05b5d188C8",
				SignAddress:     "0xc97aF8B35742058D2b97a4C534C1070efbECdEe2",
				ConfigPubKey:    "33f480f668e16c28b87c7736b0b8217ff6d909bb1e5eae344d295e7e6b525d5e",
				OffChainPubKey:  "bd80552565fad2ec7744fd84460336a593bd690ed27cb4fc72cddb9c3e8468bd",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "d19e598c99386a2b5dd84af2d96b8ad729188484b016a6d980bfa50cd6ff2d3b",
			}, {
				Id:              9,
				TransmitAddress: "0x8cE54324473A2ee30956F2d7473aBBB7455078B1",
				SignAddress:     "0x9848dBe0B2820bcd19e8421d5ceB8D61A69156A3",
				ConfigPubKey:    "b56f1759470fa4201ec49a7966d8f6688618b468b80fbc9365070baf6b850d75",
				OffChainPubKey:  "57899042eb8a1129f2c85d44d577356aa0769a86ce36d18274a8ae2873ec2134",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "a4bcbc14d1896cee462185c52772deaedcb588a083a74f962d91c9bf820f0370",
			},
		}
		nodeConfigs[BSCstBTC_BTC] = nodeConfigsBSCstBTC_BTC
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
		publicKeyFileName := "publicKeys_fb_usd.json"
		publicKeyPath = filepath.Join("/Users/greason/Documents/workspace_apro/aproOracle/apro.configs/bscMain/publicKeys/",
			publicKeyFileName)
	}

	ocrConfig := GetOffChainAggregatorConfig(BSC_FB_USD, publicKeyPath)
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
