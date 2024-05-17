package bsquare

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
	BsquareUsdt = iota
	BsquareUsdc
	BsquareBtc
	BsquareEth
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
	ocrConfig := GetOffChainAggregatorConfig(BsquareUsdc)
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
