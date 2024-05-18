package services

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type EthereumService struct {
    client *ethclient.Client
    rpcClient *rpc.Client
    privateKey *ecdsa.PrivateKey
    fromAddress common.Address
}

func NewEthereumService(url, privateKeyHex string) (*EthereumService, error) {
    rpcClient, err := rpc.Dial(url)
    if err != nil {
        return nil, err
    }

    client := ethclient.NewClient(rpcClient)

    privateKey, err := crypto.HexToECDSA(privateKeyHex)
    if err != nil {
        return nil, err
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return nil, err
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    return &EthereumService{
        client: client,
        rpcClient: rpcClient,
        privateKey: privateKey,
        fromAddress: fromAddress,
    }, nil
}

func (es *EthereumService) SendTransaction(toAddress string, amount *big.Int) (string, error) {
    nonce, err := es.client.PendingNonceAt(context.Background(), es.fromAddress)
    if err != nil {
        return "", err
    }

    gasPrice, err := es.client.SuggestGasPrice(context.Background())
    if err != nil {
        return "", err
    }

    to := common.HexToAddress(toAddress)
    tx := types.NewTransaction(nonce, to, amount, 21000, gasPrice, nil)

    chainID, err := es.client.NetworkID(context.Background())
    if err != nil {
        return "", err
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), es.privateKey)
    if err != nil {
        return "", err
    }

    err = es.client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        return "", err
    }

    return signedTx.Hash().Hex(), nil
}
