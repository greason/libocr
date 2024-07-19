package bsquare

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
	case BsquareBtc:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case BsquareEth:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case BsquareUsdt:
		// 0.1% / 86400s
		AlphaPPB = uint64(1000000)
		DeltaC = time.Hour * 24
	case BsquareUsdc:
		// 0.1% / 86400s
		AlphaPPB = uint64(1000000)
		DeltaC = time.Hour * 24
	case BsquareFDUSD:
		// 0.2% / 86400s
		AlphaPPB = uint64(2000000)
		DeltaC = time.Hour * 24
	case BsquareBSTONE:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case BsquareSTONEETH:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1

	case BsquareMBtc:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case BsquareSolvBtc:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case BsquareuBtc:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case BsquareMBtcBtcER:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case BsquareSolvBtcMBtcER:
		// 0.5%/10天
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 240
	case BsquareuBtcBtcER:
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
		F:                (numberNodes - 1) / 3,
		OracleIdentities: []ocrConfigHelper.OracleIdentityExtra{},
	}
}

const (
	BsquareUsdt = iota
	BsquareUsdc
	BsquareBtc
	BsquareEth
	BsquareFDUSD
	BsquareBSTONE

	BsquareSTONEETH

	BsquareMBtc
	BsquareSolvBtc
	BsquareuBtc

	BsquareMBtcBtcER
	BsquareSolvBtcMBtcER
	BsquareuBtcBtcER
)

func GetNodeConfigs(target int) []test.NodeOCRConfig {
	nodeConfigs := make(map[int][]test.NodeOCRConfig)

	// bsquare-main
	{
		nodeConfigsBsquareBtc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xA8455696F2DD85a6dce6b5AD683BF3602d73181B",
				SignAddress:     "0x846E69791E739C08d1bE8E1b10dFC318749960F7",
				ConfigPubKey:    "be086c11a246459a25513bf6184b735a6947945933d9ffda4d568f72a65af226",
				OffChainPubKey:  "86d64c182e77054c8a49523450e9cdac6320bae62b3bb6b45769fb491288da87",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "9457c5d905f97db0e3e4ba9e027d0779095a0dd8469d01a5bc49fd59ef206b47",
			}, {
				Id:              2,
				TransmitAddress: "0x3c892f62acfeDB80c1046F6215b31BF48A704636",
				SignAddress:     "0x34d141c513c021ca53Db902aC1E6De07F4f12E9c",
				ConfigPubKey:    "df964224813a335a499c41710eb6657d0867a0ddb879ff88a20f3f1f7609d87f",
				OffChainPubKey:  "ceb22bac76947d6adfafefca66fdfe907c32c38f3cb815ab82ea0b14b75057eb",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "ed4b92d062fda79e82ac9daa2c643cbc47ad2504945e7bc4e6feb5ec64fb64a9",
			}, {
				Id:              3,
				TransmitAddress: "0xD11D0A08445F4F2daf2412f2D665e1fAD20F4f84",
				SignAddress:     "0x4C2cA3a4Bf51Db200DEE56b8bD75F74E0479c8e7",
				ConfigPubKey:    "474c2e3919f7af59554f7d0a0e75a9139e7af1ba2c260746bc7e91ca56e93e3f",
				OffChainPubKey:  "0ef6b4a4507e9eb902b8f30a819af85f1a0e60fa562cba896796e95e275a456e",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "81c8b488aeeb526e277f260b18b225cb4514c44b40e37ce53c9b9386dba6bfb3",
			}, {
				Id:              4,
				TransmitAddress: "0x18AE11b760B3e1EAad18946cb3D64e97a565e3Ec",
				SignAddress:     "0x558e420141e3a2D8DbDf2c6BE380B98a577d1250",
				ConfigPubKey:    "6bb26dbd6d474d0fe807c69fcda879ccd60a16e69f79284dfb3ea3d74e9b5f51",
				OffChainPubKey:  "76ceb30e6e54b1a79146315a5c86501e15326d5b32697e72aef1514c17ded4f4",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "ad725fae1f2a7d4b8b709ff6487e6f06d6239a631555a362e5b3126f77847339",
			}, {
				Id:              5,
				TransmitAddress: "0x4EA989bDBC43701bAeab917c16E8d4d7f193b590",
				SignAddress:     "0x71fe355dbaA421F03dC4Acf7a383c3BfB7bF6a4e",
				ConfigPubKey:    "3ab631def9476ef333aee3ca80f2d69f5f47582796622aca1bdc630cbdde1876",
				OffChainPubKey:  "791fb1d036af620e50ce2dfcab1331a42c3b3276ec17fffbd502f794122b1c7c",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "96de9a07078c782c84edfa45ec55f51d268340c3e14e9945b7fa6cd4dab235a1",
			}, {
				Id:              6,
				TransmitAddress: "0x43020E6C38C3a8ce5247AF292858Df5cF3e664c4",
				SignAddress:     "0xac3E0Dc4287799FE39d73e226B7643F8ac58cD23",
				ConfigPubKey:    "b6b7068954a21a8a01a67f36a35893d30d1322525544decaa128b27b5210c223",
				OffChainPubKey:  "aaf7ad2d48d35bfa7b3012fbca85d8866aa7da07f102bece5f00059cf00591bf",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "a55154d51f55bd8d0efb940dcce7e0f8bc2632eaa435f5e4ffe8cb9013f91cf6",
			}, {
				Id:              7,
				TransmitAddress: "0x7BDe199140F7bdBA1e259A26ab15775a1cb5e599",
				SignAddress:     "0x2f52B998aCdD5ce38830a4fd1431462aBe72094f",
				ConfigPubKey:    "10ef53c55c9b7e22e9cc50084550bc1dfc8f361e7145ee4c6433e7552689eb1c",
				OffChainPubKey:  "f24e6d7223134514aad256422f6928a666fa58e552e7c8d2b55967fb51b1c6b6",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "c9367dbc324e477b59e227fab30c26f6c81ed8ee40aa288711a078525ea5bb36",
			}, {
				Id:              8,
				TransmitAddress: "0xa52074d8C7614cf4CaE28724c0C18476802e3434",
				SignAddress:     "0xb519098fA086F6D3e71b90b8aA56A9722CadD01F",
				ConfigPubKey:    "b9ce78e5ae961547df65a6c7d9b32381815111e9b5acfcf1334982e6dcc20a4e",
				OffChainPubKey:  "e1542c59b4f809873bf747ce86487652088758ae76e1f64dee21056dccdf839f",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "81c348e2f07dd0cbb9ef2061592dbe191b76441157797a28c39cb8ac7ce854b8",
			}, {
				Id:              9,
				TransmitAddress: "0x97cd58ED3B696C21128F395fB96f37550D201207",
				SignAddress:     "0x840bb3c2944520B5a8Fe3bB3691aD0891Cb5ce8d",
				ConfigPubKey:    "91f99f2dabad2cd62e5c1690733881be038d27125f22209f981326daac293924",
				OffChainPubKey:  "a65fdfd043f3ea78ebabefb29f66b6302c6d84f7b5b2458e969f69bcab4015a3",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "e56802824ab2bb217e58f27697aa0b445acb4377be80742f81fa040b359c8a82",
			},
		}
		nodeConfigs[BsquareBtc] = nodeConfigsBsquareBtc
	}

	{
		nodeConfigsBsquareEth := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xA8455696F2DD85a6dce6b5AD683BF3602d73181B",
				SignAddress:     "0x55A352429fa28078E7a5133511e504E543230994",
				ConfigPubKey:    "946d36a0ad6b6f788981b2ba6809412cbac67a754eacb2b636e28de31a344f59",
				OffChainPubKey:  "485ef676066faf507c7ae03a9ddd9760cada31fbdd9dd7c7557abcd821c44678",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "cb2943970c78f86dc9c90fc2556aaf7bfe91651bdc316935520b462f99ff2155",
			}, {
				Id:              2,
				TransmitAddress: "0x3c892f62acfeDB80c1046F6215b31BF48A704636",
				SignAddress:     "0xe523D140be369BF62CA2E3e6194570a08ee79d41",
				ConfigPubKey:    "0f4f05a59a28126661e8bcbe271aaa760f33e23708f215e67d34a39ca545c26f",
				OffChainPubKey:  "f617dc99e003825c69684780efaaf0cb6113738c76ed7471bfa4fa37b3d71e91",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "cee84a6acd186309de82a22cfbc817c78293416ac26321d90ee964390888e91f",
			}, {
				Id:              3,
				TransmitAddress: "0xD11D0A08445F4F2daf2412f2D665e1fAD20F4f84",
				SignAddress:     "0xBc97f6B42182E045D451D62fcD25A1F0754373dF",
				ConfigPubKey:    "b25c1c111b45a503dfe5b75ccd5c3979cce96e51c1212cc8692ae6af8e6c7f35",
				OffChainPubKey:  "7b1ef1c1bd0a1fce07efb6f694a22cfb8a8054f90e5eedd3af4810a9ed212635",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "26410b47268649903db9562f9568b24396d107a72595fe84d1764a59363015fa",
			}, {
				Id:              4,
				TransmitAddress: "0x18AE11b760B3e1EAad18946cb3D64e97a565e3Ec",
				SignAddress:     "0xAF4c1a71040C6AAEDC36F13e7169f342256eE184",
				ConfigPubKey:    "6021eadcaab482aa4546ba1101df157f85d3ca7dde1577208b9ed5141f1deb6a",
				OffChainPubKey:  "dbc639fcc3b38bf19c135d63a9b72df8d6195dc032fbbb0f8a0abf1e8649655f",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "3c23c8a212eba392127fc2e6328477437237d36e1a723cab87f6f70ebcfcc9be",
			}, {
				Id:              5,
				TransmitAddress: "0x4EA989bDBC43701bAeab917c16E8d4d7f193b590",
				SignAddress:     "0x1220d3970DFa02e5a643125421fFB5405A18f9c4",
				ConfigPubKey:    "5abbd55e28bd5db54ee1f258a91e2a2764d58cc7c48efcb4b89b87df21e7cf42",
				OffChainPubKey:  "31a876308c257f2cbef36ab288356d1bc5057a0223d10bb7f575423089b9974a",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "9721ddc7799347583bd183c7ea51d386c02de11a01ce4e6b125d1050377acb38",
			}, {
				Id:              6,
				TransmitAddress: "0x43020E6C38C3a8ce5247AF292858Df5cF3e664c4",
				SignAddress:     "0xDE4a11a260236581433bFA8ac32776513106A542",
				ConfigPubKey:    "7879bf5695bc581fecf8abd0cc4173dcfbcc43881f6057bc9323e6af0a433757",
				OffChainPubKey:  "e4bb96c037767070aca82ebb7bb0d16f4e346a62c9edea42eded6c708418e26e",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "0d183b43fe786042ed42f5888341902fdf13d98b07532a29fa0006e2e178a093",
			}, {
				Id:              7,
				TransmitAddress: "0x7BDe199140F7bdBA1e259A26ab15775a1cb5e599",
				SignAddress:     "0xcD27480a994103a2d0CF966bfcbDC2aA50912628",
				ConfigPubKey:    "d4c45ad52541aeaaf59bcf075fd931b5bfd1e696950aaa857c34096638006f2b",
				OffChainPubKey:  "e57b3bcdc3b09cb89c2ebf69a1d6839acf1216e770f9f30109b66cef31221978",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "4bffd5ac2a0fd0b4d7c198bd3e71485039ceef0e34bfcc5743f9d071b97dd89f",
			}, {
				Id:              8,
				TransmitAddress: "0xa52074d8C7614cf4CaE28724c0C18476802e3434",
				SignAddress:     "0xC71309848b77Ed71F2332Af31C95c6a1Ae7E490E",
				ConfigPubKey:    "4570c4ff01a3f00601198393c6b2640ace81337ece75929a9722729dd6277327",
				OffChainPubKey:  "58a05781fbf7fa816904c2d5be766d9d24df109d4a0a3b7e436544bc6fbb4cbb",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "de41137c1e62408a4f5e992b588b0467efe79b13423926e75f67a7956d1b40ca",
			}, {
				Id:              9,
				TransmitAddress: "0x97cd58ED3B696C21128F395fB96f37550D201207",
				SignAddress:     "0x60d6EEa3465C0833aB229cA6474EcB13043Cf706",
				ConfigPubKey:    "49504034cbf9152d9403cf58d7221248b0c70b85020908a56c42b10e10a2fb32",
				OffChainPubKey:  "b3fd9a20459ccbab37481654867e884d82898e3e785942208db38ee31f3c3046",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "eac12c03aec1e413062d2925545d6266387b7b7a9cfc5b016397961d7e840973",
			},
		}
		nodeConfigs[BsquareEth] = nodeConfigsBsquareEth
	}

	{
		nodeConfigsBsquareUsdt := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xA8455696F2DD85a6dce6b5AD683BF3602d73181B",
				SignAddress:     "0x54Fc7242dD5eAF631c113Fc109BDf14E4b763D25",
				ConfigPubKey:    "13836fda75d25ac2245b00882fa92e00d823cc3b40dd66c58031957f8367a208",
				OffChainPubKey:  "38133c19c81872740d8ad2ee5a565b7ddd863c7cac5bcc01276ce1c938c639d7",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "d9c15dd26f321b5c311d568d1e0a636db97a5b35e8f0f7dbd8df8d69b1f0067f",
			}, {
				Id:              2,
				TransmitAddress: "0x3c892f62acfeDB80c1046F6215b31BF48A704636",
				SignAddress:     "0xbCa134A6AC9efd5f3206C8791E211148147dB52E",
				ConfigPubKey:    "d4fc69429a17abaade7093c9c2e73b2981dc6a728861165dffc783deaa691963",
				OffChainPubKey:  "ca05cd5676b60d5367f2989bb6ad049df6d4a35480fce47f2b83f51e6a8ba737",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "33a9afa00a4a439d1349e64f05b9a91972b46f9a2d26abe3721b2eb84e8cbf57",
			}, {
				Id:              3,
				TransmitAddress: "0xD11D0A08445F4F2daf2412f2D665e1fAD20F4f84",
				SignAddress:     "0x4735CBf43f9d0eFd9aeE22690d48604777185Ec4",
				ConfigPubKey:    "51a0bdc07ef7556b71d703bc8cbc949351c4cc749633c21b910c666e97ae506c",
				OffChainPubKey:  "5d058d44f715ec353874e50d09182debce29c1455223efff4e714d09c890083b",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "e1e45abb99293db7f863a82d4f34c1202512ff869b04f589203ab85fbaf4aa0a",
			}, {
				Id:              4,
				TransmitAddress: "0x18AE11b760B3e1EAad18946cb3D64e97a565e3Ec",
				SignAddress:     "0x7294C1559e705866911FF74911a8F6B7032d8714",
				ConfigPubKey:    "f61f1407b858279c5b43602c2c6d814b051522b143072b05f7b4e7c34fb6cb4b",
				OffChainPubKey:  "aaf3212f38bfa9671438fd2a1880023b319fb1339da1e8f26969e59e855050c3",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "c4a3fc82cfb293de2f210170a1f231671c42ea3e79325386ed3033b1a1c52389",
			}, {
				Id:              5,
				TransmitAddress: "0x4EA989bDBC43701bAeab917c16E8d4d7f193b590",
				SignAddress:     "0x8C7194D0557da2745079bd0798C2293A765520af",
				ConfigPubKey:    "17aca6926c177d6dd6417b87e76a093b5df31516abe86400d04eb79e6beb5723",
				OffChainPubKey:  "66b1d279427d653e9edd08086095649d0a5a04c37277abbde0264c059081f570",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "45efe1e949896702032c3f7faf59546cdbcc9382482dfdd197bd2d74a492c141",
			}, {
				Id:              6,
				TransmitAddress: "0x43020E6C38C3a8ce5247AF292858Df5cF3e664c4",
				SignAddress:     "0x784a771898Cae197f6dA13ddDFbe2980236D6bBA",
				ConfigPubKey:    "372b516f51b158dc12e1e2370fc753a6dc2c17cd56c5c657d2cc166065db0326",
				OffChainPubKey:  "3330ac8119c21f82db362287615771875165cdd5ce215ca7e5656b9f0595b7ae",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "009f32217c9c795c560dcfe630a6240137380552dfd38496566558d1d6309ecd",
			}, {
				Id:              7,
				TransmitAddress: "0x7BDe199140F7bdBA1e259A26ab15775a1cb5e599",
				SignAddress:     "0x72361DB8258d3FcAf7dAc6C767d5FE13bF3B2554",
				ConfigPubKey:    "398667eace34bc7145178d2e8e6bb962f7820880f3291d4b199b7e31e15cfb09",
				OffChainPubKey:  "a2902da813c1642f82bd091dc0a9d64620514bcaf8799bba523a3236d058178c",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "2636b7d1a3b6050d0a81e870beeafb3ecf5365665b79f5d9e723a0136c649dbb",
			}, {
				Id:              8,
				TransmitAddress: "0xa52074d8C7614cf4CaE28724c0C18476802e3434",
				SignAddress:     "0x020ed9753F8a996746096Ca58f617947578d7943",
				ConfigPubKey:    "11f505800ed9b9858904856e76781aed64dce02a4af0425dca496b251847465e",
				OffChainPubKey:  "ef24e06b3814bea84a4ee84ffeed1ded4266c344de16a5c9fffa3c25ee7fd3ab",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "e46e76310d444fa1c95c67b27e649dcf68b1528d4882ee88b21080d627f4dbd5",
			}, {
				Id:              9,
				TransmitAddress: "0x97cd58ED3B696C21128F395fB96f37550D201207",
				SignAddress:     "0xa36045296AEBA27D3D975038027ADc5fe14a45ca",
				ConfigPubKey:    "5a2a1f98b601198bcd021f1c4cd1ead016cc65b04f2a5c14d72023120762d40a",
				OffChainPubKey:  "5d528cf73d09df40a6621062ce15680d05eec7ccd711533c19692072c2281382",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "2968232dc26371efafb3d7736655cf3ea89fff48a85aaf2e7a323b2e7db40f32",
			},
		}
		nodeConfigs[BsquareUsdt] = nodeConfigsBsquareUsdt
	}

	{
		nodeConfigsBsquareUsdc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xA8455696F2DD85a6dce6b5AD683BF3602d73181B",
				SignAddress:     "0x55B7b6220fcFdEDe6fbCb13836835E2f13c072Ce",
				ConfigPubKey:    "5e5987feb8abe2cb260980306058d405cea2511de0793bfbe4dc1202f4fed01f",
				OffChainPubKey:  "f23c95298dc04d071563cf015ecbff40efc9f4d6b6b4d60431d0f7ce15117a38",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "abcfbfb40ff4d7d6ec83a08da4174789dbb1e8095f63b704adc10d34067d5388",
			}, {
				Id:              2,
				TransmitAddress: "0x3c892f62acfeDB80c1046F6215b31BF48A704636",
				SignAddress:     "0xFBea4AF7543AA2F2EDe06776c9E2D66A364052a8",
				ConfigPubKey:    "0a8d1157516f9d0053c10873c6235871b49d3eb910e37d0551e73d846e042d5d",
				OffChainPubKey:  "dfe5087499ab2140b521fcb19423f4389ed75897e7c8befe3fa0fd409996f4a0",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "f48f4e52d46ffd71451f24cb7559d54a33f554891c41dc765c5d0ce1bd5b6d92",
			}, {
				Id:              3,
				TransmitAddress: "0xD11D0A08445F4F2daf2412f2D665e1fAD20F4f84",
				SignAddress:     "0xb95e6A4D16B70A0e4568211201D1AdE730F4Df39",
				ConfigPubKey:    "6a624714728884fa7dd9bf7c399c9c9b25817b8c07c770c133ba3184f7b13e3a",
				OffChainPubKey:  "885b99aa7f6bc5e8bf321b96eca442584d38704fd0856f7abf9abb1912a1fdf2",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "1d09c02863effab89713f26365a79af15f99fe04436e10f5d22b6d3ee058582e",
			}, {
				Id:              4,
				TransmitAddress: "0x18AE11b760B3e1EAad18946cb3D64e97a565e3Ec",
				SignAddress:     "0x606EBF4F903D55Ce6e339a7ea8a2D0E34E9DbE17",
				ConfigPubKey:    "380e4c0cd5132b889904dc3f3d78a2f5acfd1d0c0d22a1930dfed0d75ae1e63b",
				OffChainPubKey:  "842e2f70752eaec761c79e5e794b0d93e22f3de1000a71e5f0b3f57b0613878e",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "8264154f1b26a0ec56f17e4b753cbfe2fec11b440d98cfce2a9eab35bcfc1187",
			}, {
				Id:              5,
				TransmitAddress: "0x4EA989bDBC43701bAeab917c16E8d4d7f193b590",
				SignAddress:     "0x6503633f2f4cdD19986Ad78bDF6b14BE289CA06C",
				ConfigPubKey:    "51ac6d0824d1d75fd19afbbaf5da3105ddbe2568e60abd40ba5a2b65e151787c",
				OffChainPubKey:  "96840769875bc4960e566a2419cebf02bee9f1281a0e942e325d2517997de8b3",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "133dc54b8ac8c4c215e85a3c9cf92285ed22c0bf5cabd235ff6a46fdebda78d5",
			}, {
				Id:              6,
				TransmitAddress: "0x43020E6C38C3a8ce5247AF292858Df5cF3e664c4",
				SignAddress:     "0x0385255eE3646a833c314144E479ECAB26D72C65",
				ConfigPubKey:    "b035d6aaeae331e250e68e8350e14c4722916a775e63c7a33c3a2d818b821b2d",
				OffChainPubKey:  "8e1eba5ce77aa081d75736435bf31c74dd24c03f42b7b7760049799e665085a8",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "d37384a8ad6061fe8faca76a00f30462b10af4c0252d8cb4d9cc267a1a482a50",
			}, {
				Id:              7,
				TransmitAddress: "0x7BDe199140F7bdBA1e259A26ab15775a1cb5e599",
				SignAddress:     "0xeed6667B3fACd6F2Bb7a13cc5e9f37dDfC27B6B3",
				ConfigPubKey:    "a4bb72314d868dd7974eda406f4ec12af1d17a12b7782672df9cd42610ab3873",
				OffChainPubKey:  "0cf60c2d0bfcff434cb9c60ec92b928e2e32227b5adb86baf22f3e63ff73200e",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "63bbe63dbbbf5e636741ebcf2d2f9fba548ba7e39add5b1affa0e0f174eb9e5f",
			}, {
				Id:              8,
				TransmitAddress: "0xa52074d8C7614cf4CaE28724c0C18476802e3434",
				SignAddress:     "0xbc43A5a8dEF000e5f2d621ac9b7B3A8AaE00edCD",
				ConfigPubKey:    "a5453fda6717ac84f5f80aa01ec99af10381074a6db38c7d45532359c5ba426d",
				OffChainPubKey:  "426d449cc83710a7b533e62a654d1eca8cdc2d13659f7ae3f29c4d4778b308cd",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "5701da5c618cc4f42bde6d71a60cc9e042abab98d6a8e9a1b6aec04146f3524b",
			}, {
				Id:              9,
				TransmitAddress: "0x97cd58ED3B696C21128F395fB96f37550D201207",
				SignAddress:     "0x6b04c8Bf1c2E0e25291E92b2C054D517D9b0c063",
				ConfigPubKey:    "ca8e0a93ba78b1810f573f6c5b562c99e04c8237ef5144ad803d6588d7e76722",
				OffChainPubKey:  "613ed4354a59f912b843a72c7233f8b1aaabf19d9e8ae3e1143a626883b4b35c",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "167e623adcc792f16daa27d792a6e395a93b30d162de32902d9a3572c8eac94d",
			},
		}
		nodeConfigs[BsquareUsdc] = nodeConfigsBsquareUsdc
	}

	{
		nodeConfigsBsquareFDUSD := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xA8455696F2DD85a6dce6b5AD683BF3602d73181B",
				SignAddress:     "0xd77575128bD0a1BC01F8032c316df2d095F87e54",
				ConfigPubKey:    "ff7c1888ee5677a0148de057ad692670c92d0d00a56a382b4b391625808a710a",
				OffChainPubKey:  "3b81107a5b36b6430ceca930ed6c95d0d1941cec43167c5023ed7c40ab9c8b94",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "8b2ec0442b78b271ad307a41efe9b4e4db26b9b4784b3e25288fc405bed8594a",
			}, {
				Id:              2,
				TransmitAddress: "0x3c892f62acfeDB80c1046F6215b31BF48A704636",
				SignAddress:     "0xB300762902b98C81bADdB7b4EE64Bbbd282D6f95",
				ConfigPubKey:    "4ac00cf4a2ab0f59bb131300d0637ac86207d7fbb6930477f40c814ecba45264",
				OffChainPubKey:  "a3a6dfed46b0c070bb79c2135276f2a661cb80b2d5c894f08757d443e8c7bfe2",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "e230a938406114c514113dd26dd19ea05e235d13abdf5b010292bf51f787c0f3",
			}, {
				Id:              3,
				TransmitAddress: "0xD11D0A08445F4F2daf2412f2D665e1fAD20F4f84",
				SignAddress:     "0x4A6cdcc7f2306B9A79C38C2ffe0C5B9f1D2D6a56",
				ConfigPubKey:    "92d5cdf5dd77bbc0334469ed5f9a2a308d7ad2393799da728257a2adeb26915c",
				OffChainPubKey:  "3a011d78d03d700f16846d657b92df1325423ee09c57df24c997166160f3a40e",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "ccac0c8bb446c8e9e76430509127aa49da1db5bcb02096cae14353b29979712c",
			}, {
				Id:              4,
				TransmitAddress: "0x18AE11b760B3e1EAad18946cb3D64e97a565e3Ec",
				SignAddress:     "0x938859d80724CECFcD7E8907F389c8264d6fdfe8",
				ConfigPubKey:    "7bcbc07261762f62a47152352dac677b130d9205016cc092558e96946940ad25",
				OffChainPubKey:  "179b8e02a715327be87e21d546e0ad8e17c5f7ed5288ed70046996a1896beed5",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "a333cbcf40e0aa8325ad014eb00955f1618cb26dce5e8c09017b992a4da9de24",
			}, {
				Id:              5,
				TransmitAddress: "0x4EA989bDBC43701bAeab917c16E8d4d7f193b590",
				SignAddress:     "0xA41d148f4aEdb91692c08aEaa190B768d1A200A7",
				ConfigPubKey:    "0d8d4461d63db375ffda48fe322f35a07d07f32ca2e9bdd7e2ed73f89e26f744",
				OffChainPubKey:  "dcbeda1dc7fa15b5f665d76391136eaa6b99b92ff9d5b6d99c851469de37590c",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "583dd9700ed3e5b6f879a551dd59ba9eb40affefed39f1930549ec232b95ab45",
			}, {
				Id:              6,
				TransmitAddress: "0x43020E6C38C3a8ce5247AF292858Df5cF3e664c4",
				SignAddress:     "0x3CB8c70ae590D95E62e52f2d02608a59BC264CD1",
				ConfigPubKey:    "26f59254e304d635a4e6c9b4351e6e0befac3bd3e2f266e3e74d585160fb9702",
				OffChainPubKey:  "47017a9ddfcaf66b0aabedea974d79a65c45a6ab84609241adbdc22a21ffc959",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "d0616833db923275fb3854a320e866674b25d1b1431a0421b75d9374ffe9bb8c",
			}, {
				Id:              7,
				TransmitAddress: "0x7BDe199140F7bdBA1e259A26ab15775a1cb5e599",
				SignAddress:     "0x9E0224f1D7Aed1eF9004Fcd901A701c9e5f6a89E",
				ConfigPubKey:    "036cdaff3979d7b3d76c6dd15b47f87115bc2a9d5de80904cdef21f37a8b1621",
				OffChainPubKey:  "1888e1b9f73f910ee90219bfa9db93c7152290936d3c26eca0abdb97fea3f215",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "39ff703115379a80a6ced0ae47141ae0c6775247bf7597f9a700e1f9aed0d2e8",
			}, {
				Id:              8,
				TransmitAddress: "0xa52074d8C7614cf4CaE28724c0C18476802e3434",
				SignAddress:     "0x59AE03D9994f25f1be73BaE800d69EF0510996A8",
				ConfigPubKey:    "6fb5acdcf5618bdd7afcace2eb59fcb3b1ee642e751b9f36687a49ca072a3f25",
				OffChainPubKey:  "813f6b8667b2c9f3bdddadf43c864f0a7c690aa89f49afacfad522378adf9502",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "b77475dfb5fc2a3173f516308241e646feb9331da29d24f19dcd7589e8abb686",
			}, {
				Id:              9,
				TransmitAddress: "0x97cd58ED3B696C21128F395fB96f37550D201207",
				SignAddress:     "0xaDFa394b73339B4bAf1A07C6a64B67Ae73a856e8",
				ConfigPubKey:    "a795c8038abbce47260234e420e945cbe9754bafae4e885f0d1508505a06d37c",
				OffChainPubKey:  "5f48f08e5d76b9b9542bcb1a5cb67064194a0754d4ef9f624e49b96d0d31d45a",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "3af1aca2d14a339bd7431ed92c1a01b244245aab6ddebe38a0a2ba70da27d915",
			},
		}
		nodeConfigs[BsquareFDUSD] = nodeConfigsBsquareFDUSD
	}

	{
		nodeConfigsBsquareBSTONE := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xA8455696F2DD85a6dce6b5AD683BF3602d73181B",
				SignAddress:     "0x1a31dd237c0a7626BCAfD6bD9Ed3ba8fcC914A5F",
				ConfigPubKey:    "906c33424805cd1dc74330a06fa761fe14f1698f4bd0a4e7ae2514ff488de07b",
				OffChainPubKey:  "a6a29f8435f1a6a498506efb1800f9aa82804eea58301293303ac1661a151183",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "4bb2ce01393d7f87e3efa6c460432a8b9d5b7ec4c8f6196c007e9bf75e96848a",
			}, {
				Id:              2,
				TransmitAddress: "0x3c892f62acfeDB80c1046F6215b31BF48A704636",
				SignAddress:     "0x163b7cFdAE7A0b10E38199999b091fD97E0F88D0",
				ConfigPubKey:    "a7ac14c3c66912875213762f213399f226528a460d49a2d85e99505697fe7a3c",
				OffChainPubKey:  "1208913aa7e97ffe6ed8f69e4e7aacf3da7bb4957a923e4e2670b7174a37b19e",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "8b20094f19c13d244d77325b5be80267399983dd81362199a06c87af7de4e467",
			}, {
				Id:              3,
				TransmitAddress: "0xD11D0A08445F4F2daf2412f2D665e1fAD20F4f84",
				SignAddress:     "0x08C815035C883941F97D702Cc9789D2e38dd9478",
				ConfigPubKey:    "26f25dfd48fb59e0005d82d3d58242a1fdcffe4afda2547744a46ba972b0b117",
				OffChainPubKey:  "adc799066f170ba995eb180d68a75d9a0e5904291b174f930345e24743782773",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "df048935d94e73cae78cc1d93980275a751ef5e553882bc0c58c47bf62150eb9",
			}, {
				Id:              4,
				TransmitAddress: "0x18AE11b760B3e1EAad18946cb3D64e97a565e3Ec",
				SignAddress:     "0x0A540b92b43d3E8E45f87E997EC42A886D1AED76",
				ConfigPubKey:    "f6feb41a8931753597501095cd57f38f28d391383c397d728371fed2e214ae72",
				OffChainPubKey:  "7cb0b97f3e242108d9a32fc4f6671b671061c177b0bbd530c3c863af01748415",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "37adf91b551a973f76e70947cab2f7c7ea4ed569651d3384be57fbf993fd6970",
			}, {
				Id:              5,
				TransmitAddress: "0x4EA989bDBC43701bAeab917c16E8d4d7f193b590",
				SignAddress:     "0x9b7bbfd7b57e347CFFA156262Bb34c2a784C7893",
				ConfigPubKey:    "b432b4487e0692eba59969032688df4a048a5b2c76e2aed18d63c9201da0515f",
				OffChainPubKey:  "074f5bf90bec79270640e03db86093e9c29f8c67c4f0dab22fe4a3a6bbf075f7",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "13b69485da2a3ad3d9c7dde0c539c04cc7869d467f039895f1a7479f835013b8",
			}, {
				Id:              6,
				TransmitAddress: "0x43020E6C38C3a8ce5247AF292858Df5cF3e664c4",
				SignAddress:     "0xd44b7C401E5E567B5e3d967e8F6c19cD305A1CfD",
				ConfigPubKey:    "a7367fc3d0fa54d10c77561c5aa9e0a4ff1e906d97cf56ae9bf08ae804f07c6b",
				OffChainPubKey:  "209b0fb1db9abb15313dda5230235fdf1ec49a7510c785dbb1016e61fd459311",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "c81bcc899e09a1d82f04986eae930bf22c3b12473ae257a570f74f9aad876052",
			}, {
				Id:              7,
				TransmitAddress: "0x7BDe199140F7bdBA1e259A26ab15775a1cb5e599",
				SignAddress:     "0xEaF4aCc147A39e7Da30Ff918E7940FE46AE4548E",
				ConfigPubKey:    "0b034d30aa9bab62332b39024de84af10241d875995ca7cbe607ea5704eb1864",
				OffChainPubKey:  "8ef4fc675c1538bf3a512f39e65e5c93023d74cf7183fd1e9c804de88594fb0d",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "445b9aa0a6fd63e3f07343f351a25663687b1829c7bfae59a9317588a49f50b3",
			}, {
				Id:              8,
				TransmitAddress: "0xa52074d8C7614cf4CaE28724c0C18476802e3434",
				SignAddress:     "0x61F4bF59574876b9ac5F3EFFbD3cAFE8fD1FD257",
				ConfigPubKey:    "51b90e72001e882c4bdeeef4bb9c03bce65c7b6e03784449e6297d2dedd1bc69",
				OffChainPubKey:  "05ee3e0ff46650b0e906d59977b0783be2955cd3c959af518872e9675f4ef4a4",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "77f0cdeae73152ac7be2dd52fa68af2bd48f3c7c23fde8e955c5f0b1fd613c65",
			}, {
				Id:              9,
				TransmitAddress: "0x97cd58ED3B696C21128F395fB96f37550D201207",
				SignAddress:     "0x1B1BC605eBCd64019c862E41421cc454eB0Bfd16",
				ConfigPubKey:    "89ef010a28c4244ac5d218b5b1a946ee9227e4de5e4d417725846df3d13d707b",
				OffChainPubKey:  "0f659fea394dbb2a0e7656b2eecaf71795c5611505b00e44d24922b7be5ae083",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "f9899ee9aec5ce336f45d453e0bfe1d4c7cc2458c73e8091b420756896b123eb",
			},
		}
		nodeConfigs[BsquareBSTONE] = nodeConfigsBsquareBSTONE
	}

	{
		nodeConfigsBsquareSTONEETH := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xA8455696F2DD85a6dce6b5AD683BF3602d73181B",
				SignAddress:     "0xe3E2ca363f95d73A09004D998edc847A2B54c3BE",
				ConfigPubKey:    "11cd7ed82bad0390e050572d8802206de9c18818f7baf095ea14eaad8a012c56",
				OffChainPubKey:  "d9830855578491061f4cc28f37bc42acd1b8fa33c325dc729923b0ca5afd88db",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "2ffdb79c9f333d900627bd2aa3fc852fc239a703602fb95b806791565bbeb38c",
			}, {
				Id:              2,
				TransmitAddress: "0x3c892f62acfeDB80c1046F6215b31BF48A704636",
				SignAddress:     "0x593256dd35A30fdfD28D911354694bBCD9ffFED7",
				ConfigPubKey:    "17895dea0e314618b7ddea718a7bd800159fde702577ad17dfdb1d96f26af07e",
				OffChainPubKey:  "56ebda6aed7b818b5351030b1c745e46eb3fbfcf7fdb30a363a0caa9eb34ce6a",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "1dac70a0e9e974c252f868406fd29b505717897393a505da65364d4abbc56632",
			}, {
				Id:              3,
				TransmitAddress: "0xD11D0A08445F4F2daf2412f2D665e1fAD20F4f84",
				SignAddress:     "0xc5386D9E945b9C5d2126d54d4f48B3224040D0Ae",
				ConfigPubKey:    "4db5c3b2651d5ecd1c3c9291dd860f7d09f94e2a59c8e718f95d11b162b4182f",
				OffChainPubKey:  "5c402b260ecf345a01ed630973fdc68bbd30532a8b15688b583b7922372cefae",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "9c1b2338c02ae0c25d5e367a001c1eb61a7a5b333f961bf2ad75ec1c13b1fc97",
			}, {
				Id:              4,
				TransmitAddress: "0x18AE11b760B3e1EAad18946cb3D64e97a565e3Ec",
				SignAddress:     "0x2169f620e48B4429451711483e4aD767251ed50d",
				ConfigPubKey:    "7acfa9058250c8080eeed629ce117196f34fc7bf199341881ced2f95b4690b6b",
				OffChainPubKey:  "72a54123f72fd2f8b7305f061a6cb2467651a01fe8d58b113ad4a7882f97b1e3",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "530a4fdf85bc5019fd4da09b661559b41339cb06b39ebc3a0452d4a023fd5237",
			}, {
				Id:              5,
				TransmitAddress: "0x4EA989bDBC43701bAeab917c16E8d4d7f193b590",
				SignAddress:     "0x287311AfcD0Bd7Da0CC6D52F3443176F7512D019",
				ConfigPubKey:    "314a7b93c043f470276a2d0161925f799cd040fe5e19b7edd952a7fcc7f4541b",
				OffChainPubKey:  "755e2ef89f3dcca2c5f4226a3d4eb2ab4e23a0cff3fe799fba27dd1fc95b32f3",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "4fc86e775ea64d78594685f23e07e0fbeed103c660b7d85142d80c04ef5ff23c",
			}, {
				Id:              6,
				TransmitAddress: "0x43020E6C38C3a8ce5247AF292858Df5cF3e664c4",
				SignAddress:     "0x3e6522007ef8c967ACfb7037DdC4f1731d6A57E9",
				ConfigPubKey:    "771a1988208c48f5833c35bb91dadf36e07d20a4c4b60323d97ab1b48a264004",
				OffChainPubKey:  "1db4e970ccfb8ac9f8554d1bb404fcd5276a39bbbb25060192f687fc5594ea1c",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "b385603d39ba582cbe7b3f8b264d6ab4cbd4dedb358be69da93ec59d88e6ba48",
			}, {
				Id:              7,
				TransmitAddress: "0x7BDe199140F7bdBA1e259A26ab15775a1cb5e599",
				SignAddress:     "0xd41797bb92d89778a30418CC075D241A1f307677",
				ConfigPubKey:    "1ac781cc84c2a7c7371c5d36b29f4f45759d6030257f1707d0150bde8f11d379",
				OffChainPubKey:  "690849339e42f469732fe6762b9128c1f010a90507a95affb5317514c75a1309",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "7ad3034e7a2858926c31eb8bab64b643780bf2bd462a7bb86a63e980094d6f40",
			}, {
				Id:              8,
				TransmitAddress: "0xa52074d8C7614cf4CaE28724c0C18476802e3434",
				SignAddress:     "0xD8EBBCa70E7D0131a7B4E6E531A212a4F5B90839",
				ConfigPubKey:    "cc3384fe45fc9e0c1a90dc33adf4977a46a5704de678866d5051a24c8ea4e950",
				OffChainPubKey:  "5ad9595799d23108c4e8c9fc5511c2bf7f2dcd7a395b51ef41df8ec99f479177",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "51f35df80f194816519b0342f8ed405f1845917ea246df13d5649ffc45e5868e",
			}, {
				Id:              9,
				TransmitAddress: "0x97cd58ED3B696C21128F395fB96f37550D201207",
				SignAddress:     "0xB6a0bDa4C71913634efA4605C128018419bfA76C",
				ConfigPubKey:    "d26ceae71dd70c5e947da2677db6daca92a25f68ff97430a220f2e586d4c3410",
				OffChainPubKey:  "5f56b9463485ba5cb331a950e116c7285ad31ceb0a8e8463956b57404bbf0d0b",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "1a9b61bae308aee55bc8527431b1fc08b5c7a34a1bddcc6b22fd5442037e78c6",
			},
		}
		nodeConfigs[BsquareSTONEETH] = nodeConfigsBsquareSTONEETH
	}

	{
		nodeConfigsBsquareMBtc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xA8455696F2DD85a6dce6b5AD683BF3602d73181B",
				SignAddress:     "0xfb7AE6A9b598Ed97a3312e7DF233481Be5204a5E",
				ConfigPubKey:    "4b7f0e3026318165757ab5b04e5bc359ff09b5fbd616238c8aaa11b7ffd20950",
				OffChainPubKey:  "eee25f0ac4cc52751803ddf881bb530075dace8e6fbeb865dd80279ae90e3f06",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "2870d57332e1cad79f76b6dfe1e7681c9be2cb1551c8f81b748d3a1cda8da040",
			}, {
				Id:              2,
				TransmitAddress: "0x3c892f62acfeDB80c1046F6215b31BF48A704636",
				SignAddress:     "0xb21D333819001bF4Dc38b2bc288EFDFD605d2043",
				ConfigPubKey:    "ae3f70ac44590a7d4866c22c39f0e6926ab8add51ebf51b1478ccc417d19cd30",
				OffChainPubKey:  "171331b23c8b4f701369366852914092439d4f4e7ae807f74cb44ce2459595c5",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "db04318d3b77e78e2311c9e17bb659244673a4c73709f99c41294b31ede09161",
			}, {
				Id:              3,
				TransmitAddress: "0xD11D0A08445F4F2daf2412f2D665e1fAD20F4f84",
				SignAddress:     "0xae166365488A3041dbCFAD0632d08a3564fD8491",
				ConfigPubKey:    "6e91394dc0f30068d3e3921a826ede65cb6a6d11b58d24702544c016ce744f0d",
				OffChainPubKey:  "53e2fb2c7a37f15cb2553817231e4502a3025cc89976db544b2149b64cdb8edd",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "54151e64af89629ce8b9f01ea29cc5f1d436688d0a8716717793607ce4b46e24",
			}, {
				Id:              4,
				TransmitAddress: "0x18AE11b760B3e1EAad18946cb3D64e97a565e3Ec",
				SignAddress:     "0x57742f2915226E3eDbf5c2640C8Ae30Aae2680F7",
				ConfigPubKey:    "ec7a122ba6c4f7cb5d23c94ab1d785634450f38084645943bcc65abaa0e5c56f",
				OffChainPubKey:  "13f470d317927154d648f184432678c24a4de53e95c7e8ba18090dc0118ab40a",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "a2bd577b250e616edcee88cd877002fc8b9beb9a464e485792e2b60261f11a4d",
			}, {
				Id:              5,
				TransmitAddress: "0x4EA989bDBC43701bAeab917c16E8d4d7f193b590",
				SignAddress:     "0xF009ddC40Be8982Ca711619a37b823fD25b70D50",
				ConfigPubKey:    "dc8163d676e17745b8cb18e2048c3124b86dcb298c0bccd8a60b6d7f76eb3612",
				OffChainPubKey:  "b81e4eee83726cfd1545c98d38b7c393eac25ba90852ad8753e9c317370c7b2b",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "4a7ab2511c7c06b6f5c9def5b83d1814b14a4694ccf22054e41a340519bfab0c",
			}, {
				Id:              6,
				TransmitAddress: "0x43020E6C38C3a8ce5247AF292858Df5cF3e664c4",
				SignAddress:     "0xAF446BEb530FF5ac6869f942833078Bbdd701982",
				ConfigPubKey:    "8292e434ff116d500c32c92d6f23fcf0fbf1d08c095a69d98560b7c444816a04",
				OffChainPubKey:  "d173d6c920f73b4ebfd59094bd0a1c8ba196435d7bcec09d89d57bdc457d1ecf",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "fbfc01e5761ce14e31dfd64a08ca3ebd00b55cf166094e444d075bb7bf4f1b3d",
			}, {
				Id:              7,
				TransmitAddress: "0x7BDe199140F7bdBA1e259A26ab15775a1cb5e599",
				SignAddress:     "0xE294d3E74e1c0ac23ADd438720a6EAA11fe25D46",
				ConfigPubKey:    "08ebf0e92fc30289707472781a1ed7d0ce1e1fb29315c803480904dcd65d0a58",
				OffChainPubKey:  "4ea2773cabd25d94af9c012170f598eebd06a06172e6f40e0bb66cfbd21538b8",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "7634651aabf978be7122801fc631da6af5891f3b83f141bfabc867a53fe82d9c",
			}, {
				Id:              8,
				TransmitAddress: "0xa52074d8C7614cf4CaE28724c0C18476802e3434",
				SignAddress:     "0xda02c0448AFBe8FCBAcaFc00c893163a59F8a25A",
				ConfigPubKey:    "b07aa079bc32b52cd2bd601b84c5fce70a0d5c124501c368ab437aefce2fd713",
				OffChainPubKey:  "14c1cc761666469813d4064a44cca4396b91b8310180f99f78d91c515aef85e3",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "d6f6e263272d29bdd8d495ff8f9b52eca3ca23c32c8dfb430733cba04570aeaf",
			}, {
				Id:              9,
				TransmitAddress: "0x97cd58ED3B696C21128F395fB96f37550D201207",
				SignAddress:     "0xc752A45F3d22d534339BDA3f4ec8369C842ED4d9",
				ConfigPubKey:    "c0ca8e6e26030fe25223b7865ca2b096addbb12bfe51563c6920a1b3286ecb41",
				OffChainPubKey:  "9b441514dcec9b28a08898e8a41a3616689f14593ff085fb67189f1edd3ce129",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "9621b95e22f42bbe593576a72ad4114961e6128095e00e3c86413e8e6f19f8b4",
			},
		}
		nodeConfigs[BsquareMBtc] = nodeConfigsBsquareMBtc
	}

	{
		nodeConfigsBsquareSolvBtc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xA8455696F2DD85a6dce6b5AD683BF3602d73181B",
				SignAddress:     "0x3e843D47325A44281faaAf69C2eECC97AE15D552",
				ConfigPubKey:    "53fc049617c7755965f1a196e6b873311ffad331428b6b100e9de0eeed0df22f",
				OffChainPubKey:  "e55ccf7ce14ffb0901a39c59f50dcc44cd5fd461cac2baa0ed880992f2afcf1e",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "936bd808a19691effa3b8cf33780982e39f6feffeca449c673a1e966684a4e7b",
			}, {
				Id:              2,
				TransmitAddress: "0x3c892f62acfeDB80c1046F6215b31BF48A704636",
				SignAddress:     "0x7A08992d064d2E02dF8B0dE607348931ce265c51",
				ConfigPubKey:    "41c64089091bccd7d1d72bd893976b1cd9edca40c97e286d34fe78ff17824604",
				OffChainPubKey:  "de34e4d3c9b462c83865a4f6a1c4174d5a06eff3a0e9e8a12f0689f61b2d7fae",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "e45ad95649afd07349bb6ba346dcc6d6d6e09c753bab86464fcc7a7c423d7fed",
			}, {
				Id:              3,
				TransmitAddress: "0xD11D0A08445F4F2daf2412f2D665e1fAD20F4f84",
				SignAddress:     "0xdaE283eBc1974FCE20e368745EEB64b9B172a52f",
				ConfigPubKey:    "22ebb114ec461064c92c3fcab121817f1fd0bc0f85fe33c25007e5688b120c2b",
				OffChainPubKey:  "1531d80ab6a93ee5ac3057e03d511b3fb5e1af1bba68e93e36cb75d068f20a3d",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "50e0512e0e6ab4fb646f094555b09b8b8d2910e47e57ed94cd14fba21a0abfad",
			}, {
				Id:              4,
				TransmitAddress: "0x18AE11b760B3e1EAad18946cb3D64e97a565e3Ec",
				SignAddress:     "0xFB4CD0216DC4846ee1ccaF76AFf8d1F41b95191F",
				ConfigPubKey:    "de196ac081156db017a0b0a0117c7bc902c10155f261422b648f5df8c2a2a840",
				OffChainPubKey:  "2b0ee6873f0ded0d7264a16e358c55e568be8615e293bf42dca8109dbda919ea",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "881f79d9b022661b461061354f2584a0dcfad8e5b1c15278ca245ef0d97827eb",
			}, {
				Id:              5,
				TransmitAddress: "0x4EA989bDBC43701bAeab917c16E8d4d7f193b590",
				SignAddress:     "0xB0C8d539eE0f4CA5732c62B796CE14961DdE6e04",
				ConfigPubKey:    "386f9f1fd37fe31ab105b12e61ed7a60dc756b3996750ad411f694d9961ca33d",
				OffChainPubKey:  "8b8da6ea814b516ab34dfdc173d47eacedbcb9a8a748e0f1624f9abe9af1a6ec",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "33e9eb35c3fe714a9abff1de8cf2967c46ee64b610c3ba09cca98b2d0b8ed698",
			}, {
				Id:              6,
				TransmitAddress: "0x43020E6C38C3a8ce5247AF292858Df5cF3e664c4",
				SignAddress:     "0xd08D12E081128b07D6042C3a6460853c02cdc5f6",
				ConfigPubKey:    "ede0849dc87ac82b5d3bb5b9352430bf2de598f59444e93edaaf288a8b1bc36d",
				OffChainPubKey:  "3c127b6f246f27d80827925edae478f2764ac57ec135acdc175efab5d7fa2322",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "d96d2a01954ffc8dbbef00678bc7f2567a7b569e76b80999e5ef65a3d706a586",
			}, {
				Id:              7,
				TransmitAddress: "0x7BDe199140F7bdBA1e259A26ab15775a1cb5e599",
				SignAddress:     "0xB35a92Db52ef96a49c005c8D20d09F0cd620d30D",
				ConfigPubKey:    "a3ee715eb47da751b95ecdd07cadb0fde15f9667ec32af32dfc00b77c714c341",
				OffChainPubKey:  "dfcdbabed63505c79d0774ff822f566d4855d3d3f88bceda4d67d4ff21f1adce",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "c78d1fa6ae38d0b61f1e9ea044662a6d1a5b99acaf3a2e33192a5827fc3bac2a",
			}, {
				Id:              8,
				TransmitAddress: "0xa52074d8C7614cf4CaE28724c0C18476802e3434",
				SignAddress:     "0x07dbC95a166E3547000E622327b6E28a336D1884",
				ConfigPubKey:    "b901a971e62ebd6698ba95a77815edf3610083936b55780a454cc8c652fdeb13",
				OffChainPubKey:  "3e0a32911a9ba4bf28a4de1e4bb2491bad6aced7a08560ae66e4cfd72f8da5fa",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "2ad8e3b9ad9938d107cf291170e03d6a57e600c0878b62919884c69ec1f779de",
			}, {
				Id:              9,
				TransmitAddress: "0x97cd58ED3B696C21128F395fB96f37550D201207",
				SignAddress:     "0x63aF9B67d57c898151131BAC61619A932E7FA323",
				ConfigPubKey:    "25893edafa799c7c3e7464846d4be8ad76b787c4c2a4a240776d2b387b25b376",
				OffChainPubKey:  "945f1ed44e1f101418d820c7cfdce9bfe8a149dc5904676dd4b65d33d6483dea",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "3a38fc2a9446ee7c4280910fdce06a37fe8ac840fd388ce0b361d040928a4ab5",
			},
		}
		nodeConfigs[BsquareSolvBtc] = nodeConfigsBsquareSolvBtc
	}

	{
		nodeConfigsBsquareuBtc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xA8455696F2DD85a6dce6b5AD683BF3602d73181B",
				SignAddress:     "0x95b3EE15508e181b3e24d459f5593aA6F14c91d4",
				ConfigPubKey:    "50010e224f17e7a419de94312429ce6959b4f9413255d5e5ede99681c3a7be43",
				OffChainPubKey:  "2155a1baf75c6703d60636bfa4fabb53181e149480d124e88027eeb2cb7bbb4b",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "eb7bbd9c2e0b1829f41a61172c95ee7f30e043153b66425a5695b66bd57554c1",
			}, {
				Id:              2,
				TransmitAddress: "0x3c892f62acfeDB80c1046F6215b31BF48A704636",
				SignAddress:     "0xAC0AF6DdE944A2709D98e7130B6C3Fda83045089",
				ConfigPubKey:    "63561e763a4b27197957d62b9e8d69a111d90d665697eb259121bd51ba5f9b29",
				OffChainPubKey:  "6493a4e1e5b5bfd726249714c214af9bf4259403412959299c2e1e8f08bd0261",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "3cbb90f6cc76470ae37858726ae94d2042b2df76404d626b28032a2b270617f1",
			}, {
				Id:              3,
				TransmitAddress: "0xD11D0A08445F4F2daf2412f2D665e1fAD20F4f84",
				SignAddress:     "0x72bac9fdC7ccd487523374D2b59Fd73F9F2AdAaf",
				ConfigPubKey:    "d66489d58ac9a5a38b9b819809ada9c54f41be9ccc0d1f5862e58946a9909821",
				OffChainPubKey:  "02f3c81a90ff5f5dd78325db9e6369a3dc10b154e4ead13c0bab89f994e36e4f",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "a07fda356b7e38aeaf6e32a84dd427942f80a71220fce9cd2edbcdc5975d9383",
			}, {
				Id:              4,
				TransmitAddress: "0x18AE11b760B3e1EAad18946cb3D64e97a565e3Ec",
				SignAddress:     "0xf05FE6D50A5126918bC9f55F705dd74CCf3F0857",
				ConfigPubKey:    "fe2f11d619d43224b19373cd91d17df9a740b04f0a67d97661ac4c20b5b97135",
				OffChainPubKey:  "7a9d836d9af4fa07621bf30545bf5af25e187c30b99bd7f2f27672bfca1dd71a",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "9ebfeb4cadfed4a7aaca1a8fe47f2a8488b5e4cfda1ec02049fd14a1338a45b2",
			}, {
				Id:              5,
				TransmitAddress: "0x4EA989bDBC43701bAeab917c16E8d4d7f193b590",
				SignAddress:     "0x08E1aa58C4A1955B2AC661CD4bBD7A540B02005b",
				ConfigPubKey:    "302b1e4ebeb58289e83a0f15ef7ca10117d936b3ed09d9a5b00725a069423b44",
				OffChainPubKey:  "08700d2e057fb16c003c0ec6761351dfb52307cc07fdd6c340db1b03d89aa98b",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "a1699fce1e05e0935dff04bd44dcd402bde7b0207942b24217aa359852c2043c",
			}, {
				Id:              6,
				TransmitAddress: "0x43020E6C38C3a8ce5247AF292858Df5cF3e664c4",
				SignAddress:     "0xD0120D464391c639Ae5458609900BAe1AF542655",
				ConfigPubKey:    "73ce455aeea7d3c65b6342fb63bd8ab234322fb92d4434791e172a3878b8eb2d",
				OffChainPubKey:  "38e5f447fea0d63fdba192fbeba29bfd30d06c6e3278a85e2f6849a358abde6a",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "3a36ba1b3a8052ba123178b1fe9fc015035857c1291c2b19fc7638e6484a227b",
			}, {
				Id:              7,
				TransmitAddress: "0x7BDe199140F7bdBA1e259A26ab15775a1cb5e599",
				SignAddress:     "0x9C583678aF0B237ab8FAB6b45c978078f14EB940",
				ConfigPubKey:    "ad79899186420a898b5a263d571628878d3adc83f15684c150a17d7ebae07e32",
				OffChainPubKey:  "4e14d23e7db10aa03673a04369e562741116b58ca16022488ca166e67f97234d",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "14830a060aeec76e94b7b7dcecaf5e28616d2bba52396315f1094ee7251a8627",
			}, {
				Id:              8,
				TransmitAddress: "0xa52074d8C7614cf4CaE28724c0C18476802e3434",
				SignAddress:     "0x3cA8507fB6Fd41C4745b6BD429E08030Dde6a0b8",
				ConfigPubKey:    "8bc50d8f429f50e05d1987f2a0c62b0f9467fd1356d5bdb6aa1b510e45601536",
				OffChainPubKey:  "176e2694a2eeb459de60e277fa4155d331fca0cbfcdf5f1b32d2170d6f183886",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "d86142a4279c9cee2cbb439c8a445f90a23af424bad5787e91eced5efcaec460",
			}, {
				Id:              9,
				TransmitAddress: "0x97cd58ED3B696C21128F395fB96f37550D201207",
				SignAddress:     "0xfada468ad0143D91d803ca29b07Ab7C43646CA8b",
				ConfigPubKey:    "015a89b25b8cf6fa0b60ff87d31bfc9fc365eef7a8b3337907d18b60f06c3249",
				OffChainPubKey:  "761a13a8132fc9030d37dd3bb64c3dc8e328a6cc9097acbea503109d18dfb54d",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "036a66af6e0b90f4d117661413bd93acf9bd79de216b7b6e9d9209cfccafba79",
			},
		}
		nodeConfigs[BsquareuBtc] = nodeConfigsBsquareuBtc
	}

	{
		nodeConfigsBsquareMBtcBtcER := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xA8455696F2DD85a6dce6b5AD683BF3602d73181B",
				SignAddress:     "0xD45c32865f34C3bC3d70Ab97A6a45Dd3C1EC1eF5",
				ConfigPubKey:    "fce563b7a82c51982999a45fc9f97e7c56a9a4bd573da07be1ee0f0701ab7305",
				OffChainPubKey:  "3dd474bb78e0cc5d8ed7ba9adb92fef6189561019d4c6a208de77da5ea501cfc",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "b1fd3f14843d6b8a2afbcef3699a05d5e4d3ecdb6f7354cfcb05f294106997bc",
			}, {
				Id:              2,
				TransmitAddress: "0x3c892f62acfeDB80c1046F6215b31BF48A704636",
				SignAddress:     "0xBe7973a7C3f4835Da00B47944aE3b626732d50f1",
				ConfigPubKey:    "9a23bc2b160c61ae5116df994208de549e190f59c4c0ca843cc539864229207d",
				OffChainPubKey:  "cfee1bb3064db55a21a67ea74c42d08a07f128a71005e843e782bf66a7bbb89c",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "95ff3ed6c6592a31ac14bbd8baa358889728b36986d2bc0cefb51112fadfb120",
			}, {
				Id:              3,
				TransmitAddress: "0xD11D0A08445F4F2daf2412f2D665e1fAD20F4f84",
				SignAddress:     "0x41D9229cB25E47525756D955Ff379E03A9982C5C",
				ConfigPubKey:    "e5e9ceaeea0900c8926beb614f224227a9542b45243fc1ac13d65e69cf423a37",
				OffChainPubKey:  "dba34ceae35fd61768b03eb2b83aaad7e4b5e82241bb687f326600bd4cb79d73",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "9f7142bbd5bf87312035d837f3223b8dba3a6419109fb75c8bbe306d18c511b5",
			}, {
				Id:              4,
				TransmitAddress: "0x18AE11b760B3e1EAad18946cb3D64e97a565e3Ec",
				SignAddress:     "0x80Eb314B165445AD475927C93901F04fAACDdeB2",
				ConfigPubKey:    "399b9710704ae4b1cdc4c6ac22ff4e30f07ba5a41fa35a43ab0c67de07970f5c",
				OffChainPubKey:  "2f3b034807a0dc6026ab05bc888f3e10657538b0527ae3fc37fa55ef30f3d043",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "6832da69b121cc65b3e5defd0f80cdda4372b2e0c7ba8b17b6742c06f29eb951",
			}, {
				Id:              5,
				TransmitAddress: "0x4EA989bDBC43701bAeab917c16E8d4d7f193b590",
				SignAddress:     "0x5265e46c3527a6FAD6b6A11D700caf9A7e2F56BF",
				ConfigPubKey:    "f22627ae053bcb9e9a1ecfe9b84a65e83efa56d5660d36e74d29c3fa410b5c65",
				OffChainPubKey:  "3493a0ce0feb85541e81d002c43e79f92ff5ced697275b17bdbf1f38ab040399",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "b7eeeb3c361d3088fce11aab80121a4db020d5df6b6389ee13b98a8dc6de7bb5",
			}, {
				Id:              6,
				TransmitAddress: "0x43020E6C38C3a8ce5247AF292858Df5cF3e664c4",
				SignAddress:     "0x95e6D192DC8042921401694Fc7E5CF4901d6986b",
				ConfigPubKey:    "2551bf193666498035d99db959cc9fdd248b5dc4a0b15f3458839d886321a22f",
				OffChainPubKey:  "5047693afbe0fbfa1d44024750c166b998954fbfdb8dac176d66282f31b973b8",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "fac3b252c755efe7519ed720cb6da45daea82ef3570f3575ee55b8d9b81f9988",
			}, {
				Id:              7,
				TransmitAddress: "0x7BDe199140F7bdBA1e259A26ab15775a1cb5e599",
				SignAddress:     "0x8F8219a65C1B8cAe2b31f6cFE2AA8AB310D8f920",
				ConfigPubKey:    "5b9cd65d5f89c27ad0dd5b219f4909ec86cce6eb4e137e630587f965e22dfc44",
				OffChainPubKey:  "6c1db93098071ea38632549466579d65007447ded56a62aba6812ce9bfd36113",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "93b58297d150dd5454bc65fbd80f3ae5e225ed89f1eee38d7ac230738c851b87",
			}, {
				Id:              8,
				TransmitAddress: "0xa52074d8C7614cf4CaE28724c0C18476802e3434",
				SignAddress:     "0x752607C4AF33C1a23c9210b5e982Bf205630f069",
				ConfigPubKey:    "fca8248087ed7477efa3b979c9d29b03d29f8292ed6f57c566207e12e126a708",
				OffChainPubKey:  "454ae027262f27ac2ac3d4016ecd07237b699396aa4d67f2ee2e2d653a4afa57",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "48d3a1cd53ddf2fbe0909486b0fcc7b9b9605f8ca9f54a31733215e3a9abf146",
			}, {
				Id:              9,
				TransmitAddress: "0x97cd58ED3B696C21128F395fB96f37550D201207",
				SignAddress:     "0x533ac1Cc8374BcC60E549AbEC2c2EfB4Ff837A8c",
				ConfigPubKey:    "64e3cc8a533b7766a922266b73be94f8773ef927d92ab9a59508ba0e46a07a3a",
				OffChainPubKey:  "5800edfdf456e87b7964a39b16544459cdb01b45eb5f3bf26c84dcc719bd9287",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "532dd97fbb4b40972c5b50ff64cc178071986519e9252ccc8d9101a8f88a9a59",
			},
		}
		nodeConfigs[BsquareMBtcBtcER] = nodeConfigsBsquareMBtcBtcER
	}

	{
		nodeConfigsBsquareSolvBtcBtcER := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xA8455696F2DD85a6dce6b5AD683BF3602d73181B",
				SignAddress:     "0xFB4aDD1080b4b7CC65CE007B46A02f9FC418D3f5",
				ConfigPubKey:    "72cd706bc01c14477bb196bd31638a877a30a09ec5f4ce6bcff04cf6e7b79966",
				OffChainPubKey:  "d026bcedd10322a31ef8df1ea15380d5a2018b88438abbf6b6ea12c93e8db775",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "1ff787dbdf983f9a8ed76589cb141af46c201e2264b98532a6a4f2541242e6b2",
			}, {
				Id:              2,
				TransmitAddress: "0x3c892f62acfeDB80c1046F6215b31BF48A704636",
				SignAddress:     "0xe86EedfD7e4Fb5Aa111f72726929F929509d9ca5",
				ConfigPubKey:    "4f264e4a1f0cd9939f4ab105d764f1d206b49cee2e0a7bef27fc7d0621c10833",
				OffChainPubKey:  "3016545fd32d3680c105c0f5181f90f74483c2abf3c3efe6736fb79aac9795ac",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "c433d88690542ba0c43bc2252a28f11b5d0c1263dc88e4b7db2bbc71096c382d",
			}, {
				Id:              3,
				TransmitAddress: "0xD11D0A08445F4F2daf2412f2D665e1fAD20F4f84",
				SignAddress:     "0x4855A2C7b256eDd374933beceE7712FdBd87ea20",
				ConfigPubKey:    "47a7792b3319b10847b45a4617dca41f2d1ad2aff0ff77629913bdc85c31256a",
				OffChainPubKey:  "5efc537912c425c349928bae0a7eb2ce4b027e00793949992b66363fdd8b7157",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "52d183dadd7cf17cc44dd5c47e804c9a1d8106683eadc9f77963519b2988b9da",
			}, {
				Id:              4,
				TransmitAddress: "0x18AE11b760B3e1EAad18946cb3D64e97a565e3Ec",
				SignAddress:     "0x7c7DD390C91301A5818332bb2011A6631F5925fB",
				ConfigPubKey:    "598968bb07a069353f288d10e39eccb0c56cab00f51b523bbc4fb1469444bf06",
				OffChainPubKey:  "367627c54c5769c207df64892d99cfb70ecef95d2f60cb7e591b5285a9353dfe",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "3f522bee3a97f1af2099eb14433c559db2ca8f0b8adb662a25caedca995abf0a",
			}, {
				Id:              5,
				TransmitAddress: "0x4EA989bDBC43701bAeab917c16E8d4d7f193b590",
				SignAddress:     "0xcFE9dA5439d412498b79641a807c5D8c12bC6632",
				ConfigPubKey:    "8d633694d72de47a52139f3e4f99559722b49ea32725482e39a6709aa5b4c849",
				OffChainPubKey:  "2c00491d740d7a32d7efe2bef2809d18660eaa0decf744ff6cf201987a4fb688",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "3c6b108cb34760b421aa7d3b38d1e5ca061f067f63644f47c455eccac3434381",
			}, {
				Id:              6,
				TransmitAddress: "0x43020E6C38C3a8ce5247AF292858Df5cF3e664c4",
				SignAddress:     "0x0F967c3286eea71309490f3aa4AaeE9F8E88ef39",
				ConfigPubKey:    "f2b41ded39b5b1d7ec1c7fd0f9349596e66e8ed44fe12300022ee95957bf5f34",
				OffChainPubKey:  "15020955deff93b567bcc9596dee4801d8c6a9ac211391e0d88112dd945fad0f",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "186dabd696d6028339d0c16d4dd9d465fe2939199fb70691536f2b2ce574e807",
			}, {
				Id:              7,
				TransmitAddress: "0x7BDe199140F7bdBA1e259A26ab15775a1cb5e599",
				SignAddress:     "0xd56a6C5d0fCd697429f66b9CC3FE1B11f5ABf67F",
				ConfigPubKey:    "175cf7246de8bad3cd62e02a05efec382a9e61459198f2c37c3028bb1524bd47",
				OffChainPubKey:  "e846516ee922ea059e354a050404a067d3dd46c30cf9f07d4cd49a4a10213a7c",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "3843f9d95c737897da042e608181ec3031437d5c9b35a2b3136001e5b5d30d1f",
			}, {
				Id:              8,
				TransmitAddress: "0xa52074d8C7614cf4CaE28724c0C18476802e3434",
				SignAddress:     "0xbAc2C00611DdEA0a0961740d4D62A03D15036749",
				ConfigPubKey:    "baea5a05752bddcc83e1b6b82a62f359b9a80ca91ac57f31208d16a00fb48864",
				OffChainPubKey:  "cb53b439717af505d264eb96887fef0148823434ee63cd3a68e1d0993c9c1d93",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "402a057b66a622ae51b3df38c6cf293108abc241364929be552cc7ab5a5e3487",
			}, {
				Id:              9,
				TransmitAddress: "0x97cd58ED3B696C21128F395fB96f37550D201207",
				SignAddress:     "0x069E05082b7Fa1B019A39809a8dc0725C16DbbF3",
				ConfigPubKey:    "59ca36050fe35f7818f63a5d5ab4193445ea3d052c53e9a0969151b2a46bf367",
				OffChainPubKey:  "d2e8148b0db11cb9c81622318ffd92581494c6cd842b1a92ee9d23646bd642af",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "7ddcf8216a1af1774e4dee80e58ce27f9949d56d4a6ffa7cdeaadf50a77ad8ab",
			},
		}
		nodeConfigs[BsquareSolvBtcMBtcER] = nodeConfigsBsquareSolvBtcBtcER
	}

	{
		nodeConfigsBsquareuBtcBtcER := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0xA8455696F2DD85a6dce6b5AD683BF3602d73181B",
				SignAddress:     "0xb1eB5956dFC46231672beE8a3f256Cdb300fB2b1",
				ConfigPubKey:    "327fd067ff5b2fa5dbc13b17e8de10f93126988a8b4795a66b5d012d35d03361",
				OffChainPubKey:  "4185973160c05dd69966dbaf098b771e045c241918e8e721ad861b185a38a2cb",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "347c28b95b66982396b0bf8533e7e10b8deb6ef5547857589a88721b581628e5",
			}, {
				Id:              2,
				TransmitAddress: "0x3c892f62acfeDB80c1046F6215b31BF48A704636",
				SignAddress:     "0xFBa309B539c9c843c9B58EAd530039E532ccB7B2",
				ConfigPubKey:    "ebf653af78ede7ea11db97a05823985decd75f0e4cf704a7ec1139a0674d696c",
				OffChainPubKey:  "37361a06980b8000ec8e001c1d27da53b13d9d94a43f2e1c76751d387560eb28",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "3cf96966828976c175ba68b6f6f51149107245210c82ebd375290083e76e826e",
			}, {
				Id:              3,
				TransmitAddress: "0xD11D0A08445F4F2daf2412f2D665e1fAD20F4f84",
				SignAddress:     "0x87392fE9115EAFc483DAe45A7D6373E61C3Db1Aa",
				ConfigPubKey:    "0ccb9c325ca2758a4b322ce8e9487489a31d15846e4087d6213f51e664bb7334",
				OffChainPubKey:  "7b361654983271656d4e1bcac84f29c1d1408b521de67ffbedb1ad83b62eb2bb",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "935465d02a8d666436f4448d21db12712cf137d0078113ecaa0fbac1fdcd958f",
			}, {
				Id:              4,
				TransmitAddress: "0x18AE11b760B3e1EAad18946cb3D64e97a565e3Ec",
				SignAddress:     "0x984C55BAA7A88a5Ba1535dcFbfB2C3DC994586F9",
				ConfigPubKey:    "1fc42d91957939dfdd04e798789b37991d24234169c500286bae270b5cd0967b",
				OffChainPubKey:  "4b97c089dd9a38d03adf237c1328451ed6a81ca7485fe1c37981aad167457888",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "9ce620f14b92e807b58d5231f7bb6edaa7d3cb74bcbf9bc9db98fdbbbfd33fc9",
			}, {
				Id:              5,
				TransmitAddress: "0x4EA989bDBC43701bAeab917c16E8d4d7f193b590",
				SignAddress:     "0xB2FA37a9D591eC4D52Cf64C7B7deF09BF8aba4bE",
				ConfigPubKey:    "c3f23f9ffb03975751f33fd14ac84ab38780ca48c0bd82d6098476dbbc9b256a",
				OffChainPubKey:  "e5445219827d2589d62f99993cff953f296c0488b9628d6bc0c2bb58b0d5b26f",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "5be49c9fe9a7b38e5d432d821644f7130f8bfd62bb94bafebf7a630b62ecd376",
			}, {
				Id:              6,
				TransmitAddress: "0x43020E6C38C3a8ce5247AF292858Df5cF3e664c4",
				SignAddress:     "0x33B176c976A01f10D5f78Df434E6a165c4CF4036",
				ConfigPubKey:    "db8931799c01eaca6f466ae1017c41522dfaa99e74ee2e56f6a01ab6e8f29c16",
				OffChainPubKey:  "546196cf4c448557d301188e0e16be6e0fa17c016311054ebe3b8cfd0ed317d5",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "d62ac2385e0c70b41c4c2157f64148014904829b07cf420ae6979e62ec1c6644",
			}, {
				Id:              7,
				TransmitAddress: "0x7BDe199140F7bdBA1e259A26ab15775a1cb5e599",
				SignAddress:     "0xa4c05FE37194A7efcb5a1eDD46a7C9cA9a1fBaF1",
				ConfigPubKey:    "4dc11d483b53db6b25601bbd55703012372b2069ce90623352287aa34bfa555f",
				OffChainPubKey:  "8cca9a79ff843a61a50058dbe18677c97c28b9ff560fc77e446038760b6e5b81",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "a936f7e533766afb7f3554fdb1a9fae69a1f9f86e41d194e26f59dcdcf678e28",
			}, {
				Id:              8,
				TransmitAddress: "0xa52074d8C7614cf4CaE28724c0C18476802e3434",
				SignAddress:     "0x88Fd7F51aA97E8781073993e69Dd2FeA5b94BA1c",
				ConfigPubKey:    "924ca36010ece1540be12d7283455d3fe7d803922f9d2104f5ea8c4f3e00ee1e",
				OffChainPubKey:  "cd74be8ef32cffb4646f32356416f45a47f20b7e066f7e0e4f85027a805faea3",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "29049b91340bd341171ffa2d9b33dd96c2074f44712a2f91df13f905808ece25",
			}, {
				Id:              9,
				TransmitAddress: "0x97cd58ED3B696C21128F395fB96f37550D201207",
				SignAddress:     "0x46E16B309D16920750c9a8BCdBC92B4D27F7aC47",
				ConfigPubKey:    "3cc286bfa835aae9b0c42ef3e7c8a581e888514c86a6c9a8c9373bd850a8b711",
				OffChainPubKey:  "9acca6207bf1d0b088afdd2e76d6a297a63a9249e40a1751546eaacf2fe36345",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "b0ea3655cf739497d6fcedebf5ffc78a8467a045aced7e7b518f1c08cb1abe24",
			},
		}
		nodeConfigs[BsquareuBtcBtcER] = nodeConfigsBsquareuBtcBtcER
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
	ocrConfig := GetOffChainAggregatorConfig(BsquareBSTONE)
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
