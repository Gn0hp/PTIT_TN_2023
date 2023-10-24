package main

import (
	election_contract "PTIT_TN/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"logur.dev/logur"
	"math/big"
)

func New(logger logur.LoggerFacade, chainId *big.Int, optsW *bind.TransactOpts, optsR *bind.CallOpts, instance *election_contract.ElectionContract) *ContractInstance {
	return &ContractInstance{
		ChainID:           chainId,
		AccountBoundWrite: optsW,
		AccountBoundRead:  optsR,
		Instance:          instance,
		Logger:            logger,
	}
}

func TransactOptsAccountBinding(chainId *big.Int, prvKey string) (*bind.TransactOpts, error) {
	toEcdsa, err := crypto.HexToECDSA(prvKey)
	if err != nil {
		return nil, err
	}
	user, err := bind.NewKeyedTransactorWithChainID(toEcdsa, chainId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
