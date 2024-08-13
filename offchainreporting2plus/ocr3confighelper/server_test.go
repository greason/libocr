package ocr3confighelper

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"testing"
)

/*func startMercuryServer(t *testing.T, srv *mercuryServer, pubKeys []ed25519.PublicKey) (serverURL string) {
	// Set up the wsrpc server
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("[MAIN] failed to listen: %v", err)
	}
	serverURL = lis.Addr().String()
	s := wsrpc.NewServer(wsrpc.Creds(srv.privKey, pubKeys))

	// Register mercury implementation with the wsrpc server
	pb.RegisterMercuryServer(s, srv)

	// Start serving
	go s.Serve(lis)
	t.Cleanup(s.Stop)

	return
}*/

func TestServer(t *testing.T) {
	serverKey := MustNewV2XXXTestingOnly(big.NewInt(-1))
	serverPubKey := serverKey.PublicKey
	fmt.Printf("serverPubKey: %v\n", hexutil.Encode(serverPubKey))

}
