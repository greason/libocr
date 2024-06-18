package mode

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
	case ModeWBtc:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case ModeMBtc:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case ModeEth:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case ModeWeEth:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case ModeUsdt:
		// 0.1% / 86400s
		AlphaPPB = uint64(1000000)
		DeltaC = time.Hour * 24
	case ModeUsdc:
		// 0.1% / 86400s
		AlphaPPB = uint64(1000000)
		DeltaC = time.Hour * 24
	case ModeMode:
		// 1% / 86400s
		AlphaPPB = uint64(10000000)
		DeltaC = time.Hour * 24
	case ModeStone:
		// 0.5% / 3600s
		AlphaPPB = uint64(5000000)
		DeltaC = time.Hour * 1
	case ModeBBUSD:
		// 0.2% / 86400s
		AlphaPPB = uint64(2000000)
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
	ModeUsdt = iota
	ModeUsdc
	ModeWBtc
	ModeMBtc
	ModeEth
	ModeWeEth
	ModeMode
	ModeStone
	ModeBBUSD
)

func GetNodeConfigs(target int) []test.NodeOCRConfig {
	nodeConfigs := make(map[int][]test.NodeOCRConfig)

	// mode-main
	{
		nodeConfigsModeBtc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x625540C630ad7582FE50c03E5faf5A47616A5bf8",
				SignAddress:     "0xc65Cd9fa0b67E243a8513B410BF52b4E450D68F2",
				ConfigPubKey:    "aaac672575852813575963af9e9c1bd7aaca5eec32113dd5e42d76e625173053",
				OffChainPubKey:  "a88177e9fb035f62e49c6cd3f925ff191157f9442afc9020b737f4dce6e76661",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "84f9d65924dec2c7f1d01223de62d4bc70aa398f772a7f53f65546f24effa33b",
			}, {
				Id:              2,
				TransmitAddress: "0xBf732AA2eA42515D2ca624A003aE19d1a220d75A",
				SignAddress:     "0x54753a62Dea0BaFEB774CD55614a2ab84D1A5391",
				ConfigPubKey:    "3990a890b96f031fcffc4b0f827c0ff622dfbba5a66d02955c9fe5693bef9322",
				OffChainPubKey:  "45c56a9b2ab9790dbaa3929ed2295b573eefb1d5296351a6bc2393ceddc95784",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "8a2545fd84d45ec9c8dc5f427ec0763e31ab0276ae09f3927943032da5d4a12d",
			}, {
				Id:              3,
				TransmitAddress: "0xCeB7Fe944093a683F4A885C3DDfF774C5ee48C32",
				SignAddress:     "0x8aaE2f905c2fd3140d311e20804f9c3bE160cD92",
				ConfigPubKey:    "18787929ff0265aab677e76d44498ec5994bddd42f91ad7275261cc5c58d5678",
				OffChainPubKey:  "266bfc8577b852cd5542d50014820e794d691b8210e77c6dcfe4070fdc32998a",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "701aa8a99828977482ee185350a87a66d269f7c9c097a9634d333cee05ed388f",
			}, {
				Id:              4,
				TransmitAddress: "0x4C6aaAd9Dd9Fe03F72C261B0179F364745C3796b",
				SignAddress:     "0xE829614Cb89A01577ffbA774FF3194706CcbA198",
				ConfigPubKey:    "4cca0d86ff05f5315ad4422b194c5d3717b440f5c031b14711ce3de8cd35f819",
				OffChainPubKey:  "03b4b8775db0059068a14ba0d6d2837f12d6e7ac9d25fd7b046e73a2a0cabcd8",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "762dcf76ec27b22ffd99d1cfdedf30e6ab1c864a42e154ee07494b162c7cb7ef",
			}, {
				Id:              5,
				TransmitAddress: "0x3c8f40B157Fd9a0d7f0Ac6c4746ECC47247AD574",
				SignAddress:     "0x3DDc75E4F27a06BEDeEDE732f655C4a9d57b19a3",
				ConfigPubKey:    "b4fd77ddb3cd87fb6ac40c23566179b6925417c9e2d8d378792a9a4702b03967",
				OffChainPubKey:  "487bc07110e84d9a22ccb74812e274bccb3215303952ce624b2c3df45be8f9f3",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "4c8146965a3fc60f4b6df8671850958e7ef88b93b1ccab4181d35ad3b2c3f9b4",
			}, {
				Id:              6,
				TransmitAddress: "0xE7b320CdcB0514CFc7fACEC7FDd32Eee6d9fd2A5",
				SignAddress:     "0x339ff154b6bb9107d543E222E8ABBAcead343A34",
				ConfigPubKey:    "8b127999ee9d1c2b0dfe039bd94bf54dcf1d11223bfd382fce97a877bf834576",
				OffChainPubKey:  "6ae250b4e847bee9d044b00f7fad9ccd886c5adf95f8c3997f5f726b6dbdc991",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "38e7ebfb97ad802e8dbd0d71464f029cbfaf0da0c8127b7e14a69c12f9871db4",
			}, {
				Id:              7,
				TransmitAddress: "0xeBa46834073CeDbAE9322C677A132c7a9AE911c8",
				SignAddress:     "0x708c145f523D09312307F0276cc2A223609e3FBB",
				ConfigPubKey:    "68d5709571e240595078b43e7d9cfa7c28a380a4d2392224b4f4813b27a64a5e",
				OffChainPubKey:  "cf0d6803284930611ce5e07bca5bc5cffa3497b900c9917735f9a26332946172",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "302ddb9c3320f2f2767d8d188cf57226543a7d783b3a6a40fc1f417b8f097cf8",
			}, {
				Id:              8,
				TransmitAddress: "0xB742c3eda47205952420B314C5429664C8A8C406",
				SignAddress:     "0x06c6f554FC1cCfa709d0Ae94CA52B7B6e175c16a",
				ConfigPubKey:    "215da41977afd3317da6b2576efa014cb260e96b20d69e036350ecee3af2384d",
				OffChainPubKey:  "73811726a57934fe3b95ce91398d05aee1b28399d216fd171f2b580a37047660",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "366eb94e26e65418ed88adbced3e6ee6f7244913e21b25e4b6bf99ba7e430862",
			}, {
				Id:              9,
				TransmitAddress: "0x0f0967cE269cB73A27918F0786688FE09D7488dD",
				SignAddress:     "0x9221bEcf1Cf6054331Fa7Fd44355776aadCfb4d4",
				ConfigPubKey:    "5a1ef0ae7bcfbe0ee888cb6d35bc12237f381f99a5cbbadd521d148ead91fa33",
				OffChainPubKey:  "659ffa90003e20d0e032fb5fad96a82213f0c6c261d1bf2faa2f939a20c1f440",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "2089a2676d68514fdd15de4d6cdda8d1fec71078d9b52e8a00d34b4198c1e4e8",
			},
		}
		nodeConfigs[ModeWBtc] = nodeConfigsModeBtc
	}

	{
		nodeConfigsModeBtc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x625540C630ad7582FE50c03E5faf5A47616A5bf8",
				SignAddress:     "0x8e35906F1b59acE3Ed919F5aC6adfb92141BCa34",
				ConfigPubKey:    "9fc4d41b6c358c53a64529344e5bad1795d2bccb0473544550eb2fa460f0b53b",
				OffChainPubKey:  "efac9fa658ed32d0ea6fdf3a0911264f875507a2345f94f2daea134ca9d8b5de",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "10e83059d6ab8e2f5c24019f4ad2760873f9259920ad93fb815de1f23c12c3b4",
			}, {
				Id:              2,
				TransmitAddress: "0xBf732AA2eA42515D2ca624A003aE19d1a220d75A",
				SignAddress:     "0xcF478f929B4FEd4848DE92368F22C6e0f0a5681A",
				ConfigPubKey:    "6433cce8c1d39f22986b5ab01117fe4245b660e33e1a42175a8914e05af0c234",
				OffChainPubKey:  "bddf70985363bcc7a045ae68fcc53ff18c4d0292a7f57ba3bb3894cd0bdddf16",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "0e495381e85b865972407824fa31cb7ac0075ac9fdc9c2b92b798ad6123c583c",
			}, {
				Id:              3,
				TransmitAddress: "0xCeB7Fe944093a683F4A885C3DDfF774C5ee48C32",
				SignAddress:     "0xb55808a2326F6cC8d5cF32a6663CE0d78d001E38",
				ConfigPubKey:    "b9de76396d8c528135466a9ef1d9ab0093694da2550a0325335c9b7593d6a22e",
				OffChainPubKey:  "86d49e46c3eeb70eec3c6d7a81542404d0cf81252318465a4070b557cccf6f41",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "937c7f234fe7a0be5ddab4eae51dddb22d0883e5f00a76a733f00c58a41f016e",
			}, {
				Id:              4,
				TransmitAddress: "0x4C6aaAd9Dd9Fe03F72C261B0179F364745C3796b",
				SignAddress:     "0x35a87E63944504c978eecFFF3C45264A051BeAb9",
				ConfigPubKey:    "b1ab5c7ed2388d7d2708e88fb42964d255807f080ff5099be147c71b54fdac5f",
				OffChainPubKey:  "48cc6c75a7fa74243b76ee7a68570da5c0a51893ff6dc5b546d34bd15aa5bc7f",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "0157a6847c9661fa063d03db95600dd15c42fc5bf7a50ec046ff8b24cdfcc17b",
			}, {
				Id:              5,
				TransmitAddress: "0x3c8f40B157Fd9a0d7f0Ac6c4746ECC47247AD574",
				SignAddress:     "0x88522f053C6a145A340b85Ac9034846AeCc552a4",
				ConfigPubKey:    "cb6f4a2fec2b115aada2226630cf0749d17b16891c7935a07916205bde90f758",
				OffChainPubKey:  "26f92cb72f43406c31eab25cd5f4ba405675fc8908d60a10db9c344123b93c71",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "7d1625512f90ad64af2dcc37f3ede293f1e3a52d9671152f2314d6ea5029c293",
			}, {
				Id:              6,
				TransmitAddress: "0xE7b320CdcB0514CFc7fACEC7FDd32Eee6d9fd2A5",
				SignAddress:     "0x98bd768047244A561675923BA8deb010b2b78dE1",
				ConfigPubKey:    "d67ecacc2437aa7ae8f7eb3b32bb041e20f85d034ec21fb4c55be245182f1427",
				OffChainPubKey:  "a0516ffd3a5fc893b5dffef703a5a2fb5271cb3f5337cdcbcb4e54079145b09d",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "3b70331b3551310d899220c2c60d4939aa3a18f531ad906be0f015c629410a73",
			}, {
				Id:              7,
				TransmitAddress: "0xeBa46834073CeDbAE9322C677A132c7a9AE911c8",
				SignAddress:     "0xf01Bba1eC6f7583Ead76b17061F23dF6E3A5162F",
				ConfigPubKey:    "0c1535185335416a37f016f0e960365cea6c98c7d1b7f0f3732a8d538733b43b",
				OffChainPubKey:  "5ea372580052f1ab8f6217df0f0a9b0836588bf807da94b30cf315109ae15a23",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "49e38a192b63f072bc3c312072190134071614b787705abf7a5ba365b571a496",
			}, {
				Id:              8,
				TransmitAddress: "0xB742c3eda47205952420B314C5429664C8A8C406",
				SignAddress:     "0x545ff72F785cfe6624b879eA2cfDc7FaB82D7B1d",
				ConfigPubKey:    "707aa741a4bf4a604fb7c83e606f96a5943378f824cad4c4f8e8c4dbbd45350f",
				OffChainPubKey:  "2cf11c3f842daa13f74de9b643f09121da8a691dfff086aa96bb60423ca46655",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "9c784f53cd4ecbea5ed579d357600276721e359352b0fb4791378cac01e5d948",
			}, {
				Id:              9,
				TransmitAddress: "0x0f0967cE269cB73A27918F0786688FE09D7488dD",
				SignAddress:     "0xD735ED968736b6552B4e524250D4Ab6440b2da20",
				ConfigPubKey:    "98e5f641130a57136d1036f48c62f5f39918b1e9cd31bd6adfd8211cc4233c41",
				OffChainPubKey:  "de70428c5663b715bc9d18bda9bf64a1027e85b1a8bee7b2bd79260d0b8fb8c8",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "90b60f3e4209429995f0779486b840e96fbcf0312afa20fd5b015383a74525e0",
			},
		}
		nodeConfigs[ModeMBtc] = nodeConfigsModeBtc
	}

	{
		nodeConfigsModeEth := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x625540C630ad7582FE50c03E5faf5A47616A5bf8",
				SignAddress:     "0x6769FF20B0d1b6D2c47E5A98330a26a70F2499f7",
				ConfigPubKey:    "87310bfdf1cb411a728a012df9d815319e0ae0499339b87434246fd8a6a2e209",
				OffChainPubKey:  "869f079399ec17e3b68e04569c0cdfe3fae42592f97b92fd8b454eaa137733b4",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "f7676d73b28f5c8580431b9c4666bf44da8c0ff3a3493165978ca04f1c05b4fb",
			}, {
				Id:              2,
				TransmitAddress: "0xBf732AA2eA42515D2ca624A003aE19d1a220d75A",
				SignAddress:     "0x39B23D7b533ae76D054A853462067E9A5be27A2D",
				ConfigPubKey:    "4b8bdef76de0ec38aa1d50883357859ead776a0baa2a98c83f74713cc2024474",
				OffChainPubKey:  "7c0aec539b20a4f56ef7afb052cc53c87216c33d0aa305c32d3e098322cd1da0",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "f1a26da49465e3f5bb3a03b804dafe2d85d809436845db88a52de2e3533e4a6d",
			}, {
				Id:              3,
				TransmitAddress: "0xCeB7Fe944093a683F4A885C3DDfF774C5ee48C32",
				SignAddress:     "0x0E482f5dD7dE38aB4aa026D475Bc8868e4F23dc9",
				ConfigPubKey:    "77150cb9c56d1be841c286f04f15efe0eeefa4457f4d06896dea05c304053f7e",
				OffChainPubKey:  "9545ea04b11e244508df17fcffebeb463ee0c902e1a690c3d5726c10aeeb8361",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "5eb30805f23c0c2fd8242c0c836bf6f5511dd3bfd1ee9e1ccff44483d1d3b7c6",
			}, {
				Id:              4,
				TransmitAddress: "0x4C6aaAd9Dd9Fe03F72C261B0179F364745C3796b",
				SignAddress:     "0x55AeCD96c7651a005AbAb5007C47291eB9Dc097b",
				ConfigPubKey:    "c3885c320c40acdf9cf7db70c603afc18d98731e5a49c94ec3e2ee3822446908",
				OffChainPubKey:  "0f483e813109c83e716bfd7cf16ac6b3b39e288ba13fc036a1dd7adc75a60fac",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "4e666fa31dcd04e3eb18647f2d29d27ae6e659be2289fb6bdaef66a9210bc9d3",
			}, {
				Id:              5,
				TransmitAddress: "0x3c8f40B157Fd9a0d7f0Ac6c4746ECC47247AD574",
				SignAddress:     "0x3E36E196b743aEE533559fbEDB6355B328704903",
				ConfigPubKey:    "f4237af2de7533a88d7f0a08ea14162bb754d39bd4324ec8980ab0335d335003",
				OffChainPubKey:  "e93a5d38d4aa4cfa48e090fb852dcf8fea3a0006ff8f9465909bc46706005d48",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "f4e38b3f7a767c6afe0e5f231e781b8e4bc28107e3de73784c79db2fa154be0e",
			}, {
				Id:              6,
				TransmitAddress: "0xE7b320CdcB0514CFc7fACEC7FDd32Eee6d9fd2A5",
				SignAddress:     "0x4E8384831901F260Db0E58bE5889845f77627c89",
				ConfigPubKey:    "5d067d6253d98d3b9f3f41aa1ae425d22146ed7fe4eb431b10b47126c47a995b",
				OffChainPubKey:  "541befec9aff2ed6e56c20bcbdd34d823ce94a3492073cee9ca6b5829946f535",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "3d1c12b901cf05a29e7b0e89177b23c91d887c0fa578afcdf7b45e9660633278",
			}, {
				Id:              7,
				TransmitAddress: "0xeBa46834073CeDbAE9322C677A132c7a9AE911c8",
				SignAddress:     "0x4a8CC7F35bDf86d3b1C3c6D2B9ffC0b25427E896",
				ConfigPubKey:    "57093f740a0fa80bfd10be21adae47aa1956ac238bab7d730c25d3094ec47b6c",
				OffChainPubKey:  "5e34a6f9749fb2985d9c4779e10dc26784b471d008c5b8b699aec49089af882c",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "d2b79c8f71c26a520db30a5b279dc5c8498ed7cb226d888e6e732b4cb835d98e",
			}, {
				Id:              8,
				TransmitAddress: "0xB742c3eda47205952420B314C5429664C8A8C406",
				SignAddress:     "0x587b967023d001992060655023e7B49F4Ebfc9A2",
				ConfigPubKey:    "c53ed95a4b38eff8f8b82ffe7c8524e44da8ce67e0636779fa955e252c56b008",
				OffChainPubKey:  "1aab4e745a25cc0a402525c5d3ddbb96bbd6c46cab927fb3e4e004fa7b1fa23c",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "7d29dfb326f56ae6920c47a3a8b47ea93152facd5d705b329b02944c511d64da",
			}, {
				Id:              9,
				TransmitAddress: "0x0f0967cE269cB73A27918F0786688FE09D7488dD",
				SignAddress:     "0x8b2E27cF574599c3A6db9b21eE151D8c8C16dbb8",
				ConfigPubKey:    "9965adf66b060b40a59807b4e396202bed620c8f989167268d2357e68c00a80b",
				OffChainPubKey:  "792ce8b7781577bddc572c8940fe7b564cb638f0e6519e0ebf4a4f8c8ced4d1f",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "6ccfcebbf5bed7dfdc0c2721d775eaa794bf2b9e75e8c94f030eee9886d41e67",
			},
		}
		nodeConfigs[ModeEth] = nodeConfigsModeEth
	}

	{
		nodeConfigsModeWeEth := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x625540C630ad7582FE50c03E5faf5A47616A5bf8",
				SignAddress:     "0x6bdF5B31C3E7ddbAb9Ec94037116c441ea9be7dA",
				ConfigPubKey:    "e8718ef3685fbd830d79e2055ea44f1232ef282befc6c4cf097b59149fc2d256",
				OffChainPubKey:  "ebc1d87855545799a033e27207bd78b1837605a079519c892005402c20661470",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "18c39caa4801772fecb26241f07fe2c84fabb3bb8b8ee1a8a88ecbfdf26fe77c",
			}, {
				Id:              2,
				TransmitAddress: "0xBf732AA2eA42515D2ca624A003aE19d1a220d75A",
				SignAddress:     "0xeF8B7C70Ec0a4503ce0D6DFd4353A0546Bab3537",
				ConfigPubKey:    "9f8c094e05645644296861c96f4db04c205c137ea392e3346e5f61a6627eb320",
				OffChainPubKey:  "e799d0a4faf0b5f2142bcef5b41d00747c3286284168666a14b154086ef49d40",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "8388a8b7c2769c1543e338861c8633133e3e2eebf54bda5a5b6cf62ad1815b6a",
			}, {
				Id:              3,
				TransmitAddress: "0xCeB7Fe944093a683F4A885C3DDfF774C5ee48C32",
				SignAddress:     "0xa34944e0d71fe3923775a395dD9CB40fD683B135",
				ConfigPubKey:    "51a1eeef4feb3c3817f01324e6a7e8d32e10dd31b629f22c15b1baa7966b8602",
				OffChainPubKey:  "1edd2bb9b335d918f54ba1a10f2206e9a01695c2dd71c594c9bc687770eba46e",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "2df1d34fb9c8228da86958e779d996c044754d2cdfcedc42f038527547083e81",
			}, {
				Id:              4,
				TransmitAddress: "0x4C6aaAd9Dd9Fe03F72C261B0179F364745C3796b",
				SignAddress:     "0x8Ac03c33727eA14dd989a30Bc8F0c3EfEfA30934",
				ConfigPubKey:    "f554e737ef8375704a0b550c501ba9663fb7fa4dfc7989136c82d163fd144d5b",
				OffChainPubKey:  "f3bb0080dd20409125578377fb150971cdfbbe4e3954571f2a6218bfe0b1cdaf",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "efc00da4033adb40e91cc06919973b45d3557ba9b14ec9f1da2e6d4c223a7e4e",
			}, {
				Id:              5,
				TransmitAddress: "0x3c8f40B157Fd9a0d7f0Ac6c4746ECC47247AD574",
				SignAddress:     "0x8F9508c57f43d8a42cfaAd1d1d18C03B627b3619",
				ConfigPubKey:    "8f9c731b603a3b96dc372a97fcb5fe6c583da5d3f1cbf106e0f9f69db53bbb3d",
				OffChainPubKey:  "68cbd0d18286ee78ef167ec8645cec27dae42cb315b5cdfe0b9cf671e04fceeb",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "279f2af8d99b1e24758eb0231d15b53d9c2fb161bb1d631dbcbe104690bb2045",
			}, {
				Id:              6,
				TransmitAddress: "0xE7b320CdcB0514CFc7fACEC7FDd32Eee6d9fd2A5",
				SignAddress:     "0xD9660448c21a69f3424B1cff5f7C502ffB5a70e8",
				ConfigPubKey:    "83230bd43456a629e26eee8ff3ff7a5e919648143455817677a718780e4d2e11",
				OffChainPubKey:  "76a1953217338924394baba5c3d4e745262b6a0ae61a894fcf1553243cfface4",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "1c50e472202c2c9b0e8ca752bffb9225f0706f4bbfc39b23311905899dbdf00e",
			}, {
				Id:              7,
				TransmitAddress: "0xeBa46834073CeDbAE9322C677A132c7a9AE911c8",
				SignAddress:     "0x8CcA1dDA34D68966a626Fbf142BA39430D021b68",
				ConfigPubKey:    "46b01255e176ea0736aa5abe3e274283e0bf0d553689a1ff9f3636575708db2c",
				OffChainPubKey:  "16936813e6c6b4f0eb6351ea559961a02c7f937c1405e270a384a70856e8731e",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "844c346907f4b3dd4630e71ef5073d2cc6d6ba0fc0b5001e946e1674aac0a930",
			}, {
				Id:              8,
				TransmitAddress: "0xB742c3eda47205952420B314C5429664C8A8C406",
				SignAddress:     "0x01db058Aec6bd7223cCFa0a115AFB7D536abd72A",
				ConfigPubKey:    "0f468d35bd681f3cc163f777c63e2a1aaffd80272489812ae1c5681296e74035",
				OffChainPubKey:  "2722e387fcdfb34ee7d82f094ac070ac49379c310327fd0baf1a5b02765109f7",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "54bada6d4c04b6c58d1ef2ede0655e257d94adba1c8707ad6ff5b38139363b74",
			}, {
				Id:              9,
				TransmitAddress: "0x0f0967cE269cB73A27918F0786688FE09D7488dD",
				SignAddress:     "0x323Bcaca0Dc360CB0a5e55e1F143f79827E19d75",
				ConfigPubKey:    "8a26d3c5ee7c32dd62bc4cd635f3a593dce5d5c995b6a02883f19d92efb90e0a",
				OffChainPubKey:  "9f48c6862a1061078a960d7876124b1809cecc8022ceae03cb307adc17391898",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "2e5f390cd9d402c93494e96da018c0c0ffe823d34eaa406467f992265dad900d",
			},
		}
		nodeConfigs[ModeWeEth] = nodeConfigsModeWeEth
	}

	{
		nodeConfigsModeUsdt := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x625540C630ad7582FE50c03E5faf5A47616A5bf8",
				SignAddress:     "0x4129A0DDA8c8CFD62b89414bA22088f6Cf734C10",
				ConfigPubKey:    "8aa42379418814bfc5d07fc6302c0d00ca13356c34d05a76cbc2057a5671e039",
				OffChainPubKey:  "36f77dee2cdee75bb1382d918075e7498d4cac4bcb0149a50537b4eb49c1465f",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "fb334fff7617cb21cdbef5b1adcf3a7daf8594762b755bd25491c351a876dee1",
			}, {
				Id:              2,
				TransmitAddress: "0xBf732AA2eA42515D2ca624A003aE19d1a220d75A",
				SignAddress:     "0x0ad53D4Dd81C167E7996D05d6035716715731A8b",
				ConfigPubKey:    "a72ac8d48405da4ca6cdeed516e0ac32938185bf0ddccee73b8a6aec38e63c7d",
				OffChainPubKey:  "1737262ba9ea1bbec494bde0c0ba821980b55f1955202bb0146a7066dcccfeee",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "feaf06f311df40526e8d899b9530740ee9c4ab4e3c48c197a5159d85675bff9f",
			}, {
				Id:              3,
				TransmitAddress: "0xCeB7Fe944093a683F4A885C3DDfF774C5ee48C32",
				SignAddress:     "0xA3B5dca4CDDFed4988E7a2eecbD2Ed95e2F3121F",
				ConfigPubKey:    "f1b37f082bd51eb84ae5b85d06bddeb053c91276a1c935cc9d3a0c63686a420d",
				OffChainPubKey:  "38599ed180a7a92841ba1696df5f04676b8d002a92b8a2e64c1eaebc11c59f78",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "bbe2df1d77c05bfa0c57f3fc9e5af21c314f00165a68eb11955ab79325feedd6",
			}, {
				Id:              4,
				TransmitAddress: "0x4C6aaAd9Dd9Fe03F72C261B0179F364745C3796b",
				SignAddress:     "0x70011B20BD06fdD0265BdC1708C0C1c53d0c4CeF",
				ConfigPubKey:    "cd691d74ba58e4f2eb2d1ec10c7dd7095900b9a4660fa8728e4f1434fd71ed19",
				OffChainPubKey:  "c261c8956d617c26f987d3f0a49d38959f3c5576c1c3c57317c41a68b3b83def",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "c3cde1ac8f644ccb723ce7d6d44214bd1ff49d17eb48dd8a8f4fed8b62ff21ef",
			}, {
				Id:              5,
				TransmitAddress: "0x3c8f40B157Fd9a0d7f0Ac6c4746ECC47247AD574",
				SignAddress:     "0x402A7c1832b2c925d9eAF39E00D3C9485C0E4945",
				ConfigPubKey:    "a64a6d7ce0cf7047c9f52fa0fa681456becd4a7858f8944ce7996d613137f032",
				OffChainPubKey:  "ff5e82042ac9a4d3c6488c57ac9f39562f4081f4b35329cecd735c9c461b26fc",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "1ccf7269f0e0684f271d0728e28eae560a4bbce462dbc1b151437777bf5aa3c4",
			}, {
				Id:              6,
				TransmitAddress: "0xE7b320CdcB0514CFc7fACEC7FDd32Eee6d9fd2A5",
				SignAddress:     "0x1e221614c9d981CCB91fa50F2Ff03C722B6C088E",
				ConfigPubKey:    "f9021069f658da5c8d7d845917ed58c221e4132a38038c95760e91eadcc82d16",
				OffChainPubKey:  "d0d3be271a9bdc9498d8921d3257c52af126ff0fd1335a541a414148b5443a5c",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "d1fb80efdd436485fe15134c00883dc59146ab581517028921f560e3c25f0c7a",
			}, {
				Id:              7,
				TransmitAddress: "0xeBa46834073CeDbAE9322C677A132c7a9AE911c8",
				SignAddress:     "0xD0E460ff81f0d1244c7D7741c58eeebFBd779762",
				ConfigPubKey:    "749f4148c13503dfac434599edb5c9ad329aff2e0ada63fd4390d93b6a4ef664",
				OffChainPubKey:  "5f3eb32d57f17eae4a5533eac15eee0a172c8b290e966fbc2954cd9f5f796c78",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "bf15b0ca3f261e24fafacb767d7140d9cd976742928eeca9488040bce2c13e9e",
			}, {
				Id:              8,
				TransmitAddress: "0xB742c3eda47205952420B314C5429664C8A8C406",
				SignAddress:     "0x0B79145Ee8686f14435B5F0Fa2C4818b3B7e874A",
				ConfigPubKey:    "cac9884ae159d76c42b9747e12c03c1f54003f8fc95dd5a6a064e2a2cf222c52",
				OffChainPubKey:  "8e470d7d691e55f6c3d6203c0cc3d1972d37d61acd027770fe6e6dd5729c1cdf",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "2760add29ad910435eae9ddbd772b64f095b030fbe8df8e62fb1faadaddf4a5e",
			}, {
				Id:              9,
				TransmitAddress: "0x0f0967cE269cB73A27918F0786688FE09D7488dD",
				SignAddress:     "0x48203fC3d9E003360c77379f046002CAA49E8A7d",
				ConfigPubKey:    "7be5319e1e5c9481c0de506ae4783661b4c29c622a832ee8e351c3ab6385da77",
				OffChainPubKey:  "7623568f7d2630a4130a87f3390cccc2eb6fe0786533280066644c4d8fcdc98e",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "7556fd69d8a1422a53d1b6fcfda090be32a4623fc10b3223451abf4054500f20",
			},
		}
		nodeConfigs[ModeUsdt] = nodeConfigsModeUsdt
	}

	{
		nodeConfigsModeUsdc := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x625540C630ad7582FE50c03E5faf5A47616A5bf8",
				SignAddress:     "0x2132BDBaa1019678137126165e18C343aF61D98A",
				ConfigPubKey:    "ebadccf2e8994759cde315d33c99fe7a19a0c204175d728fa4235e8217fb1041",
				OffChainPubKey:  "9da4531a0d098085cfcbf252f1cd4dbb8ad5260fac9d0a5047e02cc9806b6bab",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "579d6b696f6c157f42ec2f7777ad9400d7fceab268d41abd92ea808bd76b5ebf",
			}, {
				Id:              2,
				TransmitAddress: "0xBf732AA2eA42515D2ca624A003aE19d1a220d75A",
				SignAddress:     "0xCA2426BB7d84654A391040FffFaaB09842c9FE6E",
				ConfigPubKey:    "dc13c34cc14fa3076191751b2127d5f8c0dc4aca0f37cb6f47fb8dda1ea5a675",
				OffChainPubKey:  "18aaa055e05303053eaf02e8d7b929362badf70a3c2392268673b1ca51dfb632",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "7f4ff557916d840d90969809884b8149e249b9807867449d9d76d5c4146aedc2",
			}, {
				Id:              3,
				TransmitAddress: "0xCeB7Fe944093a683F4A885C3DDfF774C5ee48C32",
				SignAddress:     "0xC8B69f13B774e1b919bf8Dd59eF0c254E9998672",
				ConfigPubKey:    "feb52b067e55bbfc3f7faa171a20afde57ac8aa97561195f418c9d795f792b49",
				OffChainPubKey:  "2f4a69486a0b81ced5ba51ea2984a7d861cb8b53f5ed4423da80c8d26011d60e",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "d1b7a74b4de99c3d597b86fbf7b38d63f2c5bfeaa6e7b55701ce9a2ac272b309",
			}, {
				Id:              4,
				TransmitAddress: "0x4C6aaAd9Dd9Fe03F72C261B0179F364745C3796b",
				SignAddress:     "0x08210EbD783fF855E527F7BA544A46a22a51671e",
				ConfigPubKey:    "abf1c8e0f0773b15db359a92c576bebf06e813f423e676b3443fade0039ec448",
				OffChainPubKey:  "ae7dc0c2241b243dbcccdf29829d424512f1097260be10852233ed24333ca733",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "bcb99a4cfff0c03f9abf294d30626a1c2fda8b240fb2dd1ef805a95f2528501b",
			}, {
				Id:              5,
				TransmitAddress: "0x3c8f40B157Fd9a0d7f0Ac6c4746ECC47247AD574",
				SignAddress:     "0xAB5e62a82406acadDCb993ffDE011C481875B545",
				ConfigPubKey:    "113f050e5b4f01b77778be8df39a4df63c5aaa4454e2923980e4146337056768",
				OffChainPubKey:  "b2f7c019beafacf7a2dc919727c8e4dbb413baf583fb589fdbe8e5841dd751a4",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "cde13c3a60178fe950196a23a495cd957d164e4ec9d253f7316faac8c85ca242",
			}, {
				Id:              6,
				TransmitAddress: "0xE7b320CdcB0514CFc7fACEC7FDd32Eee6d9fd2A5",
				SignAddress:     "0x0C33319114742cD880186f0049776a9bF3A3F07E",
				ConfigPubKey:    "e4d73af7b7d3ee9437b9701afcc76312bb3898fd0cd72ce7278cbc550b434a10",
				OffChainPubKey:  "af1ac657b7dca63b6e02ac70533afe717dbb5bbeb211d7c8a074000ac06ee09b",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "d00399baa6e38e4c4bac63d7a64b3633018c7741f22b0dec1ca73c2effcca48f",
			}, {
				Id:              7,
				TransmitAddress: "0xeBa46834073CeDbAE9322C677A132c7a9AE911c8",
				SignAddress:     "0x81f3BbE224D1892eB2765BB64CDE7A65E0942D97",
				ConfigPubKey:    "32d6dc9f7e02118b0487ebbad562e1709ae20f9cc354d215235d29d80631ab76",
				OffChainPubKey:  "387fe43fea4efa0f4e44676ac367672ce6de481cf620b9714087e773914cfe05",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "df2b9e9e8470855adf427d5a58febc97bd093bcdcd168fd5393f597a396e288f",
			}, {
				Id:              8,
				TransmitAddress: "0xB742c3eda47205952420B314C5429664C8A8C406",
				SignAddress:     "0x46eBee3336203A7a4d2831F2c5ca2814843B116D",
				ConfigPubKey:    "ee08010e490cc83d2d6c2c24f25f004380b1ce40a3ef2cb8c1eb75296386d462",
				OffChainPubKey:  "edcf2c337b2468a0394b0dd3c49832bac6aafd08ea8b7f267a6daed981f8b701",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "d1005589a6f5bbcc9f8320a9feb5ef2826e5058497cb22a11a86457b2ea7fdb5",
			}, {
				Id:              9,
				TransmitAddress: "0x0f0967cE269cB73A27918F0786688FE09D7488dD",
				SignAddress:     "0x2B3fA83975EEA065A2dd1c4AAb3D9Def5144AeA3",
				ConfigPubKey:    "dd7b52b500039b444a05c6bb595d6cbc3cd19f918b57dbf218de3607aed04457",
				OffChainPubKey:  "8cb6917e007bc5b2d8fb86be58e602a3fb57c49e8e5b0af215613c98e2eaf9a6",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "4ebe42407aa7123118640ee9df53c3bc2941c54bc19e9d8356f1aba2c0aa9edc",
			},
		}
		nodeConfigs[ModeUsdc] = nodeConfigsModeUsdc
	}

	{
		nodeConfigsModeMode := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x625540C630ad7582FE50c03E5faf5A47616A5bf8",
				SignAddress:     "0x14346F3D5d7A6cB6706AE683929b8C9881cae43F",
				ConfigPubKey:    "88c0004249b53a5e7e90b79719911a637ae19ecbc15d38edf32a26d849fbe722",
				OffChainPubKey:  "c40a98a028f8638f4e3fd429c84ebf554cef56c0a7000a11cfd122d275a8af18",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "2346cfb32f547ee899343d866ebeb4f52dfbc8dea604f7d8616beca4cb234abf",
			}, {
				Id:              2,
				TransmitAddress: "0xBf732AA2eA42515D2ca624A003aE19d1a220d75A",
				SignAddress:     "0x8B7D5613c1e07b5b11220904F793A7F3B851B9a2",
				ConfigPubKey:    "d4c2d312f2f2f50537756268e5e039ff77b3bbac7a66294bd21c7b1b41ea8351",
				OffChainPubKey:  "8289aba0dc12c6626e536402ee53318199934b07b387b8e23ee88924667dff17",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "cc84832ac6d47ffc9e9b755c82786d2af9f1e2d7f5b3df84ddbdfb692d20c7b9",
			}, {
				Id:              3,
				TransmitAddress: "0xCeB7Fe944093a683F4A885C3DDfF774C5ee48C32",
				SignAddress:     "0x1be8150f6B4Aac449D567fc4D3148D68E4233a8e",
				ConfigPubKey:    "d7ed6b3deeb3a4998b150c0cfe22557d9139fdc592ff29ade8d1d3f8696d2f59",
				OffChainPubKey:  "69d653055ae9f193fee21d078ff0892d968c3ae20bffd587ec82f3c5a0354f93",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "58226ba28ef0831d7223c24ae9d55f6a696e31d595f7a8a6cd222b7bb82a3631",
			}, {
				Id:              4,
				TransmitAddress: "0x4C6aaAd9Dd9Fe03F72C261B0179F364745C3796b",
				SignAddress:     "0xd9B6575dEb43bA72e729f336dfd409c398fE3f23",
				ConfigPubKey:    "a674a87e97b467e347a0231a6f16b514c60390422b181ab8fde627eb2270bc77",
				OffChainPubKey:  "7bd4c380ee98c4eba5c79b3adaee032f3b134e458850c048fb25d4e3c09e4cec",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "4a13af619309b601645086c8885f970c22293bc584dd0e26ab563f30c94c3a19",
			}, {
				Id:              5,
				TransmitAddress: "0x3c8f40B157Fd9a0d7f0Ac6c4746ECC47247AD574",
				SignAddress:     "0x345270380981e7CD48Ca394e5d800573655cB8bc",
				ConfigPubKey:    "9eb36a5dd2460178d224be69c804a88373a51981cae4e7dfb38bf0765c554c2e",
				OffChainPubKey:  "bc7e71accc58cbb73e026dfb6c41164980daeede824af6223d2f5e547d2b6a5c",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "a8813a57b0885cf1ad6aa8b35422b97b594c53be741d4b87972b437e2d44729f",
			}, {
				Id:              6,
				TransmitAddress: "0xE7b320CdcB0514CFc7fACEC7FDd32Eee6d9fd2A5",
				SignAddress:     "0xb27d48D03cfFfdE8e3D4c63Cc79A746Eb34D7F80",
				ConfigPubKey:    "29d3a1b0ffb8cca8b6dac98853ee9b12116747e72f07d36e9485105766261000",
				OffChainPubKey:  "06ea88d78b842a07a931415923c0b24f1ba9901f7171615bc172c99316ffb0c6",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "a57363e5fc984fc7c7ccf1716f70e3f8005fb5693e35072696780e30ec7705bd",
			}, {
				Id:              7,
				TransmitAddress: "0xeBa46834073CeDbAE9322C677A132c7a9AE911c8",
				SignAddress:     "0x8bE97D2f6a99Cf04fC385B1117D9dD36adbDe7b3",
				ConfigPubKey:    "9fefcc2a17c574062fb546a5df13f7d85389dde0f6d6d2fdc506e03dfedacc5e",
				OffChainPubKey:  "26029bea7bcec6db57eec9bd714348ac8fef6d6126e90686aa62b01bb331e88a",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "de9014bc1c4d14eb5484b919c58dca7f0e5b75fe0a0cfbcf513ad1ae65859d3e",
			}, {
				Id:              8,
				TransmitAddress: "0xB742c3eda47205952420B314C5429664C8A8C406",
				SignAddress:     "0x53fBf366928067a5c8d33F56F880DFDc38D191E2",
				ConfigPubKey:    "986f1bde619daa2e617a1dde24783a5ad1d67f24d7a85d41f36c8fcf8dcd9026",
				OffChainPubKey:  "afc3f67c5ea6ce80a73a7e8f8e2ceaac6737072bb1be9314803cffdbbc909ef3",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "285d6f592ec3ee935cc01607524fa4e92bda88368d77b87fece17337d6359242",
			}, {
				Id:              9,
				TransmitAddress: "0x0f0967cE269cB73A27918F0786688FE09D7488dD",
				SignAddress:     "0x90D788256eea8ff146dCeD92B8CEDa7eF8Ae658c",
				ConfigPubKey:    "44a11acbce47eb59f5f7e3090df5de0238925db9d13fc7b2fdac9afca24cb009",
				OffChainPubKey:  "71fff48bc811da5a4f27b643b5bf09e507fe04a4544741d913215a0d6eec5316",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "ddc6c4ce12600ae7f6c7166decb4bdb9a9a1f99114ad1655d160e06e21cc6dcc",
			},
		}
		nodeConfigs[ModeMode] = nodeConfigsModeMode
	}

	{
		nodeConfigsModeStone := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x625540C630ad7582FE50c03E5faf5A47616A5bf8",
				SignAddress:     "0x0272B1301b256073022833c67b8485ECAB07c69c",
				ConfigPubKey:    "6f0e7ae1bb8008a84e5bb4e7f0592acaf09aa6edca55e7628a091cbf7dc7a076",
				OffChainPubKey:  "cf5e9b53452905178ae2d2dfd33f1e6feaaa5c3f444e390d6a53d832c455314c",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "636ec29a89054e1635a6da9b105f63f66861193be3bc7a129b7b2d38f8b6f694",
			}, {
				Id:              2,
				TransmitAddress: "0xBf732AA2eA42515D2ca624A003aE19d1a220d75A",
				SignAddress:     "0x7E41F2C2963E10444f85fF8fAE0f59869852a9e7",
				ConfigPubKey:    "3cd116e4ba0d787388cfa9fdd7df86013e2bde2dec3c7a98bb583e6e510bcf23",
				OffChainPubKey:  "3b8a2b11937409c4d06763d7b1d5cc86553ff611f40ad996b5fd20a9cf784d1b",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "fc09e3ad1c5c0a185e4a1d6f6d3b923db102a6ccd3cfc9ff3f06fcac855cc6cd",
			}, {
				Id:              3,
				TransmitAddress: "0xCeB7Fe944093a683F4A885C3DDfF774C5ee48C32",
				SignAddress:     "0x545D168D0DD253314655D998208125c0412286b5",
				ConfigPubKey:    "b715e3ab9dcc7b78fe2e5586a93dbf7b68c4259d2105cc9803fc0d6044e9a675",
				OffChainPubKey:  "beeadb81ee20ea6b9eb08a87f2fe17c61d302e5dcf4a94a60d3339fe5f847952",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "de45747971646ae79b3313d15c32d80f053e65d6b6826e52881c845f9e27b2ed",
			}, {
				Id:              4,
				TransmitAddress: "0x4C6aaAd9Dd9Fe03F72C261B0179F364745C3796b",
				SignAddress:     "0x80C92726186C7841606774D758a0B4c0Ca1d495A",
				ConfigPubKey:    "e901859000afecec80f115165e9fdd772bc8476a870ed0deca78e64eeddf452e",
				OffChainPubKey:  "a35fcd75f5d5573ce33076e41d23fb4eb63af18dc6f64ed2f9fee068d7251655",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "6466a0969f282feaf030b493deaf3da1f063fb9b6ed7ae5c2d9dc5dd679b2f30",
			}, {
				Id:              5,
				TransmitAddress: "0x3c8f40B157Fd9a0d7f0Ac6c4746ECC47247AD574",
				SignAddress:     "0xD903f2Dd7945510D83E94236657F4be30e65eA24",
				ConfigPubKey:    "00ccadfb0229f8592404e99d85fac7ee9d0911edef5611098ed3dee41d853d4b",
				OffChainPubKey:  "f3e4bd9ad4f69a1e9f89b59e3e9c36bbbaf5252a23a3a44dd2b288bcd63fe7b9",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "29931cfa7b3a386356abeea7f0efed7f462f14e955e0665f1d9c115cca9da63d",
			}, {
				Id:              6,
				TransmitAddress: "0xE7b320CdcB0514CFc7fACEC7FDd32Eee6d9fd2A5",
				SignAddress:     "0xe8338709c82D489A54A8223B98fF5775f1846799",
				ConfigPubKey:    "30ad7b1296e0bcb41c844b650a644ede92f68481d9ae7c9d2f005ad6445a767c",
				OffChainPubKey:  "50c42d82510924039863e773c85081e835c8d6f8d5a1ece675a961df447ee66e",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "3cdec8826cbcf032d0737b4d85efb35d795214a4165b206b7f7eb1c5bfe54718",
			}, {
				Id:              7,
				TransmitAddress: "0xeBa46834073CeDbAE9322C677A132c7a9AE911c8",
				SignAddress:     "0x4d6d3a18b73d6219d6fca1f7049011fb3B1170EC",
				ConfigPubKey:    "dcda618538e0815ab7781a1402b2ef4fe24925a514c7a7312d4b9d1bbcb54b1f",
				OffChainPubKey:  "2cfc9b601e6701d878b955806f5701b87c90469521a54269b8f6ef1ce21132ec",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "c7dd896ab99d0bb0e0522fcefb96dddd69c54df3828b7eaf3be87f9f9d7b4e88",
			}, {
				Id:              8,
				TransmitAddress: "0xB742c3eda47205952420B314C5429664C8A8C406",
				SignAddress:     "0x23Eca224eC0F5A4915083F18088d7f7716f3c9fd",
				ConfigPubKey:    "1823c408e2758a22226013912858b42116a515c412661201a9a6b55ed2e39b4c",
				OffChainPubKey:  "7039b83b6407fccd427a9b0b08f7ddd397ef014f6bf2fc823051dd5f1d1864d3",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "eee236b77fd0b63eedb8b4768864251b0274d91a4e7a594bf0f20b3c77a6aca3",
			}, {
				Id:              9,
				TransmitAddress: "0x0f0967cE269cB73A27918F0786688FE09D7488dD",
				SignAddress:     "0x6CA121A6B3B06e17A9CF2debd2edF2e79F0abc07",
				ConfigPubKey:    "27f5f6e4188f1fcc7a2540a92b086621d5799a13111d4ca3c38a4315501f7842",
				OffChainPubKey:  "99ea86c00744db3df6324e6b925a5084a92a1aa0f31ca9482e2ba56ad69c9dfe",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "f334182269254a9fbec38fedfb1b44f3fc2963aa3281a43922bf4245e43b176c",
			},
		}
		nodeConfigs[ModeStone] = nodeConfigsModeStone
	}

	{
		nodeConfigsModeBBUSD := []test.NodeOCRConfig{
			{
				Id:              1,
				TransmitAddress: "0x625540C630ad7582FE50c03E5faf5A47616A5bf8",
				SignAddress:     "0x65B4ceB68C9bC92D8FC8c71eBD8f7e4e4E0A84C6",
				ConfigPubKey:    "2c43a665fa904839d16b06cefa6935b45437cb885d0f6d0f187fcb06b641801f",
				OffChainPubKey:  "b7f25e6be8ed5e3f3ce4403e8d2e28e20ad54d631a2e7e176cba42fe90378e91",
				PeerID:          "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:   "f6d4fa062fa6ff857d47cf5c8ed8caa38fa47d3de94204d89a75cee14c8672b8",
			}, {
				Id:              2,
				TransmitAddress: "0xBf732AA2eA42515D2ca624A003aE19d1a220d75A",
				SignAddress:     "0xEE1ed271AF5eC39Ce8e98535e04C5137dDEbe129",
				ConfigPubKey:    "a3bdb0f2aca774614d298b1cabaa0b7323b4dcb6e3bbbb51709c975464b4ab06",
				OffChainPubKey:  "427f5aeace313ad8a9b3e55cdaca219dee31a6430005afe0c085907d52f5f3ee",
				PeerID:          "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:   "9b360afcb22e1050d77713fa19fb14f08b0920cd999280972ae633d08ebb223b",
			}, {
				Id:              3,
				TransmitAddress: "0xCeB7Fe944093a683F4A885C3DDfF774C5ee48C32",
				SignAddress:     "0x2Dc1101deE0b25785dC757bF7D48ca9318A969b5",
				ConfigPubKey:    "887d72633b6b1b727a80f97c100399499cb8591a430c20285d16cf8cd99f1376",
				OffChainPubKey:  "ce3382b2f40c0ca34d23d4f0c885b3093b622f630c54b792838036e6b1086a1f",
				PeerID:          "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:   "71d4f7d428c5fb7293294c76f9613f3da23df6b9300b905012b588cbefbb2286",
			}, {
				Id:              4,
				TransmitAddress: "0x4C6aaAd9Dd9Fe03F72C261B0179F364745C3796b",
				SignAddress:     "0x056DB72630E86c7d884b9cEd618659c7a47fcBfA",
				ConfigPubKey:    "7de2bbc0c8535a19657c114826b34aed9d46a02ff10c32f410009c0c10cc6640",
				OffChainPubKey:  "c5b041bedcbdc3b4f4bb4a0e480a5912e842b1384a0eb61721d627e079294f76",
				PeerID:          "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:   "7b3f00281fc775133e67085cf8ae57d4022f6c4132dbc2a8a76c70fef6f330b2",
			}, {
				Id:              5,
				TransmitAddress: "0x3c8f40B157Fd9a0d7f0Ac6c4746ECC47247AD574",
				SignAddress:     "0x19c5fB92a03126E92fc716d07302fa727625f835",
				ConfigPubKey:    "074a26c34d36afd64c248ab9283e66a329627422c9f3a1b59038b47ec75da94c",
				OffChainPubKey:  "d048f51f078f865b5ab9caf8cb72fd0beb35ab490959320b162baf6eba451e21",
				PeerID:          "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:   "03dacf5e9b9024ad1e52596e2c98e222f4176bb827389ca6d4c0a2c5b2e3dbbe",
			}, {
				Id:              6,
				TransmitAddress: "0xE7b320CdcB0514CFc7fACEC7FDd32Eee6d9fd2A5",
				SignAddress:     "0x452465d86D6B4e14bBe80a226729a5479B811B68",
				ConfigPubKey:    "b2e4dd81e89b6a92cf8f0d823c039c90a10d8b1e8818ce34db187251e3d0f021",
				OffChainPubKey:  "abc93209e53d47ec577a96fc0094b7d530ab4535ed20fdb78a48cbc0856f6ee9",
				PeerID:          "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:   "862b211f823452987a01e413a5182b5332173cb1d7ba8f922073fc721a848b3a",
			}, {
				Id:              7,
				TransmitAddress: "0xeBa46834073CeDbAE9322C677A132c7a9AE911c8",
				SignAddress:     "0x045CdEbD89D5E1ca8251c97A7D66Ca3b2A0847ce",
				ConfigPubKey:    "148e96a63d455c65123a692ac7c126e938db5ccc71b882cc1677efa59d5ea850",
				OffChainPubKey:  "efad1d598ad1646f4eac0b15f33b48d4c9d886c7e829740cc509d9e1bf65d8f2",
				PeerID:          "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:   "a7274ce3be1a404efcc7b3e17003b007c4ddf6a0424958491368c239dfc06d46",
			}, {
				Id:              8,
				TransmitAddress: "0xB742c3eda47205952420B314C5429664C8A8C406",
				SignAddress:     "0x74f0e1Fd2396a36916529ad69927AAfF33434C02",
				ConfigPubKey:    "6f46f7f5f086a07c2aa57b066ba2b8779a9e47b684548007c5a41a653b69181f",
				OffChainPubKey:  "ff93bec24003681cbdad1af4298e024ae636c920f608b8f8a13ff64273e2f00e",
				PeerID:          "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:   "366e09697ac6d42b4135e1f465db9cec3e115da45ef26643ab9b77dc33fe43b8",
			}, {
				Id:              9,
				TransmitAddress: "0x0f0967cE269cB73A27918F0786688FE09D7488dD",
				SignAddress:     "0x0983B79424640b34110D0328d89735aa525b3700",
				ConfigPubKey:    "8cde5bc556ae50c6834c5528590f196894c1623c70f1e6ed4859cfcf57e7ca2d",
				OffChainPubKey:  "aaf78b1e90aa519a80d054a09951d065d6a62b7855cfd62a3954b50bee075679",
				PeerID:          "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:   "e222011caed972890c22783c04f36e114b3adf433cfda5706d021c1a0e28c4f7",
			},
		}
		nodeConfigs[ModeBBUSD] = nodeConfigsModeBBUSD
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
	ocrConfig := GetOffChainAggregatorConfig(ModeBBUSD)
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
