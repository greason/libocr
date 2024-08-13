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
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/curve25519"
	"math/big"
	"strings"
	"testing"
	"time"
)

var (
	// 2 ** 176 -1
	maxValue, _      = big.NewInt(1).SetString("95780971304118053647396689196894323976171195136475135", 10)
	rawOnchainConfig = OnchainConfig{
		Min: big.NewInt(1),
		Max: maxValue,
	}
	rawReportingPluginConfig = OffchainConfig{
		ExpirationWindow: 86400, //
		BaseUSDFee:       decimal.NewFromFloat32(0.64),
		//BaseUSDFee:       decimal.NewFromFloat32(0.000000000000000001),
	}
)

func AproOCR2MercuryConfig(numberNodes int, target int) PublicConfig {
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

		DeltaProgress = time.Second * 35
		DeltaRound = time.Second * 30
	case MercuryEth:
		DeltaProgress = time.Second * 35
		DeltaRound = time.Second * 30
	case MercuryLink:
		DeltaProgress = time.Second * 35
		DeltaRound = time.Second * 30
	}

	// 参考：offchainreporting2plus/internal/config/ocr3config
	return PublicConfig{
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
	MercuryLink
)

func GetNodeConfigs(target int) []NodeOCR2MercuryConfig {
	nodeConfigs := make(map[int][]NodeOCR2MercuryConfig)

	{
		nodeConfigsMercuryBtc := []NodeOCR2MercuryConfig{
			{
				Id:                1,
				ConfigPublicKey:   "a657008c2652576fbeee83449e3e464a7c3512262f34ab13588eadbf9c1aab08",
				OnchainPublicKey:  "a33d88f4283b7916de3c36ca64632b68cb74577b",
				OffchainPublicKey: "5b682bda0deae08968686eb4ce227d8162e7a09192d5f58f9a84ee12317a6e78",
				PeerID:            "12D3KooWMM4ufpHRmAuXBhwWeY6V12UMKjonkg1YR82PPLYMsYKS",
				OffChainKeyId:     "012d9a770762ade43e3ad11b448904ba747e6ba95fae0ecfe6527163338c35c3",
				CSAPublicKey:      "47c7c72b247f52897d3eae93c451035e600dd2c5b51636929961d53a3feab825",
			}, {
				Id:                2,
				ConfigPublicKey:   "e194cc8de045254f486af65dcbdd1843433036db321a7af114cfbd3f49bde82e",
				OnchainPublicKey:  "7f754db02aab92f888db92fe6bab4d4dbe1982ff",
				OffchainPublicKey: "449c08688d0f98e1317c9016c7c845554f24f8cf831ad6ce703940e6f2fc28dc",
				PeerID:            "12D3KooWPjc5CeHTBc4LL9Z3ekyfEbVHhP85MqwffJfmfP68y3h7",
				OffChainKeyId:     "90a4ceef1170a77753f34adfacf3ce87df95b6a17994822af7831f9e41da9f42",
				CSAPublicKey:      "66ed61baa2c6e8ff8d6db48ba9d6cf7f2f212bf922dff57eda2a8558d65e69c0",
			}, {
				Id:                3,
				ConfigPublicKey:   "f1f72bc1a14400d286e6cc22b6c9a0eeeb94fdb6e1da84ce744ac64fc3e8852f",
				OnchainPublicKey:  "96ac87edf0570ccf1b806dc7b6bf7afe74ceaad4",
				OffchainPublicKey: "a75e55406e33f5d8b875518a5eac470f3578165e90cb14421902f2cfe3204f20",
				PeerID:            "12D3KooWAp9d7Hn4uw4nEwpmvUVWCoRH3qbuLhgtayDLNm7FK1KH",
				OffChainKeyId:     "84147e6accb355dec97ad0f19b6debc1bdd240545c7b3fded314171c7fa40f7f",
				CSAPublicKey:      "424661e19e2aef1cfe643202d54f0e298ec1315459bc8bf46f6b9177e4b0d223",
			}, {
				Id:                4,
				ConfigPublicKey:   "c76605c57b95ba0b42688b4666296ba69caabe34a599c7442f63f6624c4fc338",
				OnchainPublicKey:  "5a650ec7b903b33fbe1c650f107f26b9b8a23ddc",
				OffchainPublicKey: "b17cd8bb04c528836a1cf0b36a813d40073db398179ce08de17b55ee0d5462e9",
				PeerID:            "12D3KooWMDF2YxveKKa7rYN4fM7Uz46basme4fAHR28uwEoeKu45",
				OffChainKeyId:     "3a811c7747c096221b024ff26b1e11232d999e91e7bc1369b43aea4a2b23af9c",
				CSAPublicKey:      "f73ad354bde5da64affdff723ff4f0ac3d9a1d667d39b2dac7f1c75fedeab0e5",
			}, //local
			/*{
				Id:                5,
				ConfigPublicKey:   "93704424f889ebbe1927b27525f4da684864a44d2eaff68b4fc445bc8fef141f",
				OnchainPublicKey:  "438bf1231c808ad81a03dbc4723970ec91600f5f",
				OffchainPublicKey: "e8de8cb3b0e5aee85ca8f1cf44d5801dbccfec33fbf8e71f4c635611dfcdcfb2",
				PeerID:            "12D3KooWREWK4CwDCDWTWWa77c7iJLzendzv5zBEX8uS61V1Xi59",
				OffChainKeyId:     "3d1117c6c53990302ad6ebc775674cd20078b94a2920b2c5b6c46f62760432c0",
				CSAPublicKey:      "e93672b0df5f282fe6ce0385074461419f7dfae7158e3d074a21b909330a6495",
			},*/
		}
		nodeConfigs[MercuryBtc] = nodeConfigsMercuryBtc
	}

	{
		nodeConfigsMercuryEth := []NodeOCR2MercuryConfig{
			{
				Id:                1,
				ConfigPublicKey:   "ocr2cfg_evm_e9d1b5739d56f5db7a6df1ef6e98abc9d7e4eb7f83ee2896277433e7a095de68",
				OnchainPublicKey:  "ocr2on_evm_455d111e29d14fdd6bafb406b7f6b4e9b8f8398b",
				OffchainPublicKey: "ocr2off_evm_bc79751f3b8038d541a628ec55c4261f2b112ed164aa50fd3919bdd959937dd7",
				PeerID:            "12D3KooWMM4ufpHRmAuXBhwWeY6V12UMKjonkg1YR82PPLYMsYKS",
				OffChainKeyId:     "0e67918800cf03e82c6b64bc29501903479e08efe61efc70404de449242616da",
				CSAPublicKey:      "47c7c72b247f52897d3eae93c451035e600dd2c5b51636929961d53a3feab825",
			}, {
				Id:                2,
				ConfigPublicKey:   "ocr2cfg_evm_5e44bae57b8c0ff9c70ae5d9cc78743728ddb3b8b85ca711472ca62d70890c6c",
				OnchainPublicKey:  "ocr2on_evm_e4b28adb88623b55810ea8de38f165efdfddd075",
				OffchainPublicKey: "ocr2off_evm_ae38f55d982dca6b31bc7ba063705dd036c4f105228319a718ff5479a737df7f",
				PeerID:            "12D3KooWPjc5CeHTBc4LL9Z3ekyfEbVHhP85MqwffJfmfP68y3h7",
				OffChainKeyId:     "3bfa46d0fb01ed08021933a801a363ffaf62e6552bc9a35312cec3689d56fd75",
				CSAPublicKey:      "66ed61baa2c6e8ff8d6db48ba9d6cf7f2f212bf922dff57eda2a8558d65e69c0",
			}, {
				Id:                3,
				ConfigPublicKey:   "ocr2cfg_evm_e47129509a26a213d10385a127aeeed724ff0050f960a81b172c35d0e29ee662",
				OnchainPublicKey:  "ocr2on_evm_1a8b51f99e67d1fa693593724222e4e445d64061",
				OffchainPublicKey: "ocr2off_evm_19d5e4baf4e60594dbf4b12fb905489dfcb93536f9a4d34eda69516e5859a8d7",
				PeerID:            "12D3KooWAp9d7Hn4uw4nEwpmvUVWCoRH3qbuLhgtayDLNm7FK1KH",
				OffChainKeyId:     "1018fe47ca4a486febabfd5ed92c6d6beae2f88402b3feacdc524bfd3f6f824d",
				CSAPublicKey:      "424661e19e2aef1cfe643202d54f0e298ec1315459bc8bf46f6b9177e4b0d223",
			}, {
				Id:                4,
				ConfigPublicKey:   "ocr2cfg_evm_8ee1ca170a6c3768013ba5af9546daeef7b625b5ca88fa506fe3d60e908b1f7a",
				OnchainPublicKey:  "ocr2on_evm_864f1e7a341599c769654c0d78b35f9990bf7dae",
				OffchainPublicKey: "ocr2off_evm_1d53341aa1d719ddb94fd4da2567aff2705b076d236170b0147a0c543b8c263b",
				PeerID:            "12D3KooWMDF2YxveKKa7rYN4fM7Uz46basme4fAHR28uwEoeKu45",
				OffChainKeyId:     "474bc10800130147f52d5c9c6ae9468d7b8e463e3e6cbaf3439157c970acdf79",
				CSAPublicKey:      "f73ad354bde5da64affdff723ff4f0ac3d9a1d667d39b2dac7f1c75fedeab0e5",
			},
		}
		nodeConfigs[MercuryEth] = nodeConfigsMercuryEth
	}

	{
		nodeConfigsMercuryLink := []NodeOCR2MercuryConfig{
			{
				Id:                1,
				ConfigPublicKey:   "ocr2cfg_evm_8d09c3d1eb07079cb4ead6bb9b66a89db292c4bbc77215c76fe0b42c0a1df375",
				OnchainPublicKey:  "ocr2on_evm_43eae3e84ced52e3d51a57a69adb89ac3d5e4d47",
				OffchainPublicKey: "ocr2off_evm_592bec4f1269f2810921343581c27e3641177753b3050ed23ac2b5c4c6f70b3d",
				PeerID:            "12D3KooWMM4ufpHRmAuXBhwWeY6V12UMKjonkg1YR82PPLYMsYKS",
				OffChainKeyId:     "c5f0ae2bd8422f7734cd1d507a45aa9ba18e7ba7f2e0bd2e40a5bc9fe9060b7d",
				CSAPublicKey:      "47c7c72b247f52897d3eae93c451035e600dd2c5b51636929961d53a3feab825",
			}, {
				Id:                2,
				ConfigPublicKey:   "ocr2cfg_evm_8aef7740beb691b5ad5c923f888899177c346adb10610159e1178e5db6e37174",
				OnchainPublicKey:  "ocr2on_evm_8664c5b40fd6491c308833d4916c48531c437110",
				OffchainPublicKey: "ocr2off_evm_246e6e252bfbf0ee58b2bf397da566a45a985feea4707f85e81e476cf7dca015",
				PeerID:            "12D3KooWPjc5CeHTBc4LL9Z3ekyfEbVHhP85MqwffJfmfP68y3h7",
				OffChainKeyId:     "07712426b38a57d77966aeef2885e9c0df18aab7ec466fbd477e586bd904fafa",
				CSAPublicKey:      "66ed61baa2c6e8ff8d6db48ba9d6cf7f2f212bf922dff57eda2a8558d65e69c0",
			}, {
				Id:                3,
				ConfigPublicKey:   "ocr2cfg_evm_a8424bb03176b5973ea21651a44e8c95ddf7819a3a354b873d9baf4fd5456d15",
				OnchainPublicKey:  "ocr2on_evm_0060c22f693a7023f4ab0f484b234a55b1fd0feb",
				OffchainPublicKey: "ocr2off_evm_b0d3c6ebe0ff2d55269149426ab0a670eadd1da41caa8a994434d2a020927681",
				PeerID:            "12D3KooWAp9d7Hn4uw4nEwpmvUVWCoRH3qbuLhgtayDLNm7FK1KH",
				OffChainKeyId:     "1d30e5c93ba705721c884165edac901fceae3f5825a52b368a65da093dc9049a",
				CSAPublicKey:      "424661e19e2aef1cfe643202d54f0e298ec1315459bc8bf46f6b9177e4b0d223",
			}, {
				Id:                4,
				ConfigPublicKey:   "ocr2cfg_evm_93436f43f26a3d775e9304f7c7d2294bc7d125d32f2f930e8a56347b3e85f779",
				OnchainPublicKey:  "ocr2on_evm_2a2edd443183162a9b001695acf6f48b7161c947",
				OffchainPublicKey: "ocr2off_evm_0f3ca5c5c1cd7baf1694320b6042ea8212b9c70b335bdf921b55add05460d552",
				PeerID:            "12D3KooWMDF2YxveKKa7rYN4fM7Uz46basme4fAHR28uwEoeKu45",
				OffChainKeyId:     "47534d0735e7f155bd7893b1e383cb0187ec2e87bd48bf98a58902306688afac",
				CSAPublicKey:      "f73ad354bde5da64affdff723ff4f0ac3d9a1d667d39b2dac7f1c75fedeab0e5",
			},
		}
		nodeConfigs[MercuryLink] = nodeConfigsMercuryLink
	}

	return nodeConfigs[target]
}

func GetAproOCR2MercuryConfig(target int) PublicConfig {
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

		oracleIdentity := confighelper.OracleIdentity{
			OffchainPublicKey: offchainPkBytesFixed,
			OnchainPublicKey:  onchainPublicKey,
			PeerID:            nodeConfig.PeerID,
			TransmitAccount:   types.Account(nodeConfig.CSAPublicKey),
		}
		oracleIdentityExtra := confighelper.OracleIdentityExtra{
			OracleIdentity:            oracleIdentity,
			ConfigEncryptionPublicKey: configPkBytesFixed,
		}
		oracles = append(oracles, oracleIdentityExtra)
	}

	onchainConfig, err := (StandardOnchainConfigCodec{}).Encode(rawOnchainConfig)
	require.NoError(t, err)

	reportingPluginConfig, err := json.Marshal(rawReportingPluginConfig)
	require.NoError(t, err)

	signers, transmitters, f, onchainConfig_, offchainConfigVersion, offchainConfig, err := ContractSetConfigArgsForTests(
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

	signerAddresses, err := OnchainPublicKeyToAddress(signers)

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
	onchainConfig, err := (StandardOnchainConfigCodec{}).Decode(onchainConfigByte)
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

	reportingPluginConfig := OffchainConfig{}
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

func TestFeedId(t *testing.T) {
	feedId := NewHash().Bytes()
	fmt.Printf("\nfeedId: 0x0003%v\n", strings.TrimPrefix(hexutil.Encode(feedId), "0x")[4:])
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
