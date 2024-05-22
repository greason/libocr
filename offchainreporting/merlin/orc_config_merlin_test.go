package merlin

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
		// 5%/10天
		AlphaPPB = uint64(50000000)
		DeltaC = time.Hour * 240
	case MerlinSolvBtc:
		// 5%/10天
		AlphaPPB = uint64(50000000)
		DeltaC = time.Hour * 240
	case MerlinMerl:
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
		F:                1,
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
	ocrConfig := GetOffChainAggregatorConfig(MerlinMerl)
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
