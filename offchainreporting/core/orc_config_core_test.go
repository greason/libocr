package core

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ocrConfigHelper "github.com/smartcontractkit/libocr/offchainreporting/confighelper"
	"github.com/smartcontractkit/libocr/offchainreporting/internal/test"
	ocrTypes "github.com/smartcontractkit/libocr/offchainreporting/types"
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
	case CoreSolvBtc:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case CoreCore:
		// 1% / 3600s
		AlphaPPB = uint64(10000000)
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
	CoreSolvBtc
	CoreCore
	CoreStCore
)

func GetNodeConfigs(target int) []test.NodeOCRConfig {
	nodeConfigs := make(map[int][]test.NodeOCRConfig)

	// core-main
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
		nodeConfigs[CoreSolvBtc] = nodeConfigsCoreSolvBtc
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
	ocrConfig := GetOffChainAggregatorConfig(CoreCore)
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
