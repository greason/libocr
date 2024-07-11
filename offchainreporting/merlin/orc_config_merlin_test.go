package merlin

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
	case MerlinBtc:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case MerlinUsdt:
		// 0.1% / 86400s
		AlphaPPB = uint64(1000000)
		DeltaC = time.Hour * 24

		DeltaProgress = time.Second * 305
		DeltaRound = time.Second * 300
	case MerlinUsdc:
		// 0.1% / 86400s
		AlphaPPB = uint64(1000000)
		DeltaC = time.Hour * 24

		DeltaProgress = time.Second * 305
		DeltaRound = time.Second * 300
	case MerlinMBtc:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case MerlinSolvBtc:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case MerlinMerl:
		// 1% / 14400s
		AlphaPPB = uint64(10000000)
		DeltaC = time.Hour * 4
	case MerlinOrdi:
		// 5%/28800s
		AlphaPPB = uint64(50000000)
		DeltaC = time.Hour * 8
	case MerlinRats:
		// 5%/28800s
		AlphaPPB = uint64(50000000)
		DeltaC = time.Hour * 8
	case MerlinSats:
		// 5%/28800s
		AlphaPPB = uint64(50000000)
		DeltaC = time.Hour * 8
	case MerlinStone:
		// 2%/28800s
		AlphaPPB = uint64(20000000)
		DeltaC = time.Hour * 8

	case MerlinMOrdi:
		// 10%/10天
		AlphaPPB = uint64(100000000)
		DeltaC = time.Hour * 240
	case MerlinMStone:
		// 10%/10天
		AlphaPPB = uint64(100000000)
		DeltaC = time.Hour * 240
	case MerlinMUsdc:
		// 10%/10天
		AlphaPPB = uint64(100000000)
		DeltaC = time.Hour * 240
	case MerlinMUsdt:
		// 10%/10天
		AlphaPPB = uint64(100000000)
		DeltaC = time.Hour * 240
	case MerlinETH:
		// 10%/10天
		AlphaPPB = uint64(100000000)
		DeltaC = time.Hour * 240
	case MerlinSTONEETH:
		// 10%/10天
		AlphaPPB = uint64(100000000)
		DeltaC = time.Hour * 240
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
	MerlinUsdt = iota
	MerlinUsdc
	MerlinBtc
	MerlinMBtc
	MerlinSolvBtc
	MerlinMerl
	MerlinOrdi
	MerlinRats
	MerlinSats
	MerlinStone

	MerlinMOrdi
	MerlinMStone
	MerlinMUsdc
	MerlinMUsdt

	MerlinETH
	MerlinSTONEETH
)

func GetNodeConfigs(target int) []test.NodeOCRConfig {
	nodeConfigs := make(map[int][]test.NodeOCRConfig)

	// merlin-main
	{
		nodeConfigsMerlinUsdt := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x3d28d50D479810A28F61fae2b6a9F0109167A14f",
				SignAddress:     "0x6d2dBB3EBC1C25Edba1f44e5894CAE63533eCBc6",
				ConfigPubKey:    "60ad33e37f5b55d4db6f590d27033555d984ac2b85ebefbee033a03b8cf20718",
				OffChainPubKey:  "d1f0f843463212512a1e0cd2eb4cf733add28d7732d3245c354d548ec2539f5a",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "f2efd67f4f2bf48eedad8e72f3b0cff842b4a8c83dada77bba096004bb2c4f2b",
			}, {
				Id:              2,
				TransmitAddress: "0x359B74A87bf37dc8341e19e065f0d7fD4d08368e",
				SignAddress:     "0x53aa4B8256856a9740776A15E3D37d297fa5b608",
				ConfigPubKey:    "2107a54c37756da9b5b015d524a203164f29efb72ac595be8eb6817e985d5926",
				OffChainPubKey:  "f82512e0650f4b338a0c97bfe37b2c502d730ccce4f06ede3ba2dbc02d8403c6",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "82c407fa3afded83cec23cade1d5671af18c7ee0ca7a1a6e0c9ee1094ec8b04d",
			}, {
				Id:              3,
				TransmitAddress: "0x6Dbc6F2Aca99fE83493236cA70a69e9Bd74c3740",
				SignAddress:     "0xeB3CED16d3b99238A75D6feB4A48CE1BEDE35bcD",
				ConfigPubKey:    "98d15f79999ba5d4bf775b6befb1781bcd751c0d1bfef847ee24890fece38058",
				OffChainPubKey:  "866b1e0029a58badda0eb9bec3d1efacb9db61dfe024411ad90c1cff2704f0ee",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "ce0b8418a2a8b8a672024c6dee2b9743ddc24e0960ab4422f7281313f34d2419",
			}, {
				Id:              4,
				TransmitAddress: "0x78a357dd83a4A05E9F93A5b80bBBAe7ADEeCfff4",
				SignAddress:     "0x5f78FdBd308B8b9F369C4E26e73ED76E58BfEeEa",
				ConfigPubKey:    "079e7784f5e4c52e75f507449ac186a35f9b7e908d4d587bd6af9d771115862a",
				OffChainPubKey:  "78b1c74379838a3d16189f04ae9ca3cab5f17d8db5dbf46e5b2b29c5fb920a2d",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "6930e509983b8c71854387b04ab26d72bdf935b589be4880762c30455253b564",
			}, {
				Id:              5,
				TransmitAddress: "0xfAcFB1b443c4f662EF61277192e972269aD9b177",
				SignAddress:     "0x5F1c9Bc57e781Fd62d45C3909CDCcE65cC9C7c25",
				ConfigPubKey:    "311d95596109676d94f02da1c0abe64e4c73410fef9a54e83e098eaf88ff8620",
				OffChainPubKey:  "0e4225a2cd80d5c68d916d0f0d52ead4583d13bc2cb12dd2b3aa930140f9fac9",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "5f5a7bf74dcd4c6a7cf5d758da18e4b00b52f2c6dc1337d06ede50f3f32a2a92",
			}, {
				Id:              6,
				TransmitAddress: "0x9deE0F70A39dCBc707F6A256CD41936c946dF98F",
				SignAddress:     "0x0F75ca44D5be353a9b000B58e8C2f744f09794Ac",
				ConfigPubKey:    "a9c0195e2e3b927655aca8f00bd7133d3585d3ee4b8e30e64296dd5e1b7c2379",
				OffChainPubKey:  "7e7d74881dab1bbf39d0272c16db4b069f024b7e576b79b64022e6836e96042e",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "b8a86b3c8158357d3e28889b807dd502110d3595d67331f38074c0d5cd59f4aa",
			}, {
				Id:              7,
				TransmitAddress: "0x8752916dB6bEe92356b3687795663cD60A17A50d",
				SignAddress:     "0x0e7Be7375aC3e0ef93106Fc3ba7499531aCA9cfD",
				ConfigPubKey:    "63d93f7409c6ba07404e332f8852638efe0b245e661c93bf727e52b8d4215071",
				OffChainPubKey:  "d851e8868e60de1322a369513ce8ddde8166bd576bcdf829adee42948db29a20",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "14283bed386b8d83d1093ab68c267feba63843ee5f33fed85b70a4d392f33ddc",
			}, {
				Id:              8,
				TransmitAddress: "0xf775f38F2d74847c7c0D627A2869FC429AfAc7f7",
				SignAddress:     "0x872be357934050F5F092533a2C1d81bd9959A7AC",
				ConfigPubKey:    "c54373ae830a2b6e1e13595df7cd68d20540a3cc89d37dfd450675d78ae4e74a",
				OffChainPubKey:  "d21b06ab7eeaf683eecb325b979f1be8f2ec0d8d58e4d7520fdbf8dae0c93d66",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "cf126310b3153c97676d840768185ce59d4482da2b4859805769ff30f33b0bff",
			}, {
				Id:              9,
				TransmitAddress: "0x4fE7f647d293ecC5cEf93a53B86696b4BB47e74f",
				SignAddress:     "0x6654dF35C728B10453B5637797bcBeAd2a7b0DA8",
				ConfigPubKey:    "74da37605d88003ca4f659a8ed532515e5d17eb52049ae1520ce36b01e4fea6b",
				OffChainPubKey:  "53c6c5da23ade37a98641adc20dd2c02e0ea91016d3f7ab4eb2b748e15477cb1",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "8dbba83ee834c09432be51195d39dd3a2d23bbd38c7583467c5b88bab6fc0da7",
			},
		}
		nodeConfigs[MerlinUsdt] = nodeConfigsMerlinUsdt
	}

	{
		nodeConfigsMerlinUsdc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x3d28d50D479810A28F61fae2b6a9F0109167A14f",
				SignAddress:     "0x3e79310c883b7EAd43A3Ef462fb3741d4e6Ad469",
				ConfigPubKey:    "209d2da28aee9480a09d08342db84c6e84bab3341f2b9dd981a9311e76d70e56",
				OffChainPubKey:  "80938bb50712bcfcb900b13dca4d8bd02d2080a09f340116ad1d888402bc21af",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "690010a698e40f1ee73828e09bd6efe618f16ad4b0ddece8543d229a432694cc",
			}, {
				Id:              2,
				TransmitAddress: "0x359B74A87bf37dc8341e19e065f0d7fD4d08368e",
				SignAddress:     "0x13C0280A622f98C47Bc66Ac19Ab49cD6eA4351f3",
				ConfigPubKey:    "9707da0e4b147022a109ba32522c5192f49a3006fe67d2caba76fa565c92dc47",
				OffChainPubKey:  "6ecd79c780cb901377da7f5ee05c6755387f4aba1801d06f18dc57932ef2e338",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "6fa43806af28311dff4f97e3653c2c48a6327cdc24814a0d3be529e4e8bf3704",
			}, {
				Id:              3,
				TransmitAddress: "0x6Dbc6F2Aca99fE83493236cA70a69e9Bd74c3740",
				SignAddress:     "0xf4050849f06548c71180EB3c411eC46fD50E9222",
				ConfigPubKey:    "780d35ff2f04e6182a75d80db01715ff58dcd70126d97e48b5995d37b0c80110",
				OffChainPubKey:  "5739d10f3292118451741f605f9bed081550f8935da006579ea044fc31db7f18",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "28d38fc1dff198011e3c9228b0ae69815a9d8e918441b9b5933cd5269a7e72a6",
			}, {
				Id:              4,
				TransmitAddress: "0x78a357dd83a4A05E9F93A5b80bBBAe7ADEeCfff4",
				SignAddress:     "0x6e286c67C4bB67530c2F8534565b679F227a433f",
				ConfigPubKey:    "396505db97edcf9344195a0c04974436d23520996b5b9b988fdef023ddb91c19",
				OffChainPubKey:  "1bbe9a6e4b4de00b45b6b4ffcdfaaae43026780a9db17909987888c2b389fce1",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "a999e30e1c91176f10b66043ea6a4dc535483241325736375c6fc8ce7651f6e1",
			}, {
				Id:              5,
				TransmitAddress: "0xfAcFB1b443c4f662EF61277192e972269aD9b177",
				SignAddress:     "0x7a2b81ac663299f9D176f6583bd2B085175Ad84e",
				ConfigPubKey:    "a7dbf53a724a6617beb90f82424edf4e625ad38159c308bc75ada77811f21e35",
				OffChainPubKey:  "295fe81268145e4e4dd3592e6a6768e06f2730fa0dc8da9550f2458ef67ec22d",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "81b5a6465aa9f2a137acb9eae9f9aad7b36092578b12c075310ac374534e372c",
			}, {
				Id:              6,
				TransmitAddress: "0x9deE0F70A39dCBc707F6A256CD41936c946dF98F",
				SignAddress:     "0x3A1140E3335b2BFa8Ca6B70E93b2eF11d02465F0",
				ConfigPubKey:    "c280642f637ba93ce3a04ca29dda234971d1e3d74298a6e19294db47d5ee4913",
				OffChainPubKey:  "c743213a9ef74636f0f075880ad2ec9310ed89548310ce87563da02ed03108e0",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "070ecd637c0604b9ad8f627cc813d53ba39acbb2522d14b5d117758be64eb572",
			}, {
				Id:              7,
				TransmitAddress: "0x8752916dB6bEe92356b3687795663cD60A17A50d",
				SignAddress:     "0xccD2DA37c61A51e8157DDE32c830781830955ab2",
				ConfigPubKey:    "d82ab19f9fd95a0cc7f80717e94d21c1f45701a7666600b78452e66168543060",
				OffChainPubKey:  "e5cce763cec8b09360c1655dc36383097ea2981e1114f313a3fa25408c9f6931",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "21f7105d43a4946b4f43c3c2c938c8945e35e2dcd60eeadcbca009abbfb7a5df",
			}, {
				Id:              8,
				TransmitAddress: "0xf775f38F2d74847c7c0D627A2869FC429AfAc7f7",
				SignAddress:     "0x743215fcECbc6Ed4162CB69E972CfB8AAa484363",
				ConfigPubKey:    "8068ab1ea3c1ad8390a3194941acdfa9c1e48c511fcdecda8695d19e9c81bb4d",
				OffChainPubKey:  "a6664662d047aa22cb835dfdecfebbfc2e5ba382727eb8d63a4383a93f9fa12f",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "86d625bc1fc0207d9a502d98fb91e8e9cdeb20ff0529bd62bf9d83208422019b",
			}, {
				Id:              9,
				TransmitAddress: "0x4fE7f647d293ecC5cEf93a53B86696b4BB47e74f",
				SignAddress:     "0x9E467bEff3a63D91F79fCdC058E4064dF33D3d81",
				ConfigPubKey:    "bd94301323af4653ebe630904951ea89c830b211ddc7547f3f0c0150e765711c",
				OffChainPubKey:  "8705145ffd6ceb5bfd1d1be5a30ab539906387d914a928d43d8ad02aa622859e",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "8317e43d5cba0c81e5026469aad3f4294a17cbf1fd66d30872e0fcfe7d287785",
			},
		}
		nodeConfigs[MerlinUsdc] = nodeConfigsMerlinUsdc
	}

	{
		nodeConfigsMerlinBtc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x3d28d50D479810A28F61fae2b6a9F0109167A14f",
				SignAddress:     "0xae1581e3c05A3A1e2f2Fb1bcE36786f6271dbd14",
				ConfigPubKey:    "279978fffa70f0ea1a2fc04f4135c82e8e2a2fbd9d876adbbb876d23980b3058",
				OffChainPubKey:  "1bb4d4b056a5b0e351ad5f6768765aacd2f5ec58234dccc2ac23840581ded057",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "8b854538e7b864dae48932fd66a87b501918eb5a8308dcd7038e30acd9bd6568",
			}, {
				Id:              2,
				TransmitAddress: "0x359B74A87bf37dc8341e19e065f0d7fD4d08368e",
				SignAddress:     "0x2197cC62056AEc8296edE1E31f9a2485Abd62300",
				ConfigPubKey:    "7439f41b4f942a7376df46bed6d5ec793bf93b486dbed6dce940af226b1a8b7e",
				OffChainPubKey:  "efcab40b68dbb6dc6a5f4c589d91f113f7c95c8b4d31da831ab2373b804438c6",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "f6bedbd98a689f4055d22401bd3e303576a41fb42a7e7401754d91809ecec52e",
			}, {
				Id:              3,
				TransmitAddress: "0x6Dbc6F2Aca99fE83493236cA70a69e9Bd74c3740",
				SignAddress:     "0x5d21E558259c762ef48fC21cFa8FC64726f616c6",
				ConfigPubKey:    "ca9346a6e4cd0fc9742e5cd5c6a47b202108c1a42c5f1625f3504c2f6cd24c06",
				OffChainPubKey:  "3be821f9e8854d2df0da8530b7ed0d9e3549eaa6414353ecc8540d1b2da4cbc2",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "cfa487fdafde4e20634d752c63dd2a1e715fb7367d996827180991762f92ca43",
			}, {
				Id:              4,
				TransmitAddress: "0x78a357dd83a4A05E9F93A5b80bBBAe7ADEeCfff4",
				SignAddress:     "0x74d6D69e28bCEd807315aDF7f0FDc1551E9e7764",
				ConfigPubKey:    "8521b39e0c1384c4215109b527d8252ad165459a76c87991689bbb852901a53c",
				OffChainPubKey:  "76011b8373eb7aada665d68bf81eb867a2a74f2021e7510baf2e646ea4ad8d19",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "5cbeb82517db1e42c25dbc3b472320b4ebc71ba0df1580e0b1fde0684c1a0c62",
			}, {
				Id:              5,
				TransmitAddress: "0xfAcFB1b443c4f662EF61277192e972269aD9b177",
				SignAddress:     "0xaC0993C300858E88278B7890f3Ce0e53e3bB189a",
				ConfigPubKey:    "56a338f5faa7eb9341110c1caaaca4612431393df23e5b51b7997ea7f1ae2678",
				OffChainPubKey:  "d228bae0bf557e3847e03dc2066947040c48a6ff38ac03ab6fdc5bd319f69f88",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "b49f103d8168737601b35095460d776ab7ab5a2e3be72f0666f9ce0298212392",
			}, {
				Id:              6,
				TransmitAddress: "0x9deE0F70A39dCBc707F6A256CD41936c946dF98F",
				SignAddress:     "0x3D4C754f9704520B59Aa6f780275f9639B265c33",
				ConfigPubKey:    "8c3df80984f37fc36d26f7086c28e769010c0c83ef96e6f9cb3797d17e0cac13",
				OffChainPubKey:  "2300a1401236ab488b8eec083091cb9deb3505554e9c903849f6051e017fc519",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "4bb086785d5d789900ee55fb9374747c592b479567a1ce4803fc9cbbfc65d6df",
			}, {
				Id:              7,
				TransmitAddress: "0x8752916dB6bEe92356b3687795663cD60A17A50d",
				SignAddress:     "0xcB1eC9a0Fac33b6c2b8f72fD86Db3Ba3F93156aC",
				ConfigPubKey:    "96dee51f53dff9872c57f7cf83aa9312d3779933fac755c8b34e52db47ce4459",
				OffChainPubKey:  "df7394f4b37b1d9c680b76ec1d30fd936499c11b4fb440856c93677c6fbaed3d",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "d56bcfb003171486100fc40e1a868f28cbb9248029d37230bd8ed3b5df2ed8b8",
			}, {
				Id:              8,
				TransmitAddress: "0xf775f38F2d74847c7c0D627A2869FC429AfAc7f7",
				SignAddress:     "0x44207CD6FFE3c697D8A8Eba6320D4521DAF9cbB8",
				ConfigPubKey:    "49fe3beca6e4cda35e7a209f4d39baf72434c9e5f3fbfe9250c12ef5251ccd73",
				OffChainPubKey:  "8d1b5b48a60b292269ed602859b81f4de5e78efdf8ae3191c4f297d437374cb3",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "6201c36a5542fb2a3c455f1d9d1e67396a92b828f4f1ec182555d0947992730c",
			}, {
				Id:              9,
				TransmitAddress: "0x4fE7f647d293ecC5cEf93a53B86696b4BB47e74f",
				SignAddress:     "0x56f01Df136dEF5a86C7d03F0c24FE0a9D2192Da7",
				ConfigPubKey:    "7559e2342fad8cd9076e2291ab976fe2380d0b7166a7588395bdf6a85aa45a59",
				OffChainPubKey:  "27b47292a99429ca480013b6a69b168dc0edb45eb5ac4a93711c2b94cb095e46",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "33b83bd849438dd04c5557b0f01f75dc7b40c870224ede4a0262a2038cf02487",
			},
		}
		nodeConfigs[MerlinBtc] = nodeConfigsMerlinBtc
	}

	{
		nodeConfigsMerlinMBtc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x3d28d50D479810A28F61fae2b6a9F0109167A14f",
				SignAddress:     "0x27a97c5AfD0A93d95F9505394af7BCA7f432beE4",
				ConfigPubKey:    "b4b8e4995e00efe549273fcd6e719f5f56f8da5b3797d5eb7cfe44d00ebb9700",
				OffChainPubKey:  "4a69874ab48988a361c1d7e9d33b2951b14aaf30a700882d1e4edbd5caf8a9d7",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "2c7656f806e0d8caf1c867c560b2951aeb4a1a2bc736192ab88520acbb9e3ed4",
			}, {
				Id:              2,
				TransmitAddress: "0x359B74A87bf37dc8341e19e065f0d7fD4d08368e",
				SignAddress:     "0x4e5127Bb8e9D7fC387d4aBf35b1391affEAD95e6",
				ConfigPubKey:    "9e6e3406e177a3c836313d83751cd025612a4161916eae80046e50fe00698d2b",
				OffChainPubKey:  "f1ac7ecb17fad8f7013ed6f8ff13a8e7beee61548779a34ad8fdc343e888dcb0",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "1967878d7e6736c6010234de3e341d02fbee5393c5f59c4d0f919ca0e7857ca2",
			}, {
				Id:              3,
				TransmitAddress: "0x6Dbc6F2Aca99fE83493236cA70a69e9Bd74c3740",
				SignAddress:     "0x8E64EC20D731781a7dFED39c22ecd107067e33Ba",
				ConfigPubKey:    "9561ecff335fa657dd228f4714e8a47e7df88f466a4c1293e7feabf3c82f840c",
				OffChainPubKey:  "90a9be5853fa7ca27538f51ff4cbcf1b461d036466934d71ed6a005511d17029",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "47efb725f45dc7da00680e3cc5c308cc1a504acb5184a471a20d116663d707ce",
			}, {
				Id:              4,
				TransmitAddress: "0x78a357dd83a4A05E9F93A5b80bBBAe7ADEeCfff4",
				SignAddress:     "0x4E89f4466d8E65c2a90eb696530F178BbA6dEb70",
				ConfigPubKey:    "b51117c37daeea636bc120403df2871eb0f0903326991d08890392ace0ee5152",
				OffChainPubKey:  "e8fd3e578494e7989ced16cc86f6ea2b5708c1506b9819356fd22ebea57293db",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "9735efe33b94e38b617a5e971909415137bac4b57df72067699aaef54bbfc332",
			}, {
				Id:              5,
				TransmitAddress: "0xfAcFB1b443c4f662EF61277192e972269aD9b177",
				SignAddress:     "0x149800Cce1a88e29Fd060742471300c123805B67",
				ConfigPubKey:    "77b1468624172468a2447c0375088744ac29cc3a7add406eae47e5cf41cce444",
				OffChainPubKey:  "e8deeba9ab6e3c654aaeb741567b5a48b70be4e7c324a4b57f0b8b1c1437514f",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "a5e114048ed94d7ba4773245bfcc2b274480a6b781aa3b315a124f3f441b5343",
			}, {
				Id:              6,
				TransmitAddress: "0x9deE0F70A39dCBc707F6A256CD41936c946dF98F",
				SignAddress:     "0x73F2d55c92a187f92fC23c6F5Ce9489BD360bCbA",
				ConfigPubKey:    "e713ce16338f0c1209edf581543092d78cb106e10a459cdefe9f1816913f2b20",
				OffChainPubKey:  "ea8e266f7826768f5c8a9669a333f4e857915f49659ab3544f1b1089807f68f8",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "e0b7ac0fafbb3fb33a36ad2372982f4c5a9e934536c8994af361a42ba6893eec",
			}, {
				Id:              7,
				TransmitAddress: "0x8752916dB6bEe92356b3687795663cD60A17A50d",
				SignAddress:     "0x0A8F3f6eD5fa14894ffeADDF35261CEFcdeE5d07",
				ConfigPubKey:    "6503e93aff0d5f312b1d7aeacfbf97dc89d5c94ad2f8204860adcd87e04d1154",
				OffChainPubKey:  "ae2793d0390b288202c5097c5465c8c5a6a29f8ecb31703693f81cdbf31999c4",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "1ed63722b8e9e30c29cd4e2451768f3203de198ef45fc4ca3845855cfe3296f9",
			}, {
				Id:              8,
				TransmitAddress: "0xf775f38F2d74847c7c0D627A2869FC429AfAc7f7",
				SignAddress:     "0x94600D4E798740C78D4AC6a6c73ad532Df4667E3",
				ConfigPubKey:    "55b390d2f5b9c367cf05fe483b96f49bcbff2523958ee6118e61f91c6800370c",
				OffChainPubKey:  "08dcbdc25686326a93f43da87e922c7e371c02bd71acaa6950cdb83960d62454",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "5279400f862735d6064af121120248d3829bb81dc7b0072a9d5ffe2adbaeac79",
			}, {
				Id:              9,
				TransmitAddress: "0x4fE7f647d293ecC5cEf93a53B86696b4BB47e74f",
				SignAddress:     "0x7BFd4F02B5A55D1E2bCBDBb4A35B7b3d3e0aFcdc",
				ConfigPubKey:    "235349cb4023ecb00f13aa2dee3ec23056b67a26764be4dff4b9969b6ce54a75",
				OffChainPubKey:  "17b248954d94f13402d170c7b11856457df4442ebe01d08ed258eb36e0389d02",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "467f8f0c268088156f0602f3fdf5853e48c637fc38b3d68171a35d121e23ae39",
			},
		}
		nodeConfigs[MerlinMBtc] = nodeConfigsMerlinMBtc
	}

	{
		nodeConfigsMerlinSolvBtc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x3d28d50D479810A28F61fae2b6a9F0109167A14f",
				SignAddress:     "0x2f33d5e003ab2B88F40063956AAd9966DBbc783B",
				ConfigPubKey:    "f850a241747444bb19ece0f53f8461ec6d47d9aed17ffcefb605c8bcfcaa8d62",
				OffChainPubKey:  "6c04fd4ed9090baf7b48ae65c200afd5234ea9c7cf7cc454c468d697a784f319",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "e9c0668976da5b73b9796f727559611b4ce38be7ad4bbb2d1e8c4ed0df1e5c16",
			}, {
				Id:              2,
				TransmitAddress: "0x359B74A87bf37dc8341e19e065f0d7fD4d08368e",
				SignAddress:     "0xa13b821AD08FB65a44E8074Fae51A15317876FC6",
				ConfigPubKey:    "382bf476e55742f0f5577c64047670d1b01ce25a67ec670c76507619edbf1d16",
				OffChainPubKey:  "48b60eb842356ba422de9d831e56e33af0b6dc6c079a9ff37b1e2dfd4c93c3ea",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "ff3375d2cdb654af818f7980ec0edda498ae892e700006a19994039b7252025b",
			}, {
				Id:              3,
				TransmitAddress: "0x6Dbc6F2Aca99fE83493236cA70a69e9Bd74c3740",
				SignAddress:     "0x6d406A8E1b69ac4c02362cf1391F3090420BBFEd",
				ConfigPubKey:    "50093a23a17c1e0df079b4409458ee24f18d82052b7f6149b0cea22ac656d027",
				OffChainPubKey:  "1e70e0465d5477245c8b3da602a4be781db57a13ead24c3f12faefb3f06cc238",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "f53eb15bfa5ccd18ebf56638a16913619491f3463f4706b81ebe571698a3729d",
			}, {
				Id:              4,
				TransmitAddress: "0x78a357dd83a4A05E9F93A5b80bBBAe7ADEeCfff4",
				SignAddress:     "0xbB9faB9DDD0f4F18eF8423259574775C404F189C",
				ConfigPubKey:    "76f2db41983bf229539635388bd2e30a48284fcb2db9b869c6c16910d5c64257",
				OffChainPubKey:  "922212d579226809d5f5666db52fcddcf504953505fa65045993092be8f86148",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "99f39da1c5ddbe7ea5044622791f01c0ae0cdcb898e53841a4e5c33dbccac557",
			}, {
				Id:              5,
				TransmitAddress: "0xfAcFB1b443c4f662EF61277192e972269aD9b177",
				SignAddress:     "0xC767af2AC583C42e12871F06e51f5D9b25fFB939",
				ConfigPubKey:    "ef65287ee0573e22a6c584bcdae4ada62dbe1e4602d1967a9d6572c03bbccd10",
				OffChainPubKey:  "77ea7425a8816234745bd8e2b11040f733c572b7f7f869e7caef4fbf4a2c6bc5",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "ad621e06919d7731aceadc0d10430d47b684ec166b30273bf385789fe0683ee6",
			}, {
				Id:              6,
				TransmitAddress: "0x9deE0F70A39dCBc707F6A256CD41936c946dF98F",
				SignAddress:     "0x832A1EE56068817eca8BF1f370cf7989CFd12B86",
				ConfigPubKey:    "863f707be832dcd1bfb4df935fedbdc9f0f277ccd0011fe397c6a9dfaba6823c",
				OffChainPubKey:  "5bbcc097b4812ab351f31cc193ef5c34d1fd19a3fc1aa8d4bb9e65f2cc7d6048",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "ea5823c983fc721b944fbc21bddb197629e71169270e48b1d161a09cc74c90e2",
			}, {
				Id:              7,
				TransmitAddress: "0x8752916dB6bEe92356b3687795663cD60A17A50d",
				SignAddress:     "0xD77d32B687e952247CAA13045B84Beea880dA5D8",
				ConfigPubKey:    "b11bdccec629399246fdfefba880dfe2ffab886ce49772e552a93430e4e78171",
				OffChainPubKey:  "6b1dce61e8b43905ca28e62cf0b5b9eb49c0acc83c315fc26ff5d9971e78738f",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "eca8ed22f16afd43f345d72635807d405137e6d7680b96cdfa1699efeff0edbf",
			}, {
				Id:              8,
				TransmitAddress: "0xf775f38F2d74847c7c0D627A2869FC429AfAc7f7",
				SignAddress:     "0xf8822c6dF4704d81d31b6f73F4C93A0Ba2b39326",
				ConfigPubKey:    "2d5245702a25a49e7c6774a0cabce6c895ecbc3dd7c085f5f7fae48a4f2b0e20",
				OffChainPubKey:  "a44b59b98ad6a8628a9bb153c16f830815affbd378c9c960ea6669a22db7080e",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "981d28916234536ec8a8e69df37489ec4f8c3438685911969298e8d1cc95d557",
			}, {
				Id:              9,
				TransmitAddress: "0x4fE7f647d293ecC5cEf93a53B86696b4BB47e74f",
				SignAddress:     "0x641eB055887d2D83504cd9076d2f3F6592125480",
				ConfigPubKey:    "2016a17e09a444555b499d7b3d7874ad762d10e5760f506715adfef299925327",
				OffChainPubKey:  "8d020dd4389117441478774ba987e0908842e10328604a14a5ca87acb4276070",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "043b17d7df9890dc14b05e876d846f88be65e5afc833d61c3b63c44cf6f0e858",
			},
		}
		nodeConfigs[MerlinSolvBtc] = nodeConfigsMerlinSolvBtc
	}

	{
		nodeConfigsMerlinMerl := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x3d28d50D479810A28F61fae2b6a9F0109167A14f",
				SignAddress:     "0xED97A08896f43501cEdF951AEd4cE936E8802F95",
				ConfigPubKey:    "ed2168866f0ad1583526fc6bdc854e8ed1861b35d7366cfa9b37aaf026be6037",
				OffChainPubKey:  "35c245074849f7f6fd5e064d8eb69ff362dcb2e4eec9bdfaf352a4c1c0eb08b7",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "2410ce56ebac198339f23cc68a6eae0cf63693b1f839bdfcd819b0a73725bbb1",
			}, {
				Id:              2,
				TransmitAddress: "0x359B74A87bf37dc8341e19e065f0d7fD4d08368e",
				SignAddress:     "0x0b4BC506A4cAAC31Fc48750DaE35d75bf51E7044",
				ConfigPubKey:    "2289571f8ea74cec4de84bd95cbe66b6f2180100a8770f8f27d7e71c4774c234",
				OffChainPubKey:  "60c0bf732c67045dcf89b26c9fe3d4a5e31bb31eedc2ec7cf7c588e23aac4757",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "98899e2d5dc48ca8cc63be25f1c5f84aa5b541b2b0521fee8abb950267b611f4",
			}, {
				Id:              3,
				TransmitAddress: "0x6Dbc6F2Aca99fE83493236cA70a69e9Bd74c3740",
				SignAddress:     "0x5D87A34501ed50302f2cF8Eed964Ad3DF894B564",
				ConfigPubKey:    "f158d29eca845820b53d2f3088d408349f905a69f364ac8094c203ad32a5b45e",
				OffChainPubKey:  "29efeb3365efd3df5f496e7a67b4fa53944cef5407d3295bd7903dd9bcdacc25",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "cd9d6cf933d2abaf2c4adc91b9889115b83ab665b62256eaa17f5f8d41e3912e",
			}, {
				Id:              4,
				TransmitAddress: "0x78a357dd83a4A05E9F93A5b80bBBAe7ADEeCfff4",
				SignAddress:     "0x6733745Cf0607a80e607d106Fe08Dc233345A5C1",
				ConfigPubKey:    "da2591bc6384144d1c8d7ad6d45efe0d9c96224fe15ddd6fa188dd7b2226a167",
				OffChainPubKey:  "e85ada6501c940ffe75322ea87b0a9f68c5b1d9d3f8eaf1d6cf18e713bfd89fe",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "40f4d0eae1a86d7a3ab8dfc87dc0e25ae071d2f116a4bf6b9448ff3fdad51798",
			}, {
				Id:              5,
				TransmitAddress: "0xfAcFB1b443c4f662EF61277192e972269aD9b177",
				SignAddress:     "0x8072139659941F354BAd5c8B00b84b9F858042da",
				ConfigPubKey:    "f3c95b1b187b7435702be170e528a63085389b45d57fefbdedda089e85d49716",
				OffChainPubKey:  "71cab561b31db7e050a6b36138af4c8662fc90e344ccf9411a43cdd7eadf81a7",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "02d57ebb571cb3e53cfe01ca64705bd36b81469603313a946d767dd4d7b2340f",
			}, {
				Id:              6,
				TransmitAddress: "0x9deE0F70A39dCBc707F6A256CD41936c946dF98F",
				SignAddress:     "0xaeAeEc4AEfC3522F592A9BC88C297d9086D47eCb",
				ConfigPubKey:    "2e23b71fa944fa2a744871c48e3ecec8ffee822e507452ab2c86d1704d092243",
				OffChainPubKey:  "6d627c8302d3c5e9df26fde85695a1d47cc8a69f5519aa4462e04b60b71c4001",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "cda31065bb94bbd47df06d33c798110dac13c9be03240eefe766b392793264f0",
			}, {
				Id:              7,
				TransmitAddress: "0x8752916dB6bEe92356b3687795663cD60A17A50d",
				SignAddress:     "0xC89E02f0634929EEf44c8980d5F65620dACca497",
				ConfigPubKey:    "cdb9d1362e98f8271ed7c23886c185f37aef2826880e47f5d5df101ff582a905",
				OffChainPubKey:  "baea156c8f271c215e77a16c1dca57e9dddd29a4b5f74db83a96afce5a1f6e30",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "cd0e6ecb62352f4fda0bd5faa5696a624512148841a8d51ea4a6919c1363f7f8",
			}, {
				Id:              8,
				TransmitAddress: "0xf775f38F2d74847c7c0D627A2869FC429AfAc7f7",
				SignAddress:     "0x45217342373b112f08093327515Cf3928Fd61DC5",
				ConfigPubKey:    "83718e908fb4d920f082209bce6bf99214808acb6036f9cf9616a68bfdcd4e49",
				OffChainPubKey:  "4e878d89a6f1eda03540cbff63bfcf527c0a5efa7e661a54401c385ccbe10b40",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "c7196daae72aa91574ef7ca3554fd6e28e0d6f8a491f0e21b26bdef300214735",
			}, {
				Id:              9,
				TransmitAddress: "0x4fE7f647d293ecC5cEf93a53B86696b4BB47e74f",
				SignAddress:     "0x52F9b8f453Fd8096B06940A381CEDd833d6F8CA1",
				ConfigPubKey:    "8a5215e732e0426c0b7bd25529485fa4068d884cbf83f286492f80e44e1cf80e",
				OffChainPubKey:  "0447bcc24e16a6f87b452baa99dfa1ead88a4ae050d774a709097889969e1b71",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "2f436d1f2001b90f3c04a1478ce4b1e99877871e8c3832b847bfabfe0701372f",
			},
		}
		nodeConfigs[MerlinMerl] = nodeConfigsMerlinMerl
	}

	{
		nodeConfigsMerlinOrdi := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x3d28d50D479810A28F61fae2b6a9F0109167A14f",
				SignAddress:     "0x9c84fc86c7F84b5b56745C3068104F53a09478E7",
				ConfigPubKey:    "03d7b81b585165cfd86440aa9dfc2560316b12df71a1040fca511120a768d321",
				OffChainPubKey:  "6be7351d5fa8215dc8d593372a3c6de4dcdc1165d12805278e56cd1b45bb6dad",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "703cde4d5bd5b334400a974f99ae20dfe1fc0ca210f1ffb4f40a93c5550acc15",
			}, {
				Id:              2,
				TransmitAddress: "0x359B74A87bf37dc8341e19e065f0d7fD4d08368e",
				SignAddress:     "0x2C5889495ba551245EE3a1142197107A0acA9c58",
				ConfigPubKey:    "da1aac6da03764f0d752e465a43d85037c128cba91dc66c6ef01329b7095cf77",
				OffChainPubKey:  "ba83602f29c943472ce346b6e1d0869e98e7bd20eb827e6050eb7b894d85a080",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "5de829dfcf7c9c62f7506f8da07a4eeb320c2c1f69c0598ac2250efe059a3a87",
			}, {
				Id:              3,
				TransmitAddress: "0x6Dbc6F2Aca99fE83493236cA70a69e9Bd74c3740",
				SignAddress:     "0xAfA73D47faCa0449eBb1e1C14C08d7Ef9237e7AA",
				ConfigPubKey:    "2825ef57203f505c5d5097f2022a3aa58d666ecff0472232eb157307441faa3e",
				OffChainPubKey:  "fa1f7d8fb19a3375d3b573c40e6c0154c313c3d5e8a2e6e7d0f498456f13e8d2",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "24913c09b22f5aeee779e1633ee2fa85e4590f13d766ed30e5cbc0254022446a",
			}, {
				Id:              4,
				TransmitAddress: "0x78a357dd83a4A05E9F93A5b80bBBAe7ADEeCfff4",
				SignAddress:     "0x3C2A9DB509a9A662d3c627A446F0608692314F11",
				ConfigPubKey:    "2bebbb613ac2cc7f30e639d9cb52875eff53b989e8f0fcddd66dfe6e86eee733",
				OffChainPubKey:  "c2c870061b9caaeb076d6366f7bf0e2772fb9d0cd58789683501fb11d66656b2",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "ece8d9cc2b0a71c9b9cbcc76ebfc56085d2bbe2f029a007aab98f6ba8bea61f4",
			}, {
				Id:              5,
				TransmitAddress: "0xfAcFB1b443c4f662EF61277192e972269aD9b177",
				SignAddress:     "0x824B0Ede13692C0e8C72937854459F0851541374",
				ConfigPubKey:    "c8e6feed83d92f0ce48c14f1fae99c7b63097054331566e0537f0648c0cbd17f",
				OffChainPubKey:  "1f97506225cd893450b1246a5dfb155ddd1b37a52cec8994c0c5c9dfed2ecba0",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "49ae8e0e38b77c987d06d21252bbafc5bf5699f63108d55b382d738c16a044e4",
			}, {
				Id:              6,
				TransmitAddress: "0x9deE0F70A39dCBc707F6A256CD41936c946dF98F",
				SignAddress:     "0x3d10AC2e884a313Af17A2d2dd71403fE6909F76F",
				ConfigPubKey:    "22d96b4131cf6eca7f4514c1a2ba555163d4bdb33f1bd19b4625d2f7de1ab663",
				OffChainPubKey:  "cf096b285be9e42995425543bfe3ae81ac9f19d4718001624406bcd9e0630f69",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "eb8f979f9a73298736cf58c58d9c8946c426000a159a27d0a58075a80cd21fe4",
			}, {
				Id:              7,
				TransmitAddress: "0x8752916dB6bEe92356b3687795663cD60A17A50d",
				SignAddress:     "0xdEa3660E140fae7Dd8F7a79A20B8E89dF77717B9",
				ConfigPubKey:    "4161416b384a6f258e9eead813303917fa8987871d43f9324d998f9c0dfcb726",
				OffChainPubKey:  "99890e2c0012b569308b39a200747fd7b6051c13846661875f3ac44370026bd7",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "7d9292e5017711be3a364cef985d3178404c3c509d48e21fd1bd69d2b2e24cc2",
			}, {
				Id:              8,
				TransmitAddress: "0xf775f38F2d74847c7c0D627A2869FC429AfAc7f7",
				SignAddress:     "0x8ef9041F5f8F9447f75fA925B72D9c35B2731b7d",
				ConfigPubKey:    "2be12145c0c62d0652ae35d29e064f560c433dd451ff44649520e913d336fa08",
				OffChainPubKey:  "42315cfab6f03c9e171ae564df5921f7412d24cc19af14bd7efb9772e9f80ed1",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "9c7c093ad1af9a893f4e642ca03512456e3bb45f37d55859877097231cb1e3bc",
			}, {
				Id:              9,
				TransmitAddress: "0x4fE7f647d293ecC5cEf93a53B86696b4BB47e74f",
				SignAddress:     "0x4462ad22135bA56160a9B3b84e5c77E93FD1eA00",
				ConfigPubKey:    "42af2514d22a578fd66be74b6b3c7b6f97aca4e44fec8fe9b3af079b2f19f575",
				OffChainPubKey:  "4f5619124e2c82de8f7f5dab60b874909e72831a494d2fc4bdea6c81bcf205cf",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "90c9a36ff1c1aaac2e579618bec302a20b6b6a553338900da90f152e94111935",
			},
		}
		nodeConfigs[MerlinOrdi] = nodeConfigsMerlinOrdi
	}

	{
		nodeConfigsMerlinRats := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x3d28d50D479810A28F61fae2b6a9F0109167A14f",
				SignAddress:     "0x814d7f8568030D66709eBe3dbc4e9b9F86bDc422",
				ConfigPubKey:    "64806214128ffda0df6feb7448806ebc741141e998ae1733e372cadb0de8de39",
				OffChainPubKey:  "3050e5706a9beaed3f504bdedfd6151b08c82f8a66ffbdd8378c37b8e3b42673",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "4f68e7f1d45223eda8e22c77244ef061fa7e12b0e6e8b458d458efdbc9d3d7b0",
			}, {
				Id:              2,
				TransmitAddress: "0x359B74A87bf37dc8341e19e065f0d7fD4d08368e",
				SignAddress:     "0xA7D9dB0F569652b104A2bA9ebec8297268835171",
				ConfigPubKey:    "ae02ad4d3a9998c36cdfe690f439d4360b5cb38b14f05ce733f4dd033bf46d10",
				OffChainPubKey:  "80a8612db5acedd320705d43519b8fd3aa470008449a88d38e654da388d007b6",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "27ff43444303c5a589d2678b129f9287153937599906317b9a50de2993170d02",
			}, {
				Id:              3,
				TransmitAddress: "0x6Dbc6F2Aca99fE83493236cA70a69e9Bd74c3740",
				SignAddress:     "0xAcA1dC92E19464121d0F81Dab5AEa9078590aA46",
				ConfigPubKey:    "eef6337e133863edd3306f97b5f24883348ddba6b2b084173311587244d94568",
				OffChainPubKey:  "bbf1bab2cde37d493ec34fbe8d995de657e45345de2f8c0b036a092de1718562",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "ac5d303dbe2385528cd71ae5e67e7e5c326a274dcf5fd61be0edb9cad3b83dcc",
			}, {
				Id:              4,
				TransmitAddress: "0x78a357dd83a4A05E9F93A5b80bBBAe7ADEeCfff4",
				SignAddress:     "0xa7413bBDF1D911f003aa6d0a49CFd55E935254c7",
				ConfigPubKey:    "3dd35f8a022aa51d208406809e26e7a14d4b0537e5da9cb756bdd67783dbe904",
				OffChainPubKey:  "3d855d88167f0987f7f9edfb3e5698f466f31729b4d4d8a9346c453557e6d08b",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "a891c945d9b8d2f5835d08700dc5fe7090d560911edf95c7475270cd4c09eeca",
			}, {
				Id:              5,
				TransmitAddress: "0xfAcFB1b443c4f662EF61277192e972269aD9b177",
				SignAddress:     "0xB4Bf00dab75E1349059E731233203eDB2Ac46579",
				ConfigPubKey:    "1fc937f7dcad5dc1e025b834a5831cfc3e887588039c1cd048dd7f70f93a3777",
				OffChainPubKey:  "f794695233d99a208026037537fa75b4193b595e26b1a80ea614b2c788fd0121",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "91f06c1ee31814d25f23d50a674380b6f4d3823bf86b08bf106d4503204b1c20",
			}, {
				Id:              6,
				TransmitAddress: "0x9deE0F70A39dCBc707F6A256CD41936c946dF98F",
				SignAddress:     "0x3fAE6953465F7433EF6C9db26dE926a498D15114",
				ConfigPubKey:    "18b2cf55624f759ea496baa9bd4c6cd306b97c2f1eb1616da80dfba3e8c12d3f",
				OffChainPubKey:  "f84cfdeba649be4cf194988437d31cf986eff1819d27a2b79bea72c25b65180c",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "814cba8c15d7f7f9d9f79bb28f8a8d22181da5656ad981a6232350b1daeb7deb",
			}, {
				Id:              7,
				TransmitAddress: "0x8752916dB6bEe92356b3687795663cD60A17A50d",
				SignAddress:     "0x3754599215Ce7cc98dE379489aaE570AA550C619",
				ConfigPubKey:    "31f4f84c7fb6389a59f7d8add0e456358364d9804dd593386a982db87047c66e",
				OffChainPubKey:  "1b55dab9fd6f1caac87608e71cd1b608b5de840561369ce4a3f1203e8d92c4a4",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "648c4bc77481691aeba7079c6716f3a52b241f69cb1556227ecd308de8fadeaa",
			}, {
				Id:              8,
				TransmitAddress: "0xf775f38F2d74847c7c0D627A2869FC429AfAc7f7",
				SignAddress:     "0xDB0b5a33aB01B0d0Ffdc4815be0468e28df3185C",
				ConfigPubKey:    "df10505c5674cd892353825d95ec9586163abf22db024640b14089311ada2066",
				OffChainPubKey:  "0c091c75f8153778aff426be8a6b0f495374ee630554a6603894bdceb8542f02",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "850337d6fe09ee815251a0936409ed43f27260a7c5ec118aa593b10b7b42a9ca",
			}, {
				Id:              9,
				TransmitAddress: "0x4fE7f647d293ecC5cEf93a53B86696b4BB47e74f",
				SignAddress:     "0x533ef012CeC6bc3C3a6cC46280c00cD3e3000dbc",
				ConfigPubKey:    "3471bda21e22b0ba537836fc1898293c567b57595448474c5ad578cb1bb61d68",
				OffChainPubKey:  "200c05ba29902d48e597be3473e6a6694576e72a25eccd1d33f9b8951b1a1ebd",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "ba1755d397633be55dfabb44643eb369907a2a59357584fb9e4d5aa0ba02f43c",
			},
		}
		nodeConfigs[MerlinRats] = nodeConfigsMerlinRats
	}

	{
		nodeConfigsMerlinSats := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x3d28d50D479810A28F61fae2b6a9F0109167A14f",
				SignAddress:     "0xDd391F553a6b484f423f3e67cE2Aa3915dd0CebF",
				ConfigPubKey:    "3053c6a0deaf72b4c8f7a3f267a058476a4d7e424d08b51afbe0c5a8d3c19f60",
				OffChainPubKey:  "3902f3d49366abb15391ca236fce3867d05e568296b2cd948d1ae73e7143fd24",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "3c20ab1d9a977601b77c125dfe9f0f832d0984d437ddfbe117b2550b58fc972e",
			}, {
				Id:              2,
				TransmitAddress: "0x359B74A87bf37dc8341e19e065f0d7fD4d08368e",
				SignAddress:     "0xfB1a2DA70BDB7fA56118e3f0dFaA67f218cbd81C",
				ConfigPubKey:    "7beb5af433d20d227dea9a1f1832a389a4abbdc2927087ca42b489b66a500f2f",
				OffChainPubKey:  "13559b52cd68c89d138a03963ee42431d2f738659fe004c4d6522816c71b198b",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "6a2d7405a9c4465d9961890c16b4c1b4fea748947421f5a353260a1387d91e00",
			}, {
				Id:              3,
				TransmitAddress: "0x6Dbc6F2Aca99fE83493236cA70a69e9Bd74c3740",
				SignAddress:     "0xA1BBBb1B9d09bE8f8e3328435a470335fF37ae60",
				ConfigPubKey:    "2d3159961afcb6d66f719f90e6ba76709b440c58e7b98804377c994c5f435946",
				OffChainPubKey:  "04322a263ee9045042352367ef7c751614cb20eb2b7033730864aaab2ae8962f",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "c40c174eb66ecb0f070f20eb5565edde7cf6e43c5d650d7e9f450ab7d02209e7",
			}, {
				Id:              4,
				TransmitAddress: "0x78a357dd83a4A05E9F93A5b80bBBAe7ADEeCfff4",
				SignAddress:     "0xc7e20ddc9910925fB4AEc66cEb7c3D103769EB60",
				ConfigPubKey:    "851149271c13eef02c1f4dc6408c834052d579486b016d81d32c3514f5e50f3c",
				OffChainPubKey:  "9818994077530b9503cdb10ea815b76f8494ac5355c57810cb9641f3fea818b0",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "6a186a1bc5682e29d8cb3a4b60a524fa83aca32310f0a2a28660d4037a2266cf",
			}, {
				Id:              5,
				TransmitAddress: "0xfAcFB1b443c4f662EF61277192e972269aD9b177",
				SignAddress:     "0x8D2AF12CA8EEcbd8aF7f95437D023cdA0BD12a87",
				ConfigPubKey:    "79d1aafeb7ddbfee4072bd57327b428863347d79217c7ebc34aafd8a2be9cf49",
				OffChainPubKey:  "09cdff3c9ff3635f7ab349206cc768fa0c41518325f57f674357c03f25532c7b",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "6ffa12e0a19a86273f179bc6e33b5b00916e6984ee0d9bb9ebb1d2e919eb3f28",
			}, {
				Id:              6,
				TransmitAddress: "0x9deE0F70A39dCBc707F6A256CD41936c946dF98F",
				SignAddress:     "0xf1c85273264431c1581aBC8EA219108D6088d70b",
				ConfigPubKey:    "cdf6359cc3cdeed5ea6eefaa845b86b7b17bd85875e68143ff74f1ff53c66a47",
				OffChainPubKey:  "f8e16a9e05305d2dc29723fbf23b99fed9438a404693c735ee8e51420063caa1",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "fd42b15f35f2ec64d9002a57908338fcda24bc0b2177aab4c1e58343f8204086",
			}, {
				Id:              7,
				TransmitAddress: "0x8752916dB6bEe92356b3687795663cD60A17A50d",
				SignAddress:     "0x13202f50D92B93695B1c1658368B59AD47a3cFD2",
				ConfigPubKey:    "567f2dd4ad3d47dd10d7a494b862356ff833e3586c797355c21eb51294f9d828",
				OffChainPubKey:  "70d2e6a9acd40c605dc276acecdd642045f99de2399d3472da00769fbe477047",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "8d9a183755cb36f2915c20213a6515c6b91b79d4e9dbec764175097b1e274936",
			}, {
				Id:              8,
				TransmitAddress: "0xf775f38F2d74847c7c0D627A2869FC429AfAc7f7",
				SignAddress:     "0x33c946cb28A85F7959Bf4074Dba04c621A254863",
				ConfigPubKey:    "3d1654c7ea3e7ed64f48dd5ac011edacda99a2c1e45321c8ff883a0cdc351957",
				OffChainPubKey:  "b7138bd3969563db9909bb99b393a1b186fe58feefe43a8fe691fbb237c7be73",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "54f4ef020a424004a913895d459cfa8c09f5739a1753ccb8718ecef05f8c301e",
			}, {
				Id:              9,
				TransmitAddress: "0x4fE7f647d293ecC5cEf93a53B86696b4BB47e74f",
				SignAddress:     "0xFBa39C9537De50AeC41745B7717649c01369e29a",
				ConfigPubKey:    "c05f8d5db95ed2646304926f7da514f0009c9b8d9f5fb33589ba866188d9f102",
				OffChainPubKey:  "121898f11fcfba6583de2bc1cb04ba13397628619479fd3bbd8662b8e30e7561",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "5ea08b830ec4491687ca1209a3f3515b488d05904c33bafaa4e2a6358ce3cc85",
			},
		}
		nodeConfigs[MerlinSats] = nodeConfigsMerlinSats
	}

	{
		nodeConfigsMerlinStone := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x3d28d50D479810A28F61fae2b6a9F0109167A14f",
				SignAddress:     "0xc98E74afeD47E23fA410774ae961e7C14e4DC9a1",
				ConfigPubKey:    "116b843912db77cc1cf1f788fa67f34bb4d7983cb3750c1009521acc1b40d077",
				OffChainPubKey:  "7ab4c681756e9d5317a78819afe46e17d2f801153e797c68fcbdd5d7663ec387",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "bc89f61d7a205bbf0a0c0b10ac6227ab2b2e3adfe84f62aa66ac717fd6939170",
			}, {
				Id:              2,
				TransmitAddress: "0x359B74A87bf37dc8341e19e065f0d7fD4d08368e",
				SignAddress:     "0x2BAF877F16Af1a45923092ebE72728F1cA0Fb09f",
				ConfigPubKey:    "bdb4f37c54e00ece6d6c93ec8fb8d0608d1433b07e9f7dd46534519d7d594250",
				OffChainPubKey:  "00b3bbd4a75fbd2627136fe0be5a99bad684869f9bb4e69ac73fba91a197be9e",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "115f792f174e8405ee57264e3116ed81ab666c92b84174d768696844248ec1ad",
			}, {
				Id:              3,
				TransmitAddress: "0x6Dbc6F2Aca99fE83493236cA70a69e9Bd74c3740",
				SignAddress:     "0x8888Bd11306A0E19520E863eA8e82315FE0291Fe",
				ConfigPubKey:    "6500efc88dca904f0436895ee2c0a11169a04127fba58233e1d0410ada76997a",
				OffChainPubKey:  "9b0a4afdf329f8b304518730000aeb4e10ab990637ae063681239d7ce9d7a213",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "62ecad7b8c256631c7c91e413bbb1270ea96e5318606b0a9b6e8720d85834263",
			}, {
				Id:              4,
				TransmitAddress: "0x78a357dd83a4A05E9F93A5b80bBBAe7ADEeCfff4",
				SignAddress:     "0xC2307fc605762FCF750E91d94120d01F1dD8d8C3",
				ConfigPubKey:    "d3f7aa8e10cf05c6efc87760514e8fbecb1a5acfe5cf1cc31cc5bd0989df2142",
				OffChainPubKey:  "70807d7b16736b160bd079dd833acd40de8d851f0158273b6f1030fc8b094904",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "80e4d9d94b02c097812381d2d0b861fbc51651f3795a5a3bb0bc78e5faaf51d2",
			}, {
				Id:              5,
				TransmitAddress: "0xfAcFB1b443c4f662EF61277192e972269aD9b177",
				SignAddress:     "0x10df5Fd0078CF490F9930129A16DCD74e35cc11C",
				ConfigPubKey:    "6da93d64117107dae5c06ef58638e89b43fa3bd6c863b752dd80d5f95a194504",
				OffChainPubKey:  "70a60fea6fdae94171b933071ce6b81705689bbf1ecb4cbf1ac0139dca5b9b46",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "826b151c60194ba165790a997ddb613003cb12d02293d22b0e59fec8e9aa1292",
			}, {
				Id:              6,
				TransmitAddress: "0x9deE0F70A39dCBc707F6A256CD41936c946dF98F",
				SignAddress:     "0xA5C21E1708395748712885D5c28d2055852700c2",
				ConfigPubKey:    "f4be1d880b1ca2a377ace8136b41c80cdbd37e178c40167a244239ff5ab28f32",
				OffChainPubKey:  "7d186bb9312e31cfde8e67009644f7406c9c29a9abbd7007aa5f5383122cbdd4",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "fc12db3c313b18b4d05fac4a402b93e924ec3886b95827f946d28f5dd1b92ddd",
			}, {
				Id:              7,
				TransmitAddress: "0x8752916dB6bEe92356b3687795663cD60A17A50d",
				SignAddress:     "0x46f53B017281C6cd8494302C5ccabbd526e29bB5",
				ConfigPubKey:    "42d2ad6f125047504ecb103f395cb543283925da47190f409c555a3b35a97f24",
				OffChainPubKey:  "8384d534d146349a865dc8c854710775012cf8b227a1ba7b3ef0db68358e4743",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "c827c18b06a74187f0fcf7c20047407f40c5234f3a035829b5e67452aef96193",
			}, {
				Id:              8,
				TransmitAddress: "0xf775f38F2d74847c7c0D627A2869FC429AfAc7f7",
				SignAddress:     "0x054c30d3FE51a0Ae308A132c7c15258E314BA0A9",
				ConfigPubKey:    "ea21836a1e8b9cfe4019ffa7e1d3f59e3d5fb198374e2f89a31247c38ef5c953",
				OffChainPubKey:  "80c0abc7833061dbf6b2941d0fffae4243740b5f61f789538ed8d34ca2aa9268",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "3170f1f3f1b1765c12305e93918fee472562384cca4fc6825ff15ba43ca2d681",
			}, {
				Id:              9,
				TransmitAddress: "0x4fE7f647d293ecC5cEf93a53B86696b4BB47e74f",
				SignAddress:     "0xC7498d1655Bf96961e1a18BBE579F6ec37D05E69",
				ConfigPubKey:    "674e98c0916c599738e9007498f931270f46b00ceb89a8aa59717f8b6e8bf632",
				OffChainPubKey:  "114ba60a3237d3e08df6c1ebe75f54c843f677988a782d65eb5bbadfaaa9c4d1",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "4eaca81ce8f125f3c01eb428295914b032570498fd239f5c235a339ad2d83b42",
			},
		}
		nodeConfigs[MerlinStone] = nodeConfigsMerlinStone
	}

	{
		nodeConfigsMerlinETH := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x3d28d50D479810A28F61fae2b6a9F0109167A14f",
				SignAddress:     "0x9318ef3e44AB0003594c3ab0c8090bE1614f57a3",
				ConfigPubKey:    "1723e93b33ac750a9597f93efec0dc03f148ac3ef06f0298e0bd32a34618f649",
				OffChainPubKey:  "ca203f844a5e43d43880cc50fefcc8d8991124344743422828e1e8753a6f515c",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "3a6d7877122b8af2a7fb65a0eab6e2027c130f58f02cbc2c6ca50b3594b1bada",
			}, {
				Id:              2,
				TransmitAddress: "0x359B74A87bf37dc8341e19e065f0d7fD4d08368e",
				SignAddress:     "0xfA04Fc46c93cf4E52FbcB840F557Ad58FFfEAC74",
				ConfigPubKey:    "e81072b4e02d2bd9e8a88c9a065965aa9126dd38e46c2639dee35f59604a1b2e",
				OffChainPubKey:  "aeee92e921be257d83931f97dff2a77911266930b6f3aae848add4a6e9c625bb",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "43a8fe804af835724cc94abbae898a65289b3feab2fe42f1cfee4c69c131c7b1",
			}, {
				Id:              3,
				TransmitAddress: "0x6Dbc6F2Aca99fE83493236cA70a69e9Bd74c3740",
				SignAddress:     "0x0D6e42412fB7c5D92fAfBCee0831187534716a90",
				ConfigPubKey:    "da73de88438b37604704cbf9c540c28a8fb2f4d3a7d8ec0988a4e6ae2b39da61",
				OffChainPubKey:  "63dedad85892d36a892fa8bd4173f9160e5fd4f091161dd1f3d10ce2628457eb",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "1d63c38a89d86ca192f92f353298b5c5abf6ac9ac4777b792e71e20000c812e9",
			}, {
				Id:              4,
				TransmitAddress: "0x78a357dd83a4A05E9F93A5b80bBBAe7ADEeCfff4",
				SignAddress:     "0xF3D6BC97dFdfEeD62c7EC1a7C1c809C17cAa29Ed",
				ConfigPubKey:    "655a6890577fb208f951aa579d202711505c56cbaa098406eb4cc24818b32830",
				OffChainPubKey:  "546fb91f68281341297406d2c0f380e55275485c3c1ee31cad94c2a8a88f8d50",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "2ec0c0b552b9a9b0093cd9330be45825b2dcbb9da3813506e5d5dc920ec16292",
			}, {
				Id:              5,
				TransmitAddress: "0xfAcFB1b443c4f662EF61277192e972269aD9b177",
				SignAddress:     "0xE5ED92913BAC4f999288E68E4D35e54d53982644",
				ConfigPubKey:    "0c476992559a6940ee3e5e27b58422dfbed9dab8136ab040cbe4c61dfdc6f662",
				OffChainPubKey:  "0abb3815cddc039587d8c4566a6f0eaf9d7d9267448d590cbe80c920e9173549",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "12d03d556f98a3540c2ed674d2c6a984d8fb7ab2911c541838bfdd9041b7bfb7",
			}, {
				Id:              6,
				TransmitAddress: "0x9deE0F70A39dCBc707F6A256CD41936c946dF98F",
				SignAddress:     "0x22d74DD5fC5a844975B2C9335086dc039e4d2416",
				ConfigPubKey:    "39688055b5d3271aa587db304a59640a7c6dd7f4fb70ab17c4a7d0654c7b2a32",
				OffChainPubKey:  "fe72be1bfd89a4712d959ae94e1e1c5c0ad9ed86b91fadc0461222553d648090",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "b450d9361e363467a03362e47c06f23981e542e28e6c8818cbede646196e410d",
			}, {
				Id:              7,
				TransmitAddress: "0x8752916dB6bEe92356b3687795663cD60A17A50d",
				SignAddress:     "0xd17A4792Dc60Cd6654bCB4E8722A675E3777B8e0",
				ConfigPubKey:    "00fa6e6cc35909c3ed39c5f16f0f2ea8e06a229254180f070020db8a117f035a",
				OffChainPubKey:  "b3bf84849a51e217a5f8a42c974204897dfe096252907735d54dfc177942a177",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "64a89e0361ec7d2b2e94ee2bcdc8df9ed16de8acaa23a932493ff3c0134f2753",
			}, {
				Id:              8,
				TransmitAddress: "0xf775f38F2d74847c7c0D627A2869FC429AfAc7f7",
				SignAddress:     "0x44EDe3546580349ffa346b03e60f8F1FD8a5DdA3",
				ConfigPubKey:    "ad4bcd8df2a753c88143806186f79d1662d00be7a577c6a46e7a40fd4fc6977e",
				OffChainPubKey:  "c891760f588cf419a8da6998ff62fd6d7873e33add1891f79e4deae983e408e4",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "77c1c8f157b6fe193b3ad31cd9a3702592519f92e6fd7b8a3ca35d73081003c3",
			}, {
				Id:              9,
				TransmitAddress: "0x4fE7f647d293ecC5cEf93a53B86696b4BB47e74f",
				SignAddress:     "0xe98F28644f2745fF819438cd0BA385f60f8C2511",
				ConfigPubKey:    "4ecc8c433d6f4f8e29366cfa936e214d1e0725c386f3f8e3a696b8194cd62a6e",
				OffChainPubKey:  "bd350758e5c64dd80d603d1fdb2a54c70ad1d1ea3dc9f4a2a77562c8bbf00d84",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "47cc653c593647c2ea09910d1a0f3c0dc5752b6866a9e8e4a6c551e1607820c1",
			},
		}
		nodeConfigs[MerlinETH] = nodeConfigsMerlinETH
	}

	{
		nodeConfigsMerlinSTONEETH := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x3d28d50D479810A28F61fae2b6a9F0109167A14f",
				SignAddress:     "0x1326b0827b0206D9A4fC1373A3a23C362Da40a8D",
				ConfigPubKey:    "c93b49d5bc3c808e20fe112a21834e2d59dce2015a6b3b8e2acee489ed125948",
				OffChainPubKey:  "8d1f5152eddff9230950018cc9505b94d507578bbdfcd88815e78f7ffcd8b1b5",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "3017fcb49bde1f1439760804ae5cee5b3bf2f1056efcb20fabaede9c621b8a0c",
			}, {
				Id:              2,
				TransmitAddress: "0x359B74A87bf37dc8341e19e065f0d7fD4d08368e",
				SignAddress:     "0x94b9C8De6A73A4F86dCEd9fB36D13058EdbEE7A4",
				ConfigPubKey:    "8b75f943d8515f46e046221e59367f9b520b9bd997bda96b5f54b043995c4303",
				OffChainPubKey:  "79d9946bb6b05bf0fca2c5de8949d10ffe5a7e06c18000bdc748e72487a81385",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "cfb4a8ef279e2fb99166311593ddb49821eeeb7e445196b73581c8967190b506",
			}, {
				Id:              3,
				TransmitAddress: "0x6Dbc6F2Aca99fE83493236cA70a69e9Bd74c3740",
				SignAddress:     "0x832170dB17c4CCe9937778abF2524d63A7c73669",
				ConfigPubKey:    "7acb22c01f5de7692a280f5784df0a1ae3f62054adf587eb4d99a28a51441075",
				OffChainPubKey:  "596ffe725e40fb88ef5bdb74dd07734e5c52daf8e62b9c61d366e77afdaf7861",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "3e7b40a7e31bf59f5a0c75f0a1a5383ec95758f1b1d7105aa225ef5e44780395",
			}, {
				Id:              4,
				TransmitAddress: "0x78a357dd83a4A05E9F93A5b80bBBAe7ADEeCfff4",
				SignAddress:     "0x96A5dC8718C9495fE61E13aD768e24D6EAB84304",
				ConfigPubKey:    "dbe451585ca12ac1abdcc9c9972897eeaa80db17c36c091d5ae71a0dca41193a",
				OffChainPubKey:  "a318ca079216b57fa4d9c26c9056cda768a1201fe1554dc156c17f1db38258b2",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "6b548e621cd8509d90e12fd3adeff2f2bf9637da8f55413836c38e4383b40393",
			}, {
				Id:              5,
				TransmitAddress: "0xfAcFB1b443c4f662EF61277192e972269aD9b177",
				SignAddress:     "0xa56e1c004490D28f70A2995cEF5F72b4D17bc7aC",
				ConfigPubKey:    "e6e6368f9ebb10a0e6da9bdf897e9f911c62a25bcdb8d80c6cb8590f4f2b9b1a",
				OffChainPubKey:  "595804575a3313a274b89078ec5ef48abe6992dcf5ec05be4762b8d398942734",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "748b0c97f19e69b9b6463af558eba6088e15a78ddde03ae4e905b169a59a72c2",
			}, {
				Id:              6,
				TransmitAddress: "0x9deE0F70A39dCBc707F6A256CD41936c946dF98F",
				SignAddress:     "0x7798edC95Acfd65954D9EDE54C6117824f899E3B",
				ConfigPubKey:    "ea5a5a0d8ad8f436f7b628d35528d86d41188e363826f0bf899a5c3430567948",
				OffChainPubKey:  "f341533e7a570eeed6ca58e0a352a17a154431389c06a37fd378ad8053e811d8",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "11b70b2a895af8ae6b6de41bd287cb6f9e9c91272f56b5da9871aabcba848f63",
			}, {
				Id:              7,
				TransmitAddress: "0x8752916dB6bEe92356b3687795663cD60A17A50d",
				SignAddress:     "0xe21b657E6CFcFcFb8e0D630b93cE752ac39Fb6B1",
				ConfigPubKey:    "777c95d0a1af649240e482ddc0524713ea109b72c7e2740df06197c0cfdec431",
				OffChainPubKey:  "8f11eac2cf295d71fff2fbbeae70952e48b1c1e54d6822e2a29461d8fee04858",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "3dce5962f5f5331c5ce4a20bc2de200de45e3d04bfed06a829ef2581e21136c2",
			}, {
				Id:              8,
				TransmitAddress: "0xf775f38F2d74847c7c0D627A2869FC429AfAc7f7",
				SignAddress:     "0x420B715BBf964e24f542b299f63c72A8638aB9F5",
				ConfigPubKey:    "a44d2b92b889afb7e1f2a2820e77269ce4e67b2d45b33edfdd3c5ee09ac09338",
				OffChainPubKey:  "14dc72123b908c26f2e6eab90214c726e482a4373fd9236399ebdc28c4d6272f",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "4445f626c8ce52f3f9934458687b0713cbc6906fa6f642500bb46900f3cc61eb",
			}, {
				Id:              9,
				TransmitAddress: "0x4fE7f647d293ecC5cEf93a53B86696b4BB47e74f",
				SignAddress:     "0x0AE0AbD89C7C5969FFaADCD4160cC55aE38bC138",
				ConfigPubKey:    "fd08bde3109bcaae6ba9fe19f2219a5575a6726442959437cdbeb78e74b62270",
				OffChainPubKey:  "f5e7b65a3410583fd6ca6153bebfe14b559bdce971ffccca55ddcb0063b6b2aa",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "9d471dd1ade43c8c72f5ee087009e87ad65b91e84467b86d070c1bc30af5817f",
			},
		}
		nodeConfigs[MerlinSTONEETH] = nodeConfigsMerlinSTONEETH
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
	ocrConfig := GetOffChainAggregatorConfig(MerlinSTONEETH)
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
