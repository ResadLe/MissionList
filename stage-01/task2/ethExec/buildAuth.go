package ethExec

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func buildAuth(privKey *ecdsa.PrivateKey, client *ethclient.Client, chainID int64, ctx context.Context) (*bind.TransactOpts, error) {
	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("invalid public key")
	}

	addr := crypto.PubkeyToAddress(*pubKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("get nonce: %w", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("get gas price: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(chainID))
	if err != nil {
		return nil, fmt.Errorf("transactor: %w", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	return auth, nil
}
