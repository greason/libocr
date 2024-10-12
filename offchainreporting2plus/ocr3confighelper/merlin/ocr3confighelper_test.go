package ocr3confighelper

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/confighelper"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/internal/config/ocr3config"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3confighelper"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strings"
	"testing"
	"time"
)

var (
	// 2 ** 176 -1
	maxValue, _      = big.NewInt(1).SetString("95780971304118053647396689196894323976171195136475135", 10)
	rawOnchainConfig = ocr3confighelper.OnchainConfig{
		Min: big.NewInt(1),
		Max: maxValue,
	}
	rawReportingPluginConfig = ocr3confighelper.OffchainConfig{
		ExpirationWindow: 86400, //
		BaseUSDFee:       decimal.NewFromFloat32(0.3),
		//BaseUSDFee:       decimal.NewFromFloat32(0.000000000000000001),
	}
)

func AproOCR2MercuryConfig(numberNodes int, target int) ocr3confighelper.PublicConfig {
	if numberNodes <= 3 {
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

	var DeltaProgress = time.Second * 2
	var DeltaRound = time.Millisecond * 250

	switch target {
	case MercuryBtc:
		//DeltaProgress = time.Second * 17
		//DeltaRound = time.Second * 15

		//DeltaProgress = time.Second * 35
		//DeltaRound = time.Second * 30

		DeltaProgress = time.Second * 17
		DeltaRound = time.Second * 15
	case MercuryEth:
		DeltaProgress = time.Second * 17
		DeltaRound = time.Second * 15
	case MercuryValueless:
		DeltaProgress = time.Second * 17
		DeltaRound = time.Second * 15
	}

	// 参考：offchainreporting2plus/internal/config/ocr3config
	return ocr3confighelper.PublicConfig{
		DeltaProgress:               DeltaProgress,
		DeltaResend:                 time.Second * 10,
		DeltaInitial:                time.Millisecond * 400,
		DeltaRound:                  DeltaRound,
		DeltaGrace:                  time.Second * 0,
		DeltaCertifiedCommitRequest: time.Second * 1,
		DeltaStage:                  time.Second * 0,
		RMax:                        25,
		S:                           s,
		F:                           (numberNodes - 1) / 3,
		OracleIdentities:            []confighelper.OracleIdentity{},

		MaxDurationQuery:                        time.Millisecond * 50,
		MaxDurationObservation:                  time.Millisecond * 250,
		MaxDurationShouldAcceptAttestedReport:   time.Millisecond * 50,
		MaxDurationShouldTransmitAcceptedReport: time.Millisecond * 50,

		ReportingPluginConfig: []byte{},
		OnchainConfig:         []byte{},
		ConfigDigest:          types.ConfigDigest{},
	}
}

const (
	MercuryBtc = iota
	MercuryEth
	MercuryValueless
)

func GetNodeConfigs(target int) []ocr3confighelper.NodeOCR2MercuryConfig {
	nodeConfigs := make(map[int][]ocr3confighelper.NodeOCR2MercuryConfig)

	{
		nodeConfigsMercuryBtc := []ocr3confighelper.NodeOCR2MercuryConfig{
			{
				Id:                1,
				ConfigPublicKey:   "ocr2cfg_evm_c1216b09e05e6714242205a1a54e66567ccf0506fafb38a392f12d12079f4e47",
				OnchainPublicKey:  "ocr2on_evm_42cce0571dbb305033b2b966bd18cd0eefc9e76f",
				OffchainPublicKey: "ocr2off_evm_ae671f6952cb1b40dfa27259db054ef35ee1bb90275ddd35305d2b2811dfe186",
				PeerID:            "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:     "762b7e8abf53cb26abe910cf927928cecfe876b4883016f8ac9d753c3be8da82",
				CSAPublicKey:      "csa_1ef812fe7ac6d4c200a6bc1394ad72eb1055892b5ddb2a47bdcda8dbae355f3e",
			}, {
				Id:                2,
				ConfigPublicKey:   "ocr2cfg_evm_6fd8d6bcccd83ebe0c7cf8f741e24226916016de63788d9cccbdb4d67c1e9c57",
				OnchainPublicKey:  "ocr2on_evm_76b4f1b612f479c3aa0ad3d01cd693ac41bcc38a",
				OffchainPublicKey: "ocr2off_evm_8067bb12e659178a148a560807034e7c8c8eab6c73e77f39d2270c107fda8e2b",
				PeerID:            "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:     "038e0c9f664cb8609ada5df80d69d5b5ccf8d39b52ce670bbebc6caa5691014b",
				CSAPublicKey:      "csa_a75f8bf55c971c0f3b807eb6173bc53830f68f3c246babb8c4112cc64ada5b83",
			}, {
				Id:                3,
				ConfigPublicKey:   "ocr2cfg_evm_f98417b26a7f78482b4b6572804bca639735c8dbcc24e203dccf5bf223a14104",
				OnchainPublicKey:  "ocr2on_evm_6c026d6c2c97d003c9c37f8eded24f59ab2bb225",
				OffchainPublicKey: "ocr2off_evm_936f857992d927d7840963e696b46376df3709dc48e93852cc41e7b3207a238c",
				PeerID:            "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:     "524fc23dcbd1213404aa761086d8bc17b6fca2c4cdd8e6b8fca527873a819d51",
				CSAPublicKey:      "csa_7c936de89983a8d5c9e7cdc596c70869882c318c05cca3a5fca5b8055564cdb7",
			},
			{
				Id:                4,
				ConfigPublicKey:   "ocr2cfg_evm_26885fcdffd37f7b1b812de5af405e286d5d9df2836d9ab5b333c83eaddf6061",
				OnchainPublicKey:  "ocr2on_evm_dcde9bb357de3a712cb1362d48812f2314810be8",
				OffchainPublicKey: "ocr2off_evm_7f3c4024cc75e1d5de25b8f1af38810ce6bc42f4691a562a5afa7845064c4764",
				PeerID:            "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:     "4e0af41c53ffe3a4eb2732a6df2b74034152813426d925420f44a9e3d8bd9c68",
				CSAPublicKey:      "csa_5a3c54a8459b1906392d3def08051b4cb8c64db3a82fbcb285fe7cc226df3a9d",
			}, {
				Id:                5,
				ConfigPublicKey:   "ocr2cfg_evm_9a3efde39de5546b70422de9ba8ae2accbf673e0a5d4ae92fdcb1525eace732b",
				OnchainPublicKey:  "ocr2on_evm_fde7894e0ad3912840300b870a6e88323e911c1c",
				OffchainPublicKey: "ocr2off_evm_6bdaadb6b7aa578e83a9890d25c8c02b92f8fa951ee27ce680d14e4d10d9c78c",
				PeerID:            "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:     "e0c57d14a2ade1dc17fbb2cafafd094c9043d93ba617452f91f2ae90512c6a00",
				CSAPublicKey:      "csa_554bae5fb8c788fa6743fc2f1934d3304d0ab4651aa1d5c94d4cce21569ba907",
			},
			{
				Id:                6,
				ConfigPublicKey:   "ocr2cfg_evm_997791d80f90bb02f2cb49084a4e213f38d2c4c51c675d01b9e08ec3e35efc4f",
				OnchainPublicKey:  "ocr2on_evm_160e74ab0c562f8f8b15ddd29dcb2e750c1e0976",
				OffchainPublicKey: "ocr2off_evm_2e8e680ac62fa80e0002ad6ac201e48681229ee59be6628ee2153fb3189126f4",
				PeerID:            "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:     "3e7cb85da0f272286ba8f13a3a788df10070257e0b128cf0b5897d5c42cddec5",
				CSAPublicKey:      "csa_e7a7d61faf6584bbbe7137d79f59cc07ad4b89b7f6cd41742a5095e58616cd9c",
			}, {
				Id:                7,
				ConfigPublicKey:   "ocr2cfg_evm_8a8bdf64d16d977cc620390fc85e2daf1ec7794b93b719edc69b8aab20384976",
				OnchainPublicKey:  "ocr2on_evm_23ff6e59a014dbbc7b184d0e928c4e7135b2a39e",
				OffchainPublicKey: "ocr2off_evm_6046bc50a955d1f2c3d0a9a29d26d49cdf48885844187756214d554c6fb386a5",
				PeerID:            "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:     "0c71fed7e2a9f45a24521c0870f151c03b1482af36adbb1b3b82a1bea404ccf3",
				CSAPublicKey:      "csa_b19ab5aebd07a73a5f2479f4795603b46216ed2d37d5cfe111b84ad975508df6",
			}, {
				Id:                8,
				ConfigPublicKey:   "ocr2cfg_evm_6172d7772cc2f259c9dcd62ee905690a9c50afaa40cef44bd95297ac8bb4bd1e",
				OnchainPublicKey:  "ocr2on_evm_98e1c6db6ad68a3a7efd81f6ee4a4b0503ea4a38",
				OffchainPublicKey: "ocr2off_evm_dfdd443d1eba96e558886ba0fdbb96a8e2188066463e83641273d3827ea1b0d5",
				PeerID:            "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:     "da2a3ab0df550bccf61162d91b15239ac7be52a95b4cc829455f49e477387c05",
				CSAPublicKey:      "csa_ecfad4b98a535fdc2c17d66373c4b92ff783f8ca039c3e19ea4bcd53259c18df",
			}, {
				Id:                9,
				ConfigPublicKey:   "ocr2cfg_evm_ae4494c68661d173baa3d859cbe87c75e042cc81ae65afb22fa6e331441ba17c",
				OnchainPublicKey:  "ocr2on_evm_6b12cda9a427c1e4bb1b728b20db8d7dd79a444d",
				OffchainPublicKey: "ocr2off_evm_84ed94761b03acefa33e8bc2a1b0de40f3e3343c27e99067ae0e8374b0126ac8",
				PeerID:            "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:     "1deab20c608f6c76e9b6023c86f73b330fdd9870f49324c1a16aae7f5d83be0f",
				CSAPublicKey:      "csa_fbc297e151feb5a1ed2a3ddf36fd359ed69290039f2daaa50f12e21a4e7850b5",
			},
		}
		nodeConfigs[MercuryBtc] = nodeConfigsMercuryBtc
	}

	{
		nodeConfigsMercuryEth := []ocr3confighelper.NodeOCR2MercuryConfig{
			{
				Id:                1,
				ConfigPublicKey:   "ocr2cfg_evm_46c204575c793f58f9a8e924a419623052310b198bcab856f100e30200aeaf69",
				OnchainPublicKey:  "ocr2on_evm_72a0e78c46fe571f15c88305fe10bf8ffc8188a4",
				OffchainPublicKey: "ocr2off_evm_100f4785d89a44b31752c7c6626d8253252abd257b8752e15481dca694bae2cb",
				PeerID:            "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:     "c4fd015478f11021ee3b9c140eac02310e32f54725fe01cf05cc3f0abe8f8b2e",
				CSAPublicKey:      "csa_1ef812fe7ac6d4c200a6bc1394ad72eb1055892b5ddb2a47bdcda8dbae355f3e",
			}, {
				Id:                2,
				ConfigPublicKey:   "ocr2cfg_evm_ab046a9bd3a7ca629d89600c869a07ab1bf541b6a3f7a48909352efb8de79a50",
				OnchainPublicKey:  "ocr2on_evm_c219dd715d800cdbcf6067efaae1dab440e97235",
				OffchainPublicKey: "ocr2off_evm_872f7e9d6550026823b95867e0bac94eca9b1d6beab5678a31e023150db783f0",
				PeerID:            "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:     "0f0db3cef7065b25e8455cb789d4fbea245c02d76b40eda06a7f87941a353a5d",
				CSAPublicKey:      "csa_a75f8bf55c971c0f3b807eb6173bc53830f68f3c246babb8c4112cc64ada5b83",
			}, {
				Id:                3,
				ConfigPublicKey:   "ocr2cfg_evm_cebbc606269adc6981852f09a06f8c0002ff433f036c933433da69686a2f9219",
				OnchainPublicKey:  "ocr2on_evm_f6d0eb2ea8206dfbb1a18124050272b283b1ab4a",
				OffchainPublicKey: "ocr2off_evm_e2dbb768763c1483e931427a8d6c1511fdb384d898d4d5810d3be6caddeab703",
				PeerID:            "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:     "c3eaa0d3e1133ee65e29086d47b346b8f4c8eda0e5a27b2fe80aab3d8597d648",
				CSAPublicKey:      "csa_7c936de89983a8d5c9e7cdc596c70869882c318c05cca3a5fca5b8055564cdb7",
			},
			{
				Id:                4,
				ConfigPublicKey:   "ocr2cfg_evm_4aac3fa00acc88ddf61f9f5f53a1ede57553d9f6e6dae7e8b308f34e0bdc2a02",
				OnchainPublicKey:  "ocr2on_evm_226946f5d647762d04bdf80425edd5dab050abab",
				OffchainPublicKey: "ocr2off_evm_b6f31803ed0adea66f8221eef84cb44e937014b440e752cdc700cb8236b20d26",
				PeerID:            "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:     "cc7e1ea8ea786be3c06d10d5ae382a17f65ea5ca3f8b3278106a888a8b590d9d",
				CSAPublicKey:      "csa_5a3c54a8459b1906392d3def08051b4cb8c64db3a82fbcb285fe7cc226df3a9d",
			}, {
				Id:                5,
				ConfigPublicKey:   "ocr2cfg_evm_1e89ba8ec467aaf5b7bfc13ef4b5afe151295ede426b3ee2cd4b119900aecf1b",
				OnchainPublicKey:  "ocr2on_evm_59867592886ab6a6f79558104251005f80c9712b",
				OffchainPublicKey: "ocr2off_evm_53a4cd74159f3a552c031aaee814cd4aea4a7eedcd830dfef78c28d84b56d3af",
				PeerID:            "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:     "2e48df6dddea173fb993c9ca9b15835a47b19897cb761350f9cb79c0b50bfe07",
				CSAPublicKey:      "csa_554bae5fb8c788fa6743fc2f1934d3304d0ab4651aa1d5c94d4cce21569ba907",
			},
			{
				Id:                6,
				ConfigPublicKey:   "ocr2cfg_evm_8e180a15af72f89c2337d8f24463511ef36e5d6d56d4443743942c2122840161",
				OnchainPublicKey:  "ocr2on_evm_8856222c5294359d356b4f2637ee964ab0260cf6",
				OffchainPublicKey: "ocr2off_evm_a1824ea42aa5d0c1b3c755087d95d077b6a3eebdf536da5cd8897a0235487a00",
				PeerID:            "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:     "197d7044c43ce19158a45def53cc9a6c75a41d87d212cdbd0468e787243e1b05",
				CSAPublicKey:      "csa_e7a7d61faf6584bbbe7137d79f59cc07ad4b89b7f6cd41742a5095e58616cd9c",
			}, {
				Id:                7,
				ConfigPublicKey:   "ocr2cfg_evm_177a776183122a531616ed6ccf28038925439f1335daae9d33b2884b61c98637",
				OnchainPublicKey:  "ocr2on_evm_9f3afc9549d96087777039ebb5be173cbfb257f4",
				OffchainPublicKey: "ocr2off_evm_e6e0b86e602b6136d7f58d382b564c4b2b665965aaa8fb08bfee0c06c27b5210",
				PeerID:            "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:     "1d24713c51bd583cc1ae71e303f38aa307e6e1aac31e7df365dd649b190566d3",
				CSAPublicKey:      "csa_b19ab5aebd07a73a5f2479f4795603b46216ed2d37d5cfe111b84ad975508df6",
			}, {
				Id:                8,
				ConfigPublicKey:   "ocr2cfg_evm_47cb1943fa801c138f7a216c3d4c80f8e90fc24c3f820fb666b1bdf8dccd7c41",
				OnchainPublicKey:  "ocr2on_evm_044f3760c2b8596a2b45c24d578f8c7008117dd4",
				OffchainPublicKey: "ocr2off_evm_6feef5a4bce265c206aa86712e68a60541a055fed365fb3fc502bcda70e2a725",
				PeerID:            "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:     "dff2d84a0e2c7f5232fbbbccbff46339c3242c67e3789c2ae1dab39553dce780",
				CSAPublicKey:      "csa_ecfad4b98a535fdc2c17d66373c4b92ff783f8ca039c3e19ea4bcd53259c18df",
			}, {
				Id:                9,
				ConfigPublicKey:   "ocr2cfg_evm_79ee4e2c65c8b9597b7ce34142b7a37a657513a4e60bce801e54a98749be443b",
				OnchainPublicKey:  "ocr2on_evm_15ec0227125c86d19907532549a2eab182ea7916",
				OffchainPublicKey: "ocr2off_evm_ed6c10bd50774db58d8eacfd18030b3c0e8bd55ccf789a486937040107ec17eb",
				PeerID:            "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:     "a170984d0d3f323fdd18ea0ddc1220ee55a3329b0451649afeb49ec001923a80",
				CSAPublicKey:      "csa_fbc297e151feb5a1ed2a3ddf36fd359ed69290039f2daaa50f12e21a4e7850b5",
			},
		}
		nodeConfigs[MercuryEth] = nodeConfigsMercuryEth
	}

	{
		nodeConfigsMercuryValueless := []ocr3confighelper.NodeOCR2MercuryConfig{
			{
				Id:                1,
				ConfigPublicKey:   "ocr2cfg_evm_755fec4c2fe6757915904d04a99a366141b4b8c868346fee3f6c782e0a877b76",
				OnchainPublicKey:  "ocr2on_evm_83f69590bccaf6c819a95f5cf11b1458fac4478f",
				OffchainPublicKey: "ocr2off_evm_2b8dbf537af1178ec25cc97877d51f56fd447384344e4504e3d718439ac8886f",
				PeerID:            "12D3KooWLioyKDAmRUbggzL6uGBfrnboNtFnYpGydcyzJURFU9mg",
				OffChainKeyId:     "0cad3ed07f326dd4fc7ef42788404307ca60917a5c17b83ddab779bc976fe991",
				CSAPublicKey:      "csa_1ef812fe7ac6d4c200a6bc1394ad72eb1055892b5ddb2a47bdcda8dbae355f3e",
			}, {
				Id:                2,
				ConfigPublicKey:   "ocr2cfg_evm_ab930235963174a223caa923b61a8321981a289975cf3da5b4666cc1994d416f",
				OnchainPublicKey:  "ocr2on_evm_e399999eef6b208780782e673cf8f1df774170b0",
				OffchainPublicKey: "ocr2off_evm_7392efb9b039fac27e4bd2d5deb5aed07394327cd03abd33a831e271e18a82db",
				PeerID:            "12D3KooWHpWYpFFryEoVAngkwtZS626SFzdawSrFNn5jpETx8PSb",
				OffChainKeyId:     "30372a2e168ab1c6f74294d778ceb59f1aa04f129a5225ead6bf96551bae7724",
				CSAPublicKey:      "csa_a75f8bf55c971c0f3b807eb6173bc53830f68f3c246babb8c4112cc64ada5b83",
			}, {
				Id:                3,
				ConfigPublicKey:   "ocr2cfg_evm_8fbf5a00357500fcdd8329701c4fd445250913cd1b094cb1b0974ff583579e4d",
				OnchainPublicKey:  "ocr2on_evm_643089ce5400868b9966efce1ac6ea876ce5246c",
				OffchainPublicKey: "ocr2off_evm_0414a83f399bab82a9683e832872a1c5930e62cb44e5abeb1ef0cacecef13c28",
				PeerID:            "12D3KooWNxnMSWMRRoUEZMq5Ta1KxDh8BYSmbtERNb2VGbmGgwqL",
				OffChainKeyId:     "da41a3db12000ad1506303a182f1f6af73e188537b1d99246bbab496265ef767",
				CSAPublicKey:      "csa_7c936de89983a8d5c9e7cdc596c70869882c318c05cca3a5fca5b8055564cdb7",
			},
			{
				Id:                4,
				ConfigPublicKey:   "ocr2cfg_evm_16773406d2a507269e8cc6b5d55cda915f0da0009a6b766c0e9e558cf10b1d37",
				OnchainPublicKey:  "ocr2on_evm_3b481def3922d9d78143ccce598408cee3dded19",
				OffchainPublicKey: "ocr2off_evm_157c9e8a36b98f0acada2f40a796623842b785f5b694c732c5903b7fdef1a5b3",
				PeerID:            "12D3KooWDVFrNKvoTFxzXQahDwxrZ9ArctBSsMpHNx8avnc58YiS",
				OffChainKeyId:     "e833391277dee7f9a564b8539fe169fdd135b05adae3b44c4cf416ed2826b9a8",
				CSAPublicKey:      "csa_5a3c54a8459b1906392d3def08051b4cb8c64db3a82fbcb285fe7cc226df3a9d",
			}, {
				Id:                5,
				ConfigPublicKey:   "ocr2cfg_evm_7d3a3dafb8f9895aa13e641f58ee022d897e9f8b3024ebaf93e5ab93c6d71a14",
				OnchainPublicKey:  "ocr2on_evm_385a9099ef0d17b08524050ccaaed45ec8dd6614",
				OffchainPublicKey: "ocr2off_evm_311308595b48531e58eae0a5201ddd3a2406951d827ef5599bb6cacc8d85f828",
				PeerID:            "12D3KooWRaSynVTYKdo7Nky3i77dN3D3W4rw7dBY7jG7E22rJzxt",
				OffChainKeyId:     "14f317529381a6172630ba5e553b449ea392d5f52454535ef35ba1dc5638ca07",
				CSAPublicKey:      "csa_554bae5fb8c788fa6743fc2f1934d3304d0ab4651aa1d5c94d4cce21569ba907",
			},
			{
				Id:                6,
				ConfigPublicKey:   "ocr2cfg_evm_cf506f962147740f2b93b82b610e232928c7bbd02c9fbca11cc10a189edc2336",
				OnchainPublicKey:  "ocr2on_evm_c83460dfc72dd5084668584ddd47a8b2a391a57c",
				OffchainPublicKey: "ocr2off_evm_202c4fc5b2335581ed4383ce16bceca3b4737cfc17b99271afe0493e72e2ce1e",
				PeerID:            "12D3KooWK5ScgUx1T2VoPkQrajH9U4Xm1ccmsXo3BFjHJ4YQymkC",
				OffChainKeyId:     "09f7615695972ad31af9d31e3417b284286b2b66d34b0e0bdc158b8127d256be",
				CSAPublicKey:      "csa_e7a7d61faf6584bbbe7137d79f59cc07ad4b89b7f6cd41742a5095e58616cd9c",
			}, {
				Id:                7,
				ConfigPublicKey:   "ocr2cfg_evm_8c3c46789493dda5aa1bbe6ba193905a77ca71c34ab34062ddbdc678da56843c",
				OnchainPublicKey:  "ocr2on_evm_ccba7e88767314bfba900b8e0ca72c288db97dcf",
				OffchainPublicKey: "ocr2off_evm_5f137c565c360cc59238c9908c10c186f3fe02c9118ffc6d50f2dd07e02e83ab",
				PeerID:            "12D3KooWEC4V2WRH1Um5rBMhmSo6V1YHvAut7ycSXf7yW38z2K1X",
				OffChainKeyId:     "0f26875e2f65b69153cd41227d68bb67f9056781064844dd4f7cfbe4a2d4c855",
				CSAPublicKey:      "csa_b19ab5aebd07a73a5f2479f4795603b46216ed2d37d5cfe111b84ad975508df6",
			}, {
				Id:                8,
				ConfigPublicKey:   "ocr2cfg_evm_e6ed4d36fe56ee2f087522a32dfbaf72666da064d12ceb85ffa879f3ee8b7b5e",
				OnchainPublicKey:  "ocr2on_evm_491b42dcd9f6ba70d2179d7fdfa19d2863450f7d",
				OffchainPublicKey: "ocr2off_evm_5002c14d4ccabca87ab8608ed4f16c4ac260084873fbdaf5501c13d4015359b9",
				PeerID:            "12D3KooWHWViS2gJ9L8zcmxzzYNiGzxKKZxoHZNUQiaRafUhmsZo",
				OffChainKeyId:     "83f4bac8c84210dbe2910cee9d9ed197b4beea9dfdb4dae0c0c51601e704c969",
				CSAPublicKey:      "csa_ecfad4b98a535fdc2c17d66373c4b92ff783f8ca039c3e19ea4bcd53259c18df",
			}, {
				Id:                9,
				ConfigPublicKey:   "ocr2cfg_evm_0bed8cc2e492b801e958c4f3d6145a4a066b59a746549afa07e8d2eca93a734b",
				OnchainPublicKey:  "ocr2on_evm_7ad69b3e1fe01bc4beb5551006c54eef931051a5",
				OffchainPublicKey: "ocr2off_evm_a70739848a1558213af29a96679493170c64490a6400ca5a6c658d8990d1abaa",
				PeerID:            "12D3KooWSYbns1kHv4dYQzYtXLm4pF1ouBZ8mXDGJqepfnvYHpQp",
				OffChainKeyId:     "91abef4325db612ea22a84d01235eed3f37502ab2705f5357326008376398074",
				CSAPublicKey:      "csa_fbc297e151feb5a1ed2a3ddf36fd359ed69290039f2daaa50f12e21a4e7850b5",
			},
		}
		nodeConfigs[MercuryValueless] = nodeConfigsMercuryValueless
	}

	return nodeConfigs[target]
}

func GetAproOCR2MercuryConfig(target int) ocr3confighelper.PublicConfig {
	nodeConfigs := GetNodeConfigs(target)
	ocrConfig := AproOCR2MercuryConfig(len(nodeConfigs), target)

	return ocrConfig
}

func TestEncodeOCR2MercuryV3Config(t *testing.T) {
	const target = MercuryBtc
	publicConfig := GetAproOCR2MercuryConfig(target)
	nodeConfigs := GetNodeConfigs(target)

	oracles := []confighelper.OracleIdentityExtra{}
	for _, nodeConfig := range nodeConfigs {
		// Need to convert the key representations
		offchainPkBytes, _ := hex.DecodeString(strings.TrimPrefix(nodeConfig.OffchainPublicKey, "ocr2off_evm_"))
		offchainPkBytesFixed := [ed25519.PublicKeySize]byte{}
		copy(offchainPkBytesFixed[:], offchainPkBytes)

		onchainPublicKey, _ := hex.DecodeString(strings.TrimPrefix(nodeConfig.OnchainPublicKey, "ocr2on_evm_"))

		configPkBytes, _ := hex.DecodeString(strings.TrimPrefix(nodeConfig.ConfigPublicKey, "ocr2cfg_evm_"))
		configPkBytesFixed := [curve25519.PointSize]byte{}
		copy(configPkBytesFixed[:], configPkBytes)

		csaPublicKey := strings.TrimPrefix(nodeConfig.CSAPublicKey, "csa_")
		oracleIdentity := confighelper.OracleIdentity{
			OffchainPublicKey: offchainPkBytesFixed,
			OnchainPublicKey:  onchainPublicKey,
			PeerID:            nodeConfig.PeerID,
			TransmitAccount:   types.Account(csaPublicKey),
		}
		oracleIdentityExtra := confighelper.OracleIdentityExtra{
			OracleIdentity:            oracleIdentity,
			ConfigEncryptionPublicKey: configPkBytesFixed,
		}
		oracles = append(oracles, oracleIdentityExtra)
	}

	onchainConfig, err := (ocr3confighelper.StandardOnchainConfigCodec{}).Encode(rawOnchainConfig)
	require.NoError(t, err)

	reportingPluginConfig, err := json.Marshal(rawReportingPluginConfig)
	require.NoError(t, err)

	signers, transmitters, f, onchainConfig_, offchainConfigVersion, offchainConfig, err := ocr3confighelper.ContractSetConfigArgsForTests(
		publicConfig.DeltaProgress,
		publicConfig.DeltaResend,
		publicConfig.DeltaInitial,
		publicConfig.DeltaRound,
		publicConfig.DeltaGrace,
		publicConfig.DeltaCertifiedCommitRequest,
		publicConfig.DeltaStage,
		publicConfig.RMax,
		publicConfig.S,
		oracles,
		reportingPluginConfig,
		publicConfig.MaxDurationQuery,
		publicConfig.MaxDurationObservation,
		publicConfig.MaxDurationShouldAcceptAttestedReport,
		publicConfig.MaxDurationShouldTransmitAcceptedReport,
		publicConfig.F,
		onchainConfig)

	signerAddresses, err := ocr3confighelper.OnchainPublicKeyToAddress(signers)

	fmt.Printf("\nerr: %v\n", err)
	fmt.Printf("signers: %v\n", PrintList(signerAddresses))
	fmt.Printf("transmitters: %v\n", PrintAccountList(transmitters))
	fmt.Printf("f: %v\n", f)
	fmt.Printf("onchainConfig_: %v\n", hexutil.Encode(onchainConfig_))
	fmt.Printf("offchainConfigVersion: %v\n", offchainConfigVersion)
	fmt.Printf("offchainConfig: %v\n", hexutil.Encode(offchainConfig))

}

func TestDecodeOCR2MercuryV3Config(t *testing.T) {
	onchainConfig_ := "0x0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000ffffffffffffffffffffffffffffffffffffffffffff"
	onchainConfigByte, _ := hex.DecodeString(strings.TrimPrefix(onchainConfig_, "0x"))
	onchainConfig, err := (ocr3confighelper.StandardOnchainConfigCodec{}).Decode(onchainConfigByte)
	fmt.Printf("\nerror: %v\n", err)
	fmt.Printf("onchainConfig: %v\n", onchainConfig)

	offchainConfig_ := "0xc80180a8d6b907d00180c8afa025d80180e59a77e00100e80100f00119fa010110820220adbdc190c436ab4b343b309743b18a64c2b7de4b4e9658391c1a9eed4cdabd398202205f11e6dac1f9da0b1e410a5e5f2a2a1573cbc26fb8fdcf542c46d63c7b5792a68202208119731aed54ea4652d069fb82572056f2006b610583ad5c8d614c0f8661cae382022024188be12cc34d68e86d957850576780908e607a7539b128c086988355d5f1618202209bc95b9d5c4a2cd2b6840e0b2fffa7c40614164ab4beb512740f42c6a302813a820220b235c7dfae283dba5621d846939221534db11206eaa437ba97b8e621e300344d820220f4371ccba90d1f7c83b5893fdb43f89e03b328f5e1567e99f69ba3b9977ffc37820220bcf86b630a1b06ee83912fe8c0d19e8435cf3a73c16a501aa88a4e83208899f1820220c156da1312684ea1f685591bcdcc730437d0a130b05ce7fc375f731a44f6ce1a8202205baa22ce196e7d5d25caabf139e4b390b1742a8fc045714ed1b402f6d4eff0e4820220f947f7747d6ab14e23c984058a4d248374132d83d5590dcbf99ba52b048583578202201ab82b86fd883b5d1425a1916a192f6cbb1f3188aeb03463a72053bc500f3a188202201bc86b584dad07be1547418173c2cd443f88762fa3c0e7622024b1ffcb1a346d820220c2816d32032c480aab56b0cb76eb2877108e607f1d4c574f9114d2a09a78e0e482022047df0aeb47d4288911b2b462a1a5d276b3dea1ab661d0d00bb1edabee7f7000f820220ada78d7a0c08b4a68fbff194d4a981cc849c0f26c8c8d6ee668a8ac4d5be18788a0234313244334b6f6f57416a7036414e6f6f645448557336616a455635377937524c637734386b43556a536f54664377354a7a4663568a0234313244334b6f6f57505a6f4c74426f56415a6d4536546974565466626d734674646e526831556d37676a6670563874444b4d4c318a0234313244334b6f6f574b543735767353474d704c4d546a724a3875373661754a524e4a524a44773472314c796b43646b3378564a468a0234313244334b6f6f57415a59517336594b393737794136617a566a4d43775869464a7a76773547445871426561515635334d6231518a0234313244334b6f6f574b396f355865624b44696b61693968664c7a636f53515150426f5033786d4d4c4437326746764836704176768a0234313244334b6f6f574458634a6d69777a4779784870596e79723736657944554231554e354d4b46466d4e354c75544c646f616a388a0234313244334b6f6f574e387667566a396866557565316d6a31696743767a6d6665384479616a677331616d727158766b364770416f8a0234313244334b6f6f574558726f744276584a78516e3853734c6e4447326f6a587177415670376a685458655766626e6b4131665a4c8a0234313244334b6f6f574a7a754154447a38454164564c37313676484432664e513455414672735967595450646732757848325679768a0234313244334b6f6f574c634538575662586267357933316b6372356e73386b7a454336654c724e5664535a7639554442685659764e8a0234313244334b6f6f574d44794e5371414a33716a74774a555533356f7351414a6851617052375a3662425a506f44593144754d6e508a0234313244334b6f6f574b486e47745553625268344d4e524c344a435a79793845665a4a74725054634344704273375a5474576846748a0234313244334b6f6f57484c50516a37565859597a434a444d6361674846544553426151756a4e4542474a74377463705461576531788a0234313244334b6f6f5753665a7639444339794656513577485543556247505673466f5870365a395a743847584a326d664c4d354a748a0234313244334b6f6f5741714d504831516e54487146525958796f626854384461595a654678474359324e7951584569386445766b628a0234313244334b6f6f574e59773556314e32425064467a62394c434c5a3859386f4538554b34546e706d38795073745a4344594b716a92022e7b2265787069726174696f6e57696e646f77223a38363430302c2262617365555344466565223a22302e3332227d980280e1eb17a00280e59a77a80280e1eb17b00280e1eb17ba02e4020a20dcdb057957c925cc06a740094bd4065ed6fcdbd953fabc312e9c71e9ea34466b1220b166d29341fc45265a75d4b3baad670fe64a2a22eaeb678cb0341d84103295201a10516c6339fd435a253262cef776daead71a10d345b89dda0070e65789a7e1c458e59a1a10c013fa2a577b5e24f71c10600217367d1a1004c3c2897246c91852e0cde2823710551a10c8948af828ab633b3810af36def7c0331a10ffefc2b4712761d105df47802c532a291a1012484092a276a4226388e8e3f449d26c1a1044eca8ecfc854275f6f1fbc9545056831a10407b07a38cd34d14f3aa8ac18caf725a1a102931bfaf345b7b6eabeb836bd1a6fdad1a108c8ed44bce681a086c433058b42625e01a10f66b43d3387ffea96c57c33eac41736b1a106c7a82fb39a1694d74dc429c8002ad311a108e86ae3bdba3f6f1625922ebc3c91c641a1080f78dfbfb6d2922695c21bc3f8993701a106b8dedce3bad3b30b0a55dcbda6a78bcc002808c8d9e02c8028094ebdc03"
	offchainConfigByte, _ := hex.DecodeString(strings.TrimPrefix(offchainConfig_, "0x"))
	offchainConfig, err := ocr3config.DeserializeOffchainConfig(offchainConfigByte)
	fmt.Printf("offchainConfig: %v\n", offchainConfig)

	//DeltaProgress                           time.Duration
	//DeltaResend                             time.Duration
	//DeltaInitial                            time.Duration
	//DeltaRound                              time.Duration
	//DeltaGrace                              time.Duration
	//DeltaCertifiedCommitRequest             time.Duration
	//DeltaStage                              time.Duration
	//RMax                                    uint64
	//S                                       []int
	//ReportingPluginConfig                   []byte
	//MaxDurationQuery                        time.Duration
	//MaxDurationObservation                  time.Duration
	//MaxDurationShouldAcceptAttestedReport   time.Duration
	//MaxDurationShouldTransmitAcceptedReport time.Duration
	fmt.Printf("DeltaProgress: %v\n", offchainConfig.DeltaProgress)
	fmt.Printf("DeltaResend: %v\n", offchainConfig.DeltaResend)
	fmt.Printf("DeltaInitial: %v\n", offchainConfig.DeltaInitial)
	fmt.Printf("DeltaRound: %v\n", offchainConfig.DeltaRound)
	fmt.Printf("DeltaGrace: %v\n", offchainConfig.DeltaGrace)
	fmt.Printf("DeltaCertifiedCommitRequest: %v\n", offchainConfig.DeltaCertifiedCommitRequest)
	fmt.Printf("DeltaStage: %v\n", offchainConfig.DeltaStage)
	fmt.Printf("RMax: %v\n", offchainConfig.RMax)
	fmt.Printf("S: %v\n", offchainConfig.S)
	fmt.Printf("MaxDurationQuery: %v\n", offchainConfig.MaxDurationQuery)
	fmt.Printf("MaxDurationObservation: %v\n", offchainConfig.MaxDurationObservation)
	fmt.Printf("MaxDurationShouldAcceptAttestedReport: %v\n", offchainConfig.MaxDurationShouldAcceptAttestedReport)
	fmt.Printf("MaxDurationShouldTransmitAcceptedReport: %v\n", offchainConfig.MaxDurationShouldTransmitAcceptedReport)

	reportingPluginConfig := ocr3confighelper.OffchainConfig{}
	_ = json.Unmarshal(offchainConfig.ReportingPluginConfig, &reportingPluginConfig)
	fmt.Printf("reportingPluginConfig: %v\n", reportingPluginConfig)

}

// NewHash return random Keccak256
func NewHash() common.Hash {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return common.BytesToHash(b)
}

func TestFeedIds(t *testing.T) {
	symbols := []string{
		"BTC/USD",
		"ETH/USD",
		"Valueless/USD",
	}

	for _, symbol := range symbols {
		symbolByte := []byte(symbol)
		feedId := sha3.NewLegacyKeccak256()
		feedId.Write(symbolByte)
		fmt.Printf("\nsymbol: %v, feedId: 0x0003%v", symbol, strings.TrimPrefix(hexutil.Encode(feedId.Sum(nil)), "0x")[4:])
	}
	fmt.Printf("\n\n")
}

func TestFeedId(t *testing.T) {
	symbol := "BTC/USD"
	symbolByte := []byte(symbol)

	feedId := sha3.NewLegacyKeccak256()
	feedId.Write(symbolByte)

	fmt.Printf("\nfeedId: 0x0003%v\n", strings.TrimPrefix(hexutil.Encode(feedId.Sum(nil)), "0x")[4:])
}

func PrintList(content []common.Address) string {
	txt := strings.Replace(fmt.Sprintf("%v", content), " ", "\",\"", -1)
	txt = strings.Replace(txt, "[", "[\"", -1)
	txt = strings.Replace(txt, "]", "\"]", -1)
	return txt
}

func PrintAccountList(content []types.Account) string {
	txt := strings.Replace(fmt.Sprintf("%v", content), " ", "\",\"0x", -1)
	txt = strings.Replace(txt, "[", "[\"0x", -1)
	txt = strings.Replace(txt, "]", "\"]", -1)
	return txt
}
