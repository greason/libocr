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
	"golang.org/x/crypto/sha3"
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
		BaseUSDFee:       decimal.NewFromFloat32(0.35),
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
	MercuryValueless
)

func GetNodeConfigs(target int) []NodeOCR2MercuryConfig {
	nodeConfigs := make(map[int][]NodeOCR2MercuryConfig)

	{
		nodeConfigsMercuryBtc := []NodeOCR2MercuryConfig{
			{
				Id:                1,
				ConfigPublicKey:   "ocr2cfg_evm_83b7e8a0b47d51464ca457cbb286c4f8a65dd80fa4bf9a0b629a2c2003344838",
				OnchainPublicKey:  "ocr2on_evm_f5561111581cff697b71d4e97c59f9848c3db18a",
				OffchainPublicKey: "ocr2off_evm_9ae13859821d7a939a122f1858fb28639d0a92f09b5e1cd89e94d26a3a97078d",
				PeerID:            "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:     "049d6485d5957b4c67d320a3f050a5dbeab26260280b7fc98f78ca823596c9b6",
				CSAPublicKey:      "78804c63374be97b450edbb37213c230a187fec0157177f5b7dce98038cab7c7",
			}, {
				Id:                2,
				ConfigPublicKey:   "ocr2cfg_evm_f751156c56472a8f8ffdc2bdc3d42932cb9defbca3e0d850cf74c961b0b8bb3b",
				OnchainPublicKey:  "ocr2on_evm_4e68cb4feedadbd4d2322cfdfca610cf99631f28",
				OffchainPublicKey: "ocr2off_evm_a5e8bde6f541ee583ce8c6bc960f84488a9473f118e0193e7e808741979dd11e",
				PeerID:            "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:     "0ce7af053f9bd94dfbece287b343a3aee22b735071bc4eb1d1874df6ddfda325",
				CSAPublicKey:      "7d1fb12ba8ea979489a6aa5a48198002ecd36035da18332acfa58c82491698c3",
			}, {
				Id:                3,
				ConfigPublicKey:   "ocr2cfg_evm_3f4ed8547032974dca1310cde2415ff22e872481ba862961ff76210a40eb844a",
				OnchainPublicKey:  "ocr2on_evm_021c60e8d51de06396909ddfe8bb492e7e2db0f9",
				OffchainPublicKey: "ocr2off_evm_0ab32a81afc1c25b4e02bffc4db2439c81063e3beebf213bf80f05ac8c052408",
				PeerID:            "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:     "f88f8fcd68ac411a81b6bfea6e3a8d9b247fdf9304c45d4c75db9ed33ac18a5c",
				CSAPublicKey:      "699dd4f2b22fb5698e3327b5b6b4e46fe9d104b88f80b06f2ff259a5c5d821c2",
			}, {
				Id:                4,
				ConfigPublicKey:   "ocr2cfg_evm_2e21ca6676a00c9972ad82870a9eabf8bb5df8e5644a14895a5d5ec6499a944a",
				OnchainPublicKey:  "ocr2on_evm_0c335747dcbd99c45a869c1b18f5bb5da4e9f748",
				OffchainPublicKey: "ocr2off_evm_a614f4d7170cfe989ff5d72ca8148af2da2e54fe933617ac715c381166f7da9a",
				PeerID:            "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:     "96aa0f3d1d1d43ed5a00623ceaf439d0bbc83c6f9a580bfeaaa1e9af9a6d7867",
				CSAPublicKey:      "0d52032dc5e6a9eee050386d3e7fd6c413a722a6490cb3fb01f17bc18b338117",
			},
			{
				Id:                5,
				ConfigPublicKey:   "ocr2cfg_evm_0aed91881e86e5f0e56d52a9fb29c56d92e6eeccddec55fc85a62f177e9f6a71",
				OnchainPublicKey:  "ocr2on_evm_9ca4e8204f780da1c7a654dfc7fccbce61829482",
				OffchainPublicKey: "ocr2off_evm_1b2f956664ec31c6aed074b2dbd57bb8da8a27f574ca0b1ace16fe6d1b5fb066",
				PeerID:            "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:     "f0023c78e719c26878f351ade1f8b80bf5ae403c953c49226155294c7305a7b7",
				CSAPublicKey:      "8831a97b40f8a9ec8aecc4778fb5b0e80eac9f7657056bea34d2b30c14f29adb",
			},
		}
		nodeConfigs[MercuryBtc] = nodeConfigsMercuryBtc
	}

	{
		nodeConfigsMercuryEth := []NodeOCR2MercuryConfig{
			{
				Id:                1,
				ConfigPublicKey:   "ocr2cfg_evm_0a0ad55b6513b8356a63572734966d5ac2292f94d0703081540a806065777a13",
				OnchainPublicKey:  "ocr2on_evm_9b0215d21af8d04d13fe0c8e498d33ce52d853f7",
				OffchainPublicKey: "ocr2off_evm_02640fef80e7a541e1e38d181c61c2908d96daaf9c99b28337ffe7bffdcf02ef",
				PeerID:            "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:     "ad99fe4d2e6be84b60a538b0f154b5b5674b967d3563e73270c13684db7a95da",
				CSAPublicKey:      "78804c63374be97b450edbb37213c230a187fec0157177f5b7dce98038cab7c7",
			}, {
				Id:                2,
				ConfigPublicKey:   "ocr2cfg_evm_986c9f09b1149d6730bf34533832fe5a5b9fb16251deaded011e53a1c16cbd76",
				OnchainPublicKey:  "ocr2on_evm_4a16ac188a18ddb26760b669fd71c669b0084156",
				OffchainPublicKey: "ocr2off_evm_b8f4ad50a4da040a5a6c06fcfd040ed78019810bf79ea9b6702d96c440c8607b",
				PeerID:            "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:     "7abfb994e8e8f378ba480a848fad6aff501cc72edb6c98078eaf825e23ce9fb5",
				CSAPublicKey:      "7d1fb12ba8ea979489a6aa5a48198002ecd36035da18332acfa58c82491698c3",
			}, {
				Id:                3,
				ConfigPublicKey:   "ocr2cfg_evm_159f8ba5a42ef80714672644815edac86ec4e31881754e211e8f5eda38d50930",
				OnchainPublicKey:  "ocr2on_evm_6f6a00822333913bfc014093b8b060b1f2212e17",
				OffchainPublicKey: "ocr2off_evm_917b2c33e698bafa885d104589dc8d93892c742d5df32126faeb797d8ef2bae5",
				PeerID:            "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:     "234881bd69588ae6b16dced968320b7984c0aae89c105e09b8459a7b919f3edb",
				CSAPublicKey:      "699dd4f2b22fb5698e3327b5b6b4e46fe9d104b88f80b06f2ff259a5c5d821c2",
			}, {
				Id:                4,
				ConfigPublicKey:   "ocr2cfg_evm_7f194cad00f293d0adac3b2a03a66240b34ac71c5ca5be5ba0e8803a1b269c49",
				OnchainPublicKey:  "ocr2on_evm_8e4e2093d5b629a0b884f567e6867c052a65205c",
				OffchainPublicKey: "ocr2off_evm_b33cc33a7f942e52aaba9dbecbcb3ce579319cf8a05e88b6be2d70235a3c1da7",
				PeerID:            "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:     "5995a0826a84d115eb1e05978e0e55ef914e229ca83e672f367f9403ef5a6c30",
				CSAPublicKey:      "0d52032dc5e6a9eee050386d3e7fd6c413a722a6490cb3fb01f17bc18b338117",
			},
			{
				Id:                5,
				ConfigPublicKey:   "ocr2cfg_evm_3696551c72cfa5150936f5bb9c462bb38e4c9f24518e36e2fd11a06756f0d206",
				OnchainPublicKey:  "ocr2on_evm_5df501ad08f0fcaecb1e1cf5dfe96bdfd3bf736f",
				OffchainPublicKey: "ocr2off_evm_4639efe838061ad3fd62ce005917ce0702d6dad7f1834a53736ad61911e3d8d6",
				PeerID:            "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:     "133c76025ac88eba10367b8eb639e2bd07339e4e63b974946f3298962422f652",
				CSAPublicKey:      "8831a97b40f8a9ec8aecc4778fb5b0e80eac9f7657056bea34d2b30c14f29adb",
			},
		}
		nodeConfigs[MercuryEth] = nodeConfigsMercuryEth
	}

	{
		nodeConfigsMercuryValueless := []NodeOCR2MercuryConfig{
			{
				Id:                1,
				ConfigPublicKey:   "ocr2cfg_evm_cbc43f6e67d206a89ee97499f533dee52895fdcd7b53da942fe421248a893d23",
				OnchainPublicKey:  "ocr2on_evm_e29fc54bc256762dfcb62318a279c379c96347f0",
				OffchainPublicKey: "ocr2off_evm_326a134166ed00d53028aff6fd2b0431795a5634d6691c9bcd4305385e759202",
				PeerID:            "12D3KooWJCEsfgchffSMFo3WWpJaeVKpb1cx5iUhax7GPGXmvpto",
				OffChainKeyId:     "ffbd4c8ac28d23b6e3de2d7583aff0dcc2f199d2fd237bcc360b82cdcf9a7c4e",
				CSAPublicKey:      "78804c63374be97b450edbb37213c230a187fec0157177f5b7dce98038cab7c7",
			}, {
				Id:                2,
				ConfigPublicKey:   "ocr2cfg_evm_806a027619e18403f8dd440031ae3f99defaf2800d7a4456498c53d45f18791d",
				OnchainPublicKey:  "ocr2on_evm_b279a37068c323a1bc5ee23ed0b2bbcd5df2e613",
				OffchainPublicKey: "ocr2off_evm_43f442c5b58ef30f64938d43cb5d537e9eac02b74613fb73f46b7efa9d9df112",
				PeerID:            "12D3KooWK2N5cverNrfdu7DswaNGFu4iCFG1dgwgotY7iVkQNE1F",
				OffChainKeyId:     "e944dcce865bc5fc871fd3129809ab55a7dc14107a702e954241f6c283aa42a1",
				CSAPublicKey:      "7d1fb12ba8ea979489a6aa5a48198002ecd36035da18332acfa58c82491698c3",
			}, {
				Id:                3,
				ConfigPublicKey:   "ocr2cfg_evm_f3c7c0523033e2bcb9188c20dd6beb14cff8fecd70de98fe1f0f7cee9f193821",
				OnchainPublicKey:  "ocr2on_evm_3509dcd08d1abbc592bb32206dade74bbeab90e3",
				OffchainPublicKey: "ocr2off_evm_9e8766aff62ef581f1890059db71f5e17ee23da4fc13ec078ffa4dbb4762b9e2",
				PeerID:            "12D3KooWDjoTCv3HBUfVGTBxo9z4zjsVYdDSPaUWZKZKFKKJ6akq",
				OffChainKeyId:     "3a59b9dc7d8d62e534a1d78fbf90af6abf9b24c66c47a474d827c938eda981a5",
				CSAPublicKey:      "699dd4f2b22fb5698e3327b5b6b4e46fe9d104b88f80b06f2ff259a5c5d821c2",
			}, {
				Id:                4,
				ConfigPublicKey:   "ocr2cfg_evm_839a5bfed37a71fdd79a8f9657ceb70671eb6b16ad155ce76aaca03c9b65c43d",
				OnchainPublicKey:  "ocr2on_evm_8b2a235cac379fe9a6e8eaba26356c142232ab72",
				OffchainPublicKey: "ocr2off_evm_9e8a51e99182c980e623a280ef93199ae515d879eff64f9644ecfaf7052b7d09",
				PeerID:            "12D3KooWKaVh29LwUq9NfvRQw8nFuzUJrPfYzpzQFWPSCwnpkhij",
				OffChainKeyId:     "21977705b40ffa4b2aa02501d19f05d0e1856d10a2105259050cae6dc7fd423c",
				CSAPublicKey:      "0d52032dc5e6a9eee050386d3e7fd6c413a722a6490cb3fb01f17bc18b338117",
			},
			{
				Id:                5,
				ConfigPublicKey:   "ocr2cfg_evm_e05de136cff63ec27c97cdc9446657f9ce39b5b13fc6cea1b208e685cf06b303",
				OnchainPublicKey:  "ocr2on_evm_9d354edf39d6d4b69679218b4c4616ffb0366e5f",
				OffchainPublicKey: "ocr2off_evm_85908c87a3a4d2f1bb2869baac645eb518c226ac7696899949db755ee2e684f6",
				PeerID:            "12D3KooWC5Bi42rp3gH9p3DmDCk4HVgyA67BcXTDhWRtp8sdwWcz",
				OffChainKeyId:     "9bc62bf48e181cad07619e75b126293f65cc920e0564f5820e92d95b1f274349",
				CSAPublicKey:      "8831a97b40f8a9ec8aecc4778fb5b0e80eac9f7657056bea34d2b30c14f29adb",
			},
		}
		nodeConfigs[MercuryValueless] = nodeConfigsMercuryValueless
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
