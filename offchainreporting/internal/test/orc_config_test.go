package test

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ocrConfigHelper "github.com/smartcontractkit/libocr/offchainreporting/confighelper"
	"github.com/smartcontractkit/libocr/offchainreporting/internal/config"
	ocrTypes "github.com/smartcontractkit/libocr/offchainreporting/types"
	"testing"
	"time"
)

func AproOffChainAggregatorConfig(numberNodes int) OffChainAggregatorConfig {
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
	/*return OffChainAggregatorConfig{
		AlphaPPB:         10000000,
		DeltaC:           time.Minute * 60,
		DeltaGrace:       time.Second * 12,
		DeltaProgress:    time.Second * 65,
		DeltaStage:       time.Second * 60,
		DeltaResend:      time.Second * 17,
		DeltaRound:       time.Second * 60,
		RMax:             6,
		S:                s,
		N:                numberNodes,
		F:                1,
		OracleIdentities: []ocrConfigHelper.OracleIdentityExtra{},
	}*/
	return OffChainAggregatorConfig{
		AlphaPPB:         10000000, // 10 ^9
		DeltaC:           time.Minute * 10,
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
	Sepolia1inch = iota
	BitlayerTest1inch
	BitlayerTestBtc
	BitlayerTestAave
	BitlayerTestComp
	BitlayerTestEns
	BitlayerTestEth
	BitlayerTestLtc
	BitlayerTestMatic
	BitlayerTestMkr
	BitlayerTestSol
	BitlayerTestUsdc
	BitlayerTestUsdt

	BitlayerTestGre

	AilayerTestDog
	AilayerTestWzrd
)

func GetNodeConfigs(target int) []NodeOCRConfig {
	nodeConfigs := make(map[int][]NodeOCRConfig)

	// bitlayer-test
	{
		nodeConfigs1inch := []NodeOCRConfig{
			{
				TransmitAddress: "0x60A1b1932BbB7E6cCe543c0630f283f39ACfF21a",
				SignAddress:     "0xd618B820FAFC1F4B98d7BCa6F125186A5fc04F8C",
				ConfigPubKey:    "8f8602b508b0c554fde124a7ef5fa4d066887010a9e9651efc11a40556c12c3e",
				OffChainPubKey:  "9ec24f8462e8d8937b676af5656ef450eadf128dbeb50d9fb00654f7f0ec8b6a",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "a41849fe497bf8f8582b75f273a98eed370204214720e30ef4c9f59295b4ca82",
			}, {
				TransmitAddress: "0x14cBF542Aa01EFF4e8a869db97aE04ba75C5D9F4",
				SignAddress:     "0xCD5808Cc2E6Cf91Cd2B77F68E33fFe3e8F95736d",
				ConfigPubKey:    "46ead7614db523fa06060ab2a7f8cba91299fc66adb7d8c63ef5a5e11a868018",
				OffChainPubKey:  "39a902f81fabcfbcbca1d60147430ea6525500424c9113e5b7aa51149b791684",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "b7f99afb0241365baddf35af1bec17447d7d08cf16b2c53608e5f78693c9e34f",
			}, {
				TransmitAddress: "0x4e1fa23140017d34F9904e6A2a8109F9C0b672D9",
				SignAddress:     "0xEf9AA3f276C5246A5736dF40CaeB5ebc116747eA",
				ConfigPubKey:    "3a59c3bf117db4183ee3f980dfc8eca5ca51bb5b82c4ce9aaf8cc540ca12e464",
				OffChainPubKey:  "99c3e9b464318991c88b07bf109e8f9994cf5b905aa07980c64f8e48f5c55d17",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "1d0e022e45a24a3820a867d7c1e3e8f0f12c1738c2bed25f55749b4e6e355fd1",
			}, {
				TransmitAddress: "0x4B8cf000ccd6FefEFf586E7E50406E2845d83080",
				SignAddress:     "0x988722c55d231ba13Feb39F53a402F3a35EA079c",
				ConfigPubKey:    "3b1d8e27c64d21491647471d3f8b81b0d8673deba243b6d24e1f812906185b51",
				OffChainPubKey:  "01a37258fa071968ae944b5f06b3b7aae3123037029fec784fd0a22f299670ec",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "0eb1ed0a3167d4386b0ac76e9fa0450c37a8e8c7e58553ac77a3f559d430f0bc",
			}, {
				TransmitAddress: "0x2B979C416BF7D37920b61C4E266d2da72Bd0c772",
				SignAddress:     "0x79f1D85B77ac14D901C83C015A7f696107d6FC5f",
				ConfigPubKey:    "869d1342208dfcaed5c596fae64290b4a669c07c3bce45b5357d0264b3c64312",
				OffChainPubKey:  "3f930dc6e25223622d1e12e70c506b8caeff61a116c2f3d1ee7672d25eb6bbdf",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "d516715be4e4933e790c904a93b300e89408464f544b1dc647e5fe9f9dcda905",
			},
		}
		nodeConfigs[BitlayerTest1inch] = nodeConfigs1inch

		nodeConfigsBtc := []NodeOCRConfig{
			{
				TransmitAddress: "0x60A1b1932BbB7E6cCe543c0630f283f39ACfF21a",
				SignAddress:     "0x78Ec0b787d4baEc3b9970D57De3894a05ebb7f3a",
				ConfigPubKey:    "8528cafe173b7d0904f40e5822b383ec3415eae1ff1818027b3c9f974232cd00",
				OffChainPubKey:  "4b3b101552011af5a53914f59a8a79051f66f818a31d3a1cc392764c5df056e8",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "b810f903f18f3b823444af4913430c8f34d5fd508e5a7fae461a1d40b7054ceb",
			}, {
				TransmitAddress: "0x14cBF542Aa01EFF4e8a869db97aE04ba75C5D9F4",
				SignAddress:     "0x020C371e6eB32111e390370D16561d5cF54feb40",
				ConfigPubKey:    "885dacf4e88fa97ce7f9c9dc08bcafcc4a1806e4c5aedd8175b6c6fe57742358",
				OffChainPubKey:  "e976613c8dcd0c0e280ad00f5d59c71955acaa6a9b3e2cddf216da75c4153ce4",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "c02a4608e0eee6aa4aa57af482b53e1083cea54a49c9376ec65048029bf6471c",
			}, {
				TransmitAddress: "0x4e1fa23140017d34F9904e6A2a8109F9C0b672D9",
				SignAddress:     "0xD1308da16391e0ffAB05cCFfa6ad1d713Ded2527",
				ConfigPubKey:    "8c1235d463ae93adfaab17747d01f5c5ed0248025c30a3fa98434807dcd84965",
				OffChainPubKey:  "cec8fc656019f4ff1e26415af59a04fea8bab885df80adbfc5aa156b4642f7e0",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "671cffdb485cb4994bbde308c71826c522f2d067b63be136b081effcdc86320e",
			}, {
				TransmitAddress: "0x4B8cf000ccd6FefEFf586E7E50406E2845d83080",
				SignAddress:     "0x8D48f8947A5DEA81CdC8Cff06C497A6e1A1152F9",
				ConfigPubKey:    "d0ea1e7ebb2cd85fb00432097c9e97443e421ae1c6b54cfe4d1a1a085f79cc71",
				OffChainPubKey:  "cce6805f579cf71ded68cb7494bfbbf12a6bbdf4084b532a8133618c820678f0",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "c6f5358e24cafceb71b2c31d20e80e5ac14ac8e638806c6635c6275756de3f17",
			}, {
				TransmitAddress: "0x2B979C416BF7D37920b61C4E266d2da72Bd0c772",
				SignAddress:     "0x3b0823246b84DE11f4A55D8b5cb468fC2dF3b61f",
				ConfigPubKey:    "b49d355590d6be1c0aeb705895221a1f0fa14a75f69c6a4f9006ccb59d3fe724",
				OffChainPubKey:  "a0bd79b3daf0b86f9a0530b8109b9622a9eb39d5ab2f4e7fd0c8c4bdab712526",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "64422f0f3f9195ca04064dd5c4958affaf4d609d20afc452b815ba267cdfee26",
			}, {
				TransmitAddress: "0x03687ca64975dF0cAed5742dccBE0c394716DCb3",
				SignAddress:     "0x6F36e78e0938bEB49adB6Ed4dE73EcD12f97D1AE",
				ConfigPubKey:    "b32443f882f94f97f5192ad1aadfbdf0c0b6b8ab21fb82974851491300615800",
				OffChainPubKey:  "323934c09b0318fcf541c63b25da539735e85e0581f3a88940cdc555c1014e27",
				PeerID:          "12D3KooWMM4ufpHRmAuXBhwWeY6V12UMKjonkg1YR82PPLYMsYKS",
				OffChainKeyId:   "c6da35b809ce4c2cca65121e3ac6bfb651a70e814515e44b23eadfe63a3aa416",
			}, {
				TransmitAddress: "0x9C0E25212f38DC9A8ED2B063B21fBe8775D86951",
				SignAddress:     "0xe8aEcac536e14E085f42cA75E590B3945a4238dC",
				ConfigPubKey:    "2426d64a2e4b7bd986d4098e923d2bf74ec009e3856027fa0118929168adf814",
				OffChainPubKey:  "28cb12c95af6215ac4e3377874bf5674df45a14423fd761409ecd495df439daa",
				PeerID:          "12D3KooWPjc5CeHTBc4LL9Z3ekyfEbVHhP85MqwffJfmfP68y3h7",
				OffChainKeyId:   "6ca6170391cb7d0ee7fe55c71dacd2195f42cc03fce268377db3ec22d8b5e2f8",
			}, {
				TransmitAddress: "0xDB8d67c0c79183ed3C3077FC7d919A61338FeED9",
				SignAddress:     "0xe9bB991be5E201Cd55aC4BcC2e68DB4daD294Eac",
				ConfigPubKey:    "2d2644928d74f52f2fd1347465b5690a018fb18501ded1bde85bba3f35456447",
				OffChainPubKey:  "37ef7792d0dde9ae484fdbb542942cd4acd74c2a497f550849f02fc62f2d62de",
				PeerID:          "12D3KooWAp9d7Hn4uw4nEwpmvUVWCoRH3qbuLhgtayDLNm7FK1KH",
				OffChainKeyId:   "8a7a0a730946bf1bdab815ccd5d12e2a2640e0c32e309f4f76a6eaf64b252126",
			}, {
				TransmitAddress: "0x947CF3d11fEE3b1dE4EA717F2e192614947882b8",
				SignAddress:     "0xCE778BE42d4Dd7F088Cd7996a363BBe6b41bB684",
				ConfigPubKey:    "c5ae524ad5834f6deddaa0bf91bcd30b7bc17c3728d1c9e34e96bd661b5af14e",
				OffChainPubKey:  "c4e121625e5701669a13f600795b492571d0bc9707956c4ef8ad01931f743073",
				PeerID:          "12D3KooWMDF2YxveKKa7rYN4fM7Uz46basme4fAHR28uwEoeKu45",
				OffChainKeyId:   "47f123496ed2c8bb00fb11576b7396f5cff91ebddbb314d6efbc6c226f3cdfdb",
			},
		}
		nodeConfigs[BitlayerTestBtc] = nodeConfigsBtc

		nodeConfigsAave := []NodeOCRConfig{
			{
				TransmitAddress: "0x60A1b1932BbB7E6cCe543c0630f283f39ACfF21a",
				SignAddress:     "0xBb65A9e7f0e93815585eA6F12FCb4FAa121743b0",
				ConfigPubKey:    "14906e0b8421248ee0c4271025ad7dc1cb9a41472722ea4fe0cd70db1553071e",
				OffChainPubKey:  "cbe1f03de98399205556dc826d26a0b3420e491b6f23f614fec9ff119158f80a",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "e8df6698554922a5fabeeef8865d6563eb04b09924248c91d69d2cff39799d89",
			}, {
				TransmitAddress: "0x14cBF542Aa01EFF4e8a869db97aE04ba75C5D9F4",
				SignAddress:     "0xD7Dd8287e30375d091fC124c5018Cb80d38EaBDA",
				ConfigPubKey:    "6a7a986b008e33984eedaaa1150b51a7fd6eda1a523baf71b6970b5943763a61",
				OffChainPubKey:  "281423c44af78c3cdc0c75d3872b080bcb088385b312f92fde046593c0e250f9",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "c24a594ccc174cc7946d6cd09f628743b9c16d9327eedcb7348e6d973c29d983",
			}, {
				TransmitAddress: "0x4e1fa23140017d34F9904e6A2a8109F9C0b672D9",
				SignAddress:     "0x2731d218ee62FBdbC9d4902C1FE66dd4B839F142",
				ConfigPubKey:    "115fab014e81bf10935f34df7c4c212fbf4566cc3f7a746e6cdf612bf2aab912",
				OffChainPubKey:  "83eccdf294d2fbf7898a1d4ef6de635f3953c8613974542cd6b3c7455d295578",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "83063f80b735d0001e8b09ecb03295c147415aabac8073969bac8a5c9b2a4021",
			}, {
				TransmitAddress: "0x4B8cf000ccd6FefEFf586E7E50406E2845d83080",
				SignAddress:     "0x1b055516c5e32D884DC0CC952d63401ae1827FA5",
				ConfigPubKey:    "dfd35883dc38c087102084e5fa83507b23a8e48a3474885a493c6b67a70c9e5b",
				OffChainPubKey:  "e80c2e725600acee8d61f223c3877335c050a10f9af8231683ca49c708382948",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "36cdb9e5bf9ee81951de5ff39cd069d07fdced2677e237bd895bbf1753bde151",
			}, {
				TransmitAddress: "0x2B979C416BF7D37920b61C4E266d2da72Bd0c772",
				SignAddress:     "0xDdad12E7b9eD998C2052E3991EA18fc3b292d7e8",
				ConfigPubKey:    "0a3394594c77c9f240ce077534a50cbdbe8a15c2080a42d73867c483bcfa2d11",
				OffChainPubKey:  "ffe9fe33f8dd568cb3fd1df105d94e7698f2dab9e8aa4d0635ba020d560a64f6",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "12d6236ef5d0992da4cf99735fb7fc43ac92ccb1e112cbccd3d4a81ede90ba34",
			},
		}
		nodeConfigs[BitlayerTestAave] = nodeConfigsAave

		nodeConfigsComp := []NodeOCRConfig{
			{
				TransmitAddress: "0x60A1b1932BbB7E6cCe543c0630f283f39ACfF21a",
				SignAddress:     "0xf9bBfb9Ac86e985B7b2D6F7574056ad119381c6a",
				ConfigPubKey:    "9408eb36174f8f0ed989634c1368d404ada43c464918e3ef79540870d125040f",
				OffChainPubKey:  "291488c677b91b3bb2577af217730b6ca24b455142a975da19a78df55e39d3b4",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "f923d9d5f926aca38f7aa3bdfcdbead047cc938b769a9846e2c06c1ec76a9c1c",
			}, {
				TransmitAddress: "0x14cBF542Aa01EFF4e8a869db97aE04ba75C5D9F4",
				SignAddress:     "0x9B4c5576726F156ed83FF8c86313d65F25980F4e",
				ConfigPubKey:    "0e92d96d53ad03284659fd78f69da24d72e219d7238cac3970bea362791bd80c",
				OffChainPubKey:  "3c5450ba467e76c40ee3bf2e36bb715574d46d6f5d4f15ed0d313f559b67bf0d",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "9a8ff108bbd71687059f1504cac621abd594db9d1070a58f91a139d8f6745a98",
			}, {
				TransmitAddress: "0x4e1fa23140017d34F9904e6A2a8109F9C0b672D9",
				SignAddress:     "0xb5a5246B1F36Efab3DEF22639B64F1Ca37aC8D91",
				ConfigPubKey:    "08f6b50b1b539f2f8107828225e2d27e3f97225daba240f88f7c1c776bc43266",
				OffChainPubKey:  "675366dcdb6a40ad2a6ca3a7a320d4fce40b0512b693b3a86ea608b48a4922dc",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "a03395252cb040772e5787868ac2b502155e335e5127f427ed2e5d0353d3f10d",
			}, {
				TransmitAddress: "0x4B8cf000ccd6FefEFf586E7E50406E2845d83080",
				SignAddress:     "0xFAf78B7c40B5a5124d5aABc4A993949b8e182cE5",
				ConfigPubKey:    "156c8319f5620dc2701b473035d05f3c26fb3764d051e3c427aa5bffe5aaa32f",
				OffChainPubKey:  "74c18de4643d5d85084c72792248e965fe5dc7b2096cbab02c4334d79aa63283",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "28387f2a9bb62c3a3c838f5133ff2b68a33ea4c15c9211da5c8c4e15e41a9c2d",
			}, {
				TransmitAddress: "0x2B979C416BF7D37920b61C4E266d2da72Bd0c772",
				SignAddress:     "0x5F954859aD280Dce018f471b2B5a877e30DA9d5e",
				ConfigPubKey:    "4ffeb54f96b954c2bd0ba78136dd6041a1b5ed438fe15eaeb036ce3d848cda0f",
				OffChainPubKey:  "1c3e5393487cd56adce13a9bf8ced98d74757d291f51d5de3ba402db98d5a654",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "d5f7a4769a5e4748e1a446d822ca0bb03388fb09f9acc9a0aae8f7ea5f507378",
			},
		}
		nodeConfigs[BitlayerTestComp] = nodeConfigsComp

		nodeConfigsEns := []NodeOCRConfig{
			{
				TransmitAddress: "0x60A1b1932BbB7E6cCe543c0630f283f39ACfF21a",
				SignAddress:     "0x1762688F2245b8B9Eb530959F38a7CdD83245669",
				ConfigPubKey:    "af6abee6c0ee7390fd0dec30921028bc436f339bb933b4a922523b1332397619",
				OffChainPubKey:  "8efb1eee43d0e11aec6d61c0758039d354e472ff3836ce2ac94f651aeea2dcc9",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "e4d07ee0820b0b44a775513f24083b6b3fea3236aa980e2c4ac7adce9cf459c0",
			}, {
				TransmitAddress: "0x14cBF542Aa01EFF4e8a869db97aE04ba75C5D9F4",
				SignAddress:     "0x0CD37d426684EfCDe263ba9f6E3b31Bf7216432F",
				ConfigPubKey:    "179126f7d9e497e62f2b88db09adabac4f077944e2058851b0f20a900148da4c",
				OffChainPubKey:  "957edbb77eca7a7c3d6c138f55c631cbae5ff484f411fb33f85fed587df91b60",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "0a19a94b9e098c78f5addf738aa0a73a083d3dbb9bf89648a7bad15690bcf581",
			}, {
				TransmitAddress: "0x4e1fa23140017d34F9904e6A2a8109F9C0b672D9",
				SignAddress:     "0x2F76D1C94e7569f8aF5f659c676b81bd0c4bF34a",
				ConfigPubKey:    "13c492b301d89a70ffc0df3ef4c883e1427d4b0701ee32b3a347db78e0970946",
				OffChainPubKey:  "64270f473b5cfdba203d2ec79e0f84b848024272c7e43d4ce4e245880b013520",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "d0d2a5d4efe1fc6ed200c4114b56ed8fb292567232ec34295d94be1183e3fa9e",
			}, {
				TransmitAddress: "0x4B8cf000ccd6FefEFf586E7E50406E2845d83080",
				SignAddress:     "0x3d4237e4B10B7E18299eb8AA757352387e4397a9",
				ConfigPubKey:    "3b31d7c1e154987761585ec0c6944f0fac251c924c4e4d7eca4ed3ec2e4ffc66",
				OffChainPubKey:  "48f6088388342c037e6b917cc7eed8a5a840c17cfe257688e64f9867bdf6f0d7",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "ea95a7955defe6f9293926fe6f980bdea482a928d1018d49d453c1c4b7445f5c",
			}, {
				TransmitAddress: "0x2B979C416BF7D37920b61C4E266d2da72Bd0c772",
				SignAddress:     "0x2EC4296c2c7Dff2205c1E67cCE04Ade6264CB012",
				ConfigPubKey:    "218873e79abcd211373e875056681dac8f9aa701d7c7603baa46c3f9eaffb814",
				OffChainPubKey:  "d12a3de218c2492856bb8d74eb79a7eaf5ed0b26a2018b85569a465605e18580",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "bcd4b44de562803d5ddac2a731c21d5429e8548f892f0f1af114d366308d22bf",
			},
		}
		nodeConfigs[BitlayerTestEns] = nodeConfigsEns

		nodeConfigsEth := []NodeOCRConfig{
			{
				TransmitAddress: "0x60A1b1932BbB7E6cCe543c0630f283f39ACfF21a",
				SignAddress:     "0x9f19A278FdE98ea543F8138e56c22cebcF882f86",
				ConfigPubKey:    "8160f9bfcfecdcdcf6206e3b63178fd56cbd056668680d15ba6d18bf1df7ed2f",
				OffChainPubKey:  "50c9fb4cd9ce0f5bb1d79e8397956add39378175bc50bde9527433e3e233505c",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "aea5ba1228c30719edb1c3cec70cdd407a76840aaa1fcd352e3cbc9336a82ad1",
			}, {
				TransmitAddress: "0x14cBF542Aa01EFF4e8a869db97aE04ba75C5D9F4",
				SignAddress:     "0x10201086658a11fb6E0364a8a8C603F9206797D0",
				ConfigPubKey:    "364415534fccd3fd1e10d4ffb837abc43ae8c6789084aa7a45469af444b06607",
				OffChainPubKey:  "5256f6e7c6d06027e78ac98de961248218504521c41bffc3e67ccce50e00f1fe",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "5cdd470b532b6186ebe538179a26cc8ab9762d31417115261dce406f472590a2",
			}, {
				TransmitAddress: "0x4e1fa23140017d34F9904e6A2a8109F9C0b672D9",
				SignAddress:     "0xe83e5D475630bD1De017F7A124900039D3B0550E",
				ConfigPubKey:    "e7a66556b46f6b8dcc6a481bcdedffb1924ac3a984a517f80cf7308b0a815124",
				OffChainPubKey:  "0bbdae995584fc0b4b7f0e0ee1f7c9f30ec15444a1ce394c8461fcdd2a5d48c1",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "c23013d72281b63dd35c9abbbca80f39ea1ea8f0fadd719d176cd5e25e21c3c9",
			}, {
				TransmitAddress: "0x4B8cf000ccd6FefEFf586E7E50406E2845d83080",
				SignAddress:     "0x1F8b9dAbC837242cf3b1C30a5508a2073CdC8bF2",
				ConfigPubKey:    "359e142d12bf0275a2e902cce22523dfcd47400be3df0a14cae51819c576c137",
				OffChainPubKey:  "2afa483a35d2a4fc618660e93fbd5d2498603f37bfec66a3cdb6eeac1e575e18",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "ebcff2cecce37bfcd4d26d51ec88bb0f3fcf4337cb9f2c25d4bac37ff4ac68d4",
			}, {
				TransmitAddress: "0x2B979C416BF7D37920b61C4E266d2da72Bd0c772",
				SignAddress:     "0x53dB6D941B682209F137636676887c68B6355776",
				ConfigPubKey:    "fa808efb08688a7819686d747977558eec21a3ccf03feecae181288e86345a53",
				OffChainPubKey:  "14a4945dc26f59d20a8aef49e6b097b373a6ae420bfdac1194c2587e171abdf9",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "8e6ccc14ef0079ca1c39b9bf82ad0326e89a6bb9888c289d767562b841c61480",
			},
		}
		nodeConfigs[BitlayerTestEth] = nodeConfigsEth

		nodeConfigsLtc := []NodeOCRConfig{
			{
				TransmitAddress: "0x60A1b1932BbB7E6cCe543c0630f283f39ACfF21a",
				SignAddress:     "0x40E2c0d45e65461872724134e0916b4b3f493FE9",
				ConfigPubKey:    "c72f93fa42ad2caeff4184d345567e7200393c45d687a7a132c6530c893e2312",
				OffChainPubKey:  "fd6403ea4b63fad90ebf27c71641cb6a5f8065207f86251db707a3ef8a46515d",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "ab9dad4e92d964ede30ea348f06d357054d393cad5adb5e7d391d2aa432c4d52",
			}, {
				TransmitAddress: "0x14cBF542Aa01EFF4e8a869db97aE04ba75C5D9F4",
				SignAddress:     "0xbD84A865F58cc86c06229777442E6744171E9376",
				ConfigPubKey:    "741055ab62816249131458934f458a3ffb8e2e9331cef112f3a64e20b5a05359",
				OffChainPubKey:  "42674b9a631ad7c3d6aec7e86c275c44202ae0d477f1ea7c1a64c119f2bd2eee",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "d0324b4d4efe4b9adcba270e26a75c805fdfc2986d268aa08a73c711500e5c64",
			}, {
				TransmitAddress: "0x4e1fa23140017d34F9904e6A2a8109F9C0b672D9",
				SignAddress:     "0x611e795e9d60AE07CDF3fB8aCf8113BE9030243D",
				ConfigPubKey:    "5d10f8978a6d29a08df7d117257a0f089c66bca081463fa47dd0e27107e3c373",
				OffChainPubKey:  "f7d48957566af5ca0b7d4028e2c6485dcd0e2f0c6d7d95a6c6b2d30f0a1e5fa3",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "2b87101880ca708b984d86f38a98e391141c18608c4d9bdd90b352b92ea7ddf4",
			}, {
				TransmitAddress: "0x4B8cf000ccd6FefEFf586E7E50406E2845d83080",
				SignAddress:     "0xc0E18582396d832Ff9E5467e9c9245993c1A76e9",
				ConfigPubKey:    "c0aa75321fbd7aaaa110599b6d12d0955364a5cd63ad2907d2faa3b068b1e071",
				OffChainPubKey:  "d6682d0e085daacdffc387070c6415e061257561d5a967cc32872c26d38f9d2b",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "14928c9750522e615e346c47f69696d78e0cd7a8ee1c8d0840ea77f5819a7fee",
			}, {
				TransmitAddress: "0x2B979C416BF7D37920b61C4E266d2da72Bd0c772",
				SignAddress:     "0xe5B8917146D4F3F1D6d22CaAe855676Ea59e2784",
				ConfigPubKey:    "f5efbcbef13f6a65f2eee0bd8524c7313d011e57e5deb0bcfbf13eb63d990a72",
				OffChainPubKey:  "5cbe347f99cc821f938c6c7b581ada89dc20280bbb869cf882eae9514bd271cc",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "91dcb0022afde8f12e98efe8c1cd06e403e06d9955527a9ef7e6f42b3d30ee2c",
			},
		}
		nodeConfigs[BitlayerTestLtc] = nodeConfigsLtc

		nodeConfigsMatic := []NodeOCRConfig{
			{
				TransmitAddress: "0x60A1b1932BbB7E6cCe543c0630f283f39ACfF21a",
				SignAddress:     "0x29c395d783d69aCa7798584dCC01B4b4b78fa643",
				ConfigPubKey:    "55767d0108ca27ace8bd0a261f9340d9396bff15a1bef406b48d4852648d745b",
				OffChainPubKey:  "f54abf46e07a024bb6d99ee3d61c5c829b0c7b4b2f5f3dfc18e9da9fadfe1b69",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "72f76d502f88117f669498beb2e975330bf0a0be104670155266e26f8f63e838",
			}, {
				TransmitAddress: "0x14cBF542Aa01EFF4e8a869db97aE04ba75C5D9F4",
				SignAddress:     "0x933fFE628da3813baa470e9590e12c79E2E38C9C",
				ConfigPubKey:    "f381446b31dbccb6ea931868482551350a133e131bae86b913d0c92d674dfe4b",
				OffChainPubKey:  "1e685d5cc9b5873627b3c1d54b1b5d5f442f0e21b7903f64f5dac3615ce7255d",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "236c0920f13667dc29cfb413b322b965e974b7e9c83e5756949f9204178df582",
			}, {
				TransmitAddress: "0x4e1fa23140017d34F9904e6A2a8109F9C0b672D9",
				SignAddress:     "0x8842f5b4c63286FA9fA77C3A301b79a6d5058917",
				ConfigPubKey:    "7f579aadd9d1edc0cb0f73fdda558686ea6765fa7932983c5ec83716466ba65e",
				OffChainPubKey:  "54dfd4406dee70bc3abbb151ea5ca5bf2f811083f6d5e9f5cf786d23a807c675",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "674f0504c472714fc8968b02f2a8bed03cdfb04f28c5bba5a08f56ef7f380411",
			}, {
				TransmitAddress: "0x4B8cf000ccd6FefEFf586E7E50406E2845d83080",
				SignAddress:     "0xa1A31e259BE433d97e07AeFF832dd1e053523438",
				ConfigPubKey:    "431b5528d07d2eb37ebd1c13238338b2bbac28ebf57d89f1f96f527b49829953",
				OffChainPubKey:  "49f830cafd47f580401e6a58fe8fc58a18d24a423964ddbd1b2436e67891dc99",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "a45034423d87e83d04570da20cc4de34c033fb2e5395f47971219be5f37b8874",
			}, {
				TransmitAddress: "0x2B979C416BF7D37920b61C4E266d2da72Bd0c772",
				SignAddress:     "0x50f9d644bbA654370f158Fc041dfb49f6af6ecA4",
				ConfigPubKey:    "eab44cc9ab70b9b490d27e9522d0e02f45d447bbfbc0c51cf9f99f9bee74b962",
				OffChainPubKey:  "da276776ce1ea2085887fd76b3d761635504b9ce9b1fe88eb420595560e41b38",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "fc7e42b487df46bfe22d9d9a576286cf8803519bf79c65972d2c47455129ca9d",
			},
		}
		nodeConfigs[BitlayerTestMatic] = nodeConfigsMatic

		nodeConfigsMkr := []NodeOCRConfig{
			{
				TransmitAddress: "0x60A1b1932BbB7E6cCe543c0630f283f39ACfF21a",
				SignAddress:     "0xF16AA98A0cF1dDe65df3Cd4fa54aC4DF8Bb44c66",
				ConfigPubKey:    "9b9072a10420b1822b3ec63d1bb39607b91313a13d70c2875604c6ef912fbe70",
				OffChainPubKey:  "5c9686f5e54fd932b23499be228394c4ccc581af7a63c75633d78ecd3899f90b",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "c01fd0a6f238bcbf4bbf669f21ef7f27265b4983f6a0d6a46fd49d56b24d5404",
			}, {
				TransmitAddress: "0x14cBF542Aa01EFF4e8a869db97aE04ba75C5D9F4",
				SignAddress:     "0x6FA551fc850141EA80Cd9a2AeF90afDA1C5eFaf3",
				ConfigPubKey:    "670f8601fe8238ef15fca9f7f8960d2251ec944f513016b6ae7e3f0af09f8b2c",
				OffChainPubKey:  "d704e65f7e5c318c66afcd74f1710d6fd20b3883ae9c67e0ba082ef17e7aad60",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "996640724999b8c36154bedcc133b116ca4796a58391366868e92357b1939531",
			}, {
				TransmitAddress: "0x4e1fa23140017d34F9904e6A2a8109F9C0b672D9",
				SignAddress:     "0x8cEd1e465D3afAB2C425Eff27C4B2F5eB78178E9",
				ConfigPubKey:    "bbf4f29197daa7d5e79d305476b98f4fde36b7573305315f5806b84f271bee45",
				OffChainPubKey:  "52c2f4defbaf0ca2883bed3b83198cb9967052970b11f89b535b9076df0bd38a",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "beac70f1a61a4120beed768d7c0c3e30b6c0d266fb48f27e42ef362ce4880fb8",
			}, {
				TransmitAddress: "0x4B8cf000ccd6FefEFf586E7E50406E2845d83080",
				SignAddress:     "0xc8F188Bb5CBc67dbb1654c80E0baAD000A090a96",
				ConfigPubKey:    "b1587bb251117751e0d8c2f85f26a0a59c28db37301e6b6c12f139d9eb24020e",
				OffChainPubKey:  "fecd1d5c32a17d5d413ce9be214a02ed9d15ecc0b8781156b314dbf687decbaa",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "522502d05c37c96babce7e5930cdc9e83cf3f3731d576a24b3dc931bfd8608af",
			}, {
				TransmitAddress: "0x2B979C416BF7D37920b61C4E266d2da72Bd0c772",
				SignAddress:     "0x508509e4B11C5D82441b69F9bD0D9fa08EF964a7",
				ConfigPubKey:    "ffe77dd0e469296e8770c4d699ea98a2e9db23eeb3412b656e2d3004e1abe55b",
				OffChainPubKey:  "5a82a751f1d9dc1959fd04493e2f38826620e17f19c0f77a7e35cf8cb24f4cfa",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "750febeb8efede5ccda0f72cf20658e81a45acfcade81e4be3c9db109168b3ab",
			},
		}
		nodeConfigs[BitlayerTestMkr] = nodeConfigsMkr

		nodeConfigsSol := []NodeOCRConfig{
			{
				TransmitAddress: "0x60A1b1932BbB7E6cCe543c0630f283f39ACfF21a",
				SignAddress:     "0x03bB20b511e64aF53C30E0F12a446763040A8088",
				ConfigPubKey:    "761f0e51b246d6157364927897ccf878226c2b9e4371e43d6909971b344a1c48",
				OffChainPubKey:  "b6fd71e8194e0d743fbad3dbb993ddcd85c13c4dc7ce61162f1709988156d34b",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "f457a54c1cc48edb0c48160e23355a447293c6f9166113288ec4182fc0b1f4ba",
			}, {
				TransmitAddress: "0x14cBF542Aa01EFF4e8a869db97aE04ba75C5D9F4",
				SignAddress:     "0x8FAF6e5D31189607B02AD44E46891f7FeeBAF896",
				ConfigPubKey:    "0e8913d64de57d144b1e3798626c7ade683ffd12709a63ba3366203df5763577",
				OffChainPubKey:  "d83fee20dba1d369768a3587e98ce8d7250bd22373c4e624ddd1afc1ecd314f2",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "d3483082e289b47de7a8fc6d60508e0ae61cd3e6ff15cdf1eabb4cd5e101f1db",
			}, {
				TransmitAddress: "0x4e1fa23140017d34F9904e6A2a8109F9C0b672D9",
				SignAddress:     "0x9dbA30655F5a93A5E78389d01919A723bc78f13D",
				ConfigPubKey:    "18e93c239740de5757aab92e5df9c42b37cd2b282d6f281685f5aeadc97eb431",
				OffChainPubKey:  "31c36eaf884ad02773b26779151be20b2e71a18ba8f91cc5d4af0a7973928b34",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "c50306d7df689dc9e5cca3a07f51754531f710b51dcb9d17ff541ee1171cc1cc",
			}, {
				TransmitAddress: "0x4B8cf000ccd6FefEFf586E7E50406E2845d83080",
				SignAddress:     "0x0c7a86212E31C658A7869D9eD4699602F436e059",
				ConfigPubKey:    "e8750cb1fe13b11a614fd771efd8136c3bd798136d4f7b5497e971e47d6a872d",
				OffChainPubKey:  "769f25fd09292587d92627745c630a33d0999467cf426c1da77d3f73eda82914",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "a3c57f74b32157485621db7d5c3567eceedea5b7cf81681e0a35c88db93f3910",
			}, {
				TransmitAddress: "0x2B979C416BF7D37920b61C4E266d2da72Bd0c772",
				SignAddress:     "0x296B5b1782370882f6c34E618F9e365AA62195Fd",
				ConfigPubKey:    "7a6f7b06d8a926ff2277c66e27e5aeba3ced437a8e391093a9af7af3b53a2f29",
				OffChainPubKey:  "0915541f64cda1684f8d2b92bd176bf386af4ef64954a55f840083a966afdea0",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "532b2d0d11264868c314e20a9bc66f03895376a707a445c9e28a56df675a17f9",
			},
		}
		nodeConfigs[BitlayerTestSol] = nodeConfigsSol

		nodeConfigsUsdc := []NodeOCRConfig{
			{
				TransmitAddress: "0x60A1b1932BbB7E6cCe543c0630f283f39ACfF21a",
				SignAddress:     "0xD54505dFD52CFFE56c1B408d4b3301B2D8C7E968",
				ConfigPubKey:    "1b24e956b5905f2e62537eb1d3d40437af2245ad5fa25bd95fd698805148606d",
				OffChainPubKey:  "86fca114e95461769d768c4c12dac8b1b0e1f22d912c1d95c8c5beca9f40ba54",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "7e3f35b8b26d021ecb3ccf015cf5b98a1abed39cae287064e5cb6fe475474903",
			}, {
				TransmitAddress: "0x14cBF542Aa01EFF4e8a869db97aE04ba75C5D9F4",
				SignAddress:     "0x45AFE0f8DeD3254b478a03513E53fBAda72B799C",
				ConfigPubKey:    "0c561ea7e00a63660d2dc58cdcbb55b609c0f65d06cf28828bfe206286f85f0b",
				OffChainPubKey:  "975714aa9cd6a9bd46f3db10ce0346a89e5489665b4365ac6c4ed28317f2a23e",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "76f3cf7dedec03cfae58e0922031bd31850c4e3d466b36a35860106d17cda103",
			}, {
				TransmitAddress: "0x4e1fa23140017d34F9904e6A2a8109F9C0b672D9",
				SignAddress:     "0x4ac2829B6971E573615F17221d462745eF52572E",
				ConfigPubKey:    "89727b60a7468e27a67a0e682f3d9c4cc23274e39d0ab17df43837229265b608",
				OffChainPubKey:  "6f8a832328ad921a6d90c9bbd778612b599c7c00a5047ef4351734f7e2abc70f",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "28df9207dfa62c95d670fc0b2180d6191b70edc4b5966094f66b55a5b076abd3",
			}, {
				TransmitAddress: "0x4B8cf000ccd6FefEFf586E7E50406E2845d83080",
				SignAddress:     "0x52B173F77a90e78E86d2Bd3f45CA8CE79E4da722",
				ConfigPubKey:    "d6cb05ebc21368dd00507084396507954addb7e1337d43b9b666b27cf4173b75",
				OffChainPubKey:  "2eb7587b8e42f22944265abdeeecd112c65023b37aa7793864bd2a2bbf655669",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "83895f56adc3b75c9ac9a2fee58d5cf1df15affa27cb32351dfb854a6fd6cf89",
			}, {
				TransmitAddress: "0x2B979C416BF7D37920b61C4E266d2da72Bd0c772",
				SignAddress:     "0xF71D0D12b64Cf383c3e1D1A0941e0FCCc06F5047",
				ConfigPubKey:    "1f36000d8bb1a2201b4469bc315076ada14b5efb094cb0722493603800616142",
				OffChainPubKey:  "e0f0b6e3fe760b70e70403c86fdf14801b3a3869cf3d9cb5ffdeac9153677a03",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "b88b4e86c9a1646395354c0d6217b6698c4eba9e5dbb39aab06b4c054f3ed446",
			},
		}
		nodeConfigs[BitlayerTestUsdc] = nodeConfigsUsdc

		nodeConfigsUsdt := []NodeOCRConfig{
			{
				TransmitAddress: "0x60A1b1932BbB7E6cCe543c0630f283f39ACfF21a",
				SignAddress:     "0xE87d452a97d5Ec56311792C66272B9C2AEaFC481",
				ConfigPubKey:    "79ad877e103f36dda00ebfa2d35eb07e991dc34955235ed36835dee25cfca004",
				OffChainPubKey:  "f31df1a9bd1909419c39047a8ed76ca5b47d99fbffd45180bbfadf830204f1f5",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "cd8dbb70d1ea5199a483f7a9e6bd3adb9c57005dfc0cd3cfc6970c8758b130d9",
			}, {
				TransmitAddress: "0x14cBF542Aa01EFF4e8a869db97aE04ba75C5D9F4",
				SignAddress:     "0x86e38BAe9977caCE5e935e35bdD9F237674A8317",
				ConfigPubKey:    "cc06a777133e760ad88df63ba77cf118d339a229ce7d1c08281f0a9907b90f5a",
				OffChainPubKey:  "de4e8b5f61434880a2fb08ccf64e932283e9c779dbccdbb471d6053b861ea879",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "cca371972848d136249b4c5cf50e6a1b860b381010b9ab5bacb3b9fd1c5a95ae",
			}, {
				TransmitAddress: "0x4e1fa23140017d34F9904e6A2a8109F9C0b672D9",
				SignAddress:     "0x63e037Ee659Cf57Fd0fC4C5fAB06240B1919feE6",
				ConfigPubKey:    "6de1088278dfeab0701fbab5212f86ef48adae1bcbf284779ddee8a37c454d30",
				OffChainPubKey:  "5c03cd5e490a1e14cfb8db16b14f7ade5e61e5b2122c17e100b9c1dcb0370994",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "cb5651d4cab5395f8e628b828c7baa3637aa1ff7005448db4919193efa4d4d79",
			}, {
				TransmitAddress: "0x4B8cf000ccd6FefEFf586E7E50406E2845d83080",
				SignAddress:     "0xF9315c8EE78E1dfe2cCBCe135441995957d3D227",
				ConfigPubKey:    "b4730443f88c74e0b8370389d8050fd4c0920b010778d41fe171c5be57494754",
				OffChainPubKey:  "f273e42be532c0b2d7754fbc51509292e144751056fb13c2e5916a1675e24ce7",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "60d074af1b022b114f3985a43351cd6235f9d619564bb7e51bede43bab6b764f",
			}, {
				TransmitAddress: "0x2B979C416BF7D37920b61C4E266d2da72Bd0c772",
				SignAddress:     "0x762460b2fFbF5344Df2eBB55DcB9B223aa7DFD51",
				ConfigPubKey:    "09fabf5c3736079badbd62107e5558dc49bca92f896f3700530c56d46797a12e",
				OffChainPubKey:  "3a44cc54c9c9653ae91da67356000f2b8284bc53696ce98bf2c20c7e98874388",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "8e43614b10ec91b53ec8cfd02cb90518215fa262ad04cf2c929c7c1958ff6e9d",
			},
		}
		nodeConfigs[BitlayerTestUsdt] = nodeConfigsUsdt

		nodeConfigsGre := []NodeOCRConfig{
			{
				TransmitAddress: "0x60A1b1932BbB7E6cCe543c0630f283f39ACfF21a",
				SignAddress:     "0xa3b9c2B6732101d8CE4dCFBD943A99b750372a6c",
				ConfigPubKey:    "df71f148023449eefec06ea003873cdcabdec81201a261258bc4b00ad253933a",
				OffChainPubKey:  "d84c9f1592642e9d239c757e27b11dac63ec9526181151084691e10f8c8398b0",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "63f8ac5cec0a660c0df6edd7fc4d1760367cdd269a82a7740a7aa29a38024b73",
			}, {
				TransmitAddress: "0x14cBF542Aa01EFF4e8a869db97aE04ba75C5D9F4",
				SignAddress:     "0xF62647A2c7E4B593F0192B26e42711D324Ce6e5C",
				ConfigPubKey:    "92bd3d7dfe2925a530b365c420440dc4e4032757813c6b7339a3dcb8464b500f",
				OffChainPubKey:  "a3169e227c71efd3a3efdbe7e0421fc4353b2c7038223a36c7118d3d14d8fd38",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "e3a884fc72b042a3fc942fedcfde201e03cf918bdf3ec865d96ae65f0d42bde4",
			}, {
				TransmitAddress: "0x4e1fa23140017d34F9904e6A2a8109F9C0b672D9",
				SignAddress:     "0x46E6Cb729966C63Ef48a2ab5CDF82de53396A726",
				ConfigPubKey:    "f54076311150ef491e96b282348c502647b5c66dc2a39072c84fea7692936807",
				OffChainPubKey:  "bee2959e51f79398f561a8bba6646a64821c427994af219f7fad4cfff201835b",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "34bb191e084a84e794a5e51b0938b09c7ffc87aa08d992f00c0c7311fa031917",
			}, {
				TransmitAddress: "0x4B8cf000ccd6FefEFf586E7E50406E2845d83080",
				SignAddress:     "0x3C6ce1608D60B1939A6c14b5e6120aD668D06013",
				ConfigPubKey:    "9657915be31b07cf085b45d4f431c637a4842f7ed10c728877db415fc1c04c5d",
				OffChainPubKey:  "a518d003bfb9e18d3f8d330f67f04fbd7d3edc4a0d92850cd8379d33d30981c4",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "a856b28b3cc0a35522e367a48dd205f16cbe8a4d532604cdddbd41303192a448",
			}, {
				TransmitAddress: "0x2B979C416BF7D37920b61C4E266d2da72Bd0c772",
				SignAddress:     "0x8aea479Cb63485FE5efA4AE2896bA17FA4c51a9c",
				ConfigPubKey:    "e0ae8fa33c8c4ba83c1a2a2cbb106aafce1e02542aa41030614ff403cc563339",
				OffChainPubKey:  "1841d44c6543c6e4fef394f6a6111066413bc8b46a2f1f3967ed034306ba5533",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "a45b72dc16b85b70703793684011adb2880dc4f4147bd098129d5cb84d8a64bb",
			}, {
				TransmitAddress: "0x03687ca64975dF0cAed5742dccBE0c394716DCb3",
				SignAddress:     "0xBB1105F45bAb872f22C4fFDAfa80c54105159dff",
				ConfigPubKey:    "b42c4af7f18c8ae3f9982843ca1d62e5f229826ad59f8ba49ce843956292d521",
				OffChainPubKey:  "d094aee9ef3f0ec77314b11b0b317ef3adf850398c7a221d3755b0d6a33ff5b4",
				PeerID:          "12D3KooWMM4ufpHRmAuXBhwWeY6V12UMKjonkg1YR82PPLYMsYKS",
				OffChainKeyId:   "148ffe08df00e6bda67a712a88c0841d61688356b1c5fd81a2ef77c2e9501b6e",
			}, {
				TransmitAddress: "0x9C0E25212f38DC9A8ED2B063B21fBe8775D86951",
				SignAddress:     "0x57a3F4069cf20e76275911cc51A38D42E50c5D54",
				ConfigPubKey:    "c56600663003ba233c2e2cbdde2a80909d104824fd0bb2b890f0c68318b07a31",
				OffChainPubKey:  "d4a949beea5b6c31171b1964d6afb3fc47b74eef6e4e2ac170d60fcd7ceb7ebb",
				PeerID:          "12D3KooWPjc5CeHTBc4LL9Z3ekyfEbVHhP85MqwffJfmfP68y3h7",
				OffChainKeyId:   "4e26e547362ae82af8cc2f81df05913341273488cccf294fc2056d31b9db84fe",
			}, {
				TransmitAddress: "0xDB8d67c0c79183ed3C3077FC7d919A61338FeED9",
				SignAddress:     "0x08C2ebCE344396aa7d0439947B727374D39c83E8",
				ConfigPubKey:    "299c07814114af463626088ea97c7156690743a71f46b18f0d170776e775617e",
				OffChainPubKey:  "87727d03e73c9eec0a1020702329ca97d8d42170dde8a9f427cbc2c9c28d390c",
				PeerID:          "12D3KooWAp9d7Hn4uw4nEwpmvUVWCoRH3qbuLhgtayDLNm7FK1KH",
				OffChainKeyId:   "0bc034b0fc81f72b0db4c8dcf050850cc46440b91e0f1f1926c97eaa44bdee09",
			}, {
				TransmitAddress: "0x947CF3d11fEE3b1dE4EA717F2e192614947882b8",
				SignAddress:     "0xa036fd28D7a5515Fb491E0Ac9f909178C9Eb07db",
				ConfigPubKey:    "ea1a3054bff7c042aec304c4589c398aa95026152f6c4123d54ecf7850ad9a49",
				OffChainPubKey:  "aecd4aeafe86fa61c39c66059d2be05d8dffb32609d4c52c9ff567d7311353b9",
				PeerID:          "12D3KooWMDF2YxveKKa7rYN4fM7Uz46basme4fAHR28uwEoeKu45",
				OffChainKeyId:   "822d0a640c94f9cbe713fa650c482fc0655b0f489bd6334d82283eb901648cca",
			},
		}
		nodeConfigs[BitlayerTestGre] = nodeConfigsGre

		nodeConfigsDog := []NodeOCRConfig{
			{
				TransmitAddress: "0xD9FE85C6Adc5a2DDB092E322Fa2C9797B4ADB625",
				SignAddress:     "0x52F9B248F52A09Db09992Aaaaf470f7D0a3feDA9",
				ConfigPubKey:    "46430915beeb377bdf536f0539eb9092150551c72e6abb1ecde02ef6636c7d23",
				OffChainPubKey:  "42fe0718e634869982105951a610a10efca4af23e8c57e9d3451974554d97e05",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "38cc71fa2059ec7acf95ad8d2cf238d56d0f718455d29a46965b405315c154ec",
			}, {
				TransmitAddress: "0xDEe1D23Cb98D01f79549aDc0E9794e1A85930962",
				SignAddress:     "0x6AA137D66B023bf69BCE06aDe87274b32992EE8B",
				ConfigPubKey:    "8877a5cf6b46b539948dffd7b93ea49f7a6934af6c17452caede452be9a8007b",
				OffChainPubKey:  "5a2ca8198f25d3d5d786d34f220375b5f0a9a4ced26dbc6fe8ecee8220fcd3ce",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "2875e29ad45e282750246ac28464b0208d203b39881e55c54bf6bc11a41b3232",
			}, {
				TransmitAddress: "0xa485f6c06bE52Bc0258F16082126fd7817f51FF7",
				SignAddress:     "0x07ba42CA8DfbB2BF199d98131593A2D05248f7F1",
				ConfigPubKey:    "ce9acaf32d4bf335b7a11a69fe67af47f68b73d7a0d9deafe08fbaad46d10b70",
				OffChainPubKey:  "e915fa720f8769a921c9998596fa21ce5c8370333c1fa2a082dc2a5bef3c4257",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "b640152c10b39838713a5c334babaecc1424eb0d90817cca2840d950708f4289",
			}, {
				TransmitAddress: "0x7bd85D2a76EaAf975Eb651E60b4FAE8B903Fb71F",
				SignAddress:     "0xF0bC14bEC3553Fa0e3c9816600A94f24AB8d52BD",
				ConfigPubKey:    "7cea73e0af2d6241474e2d962797d543c26c4f84d01700215f2d841b0e48be6b",
				OffChainPubKey:  "1e94710d9eb26e26c2e987258296d0ee705406edcc496baac20c98b63a79c48d",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "ae230dbef255d4509f1abe960b289af38bdf884edef981c233983a3b5553bf42",
			}, {
				TransmitAddress: "0xF0035932Ed79722D1715961667ad66830a882125",
				SignAddress:     "0x6103E81bDd8A6c1Ef074d3F4AA26a82EC286661B",
				ConfigPubKey:    "d245c4a02df5260c48a6929c07802bbd1c0a4c2359893050ab4ffa046bfda152",
				OffChainPubKey:  "8c6cc1e9565354d73121f51bc03e861d64b63d58a4df621ec06b69cf23a73d3f",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "b4b011815be193c9eb345cbbd8a1c9ba87dfab31b02caa62d95e7728da5f6300",
			},
		}
		nodeConfigs[AilayerTestDog] = nodeConfigsDog

		nodeConfigsWzrd := []NodeOCRConfig{
			{
				TransmitAddress: "0xD9FE85C6Adc5a2DDB092E322Fa2C9797B4ADB625",
				SignAddress:     "0x8070326c3C71F746dD564A589656dcc1E0921320",
				ConfigPubKey:    "58357c689bc0dc4b87a7aba987afea2f932ecef2f258d82ee394b19efad9f45b",
				OffChainPubKey:  "33e806234468655b2de53acc378ef402aa7d8e6f9e5bc1964b4cd455206945f7",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:   "dceaa8717764ce927734f8277ec9cb4bf591aff2ee6fae42fdb75dbeba72e613",
			}, {
				TransmitAddress: "0xDEe1D23Cb98D01f79549aDc0E9794e1A85930962",
				SignAddress:     "0x1281caBa4cE8D29E2C2c33aB4028982c91675260",
				ConfigPubKey:    "6a395fab95859b144eb734b5d6635d98ee9cf6b2875b40f9360c53f723ef911e",
				OffChainPubKey:  "70b407f55c6cabcd5f4948bfc062c664e4953c8e0f579b7829856399d5fea605",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:   "fd94d7788620d6e322b1fc3dca5863d334054436b3b2cfb872601ece92df05ee",
			}, {
				TransmitAddress: "0xa485f6c06bE52Bc0258F16082126fd7817f51FF7",
				SignAddress:     "0x136eA669cFFdeb2c1bDAc64860b51Ec2C4889EAC",
				ConfigPubKey:    "0134052578d892fe83589071d8f2456d674de64885c0c98244657d4e6604d231",
				OffChainPubKey:  "ac579b7ea7827f5261729ade06fce5e9a1fc7b5490c37429597ce369828e5aa5",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:   "f8077b5c0f9761de1dd857b8b8c51afa55dcfd00d5438ee810d59344c17f5b87",
			}, {
				TransmitAddress: "0x7bd85D2a76EaAf975Eb651E60b4FAE8B903Fb71F",
				SignAddress:     "0x16F087B13a55d8620b6eC4028F6Ed67B5A35f843",
				ConfigPubKey:    "79d5aee9e9edeeea38b9d87b6472a61f4e2ac34f4c33ac0303e16cdebb2ef23e",
				OffChainPubKey:  "53f38d91fa526ea1f9592d32324a14ce394fed49167da1768c6682b625b783e4",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:   "896d10d7b086d59c5bb4499f5b0811c294cf3bd434201796cec28536b07a90de",
			}, {
				TransmitAddress: "0xF0035932Ed79722D1715961667ad66830a882125",
				SignAddress:     "0x29Bf56c0BD2BF68256690D2aB4226806f2C05d11",
				ConfigPubKey:    "1777245e3631115b4f8f9c3dd2eb36b05ece911ecd0cb84a0f3a1e0746e2c76c",
				OffChainPubKey:  "47bb80b90afed6779b2fc516ec847a70284df585690e39b19481f53bdbb97c8b",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:   "1a17c95e916bf4002ede10eebea01e70a8ed7e92dafb8668408ba34860b8ded8",
			},
		}
		nodeConfigs[AilayerTestWzrd] = nodeConfigsWzrd
	}

	// sepolia
	{
		nodeConfigsSepolia := []NodeOCRConfig{
			{
				TransmitAddress: "0xE0caa08142096583C4E7Be197885ffd88D07d079",
				SignAddress:     "0xd618B820FAFC1F4B98d7BCa6F125186A5fc04F8C",
				ConfigPubKey:    "8f8602b508b0c554fde124a7ef5fa4d066887010a9e9651efc11a40556c12c3e",
				OffChainPubKey:  "9ec24f8462e8d8937b676af5656ef450eadf128dbeb50d9fb00654f7f0ec8b6a",
				PeerID:          "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
			}, {
				TransmitAddress: "0x591F1555E4aEeA2B9bde7C829C56208b6F5e0e38",
				SignAddress:     "0xCD5808Cc2E6Cf91Cd2B77F68E33fFe3e8F95736d",
				ConfigPubKey:    "46ead7614db523fa06060ab2a7f8cba91299fc66adb7d8c63ef5a5e11a868018",
				OffChainPubKey:  "39a902f81fabcfbcbca1d60147430ea6525500424c9113e5b7aa51149b791684",
				PeerID:          "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
			}, {
				TransmitAddress: "0xF59C40698a3002EEC8eBa6C49AF86B803a222C55",
				SignAddress:     "0xEf9AA3f276C5246A5736dF40CaeB5ebc116747eA",
				ConfigPubKey:    "3a59c3bf117db4183ee3f980dfc8eca5ca51bb5b82c4ce9aaf8cc540ca12e464",
				OffChainPubKey:  "99c3e9b464318991c88b07bf109e8f9994cf5b905aa07980c64f8e48f5c55d17",
				PeerID:          "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
			}, {
				TransmitAddress: "0xA70c43ba08f77c5f2e1D5BeEE62d4559CCc01EA3",
				SignAddress:     "0x988722c55d231ba13Feb39F53a402F3a35EA079c",
				ConfigPubKey:    "3b1d8e27c64d21491647471d3f8b81b0d8673deba243b6d24e1f812906185b51",
				OffChainPubKey:  "01a37258fa071968ae944b5f06b3b7aae3123037029fec784fd0a22f299670ec",
				PeerID:          "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
			}, {
				TransmitAddress: "0x85C3ea9c83c18FE173B93CCC9abB1B9540CA9bd7",
				SignAddress:     "0x79f1D85B77ac14D901C83C015A7f696107d6FC5f",
				ConfigPubKey:    "869d1342208dfcaed5c596fae64290b4a669c07c3bce45b5357d0264b3c64312",
				OffChainPubKey:  "3f930dc6e25223622d1e12e70c506b8caeff61a116c2f3d1ee7672d25eb6bbdf",
				PeerID:          "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
			},
		}
		nodeConfigs[Sepolia1inch] = nodeConfigsSepolia
	}

	return nodeConfigs[target]
}

func GetOffChainAggregatorConfig(target int) OffChainAggregatorConfig {
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
	ocrConfig := GetOffChainAggregatorConfig(BitlayerTestGre)
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

// https://etherscan.io/tx/0x507a2a9524b3efe4b2b1ede7be75b5290fec93dbb2144c7416699d6ac5737c12#eventlog
func TestDecodeOCRConfig(t *testing.T) {
	encoded := "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000826299e0000000000000000000000000000000000000000000000000000000003f5476a0000000000000000000000000000000000000000000000000000000006fc23ac0000000000000000000000000000000000000000000000000000000002cb4178000000000000000000000000000000000000000000000000000000034630b8a0000000000000000000000000000000000000000000000000000000000001312d000000000000000000000000000000000000000000000000000000000df84758000000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000003a000000000000000000000000000000000000000000000000000000000000005c000000000000000000000000000000000000000000000000000000000000009400000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000101000095a49c27564f213a6fb0e3f2b857fd80afb8ccafd7849d45fb2423448560a5b04cf86dd13e69dd2af007b2e643639fb13ce2c013c61feb1eaa6c2038a47731264588417fc01b23ec2aa5dcba22748fab7cee5783df63dce7745063fdce4b8476abd625ce4d9d901db0a640aa3ec00b3d5c85cb6eef4399f2d5d0439d6128dc3176c20b0995b5caa04acdb423e2acf0ccb09ee73fd6df255efc2a5335dc9c9eb29778c052d6abbf3daeb78ac9324f4f4a866d394e8803c7ce208c8aa445e207b131fa2e3c7c43ebdd59a96018fd9996034b4ef188933a509c289066d95b18b0ce5b43d632b2a54f362c062869def58ff0216134eebde40c58853e79dce4437e4e16194221a496c3ba7f6874b15949fc339c2b34b2f88572d17028cf373189862becd6be87a521e8d4ae3359a2ebc2f88f76b7eef7a703e6ae5ef65bc189ad40d0eb62c06d81b31032ed9d4b9f4c960c07e1f1d8de52f39a958e8619632f49b3561165c51623db2aa47df522738f21adf1145223ba2eec3aa49021e7f78b597a5aeac952b584fd3428c62d7426346edeea969b2631b599f68df4b3020b1580fab529c2dee912404d5d2b0c6ea9da1102bc60e65b278d9ce0058a599edeeda949d496e74107370390b6897005b111c18d88f1725b408caa9bdbebc85f49e5539c4b3f13a61dc96be4896129726f08c9a2ded7126db6d34be923175df457100000000000000000000000000000000000000000000000000000000000000034f313244334b6f6f57433631686a66327a7a4544733359737032656438664c5534793661366d517934425339536235794b317a644c2c313244334b6f6f574348703956644b54626b384137794170523937336334554a3646556357665443626e7a5a6d6d50554b6333502c313244334b6f6f57484346714c775361346353526a735a757a42746a486852774a57757036704e624436557933574a59617070572c313244334b6f6f57527773763538616355715167355146384447424854776365447a4a684b5879734c744d414e73714e6a3555592c313244334b6f6f5745514245314841334b66414d7851794c475478477a46347043737a4a7072317864645141367a3476683756732c313244334b6f6f574248537a526859506a6b786f474d38535552584645487941344c7177484259674a5a7a686452484a414b484b2c313244334b6f6f574b4e536b6d7a4d446144503137354668717261754d6f33465871756b3663726841686e666d74326374334d502c313244334b6f6f574c72394d4b53467172683439436b356f78766955346f4b655a6d4473335a714c347a5671436b4d75333776792c313244334b6f6f574c705361575643787136525342756d44796f726b674c784a6d445a4c786659426e69317a79544e52644637332c313244334b6f6f5748506577317a67705a67726947436175666568633343367436484467374b79557a7279705a5369355958654e2c313244334b6f6f5743476679374c326555545a37644a4444786146557a7432523334506467554d797a7a64726161746b506877472c313244334b6f6f5753704c4d6f743366553751684a7731616b7179637a6f3443485172706f6e547144326a3756677942596f446e2c313244334b6f6f574e4c586f4b5a475a46565441673946323773794542596761354569784d66686a4d523867557a4770734e32572c313244334b6f6f574158426f437a536b4d547a7454726e4d6f6847464c506f655938425252795758517356344a6b615a5177434d2c313244334b6f6f574e5861486b697675364a43556e3852424d5a6d624343666a33517573706e43574d6a50665943474b5069455a2c313244334b6f6f57416b743236773339695769746b57795867797259385341415765705576316a55504c5a3639696a4d645266520000000000000000000000000000000000d9e9dfcc1e34a5f4947e74dd9783a5e045b1b1665e4d34cb35e282bb760d76794ad3c99c41279370538e5942de879803f3d2a4b76095ba11ddd176231457b6be00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000010158f2a8b7447e42eb37622aa6603a47a00000000000000000000000000000000810a4888d1c50711c67435a5c5213787000000000000000000000000000000003b109cde9c14d9c11631d6d4b97f68af000000000000000000000000000000007bec6dac5ff1388422983fe505c09d25000000000000000000000000000000008cb1a0415baf3eccfbedccd702f5795d000000000000000000000000000000000839ae871bc8241027857140b9d2a12400000000000000000000000000000000e08d8f472bb3e09a28f67ac83f6a30ff000000000000000000000000000000003206c206e2af29ac382571df2255af9b00000000000000000000000000000000900b8ce285710e8c146f1df195e377580000000000000000000000000000000045ac4248e69fcabc088c0cf1df0671e0000000000000000000000000000000008e23256184deddac5f209674c266be4c00000000000000000000000000000000bdc18b21f15f49c6bbd7cb302a2839330000000000000000000000000000000065011b87413b7197006bd5d42ffb34b9000000000000000000000000000000005b0fb30f666bc4239261c06487337967000000000000000000000000000000000d5f7feb99037cbefb3641b4c7dae3ee00000000000000000000000000000000abb2ed6271b2473c106b47ae1867cacb00000000000000000000000000000000"
	code, _ := hexutil.Decode(encoded)
	oc, err := config.DecodeContractSetConfigEncodedComponents(code)

	t.Logf("signers: %v, err: %v", oc, err)

	/*identities := []config.OracleIdentity{}
	for i := range oc.Signers {
		identities = append(identities, config.OracleIdentity{
			oc.PeerIDs[i],
			oc.OffchainPublicKeys[i],
			types.OnChainSigningAddress(change.Signers[i]),
			change.Transmitters[i],
		})
	}*/

	//ocrConfig := GetOffChainAggregatorConfig(BitlayerTestBtc)
	var configDigest ocrTypes.ConfigDigest
	binary.BigEndian.PutUint16(configDigest[:], uint16(0))
	cfg := config.PublicConfig{
		oc.DeltaProgress,
		oc.DeltaResend,
		oc.DeltaRound,
		oc.DeltaGrace,
		oc.DeltaC,
		oc.AlphaPPB,
		oc.DeltaStage,
		oc.RMax,
		oc.S,
		nil,
		1,
		configDigest,
	}
	fmt.Printf("\ncfg: %v", cfg)

	/*if err := config.CheckPublicConfigParameters(cfg); err != nil {
		fmt.Printf("\nerr: %v", err)
	}*/

}

func TestCode(t *testing.T) {
	a := []byte{100, 109, 96, 26, 88, 1, 194, 196, 16, 26, 193, 168, 155, 87, 102, 220}
	publicKey := hexutil.Encode(a)
	fmt.Printf("\npublicKey: %v", publicKey)

	b := []byte{255, 35, 91, 9, 175, 185, 120, 116, 128, 136, 119, 225, 217, 183, 240, 169}
	privateKey := hexutil.Encode(b)
	fmt.Printf("\nprivateKey: %v", privateKey)
}
