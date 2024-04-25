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

func AproOffChainAggregatorConfig(numberNodes int) test.OffChainAggregatorConfig {
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
	// USDT & USDC，0.1%，86400s
	return test.OffChainAggregatorConfig{
		AlphaPPB:         1000000, // 10 ^9
		DeltaC:           time.Hour * 24,
		DeltaGrace:       time.Second * 12,
		DeltaProgress:    time.Second * 35,
		DeltaStage:       time.Second * 60,
		DeltaResend:      time.Second * 17,
		DeltaRound:       time.Second * 30,
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
	return nodeConfigs[target]
}

func GetOffChainAggregatorConfig(target int) test.OffChainAggregatorConfig {
	nodeConfigs := GetNodeConfigs(target)
	ocrConfig := AproOffChainAggregatorConfig(len(nodeConfigs))
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
	ocrConfig := GetOffChainAggregatorConfig(MerlinUsdt)
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
