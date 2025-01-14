package signerv2_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/arithmic/eigensdk-go/signerv2"
	"github.com/arithmic/eigensdk-go/testutils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestPrivateKeySignerFn(t *testing.T) {
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	require.NoError(t, err)
	chainID := big.NewInt(1)

	signer, err := signerv2.PrivateKeySignerFn(privateKey, chainID)
	require.NoError(t, err)

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:   0,
		Value:   big.NewInt(0),
		ChainID: chainID,
		Data:    common.Hex2Bytes("6057361d00000000000000000000000000000000000000000000000000000000000f4240"),
	})
	signedTx, err := signer(address, tx)
	require.NoError(t, err)

	// Verify the sender address of the signed transaction
	from, err := types.Sender(types.LatestSignerForChainID(chainID), signedTx)
	require.NoError(t, err)
	require.Equal(t, address, from)
}

func TestKeyStoreSignerFn(t *testing.T) {
	keystorePath := "mockdata/dummy.key.json"
	keystorePassword := "testpassword"
	chainID := big.NewInt(1)
	signer, err := signerv2.KeyStoreSignerFn(keystorePath, keystorePassword, chainID)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("7a28b5ba57c53603b0b07b56bba752f7784bf506fa95edc395f5cf6c7514fe9d")
	require.NoError(t, err)

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:   0,
		Value:   big.NewInt(0),
		ChainID: chainID,
		Data:    common.Hex2Bytes("6057361d00000000000000000000000000000000000000000000000000000000000f4240"),
	})
	signedTx, err := signer(address, tx)
	require.NoError(t, err)

	// Verify the sender address of the signed transaction
	from, err := types.Sender(types.LatestSignerForChainID(chainID), signedTx)
	require.NoError(t, err)
	require.Equal(t, address, from)
}

func TestWeb3SignerFn(t *testing.T) {
	anvilC, err := testutils.StartAnvilContainer(testutils.GetDefaultTestConfig().AnvilStateFileName)
	require.NoError(t, err)

	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)

	signer, err := signerv2.Web3SignerFn(anvilHttpEndpoint)
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	require.NoError(t, err)
	anvilChainID := big.NewInt(31337)
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:   0,
		Value:   big.NewInt(0),
		To:      &address,
		ChainID: anvilChainID,
		Data:    common.Hex2Bytes("6057361d00000000000000000000000000000000000000000000000000000000000f4240"),
	})

	signedTx, err := signer(address, tx)
	require.NoError(t, err)

	// Verify the sender address of the signed transaction
	from, err := types.Sender(types.LatestSignerForChainID(anvilChainID), signedTx)
	require.NoError(t, err)
	require.Equal(t, address, from)
}
